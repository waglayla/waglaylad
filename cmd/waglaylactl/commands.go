package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/waglayla/waglaylad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.WaglayladMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.WaglayladMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.WaglayladMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.WaglayladMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.WaglayladMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.WaglayladMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.WaglayladMessage_BanRequest{}),
	reflect.TypeOf(protowire.WaglayladMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
