package memory

import (
	"github.com/hashicorp/go-plugin"
	"github.com/spiffe/control-plane/plugins/upstream_ca"
	"github.com/spiffe/control-plane/plugins/upstream_ca/proto"
)

type MemoryPlugin struct{}

func (MemoryPlugin) SubmitCSR(csr []byte) (*proto.SubmitCSRResponse, error) {
	return &proto.SubmitCSRResponse{}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: upstreamca.Handshake,
		Plugins: map[string]plugin.Plugin{
			"upstreamca": upstreamca.UpstreamCaPlugin{UpstreamCaImpl: &MemoryPlugin{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
