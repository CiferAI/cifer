package cifer

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "cifer/api/cifer/cifer"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "MintdataAll",
					Use:       "list-mintdata",
					Short:     "List all mintdata",
				},
				{
					RpcMethod:      "Mintdata",
					Use:            "show-mintdata [id]",
					Short:          "Shows a mintdata by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateMintdata",
					Use:            "create-mintdata [creatorAddress] [metadata] [amount]",
					Short:          "Create mintdata",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "creatorAddress"}, {ProtoField: "metadata"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "UpdateMintdata",
					Use:            "update-mintdata [id] [creatorAddress] [metadata] [amount]",
					Short:          "Update mintdata",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "creatorAddress"}, {ProtoField: "metadata"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "DeleteMintdata",
					Use:            "delete-mintdata [id]",
					Short:          "Delete mintdata",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
