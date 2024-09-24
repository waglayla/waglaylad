package rpchandlers

import (
	"github.com/waglayla/waglaylad/app/appmessage"
	"github.com/waglayla/waglaylad/app/rpc/rpccontext"
	"github.com/waglayla/waglaylad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
