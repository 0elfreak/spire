package peertracker

import (
	"errors"
	"io/ioutil"
	"net"
	"os"
	"path"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/suite"
)

var errMockWatcherFailed = errors.New("create new watcher failed")

type failingMockTracker struct{}

func (failingMockTracker) Close() {}
func (failingMockTracker) NewWatcher(CallerInfo) (Watcher, error) {
	return nil, errMockWatcherFailed
}

func newFailingMockTracker() (PeerTracker, error) {
	return failingMockTracker{}, nil
}

func TestListenerTestSuite(t *testing.T) {
	suite.Run(t, new(ListenerTestSuite))
}

type ListenerTestSuite struct {
	suite.Suite

	tempDir  string
	ul       *Listener
	unixAddr *net.UnixAddr
}

func (p *ListenerTestSuite) SetupTest() {
	var err error

	p.tempDir, err = ioutil.TempDir("", "spire-listener-test")
	p.Require().NoError(err)

	p.unixAddr = &net.UnixAddr{
		Net:  "unix",
		Name: path.Join(p.tempDir, "test.sock"),
	}
}

func (p *ListenerTestSuite) TearDownTest() {
	// only close the listener if we haven't already
	if p.ul != nil {
		err := p.ul.Close()
		p.NoError(err)
		p.ul = nil
	}
	err := os.Remove(p.tempDir)
	p.NoError(err)
}

func (p *ListenerTestSuite) TestAcceptDoesntFailWhenTrackerFails() {
	var err error
	logger, hook := test.NewNullLogger()
	logger.Level = logrus.WarnLevel
	lf := ListenerFactory{
		NewTracker: newFailingMockTracker,
		Log:        logger,
	}
	p.ul, err = lf.ListenUnix(p.unixAddr.Network(), p.unixAddr)
	p.Require().NoError(err)

	// Connect to ourselves
	clientDone := make(chan struct{})
	serverDone := make(chan struct{})
	peer := newFakeUDSPeer(p.T())

	peer.connect(p.unixAddr, clientDone)

	go func() {
		conn, err := p.ul.Accept()
		p.Require().Error(err)
		p.Require().Nil(conn)
		close(serverDone)
	}()

	// wait for client to connect
	<-clientDone

	// wait for server to shut down
	p.Require().NoError(p.ul.Close())
	p.ul = nil
	<-serverDone

	logEntry := hook.LastEntry()
	p.Require().NotNil(logEntry)
	p.Require().Equal("Connection failed during accept.", logEntry.Message)
	logErr := logEntry.Data["error"]
	p.Require().IsType(errors.New(""), logErr)
	p.Require().EqualError(logErr.(error), "create new watcher failed")
}

func (p *ListenerTestSuite) TestAcceptFailsWhenUnderlyingAcceptFails() {
	lf := ListenerFactory{
		NewUnixListener: newFailingMockListenUnix,
		NewTracker:      newFailingMockTracker,
	}
	ul, err := lf.ListenUnix(p.unixAddr.Network(), p.unixAddr)
	p.Require().NoError(err)

	_, err = ul.Accept()
	p.Require().Error(err)
}

// returns an empty unix listener that will fail any call to Accept()
func newFailingMockListenUnix(network string, laddr *net.UnixAddr) (*net.UnixListener, error) {
	return &net.UnixListener{}, nil
}
