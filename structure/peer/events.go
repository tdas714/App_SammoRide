package peer

import "github.com/App-SammoRide/structure"

// FilteredBlock is a minimal set of information about a block
type FilteredBlock struct {
	ChannelId            string
	Number               uint64
	FilteredTransactions []*FilteredTransaction
}

func (m *FilteredBlock) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *FilteredBlock) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *FilteredBlock) GetFilteredTransactions() []*FilteredTransaction {
	if m != nil {
		return m.FilteredTransactions
	}
	return nil
}

// FilteredTransaction is a minimal set of information about a transaction
// within a block
type FilteredTransaction struct {
	Txid             string
	Type             structure.HeaderType
	TxValidationCode TxValidationCode
	// Types that are valid to be assigned to Data:
	//	*FilteredTransaction_TransactionActions
	Data isFilteredTransaction_Data
}

func (m *FilteredTransaction) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *FilteredTransaction) GetType() structure.HeaderType {
	if m != nil {
		return m.Type
	}
	return structure.HeaderType_MESSAGE
}

func (m *FilteredTransaction) GetTxValidationCode() TxValidationCode {
	if m != nil {
		return m.TxValidationCode
	}
	return TxValidationCode_VALID
}

type isFilteredTransaction_Data interface {
	isFilteredTransaction_Data()
}

type FilteredTransaction_TransactionActions struct {
	TransactionActions *FilteredTransactionActions
}

func (*FilteredTransaction_TransactionActions) isFilteredTransaction_Data() {}

func (m *FilteredTransaction) GetData() isFilteredTransaction_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FilteredTransaction) GetTransactionActions() *FilteredTransactionActions {
	if x, ok := m.GetData().(*FilteredTransaction_TransactionActions); ok {
		return x.TransactionActions
	}
	return nil
}

// FilteredTransactionActions is a wrapper for array of TransactionAction
// message from regular block
type FilteredTransactionActions struct {
	ChaincodeActions []*FilteredChaincodeAction
}

func (m *FilteredTransactionActions) GetChaincodeActions() []*FilteredChaincodeAction {
	if m != nil {
		return m.ChaincodeActions
	}
	return nil
}

// FilteredChaincodeAction is a minimal set of information about an action
// within a transaction
type FilteredChaincodeAction struct {
	ChaincodeEvent *ChaincodeEvent
}

func (m *FilteredChaincodeAction) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

// BlockAndPrivateData contains Block and a map from tx_seq_in_block to rwset.TxPvtReadWriteSet
type BlockAndPrivateData struct {
	Block *structure.Block
	// map from tx_seq_in_block to rwset.TxPvtReadWriteSet
	PrivateDataMap map[uint64]*rwset.TxPvtReadWriteSet
}

func (m *BlockAndPrivateData) GetBlock() *structure.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockAndPrivateData) GetPrivateDataMap() map[uint64]*rwset.TxPvtReadWriteSet {
	if m != nil {
		return m.PrivateDataMap
	}
	return nil
}

// DeliverResponse
type DeliverResponse struct {
	// Types that are valid to be assigned to Type:
	//	*DeliverResponse_Status
	//	*DeliverResponse_Block
	//	*DeliverResponse_FilteredBlock
	//	*DeliverResponse_BlockAndPrivateData
	Type isDeliverResponse_Type
}

type isDeliverResponse_Type interface {
	isDeliverResponse_Type()
}

type DeliverResponse_Status struct {
	Status structure.Status
}

type DeliverResponse_Block struct {
	Block *structure.Block
}

type DeliverResponse_FilteredBlock struct {
	FilteredBlock *FilteredBlock
}

type DeliverResponse_BlockAndPrivateData struct {
	BlockAndPrivateData *BlockAndPrivateData
}

func (*DeliverResponse_Status) isDeliverResponse_Type() {}

func (*DeliverResponse_Block) isDeliverResponse_Type() {}

func (*DeliverResponse_FilteredBlock) isDeliverResponse_Type() {}

func (*DeliverResponse_BlockAndPrivateData) isDeliverResponse_Type() {}

func (m *DeliverResponse) GetType() isDeliverResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DeliverResponse) GetStatus() structure.Status {
	if x, ok := m.GetType().(*DeliverResponse_Status); ok {
		return x.Status
	}
	return structure.Status_UNKNOWN
}

func (m *DeliverResponse) GetBlock() *structure.Block {
	if x, ok := m.GetType().(*DeliverResponse_Block); ok {
		return x.Block
	}
	return nil
}

func (m *DeliverResponse) GetFilteredBlock() *FilteredBlock {
	if x, ok := m.GetType().(*DeliverResponse_FilteredBlock); ok {
		return x.FilteredBlock
	}
	return nil
}

func (m *DeliverResponse) GetBlockAndPrivateData() *BlockAndPrivateData {
	if x, ok := m.GetType().(*DeliverResponse_BlockAndPrivateData); ok {
		return x.BlockAndPrivateData
	}
	return nil
}
