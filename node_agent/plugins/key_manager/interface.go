package keymanager

import (
	"net/rpc"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-plugin"
	common "github.com/spiffe/sri/common/plugins/common/proto"
	"github.com/spiffe/sri/node_agent/plugins/key_manager/proto"
)

//Handshake is a common handshake that is shared by the plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "KeyManager",
	MagicCookieValue: "KeyManager",
}


type KeyManager interface {
	GenerateKeyPair(*sri_proto.GenerateKeyPairRequest) (*sri_proto.GenerateKeyPairResponse, error)
	FetchPrivateKey(*sri_proto.FetchPrivateKeyRequest) (*sri_proto.FetchPrivateKeyResponse, error)
	Configure(*common.ConfigureRequest) (*common.ConfigureResponse, error)
	GetPluginInfo(*common.GetPluginInfoRequest) (*common.GetPluginInfoResponse, error)
}

type KeyManagerPlugin struct {
	KeyManagerImpl KeyManager
}

func (p KeyManagerPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return empty.Empty{}, nil
}

func (p KeyManagerPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return empty.Empty{}, nil
}

func (p KeyManagerPlugin) GRPCServer(s *grpc.Server) error {
	sri_proto.RegisterKeyManagerServer(s, &GRPCServer{KeyManagerImpl: p.KeyManagerImpl})
	return nil
}

func (p KeyManagerPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: sri_proto.NewKeyManagerClient(c)}, nil
}
