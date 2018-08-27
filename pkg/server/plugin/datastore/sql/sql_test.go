package sql

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spiffe/spire/proto/common"
	spi "github.com/spiffe/spire/proto/common/plugin"
	"github.com/spiffe/spire/proto/server/datastore"
	testutil "github.com/spiffe/spire/test/util"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	ctx = context.Background()
)

func TestPlugin(t *testing.T) {
	suite.Run(t, new(PluginSuite))
}

type PluginSuite struct {
	suite.Suite
	dir string

	nextId int
	ds     datastore.Plugin
}

func (s *PluginSuite) SetupSuite() {
	var err error
	s.dir, err = ioutil.TempDir("", "spire-datastore-sql-tests")
	s.Require().NoError(err)
}

func (s *PluginSuite) SetupTest() {
	s.ds = s.newPlugin()
}

func (s *PluginSuite) TearDownSuite() {
	os.RemoveAll(s.dir)
}

func (s *PluginSuite) newPlugin() datastore.Plugin {
	p := New()

	s.nextId++
	dbPath := filepath.Join(s.dir, fmt.Sprintf("db%d.sqlite3", s.nextId))

	_, err := p.Configure(context.Background(), &spi.ConfigureRequest{
		Configuration: fmt.Sprintf(`
		database_type = "sqlite3"
		log_sql = true
		connection_string = "file://%s"
		`, dbPath),
	})
	s.Require().NoError(err)

	return p
}

func (s *PluginSuite) TestInvalidPluginConfiguration() {
	_, err := s.ds.Configure(context.Background(), &spi.ConfigureRequest{
		Configuration: `
		database_type = "wrong"
		connection_string = "bad"
		`,
	})
	s.Require().EqualError(err, "datastore-sql: unsupported database_type: wrong")
}

func (s *PluginSuite) TestBundleCRUD() {
	cert, _, err := testutil.LoadSVIDFixture()
	s.Require().NoError(err)

	bundle := &datastore.Bundle{
		TrustDomain: "spiffe://foo",
		CaCerts:     cert.Raw,
	}

	// create
	_, err = s.ds.CreateBundle(ctx, &datastore.CreateBundleRequest{
		Bundle: bundle,
	})
	s.Require().NoError(err)

	// fetch
	fresp, err := s.ds.FetchBundle(ctx, &datastore.FetchBundleRequest{TrustDomain: "spiffe://foo"})
	s.Require().NoError(err)
	s.Equal(bundle, fresp.Bundle)

	// list
	lresp, err := s.ds.ListBundles(ctx, &datastore.ListBundlesRequest{})
	s.Require().NoError(err)
	s.Equal(1, len(lresp.Bundles))
	s.Equal(bundle, lresp.Bundles[0])

	cert, _, err = testutil.LoadCAFixture()
	s.Require().NoError(err)

	bundle2 := &datastore.Bundle{
		TrustDomain: bundle.TrustDomain,
		CaCerts:     cert.Raw,
	}

	// append
	aresp, err := s.ds.AppendBundle(ctx, &datastore.AppendBundleRequest{
		Bundle: bundle2,
	})
	s.Require().NoError(err)
	certs := append(bundle.CaCerts, cert.Raw...)
	s.Require().NotNil(aresp.Bundle)
	s.Equal(certs, aresp.Bundle.CaCerts)

	// append identical
	aresp, err = s.ds.AppendBundle(ctx, &datastore.AppendBundleRequest{
		Bundle: bundle2,
	})
	s.Require().NoError(err)
	s.Require().NotNil(aresp.Bundle)
	s.Equal(certs, aresp.Bundle.CaCerts)

	// append on a new bundle
	bundle3 := &datastore.Bundle{
		TrustDomain: "spiffe://bar",
		CaCerts:     cert.Raw,
	}
	anresp, err := s.ds.AppendBundle(ctx, &datastore.AppendBundleRequest{
		Bundle: bundle3,
	})
	s.Require().NoError(err)
	s.Equal(bundle3, anresp.Bundle)

	// update
	uresp, err := s.ds.UpdateBundle(ctx, &datastore.UpdateBundleRequest{
		Bundle: bundle2,
	})
	s.Require().NoError(err)
	s.Equal(bundle2, uresp.Bundle)

	lresp, err = s.ds.ListBundles(ctx, &datastore.ListBundlesRequest{})
	s.Require().NoError(err)
	s.Equal(2, len(lresp.Bundles))
	s.Equal([]*datastore.Bundle{bundle2, bundle3}, lresp.Bundles)

	// delete
	dresp, err := s.ds.DeleteBundle(ctx, &datastore.DeleteBundleRequest{
		TrustDomain: bundle.TrustDomain,
	})
	s.Require().NoError(err)
	s.Equal(bundle2, dresp.Bundle)

	lresp, err = s.ds.ListBundles(ctx, &datastore.ListBundlesRequest{})
	s.Require().NoError(err)
	s.Equal(1, len(lresp.Bundles))
}

func (s *PluginSuite) TestCreateAttestedNode() {
	node := &datastore.AttestedNode{
		SpiffeId:            "foo",
		AttestationDataType: "aws-tag",
		CertSerialNumber:    "badcafe",
		CertNotAfter:        time.Now().Add(time.Hour).Unix(),
	}

	cresp, err := s.ds.CreateAttestedNode(ctx, &datastore.CreateAttestedNodeRequest{Node: node})
	s.Require().NoError(err)
	s.Equal(node, cresp.Node)

	fresp, err := s.ds.FetchAttestedNode(ctx, &datastore.FetchAttestedNodeRequest{SpiffeId: node.SpiffeId})
	s.Require().NoError(err)
	s.Equal(node, fresp.Node)

	sresp, err := s.ds.ListAttestedNodes(ctx, &datastore.ListAttestedNodesRequest{
		ByExpiresBefore: &wrappers.Int64Value{
			Value: time.Now().Unix(),
		},
	})
	s.Require().NoError(err)
	s.Empty(sresp.Nodes)
}

func (s *PluginSuite) TestFetchAttestedNodeMissing() {
	fresp, err := s.ds.FetchAttestedNode(ctx, &datastore.FetchAttestedNodeRequest{SpiffeId: "missing"})
	s.Require().NoError(err)
	s.Require().Nil(fresp.Node)
}

func (s *PluginSuite) TestFetchStaleNodes() {
	efuture := &datastore.AttestedNode{
		SpiffeId:            "foo",
		AttestationDataType: "aws-tag",
		CertSerialNumber:    "badcafe",
		CertNotAfter:        time.Now().Add(time.Hour).Unix(),
	}

	epast := &datastore.AttestedNode{
		SpiffeId:            "bar",
		AttestationDataType: "aws-tag",
		CertSerialNumber:    "deadbeef",
		CertNotAfter:        time.Now().Add(-time.Hour).Unix(),
	}

	_, err := s.ds.CreateAttestedNode(ctx, &datastore.CreateAttestedNodeRequest{Node: efuture})
	s.Require().NoError(err)

	_, err = s.ds.CreateAttestedNode(ctx, &datastore.CreateAttestedNodeRequest{Node: epast})
	s.Require().NoError(err)

	sresp, err := s.ds.ListAttestedNodes(ctx, &datastore.ListAttestedNodesRequest{
		ByExpiresBefore: &wrappers.Int64Value{
			Value: time.Now().Unix(),
		},
	})
	s.Require().NoError(err)
	s.Equal([]*datastore.AttestedNode{epast}, sresp.Nodes)
}

func (s *PluginSuite) TestUpdateAttestedNode() {
	node := &datastore.AttestedNode{
		SpiffeId:            "foo",
		AttestationDataType: "aws-tag",
		CertSerialNumber:    "badcafe",
		CertNotAfter:        time.Now().Add(time.Hour).Unix(),
	}

	userial := "deadbeef"
	uexpires := time.Now().Add(time.Hour * 2).Unix()

	_, err := s.ds.CreateAttestedNode(ctx, &datastore.CreateAttestedNodeRequest{Node: node})
	s.Require().NoError(err)

	uresp, err := s.ds.UpdateAttestedNode(ctx, &datastore.UpdateAttestedNodeRequest{
		SpiffeId:         node.SpiffeId,
		CertSerialNumber: userial,
		CertNotAfter:     uexpires,
	})
	s.Require().NoError(err)

	unode := uresp.Node
	s.Require().NotNil(unode)

	s.Equal(node.SpiffeId, unode.SpiffeId)
	s.Equal(node.AttestationDataType, unode.AttestationDataType)
	s.Equal(userial, unode.CertSerialNumber)
	s.Equal(uexpires, unode.CertNotAfter)

	fresp, err := s.ds.FetchAttestedNode(ctx, &datastore.FetchAttestedNodeRequest{SpiffeId: node.SpiffeId})
	s.Require().NoError(err)

	fnode := fresp.Node
	s.Require().NotNil(fnode)

	s.Equal(node.SpiffeId, fnode.SpiffeId)
	s.Equal(node.AttestationDataType, fnode.AttestationDataType)
	s.Equal(userial, fnode.CertSerialNumber)
	s.Equal(uexpires, fnode.CertNotAfter)
}

func (s *PluginSuite) TestDeleteAttestedNode() {
	entry := &datastore.AttestedNode{
		SpiffeId:            "foo",
		AttestationDataType: "aws-tag",
		CertSerialNumber:    "badcafe",
		CertNotAfter:        time.Now().Add(time.Hour).Unix(),
	}

	_, err := s.ds.CreateAttestedNode(ctx, &datastore.CreateAttestedNodeRequest{Node: entry})
	s.Require().NoError(err)

	dresp, err := s.ds.DeleteAttestedNode(ctx, &datastore.DeleteAttestedNodeRequest{SpiffeId: entry.SpiffeId})
	s.Require().NoError(err)
	s.Equal(entry, dresp.Node)

	fresp, err := s.ds.FetchAttestedNode(ctx, &datastore.FetchAttestedNodeRequest{SpiffeId: entry.SpiffeId})
	s.Require().NoError(err)
	s.Nil(fresp.Node)
}

func (s *PluginSuite) TestNodeSelectors() {
	foo1 := []*common.Selector{
		{Type: "FOO1", Value: "1"},
	}
	foo2 := []*common.Selector{
		{Type: "FOO2", Value: "1"},
	}
	bar := []*common.Selector{
		{Type: "BAR", Value: "FIGHT"},
	}

	// assert there are no selectors for foo
	selectors := s.getNodeSelectors("foo")
	s.Require().Empty(selectors)

	// set selectors on foo and bar
	s.setNodeSelectors("foo", foo1)
	s.setNodeSelectors("bar", bar)

	// get foo selectors
	selectors = s.getNodeSelectors("foo")
	s.Require().Equal(foo1, selectors)

	// replace foo selectors
	s.setNodeSelectors("foo", foo2)
	selectors = s.getNodeSelectors("foo")
	s.Require().Equal(foo2, selectors)

	// delete foo selectors
	s.setNodeSelectors("foo", nil)
	selectors = s.getNodeSelectors("foo")
	s.Require().Empty(selectors)

	// get bar selectors (make sure they weren't impacted by deleting foo)
	selectors = s.getNodeSelectors("bar")
	s.Require().Equal(bar, selectors)
}

func (s *PluginSuite) TestCreateRegistrationEntry() {
	var validRegistrationEntries []*common.RegistrationEntry
	s.getTestDataFromJsonFile(filepath.Join("testdata", "valid_registration_entries.json"), &validRegistrationEntries)

	for _, validRegistrationEntry := range validRegistrationEntries {
		resp, err := s.ds.CreateRegistrationEntry(ctx, &datastore.CreateRegistrationEntryRequest{Entry: validRegistrationEntry})
		s.Require().NoError(err)
		s.NotNil(resp)
		s.Require().NotNil(resp.Entry)
		s.NotEmpty(resp.Entry.EntryId)
		resp.Entry.EntryId = ""
		s.Require().Equal(resp.Entry, validRegistrationEntry)
	}
}

func (s *PluginSuite) TestCreateInvalidRegistrationEntry() {
	var invalidRegistrationEntries []*common.RegistrationEntry
	s.getTestDataFromJsonFile(filepath.Join("testdata", "invalid_registration_entries.json"), &invalidRegistrationEntries)

	for _, invalidRegistrationEntry := range invalidRegistrationEntries {
		createRegistrationEntryResponse, err := s.ds.CreateRegistrationEntry(ctx, &datastore.CreateRegistrationEntryRequest{Entry: invalidRegistrationEntry})
		s.Require().Error(err)
		s.Require().Nil(createRegistrationEntryResponse)
	}

	// TODO: Check that no entries have been created
}

func (s *PluginSuite) TestFetchRegistrationEntry() {
	registeredEntry := &common.RegistrationEntry{
		Selectors: []*common.Selector{
			{Type: "Type1", Value: "Value1"},
			{Type: "Type2", Value: "Value2"},
			{Type: "Type3", Value: "Value3"},
		},
		SpiffeId: "SpiffeId",
		ParentId: "ParentId",
		Ttl:      1,
	}

	createRegistrationEntryResponse, err := s.ds.CreateRegistrationEntry(ctx, &datastore.CreateRegistrationEntryRequest{Entry: registeredEntry})
	s.Require().NoError(err)
	s.Require().NotNil(createRegistrationEntryResponse)
	createdEntry := createRegistrationEntryResponse.Entry

	fetchRegistrationEntryResponse, err := s.ds.FetchRegistrationEntry(ctx, &datastore.FetchRegistrationEntryRequest{EntryId: createdEntry.EntryId})
	s.Require().NoError(err)
	s.Require().NotNil(fetchRegistrationEntryResponse)
	s.Equal(createdEntry, fetchRegistrationEntryResponse.Entry)
}

func (s *PluginSuite) TestFetchInexistentRegistrationEntry() {
	fetchRegistrationEntryResponse, err := s.ds.FetchRegistrationEntry(ctx, &datastore.FetchRegistrationEntryRequest{EntryId: "INEXISTENT"})
	s.Require().NoError(err)
	s.Require().Nil(fetchRegistrationEntryResponse.Entry)
}

func (s *PluginSuite) TestFetchRegistrationEntries() {
	entry1 := s.createRegistrationEntry(&common.RegistrationEntry{
		Selectors: []*common.Selector{
			{Type: "Type1", Value: "Value1"},
			{Type: "Type2", Value: "Value2"},
			{Type: "Type3", Value: "Value3"},
		},
		SpiffeId: "spiffe://example.org/foo",
		ParentId: "spiffe://example.org/bar",
		Ttl:      1,
	})

	entry2 := s.createRegistrationEntry(&common.RegistrationEntry{
		Selectors: []*common.Selector{
			{Type: "Type3", Value: "Value3"},
			{Type: "Type4", Value: "Value4"},
			{Type: "Type5", Value: "Value5"},
		},
		SpiffeId: "spiffe://example.org/baz",
		ParentId: "spiffe://example.org/bat",
		Ttl:      2,
	})

	resp, err := s.ds.ListRegistrationEntries(ctx, &datastore.ListRegistrationEntriesRequest{})
	s.Require().NoError(err)
	s.Require().NotNil(resp)

	expectedResponse := &datastore.ListRegistrationEntriesResponse{
		Entries: []*common.RegistrationEntry{entry2, entry1},
	}
	s.Equal(expectedResponse, resp)
}

func (s *PluginSuite) TestUpdateRegistrationEntry() {
	entry := s.createRegistrationEntry(&common.RegistrationEntry{
		Selectors: []*common.Selector{
			{Type: "Type1", Value: "Value1"},
			{Type: "Type2", Value: "Value2"},
			{Type: "Type3", Value: "Value3"},
		},
		SpiffeId: "spiffe://example.org/foo",
		ParentId: "spiffe://example.org/bar",
		Ttl:      1,
	})

	entry.Ttl = 2
	updateRegistrationEntryResponse, err := s.ds.UpdateRegistrationEntry(ctx, &datastore.UpdateRegistrationEntryRequest{
		Entry: entry,
	})
	s.Require().NoError(err)
	s.Require().NotNil(updateRegistrationEntryResponse)

	fetchRegistrationEntryResponse, err := s.ds.FetchRegistrationEntry(ctx, &datastore.FetchRegistrationEntryRequest{EntryId: entry.EntryId})
	s.Require().NoError(err)
	s.Require().NotNil(fetchRegistrationEntryResponse)

	expectedResponse := &datastore.FetchRegistrationEntryResponse{Entry: entry}
	s.Equal(expectedResponse, fetchRegistrationEntryResponse)
}

func (s *PluginSuite) TestDeleteRegistrationEntry() {
	entry1 := s.createRegistrationEntry(&common.RegistrationEntry{
		Selectors: []*common.Selector{
			{Type: "Type1", Value: "Value1"},
			{Type: "Type2", Value: "Value2"},
			{Type: "Type3", Value: "Value3"},
		},
		SpiffeId: "spiffe://example.org/foo",
		ParentId: "spiffe://example.org/bar",
		Ttl:      1,
	})

	s.createRegistrationEntry(&common.RegistrationEntry{
		Selectors: []*common.Selector{
			{Type: "Type3", Value: "Value3"},
			{Type: "Type4", Value: "Value4"},
			{Type: "Type5", Value: "Value5"},
		},
		SpiffeId: "spiffe://example.org/baz",
		ParentId: "spiffe://example.org/bat",
		Ttl:      2,
	})

	// Make sure we deleted the right one
	delRes, err := s.ds.DeleteRegistrationEntry(ctx, &datastore.DeleteRegistrationEntryRequest{EntryId: entry1.EntryId})
	s.Require().NoError(err)
	s.Require().Equal(entry1, delRes.Entry)
}

func (s *PluginSuite) TestListParentIDEntries() {
	allEntries := testutil.GetRegistrationEntries("entries.json")
	tests := []struct {
		name                string
		registrationEntries []*common.RegistrationEntry
		parentID            string
		expectedList        []*common.RegistrationEntry
	}{
		{

			name:                "test_parentID_found",
			registrationEntries: allEntries,
			parentID:            "spiffe://parent",
			expectedList:        allEntries[:2],
		},
		{
			name:                "test_parentID_notfound",
			registrationEntries: allEntries,
			parentID:            "spiffe://imnoparent",
			expectedList:        nil,
		},
	}
	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			ds := s.newPlugin()
			for _, entry := range test.registrationEntries {
				r, err := ds.CreateRegistrationEntry(ctx, &datastore.CreateRegistrationEntryRequest{Entry: entry})
				require.NoError(t, err)
				require.NotNil(t, r)
				require.NotNil(t, r.Entry)
				entry.EntryId = r.Entry.EntryId
			}
			result, err := ds.ListRegistrationEntries(ctx, &datastore.ListRegistrationEntriesRequest{
				ByParentId: &wrappers.StringValue{
					Value: test.parentID,
				},
			})
			s.Require().NoError(err)
			s.Equal(test.expectedList, result.Entries)
		})
	}
}

func (s *PluginSuite) TestListSelectorEntries() {
	allEntries := testutil.GetRegistrationEntries("entries.json")
	tests := []struct {
		name                string
		registrationEntries []*common.RegistrationEntry
		selectors           []*common.Selector
		expectedList        []*common.RegistrationEntry
	}{
		{
			name:                "entries_by_selector_found",
			registrationEntries: allEntries,
			selectors: []*common.Selector{
				{Type: "a", Value: "1"},
				{Type: "b", Value: "2"},
				{Type: "c", Value: "3"},
			},
			expectedList: []*common.RegistrationEntry{allEntries[0]},
		},
		{
			name:                "entries_by_selector_not_found",
			registrationEntries: allEntries,
			selectors: []*common.Selector{
				{Type: "e", Value: "0"},
			},
			expectedList: nil,
		},
	}
	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			ds := s.newPlugin()
			for _, entry := range test.registrationEntries {
				r, err := ds.CreateRegistrationEntry(ctx, &datastore.CreateRegistrationEntryRequest{Entry: entry})
				require.NoError(t, err)
				require.NotNil(t, r)
				require.NotNil(t, r.Entry)
				entry.EntryId = r.Entry.EntryId
			}
			result, err := ds.ListRegistrationEntries(ctx, &datastore.ListRegistrationEntriesRequest{
				BySelectors: &datastore.BySelectors{
					Selectors: test.selectors,
				},
			})
			require.NoError(t, err)
			require.Equal(t, test.expectedList, result.Entries)
		})
	}
}

func (s *PluginSuite) TestListMatchingEntries() {
	allEntries := testutil.GetRegistrationEntries("entries.json")
	tests := []struct {
		name                string
		registrationEntries []*common.RegistrationEntry
		selectors           []*common.Selector
		expectedList        []*common.RegistrationEntry
	}{
		{
			name:                "test1",
			registrationEntries: allEntries,
			selectors: []*common.Selector{
				{Type: "a", Value: "1"},
				{Type: "b", Value: "2"},
				{Type: "c", Value: "3"},
			},
			expectedList: []*common.RegistrationEntry{
				allEntries[0],
				allEntries[1],
				allEntries[2],
			},
		},
		{
			name:                "test2",
			registrationEntries: allEntries,
			selectors: []*common.Selector{
				{Type: "d", Value: "4"},
			},
			expectedList: nil,
		},
	}
	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			ds := s.newPlugin()
			for _, entry := range test.registrationEntries {
				r, err := ds.CreateRegistrationEntry(ctx, &datastore.CreateRegistrationEntryRequest{Entry: entry})
				s.Require().NoError(err)
				require.NoError(t, err)
				require.NotNil(t, r)
				require.NotNil(t, r.Entry)
				entry.EntryId = r.Entry.EntryId
			}
			result, err := ds.ListRegistrationEntries(ctx, &datastore.ListRegistrationEntriesRequest{
				BySelectors: &datastore.BySelectors{
					Selectors: test.selectors,
					Match:     datastore.BySelectors_MATCH_SUBSET,
				},
			})
			s.Require().NoError(err)
			s.Equal(test.expectedList, result.Entries)
		})
	}
}

func (s *PluginSuite) TestCreateJoinToken() {
	now := time.Now().Unix()
	req := &datastore.CreateJoinTokenRequest{
		JoinToken: &datastore.JoinToken{
			Token:  "foobar",
			Expiry: now,
		},
	}
	_, err := s.ds.CreateJoinToken(ctx, req)
	s.Require().NoError(err)

	// Make sure we can't re-register
	_, err = s.ds.CreateJoinToken(ctx, req)
	s.NotNil(err)
}

func (s *PluginSuite) TestCreateAndFetchJoinToken() {
	now := time.Now().Unix()
	joinToken := &datastore.JoinToken{
		Token:  "foobar",
		Expiry: now,
	}
	_, err := s.ds.CreateJoinToken(ctx, &datastore.CreateJoinTokenRequest{
		JoinToken: joinToken,
	})
	s.Require().NoError(err)

	res, err := s.ds.FetchJoinToken(ctx, &datastore.FetchJoinTokenRequest{
		Token: joinToken.Token,
	})
	s.Require().NoError(err)
	s.Equal("foobar", res.JoinToken.Token)
	s.Equal(now, res.JoinToken.Expiry)
}

func (s *PluginSuite) TestDeleteJoinToken() {
	now := time.Now().Unix()
	joinToken1 := &datastore.JoinToken{
		Token:  "foobar",
		Expiry: now,
	}
	_, err := s.ds.CreateJoinToken(ctx, &datastore.CreateJoinTokenRequest{
		JoinToken: joinToken1,
	})
	s.Require().NoError(err)

	joinToken2 := &datastore.JoinToken{
		Token:  "batbaz",
		Expiry: now,
	}
	_, err = s.ds.CreateJoinToken(ctx, &datastore.CreateJoinTokenRequest{
		JoinToken: joinToken2,
	})
	s.Require().NoError(err)

	_, err = s.ds.DeleteJoinToken(ctx, &datastore.DeleteJoinTokenRequest{
		Token: joinToken1.Token,
	})
	s.Require().NoError(err)

	// Should not be able to fetch after delete
	resp, err := s.ds.FetchJoinToken(ctx, &datastore.FetchJoinTokenRequest{
		Token: joinToken1.Token,
	})
	s.Require().NoError(err)
	s.Nil(resp.JoinToken)

	// Second token should still be present
	resp, err = s.ds.FetchJoinToken(ctx, &datastore.FetchJoinTokenRequest{
		Token: joinToken2.Token,
	})
	s.Require().NoError(err)
	s.Equal(joinToken2, resp.JoinToken)
}

func (s *PluginSuite) TestPruneJoinTokens() {
	now := time.Now().Unix()
	joinToken := &datastore.JoinToken{
		Token:  "foobar",
		Expiry: now,
	}
	_, err := s.ds.CreateJoinToken(ctx, &datastore.CreateJoinTokenRequest{
		JoinToken: joinToken,
	})
	s.Require().NoError(err)

	// Ensure we don't prune valid tokens, wind clock back 10s
	_, err = s.ds.PruneJoinTokens(ctx, &datastore.PruneJoinTokensRequest{
		ExpiresBefore: now - 10,
	})
	s.Require().NoError(err)
	resp, err := s.ds.FetchJoinToken(ctx, &datastore.FetchJoinTokenRequest{
		Token: joinToken.Token,
	})
	s.Require().NoError(err)
	s.Equal("foobar", resp.JoinToken.Token)

	// Ensure we prune old tokens
	joinToken.Expiry = (now + 10)
	_, err = s.ds.PruneJoinTokens(ctx, &datastore.PruneJoinTokensRequest{
		ExpiresBefore: now + 10,
	})
	s.Require().NoError(err)
	resp, err = s.ds.FetchJoinToken(ctx, &datastore.FetchJoinTokenRequest{
		Token: joinToken.Token,
	})
	s.Require().NoError(err)
	s.Nil(resp.JoinToken)
}

func (s *PluginSuite) TestGetPluginInfo() {
	resp, err := s.ds.GetPluginInfo(ctx, &spi.GetPluginInfoRequest{})
	s.Require().NoError(err)
	s.Require().NotNil(resp)
}

func (s *PluginSuite) TestMigration() {
	for i := 0; i < codeVersion; i++ {
		dbName := fmt.Sprintf("v%d.sqlite3", i)
		dbPath := filepath.Join(s.dir, "migration-"+dbName)
		// copy the database file from the test data
		s.Require().NoError(copyFile(filepath.Join("testdata", "migration", dbName), dbPath))

		// configure the datastore to use the new database
		_, err := s.ds.Configure(context.Background(), &spi.ConfigureRequest{
			Configuration: fmt.Sprintf(`
				database_type = "sqlite3"
				connection_string = "file://%s"
			`, dbPath),
		})
		s.Require().NoError(err)

		switch i {
		case 0:
			// the v0 database has two bundles. the spiffe://otherdomain.org
			// bundle has been soft-deleted. after migration, it should not
			// exist. if we try and create a bundle with the same id, it should
			// fail if the migration did not run, due to uniqueness
			// constraints.
			_, err := s.ds.CreateBundle(context.Background(), &datastore.CreateBundleRequest{
				Bundle: &datastore.Bundle{
					TrustDomain: "spiffe://otherdomain.org",
				},
			})
			s.Require().NoError(err)
		default:
			s.T().Fatalf("no migration test added for version %d", i)
		}
	}
}

func (s *PluginSuite) TestRace() {
	next := int64(0)
	exp := time.Now().Add(time.Hour).Unix()

	testutil.RaceTest(s.T(), func(t *testing.T) {
		node := &datastore.AttestedNode{
			SpiffeId:            fmt.Sprintf("foo%d", atomic.AddInt64(&next, 1)),
			AttestationDataType: "aws-tag",
			CertSerialNumber:    "badcafe",
			CertNotAfter:        exp,
		}

		_, err := s.ds.CreateAttestedNode(ctx, &datastore.CreateAttestedNodeRequest{Node: node})
		require.NoError(t, err)
		_, err = s.ds.FetchAttestedNode(ctx, &datastore.FetchAttestedNodeRequest{SpiffeId: node.SpiffeId})
		require.NoError(t, err)
	})
}

func (s *PluginSuite) getTestDataFromJsonFile(filePath string, jsonValue interface{}) {
	invalidRegistrationEntriesJson, err := ioutil.ReadFile(filePath)
	s.Require().NoError(err)

	err = json.Unmarshal(invalidRegistrationEntriesJson, &jsonValue)
	s.Require().NoError(err)
}

func (s *PluginSuite) createRegistrationEntry(entry *datastore.RegistrationEntry) *datastore.RegistrationEntry {
	resp, err := s.ds.CreateRegistrationEntry(ctx, &datastore.CreateRegistrationEntryRequest{
		Entry: entry,
	})
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Require().NotNil(resp.Entry)
	return resp.Entry
}

func (s *PluginSuite) getNodeSelectors(spiffeID string) []*common.Selector {
	resp, err := s.ds.GetNodeSelectors(ctx, &datastore.GetNodeSelectorsRequest{
		SpiffeId: spiffeID,
	})
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Require().NotNil(resp.Selectors)
	s.Require().Equal(spiffeID, resp.Selectors.SpiffeId)
	return resp.Selectors.Selectors
}

func (s *PluginSuite) setNodeSelectors(spiffeID string, selectors []*common.Selector) {
	resp, err := s.ds.SetNodeSelectors(ctx, &datastore.SetNodeSelectorsRequest{
		Selectors: &datastore.NodeSelectors{
			SpiffeId:  spiffeID,
			Selectors: selectors,
		},
	})
	s.Require().NoError(err)
	s.Require().Equal(&datastore.SetNodeSelectorsResponse{}, resp)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
