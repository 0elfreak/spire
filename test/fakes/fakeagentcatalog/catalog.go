package fakeagentcatalog

import (
	"github.com/spiffe/spire/pkg/agent/catalog"
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor"
)

type Catalog struct {
	catalog.Plugins
}

func New() *Catalog {
	return &Catalog{}
}

func (c *Catalog) SetKeyManager(keyManager keymanager.KeyManager) {
	c.KeyManager = keyManager
}

func (c *Catalog) SetNodeAttestor(nodeAttestor nodeattestor.NodeAttestor) {
	c.NodeAttestor = nodeAttestor
}

func (c *Catalog) SetWorkloadAttestors(workloadAttestors ...catalog.WorkloadAttestor) {
	c.WorkloadAttestors = workloadAttestors
}

func WorkloadAttestor(name string, workloadAttestor workloadattestor.WorkloadAttestor) catalog.WorkloadAttestor {
	return catalog.WorkloadAttestor{
		PluginInfo:       pluginInfo{name: name, typ: workloadattestor.Type},
		WorkloadAttestor: workloadAttestor,
	}
}

type pluginInfo struct {
	name string
	typ  string
}

func (pi pluginInfo) Name() string {
	return pi.name
}

func (pi pluginInfo) Type() string {
	return pi.typ
}

func (pi pluginInfo) BuiltIn() bool {
	return true
}
