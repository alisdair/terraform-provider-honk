package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"

	"github.com/alisdair/terraform-provider-honk/internal/server"
)

func New(version string) func() tfprotov5.ProviderServer {
	return func() tfprotov5.ProviderServer {
		s := server.MustNew(func() server.Provider {
			return &provider{}
		})

		// resources
		s.MustRegisterResource("honk_blocks", newResourceBlocks)

		return s
	}
}

type provider struct {
}

var _ server.Provider = (*provider)(nil)

func (p *provider) Schema(context.Context) *tfprotov5.Schema {
	return &tfprotov5.Schema{
		Block: &tfprotov5.SchemaBlock{},
	}
}

func (p *provider) Validate(ctx context.Context, config map[string]tftypes.Value) ([]*tfprotov5.Diagnostic, error) {
	return nil, nil
}

func (p *provider) Configure(ctx context.Context, config map[string]tftypes.Value) ([]*tfprotov5.Diagnostic, error) {
	return nil, nil
}
