package client

import (
	"context"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/common/bundleutil"
	"github.com/spiffe/spire/pkg/common/util"
	"github.com/spiffe/spire/proto/spire/server/datastore"
)

type TrustDomainConfig struct {
	EndpointAddress  string
	EndpointSpiffeID string
}

type ManagerConfig struct {
	Log          logrus.FieldLogger
	DataStore    datastore.DataStore
	Clock        clock.Clock
	TrustDomains map[string]TrustDomainConfig

	// newBundleUpdater is a test hook to inject updater behavior
	newBundleUpdater func(BundleUpdaterConfig) BundleUpdater
}

type Manager struct {
	clock    clock.Clock
	updaters []BundleUpdater
}

func NewManager(config ManagerConfig) *Manager {
	if config.Clock == nil {
		config.Clock = clock.New()
	}
	if config.newBundleUpdater == nil {
		config.newBundleUpdater = NewBundleUpdater
	}

	var updaters []BundleUpdater
	for trustDomain, trustDomainConfig := range config.TrustDomains {
		updaters = append(updaters, config.newBundleUpdater(BundleUpdaterConfig{
			TrustDomainConfig: trustDomainConfig,
			TrustDomain:       trustDomain,
			Log:               config.Log.WithField("trust_domain", trustDomain),
			DataStore:         config.DataStore,
		}))
	}

	return &Manager{
		clock:    config.Clock,
		updaters: updaters,
	}
}

func (m *Manager) Run(ctx context.Context) error {
	var tasks []func(context.Context) error
	for _, updater := range m.updaters {
		tasks = append(tasks, func(ctx context.Context) error {
			return m.runUpdater(ctx, updater)
		})
	}

	return util.RunTasks(ctx, tasks...)
}

func (m *Manager) runUpdater(ctx context.Context, updater BundleUpdater) error {
	var refreshHint time.Duration
	for {
		if r, err := updater.UpdateBundle(ctx); err == nil {
			refreshHint = r
		}
		timer := m.newRefreshTimer(refreshHint)
		select {
		case <-timer.C:
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		}
	}
}

func (m *Manager) newRefreshTimer(refreshHint time.Duration) *clock.Timer {
	if refreshHint == 0 {
		refreshHint = bundleutil.DefaultRefreshHint
	}
	return m.clock.Timer(refreshHint)
}
