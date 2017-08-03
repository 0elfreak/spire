package memory

import (
	"github.com/hashicorp/go-plugin"
	common "github.com/spiffe/node-agent/plugins/common/proto"
	"github.com/spiffe/node-agent/plugins/key_manager"
	"github.com/spiffe/node-agent/plugins/key_manager/proto"
)

type MemoryPlugin struct{}

func (MemoryPlugin) GenerateKeyPair(*proto.GenerateKeyPairRequest) (*proto.GenerateKeyPairResponse, error) {
	return &proto.GenerateKeyPairResponse{}, nil
}

func (MemoryPlugin) FetchPrivateKey(*proto.FetchPrivateKeyRequest) (*proto.FetchPrivateKeyResponse, error) {
	return &proto.FetchPrivateKeyResponse{}, nil
}

func (MemoryPlugin) Configure(*common.ConfigureRequest) (*common.ConfigureResponse, error) {
	return &common.ConfigureResponse{}, nil
}

func (MemoryPlugin) GetPluginInfo(*common.GetPluginInfoRequest) (*common.GetPluginInfoResponse, error) {
	return &common.GetPluginInfoResponse{}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: keymanager.Handshake,
		Plugins: map[string]plugin.Plugin{
			"keymanager": keymanager.KeyManagerPlugin{KeyManagerImpl: &MemoryPlugin{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
