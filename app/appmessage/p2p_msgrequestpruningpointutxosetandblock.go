package appmessage

import (
	"github.com/waglayla/waglaylad/domain/consensus/model/externalapi"
)

// MsgRequestPruningPointUTXOSet represents a waglayla RequestPruningPointUTXOSet message
type MsgRequestPruningPointUTXOSet struct {
	baseMessage
	PruningPointHash *externalapi.DomainHash
}

// Command returns the protocol command string for the message
func (msg *MsgRequestPruningPointUTXOSet) Command() MessageCommand {
	return CmdRequestPruningPointUTXOSet
}

// NewMsgRequestPruningPointUTXOSet returns a new MsgRequestPruningPointUTXOSet
func NewMsgRequestPruningPointUTXOSet(pruningPointHash *externalapi.DomainHash) *MsgRequestPruningPointUTXOSet {
	return &MsgRequestPruningPointUTXOSet{
		PruningPointHash: pruningPointHash,
	}
}
