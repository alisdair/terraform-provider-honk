package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"

	"github.com/alisdair/terraform-provider-honk/internal/server"
)

type resourceBlocks struct {
}

var _ server.Resource = (*resourceBlocks)(nil)
var _ server.ResourceUpdater = (*resourceBlocks)(nil)

func newResourceBlocks() (*resourceBlocks, error) {
	return &resourceBlocks{}, nil
}

var (
	_ server.Resource        = (*resourceBlocks)(nil)
	_ server.ResourceUpdater = (*resourceBlocks)(nil)
)

func (r *resourceBlocks) Schema(ctx context.Context) *tfprotov5.Schema {
	return &tfprotov5.Schema{
		Block: &tfprotov5.SchemaBlock{
			BlockTypes: []*tfprotov5.SchemaNestedBlock{
				{
					TypeName: "target",
					Nesting:  tfprotov5.SchemaNestedBlockNestingModeMap,
					Block: &tfprotov5.SchemaBlock{
						Attributes: []*tfprotov5.SchemaAttribute{
							{
								Name:            "honks",
								Required:        true,
								Description:     "How many times to honk.",
								DescriptionKind: tfprotov5.StringKindMarkdown,
								Type:            tftypes.Number,
							},
							{
								Name:            "loud",
								Optional:        true,
								Description:     "Whether or not to honk loudly.",
								DescriptionKind: tfprotov5.StringKindMarkdown,
								Type:            tftypes.Bool,
							},
						},
					},
				},
			},
			Attributes: []*tfprotov5.SchemaAttribute{},
		},
	}
}

func (r *resourceBlocks) Validate(ctx context.Context, config map[string]tftypes.Value) ([]*tfprotov5.Diagnostic, error) {
	t := config["target"]

	if !t.IsFullyKnown() {
		return nil, nil
	}

	var targetValues map[string]tftypes.Value
	err := t.As(&targetValues)
	if err != nil {
		return nil, err
	}

	if len(targetValues) == 0 {
		return []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "At least one target is required.",
			},
		}, nil
	}

	return nil, nil
}

func (r *resourceBlocks) PlanCreate(ctx context.Context, proposed map[string]tftypes.Value, config map[string]tftypes.Value) (map[string]tftypes.Value, []*tfprotov5.Diagnostic, error) {
	return r.plan(ctx, proposed)
}

func (r *resourceBlocks) PlanUpdate(ctx context.Context, proposed map[string]tftypes.Value, config map[string]tftypes.Value, prior map[string]tftypes.Value) (map[string]tftypes.Value, []*tfprotov5.Diagnostic, error) {
	return r.plan(ctx, proposed)
}

func (r *resourceBlocks) plan(ctx context.Context, proposed map[string]tftypes.Value) (map[string]tftypes.Value, []*tfprotov5.Diagnostic, error) {
	return map[string]tftypes.Value{
		"id":     tftypes.NewValue(tftypes.String, "static-id"),
		"target": proposed["target"],
	}, nil, nil
}

func (r *resourceBlocks) Read(ctx context.Context, current map[string]tftypes.Value) (map[string]tftypes.Value, []*tfprotov5.Diagnostic, error) {
	// roundtrip current state as the source of applied migrations
	return current, nil, nil
}

func (r *resourceBlocks) Create(ctx context.Context, planned map[string]tftypes.Value, config map[string]tftypes.Value, prior map[string]tftypes.Value) (map[string]tftypes.Value, []*tfprotov5.Diagnostic, error) {
	return planned, nil, nil
}

func (r *resourceBlocks) Update(ctx context.Context, planned map[string]tftypes.Value, config map[string]tftypes.Value, prior map[string]tftypes.Value) (map[string]tftypes.Value, []*tfprotov5.Diagnostic, error) {
	return planned, nil, nil
}

func (r *resourceBlocks) Destroy(ctx context.Context, prior map[string]tftypes.Value) ([]*tfprotov5.Diagnostic, error) {
	return nil, nil
}
