package resources

import "github.com/quix/tforge/internal/core/state"

type RowKind string

const (
	RowModule   RowKind = "module"
	RowResource RowKind = "resource"
)

type Row struct {
	Kind       RowKind
	Address    string
	Parent     string
	TreePrefix string
	Expanded   bool
	Resource   *state.Resource
}

func DemoRows() []Row {
	return []Row{
		resource("", "", state.Resource{
			Address: "aws_db_instance.main",
			Action:  state.ActionReplace,
			Reason:  "cannot_update",
		}),
		resource("", "", state.Resource{
			Address:  "aws_s3_bucket.assets",
			Action:   state.ActionUpdate,
			Selected: true,
		}),
		resource("", "", state.Resource{
			Address:  "aws_s3_bucket.logs",
			Action:   state.ActionUpdate,
			Selected: true,
		}),
		resource("", "", state.Resource{
			Address: "aws_s3_bucket.uploads",
			Action:  state.ActionNoOp,
		}),
		resource("", "", state.Resource{
			Address: "data.aws_region.current",
			Action:  state.ActionNoOp,
		}),

		module("", "module.api", true),

		resource("module.api", "├──", state.Resource{
			Address: "module.api.aws_cloudwatch_log_group.api",
			Action:  state.ActionNoOp,
		}),
		resource("module.api", "├──", state.Resource{
			Address: "module.api.aws_cloudwatch_log_group.api_v2",
			Action:  state.ActionCreate,
		}),
		resource("module.api", "├──", state.Resource{
			Address: "module.api.aws_iam_policy.lambda_exec",
			Action:  state.ActionNoOp,
		}),
		resource("module.api", "├──", state.Resource{
			Address: "module.api.aws_iam_role.api_lambda",
			Action:  state.ActionDelete,
			Reason:  "delete_because_no_resource_config",
		}),
		resource("module.api", "├──", state.Resource{
			Address:  "module.api.aws_lambda_function.api",
			Action:   state.ActionUpdate,
			Selected: true,
		}),
		resource("module.api", "└──", state.Resource{
			Address: "module.api.aws_route53_record.api",
			Action:  state.ActionNoOp,
		}),

		module("", "module.networking", true),

		resource("module.networking", "├──", state.Resource{
			Address:  "module.networking.aws_security_group.web",
			Action:   state.ActionUpdate,
			Selected: true,
		}),
		resource("module.networking", "└──", state.Resource{
			Address: "module.networking.aws_vpc.main",
			Action:  state.ActionNoOp,
		}),
	}
}

func module(prefix, address string, expanded bool) Row {
	return Row{
		Kind:       RowModule,
		Address:    address,
		TreePrefix: prefix,
		Expanded:   expanded,
	}
}

func resource(parent, prefix string, r state.Resource) Row {
	return Row{
		Kind:       RowResource,
		Address:    r.Address,
		Parent:     parent,
		TreePrefix: prefix,
		Resource:   &r,
	}
}
