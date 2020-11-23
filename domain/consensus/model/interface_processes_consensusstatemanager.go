package model

import "github.com/kaspanet/kaspad/domain/consensus/model/externalapi"

// ConsensusStateManager manages the node's consensus state
type ConsensusStateManager interface {
	AddBlockToVirtual(blockHash *externalapi.DomainHash) error
	PopulateTransactionWithUTXOEntries(transaction *externalapi.DomainTransaction) error
	SetPruningPointUTXOSet(serializedUTXOSet []byte) error
	RestorePastUTXOSetIterator(blockHash *externalapi.DomainHash) (ReadOnlyUTXOSetIterator, error)
	HeaderTipsPruningPoint() (*externalapi.DomainHash, error)
	CalculatePastUTXOAndAcceptanceData(blockHash *externalapi.DomainHash) (*UTXODiff, AcceptanceData, Multiset, error)
}

// TestConsensusStateManager  adds to the main ConsensusStateManager methods required by tests
type TestConsensusStateManager interface {
	ConsensusStateManager
	AddUTXOToMultiset(multiset Multiset, entry *externalapi.UTXOEntry,
		outpoint *externalapi.DomainOutpoint) error
	ResolveBlockStatus(blockHash *externalapi.DomainHash) (externalapi.BlockStatus, error)
	VirtualFinalityPoint() (*externalapi.DomainHash, error)
}