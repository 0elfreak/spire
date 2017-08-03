package nodeattestor

import (
	common "github.com/spiffe/node-agent/plugins/common/proto"
	"github.com/spiffe/node-agent/plugins/node_attestor/proto"
	"golang.org/x/net/context"
)

type GRPCServer struct {
	NodeAttestorImpl NodeAttestor
}

func (m *GRPCServer) FetchAttestationData(ctx context.Context, req *proto.FetchAttestationDataRequest) (*proto.FetchAttestationDataResponse, error) {
	response, err := m.NodeAttestorImpl.FetchAttestationData(req)
	return response, err
}

func (m *GRPCServer) Configure(ctx context.Context, req *common.ConfigureRequest) (*common.ConfigureResponse, error) {
	response, err := m.NodeAttestorImpl.Configure(req)
	return response, err
}

func (m *GRPCServer) GetPluginInfo(ctx context.Context, req *common.GetPluginInfoRequest) (*common.GetPluginInfoResponse, error) {
	response, err := m.NodeAttestorImpl.GetPluginInfo(req)
	return response, err
}

type GRPCClient struct {
	client proto.NodeAttestorClient
}

func (m *GRPCClient) FetchAttestationData(req *proto.FetchAttestationDataRequest) (*proto.FetchAttestationDataResponse, error) {
	res, err := m.client.FetchAttestationData(context.Background(), req)
	return res, err
}

func (m *GRPCClient) Configure(req *common.ConfigureRequest) (*common.ConfigureResponse, error) {
	res, err := m.client.Configure(context.Background(), req)
	return res, err
}

func (m *GRPCClient) GetPluginInfo(req *common.GetPluginInfoRequest) (*common.GetPluginInfoResponse, error) {
	res, err := m.client.GetPluginInfo(context.Background(), req)
	return res, err
}
