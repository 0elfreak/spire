package keymanager

import (
	common "github.com/spiffe/sri/node_agent/plugins/common/proto"
	"github.com/spiffe/sri/node_agent/plugins/key_manager/proto"
	"golang.org/x/net/context"
)

type GRPCServer struct {
	KeyManagerImpl KeyManager
}

func (m *GRPCServer) GenerateKeyPair(ctx context.Context, req *node_agent_proto.GenerateKeyPairRequest) (*node_agent_proto.GenerateKeyPairResponse, error) {
	response, err := m.KeyManagerImpl.GenerateKeyPair(req)
	return response, err
}

func (m *GRPCServer) FetchPrivateKey(ctx context.Context, req *node_agent_proto.FetchPrivateKeyRequest) (*node_agent_proto.FetchPrivateKeyResponse, error) {
	response, err := m.KeyManagerImpl.FetchPrivateKey(req)
	return response, err
}

func (m *GRPCServer) Configure(ctx context.Context, req *common.ConfigureRequest) (*common.ConfigureResponse, error) {
	response, err := m.KeyManagerImpl.Configure(req)
	return response, err
}

func (m *GRPCServer) GetPluginInfo(ctx context.Context, req *common.GetPluginInfoRequest) (*common.GetPluginInfoResponse, error) {
	response, err := m.KeyManagerImpl.GetPluginInfo(req)
	return response, err
}

type GRPCClient struct {
	client node_agent_proto.KeyManagerClient
}

func (m *GRPCClient) GenerateKeyPair(req *node_agent_proto.GenerateKeyPairRequest) (*node_agent_proto.GenerateKeyPairResponse, error) {
	res, err := m.client.GenerateKeyPair(context.Background(), req)
	return res, err
}

func (m *GRPCClient) FetchPrivateKey(req *node_agent_proto.FetchPrivateKeyRequest) (*node_agent_proto.FetchPrivateKeyResponse, error) {
	res, err := m.client.FetchPrivateKey(context.Background(), req)
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
