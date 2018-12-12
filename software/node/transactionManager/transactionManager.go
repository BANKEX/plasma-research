package transactionManager

import (
	. "../alias"
	"../blockchain"
	"../utils"
	"bytes"
	"encoding/hex"
	"fmt"
	"sync"
)

type TransactionManager struct {
	utxoIndex        map[string]*blockchain.Input // map key is "BlockNumber:TransactionNumber:OutputNumber"
	transactionQueue []*blockchain.Transaction
	lastBlock        uint32
	lastHash         Uint256
	lastAccumulator  Uint2048
	mutex            sync.Mutex
}

func NewTransactionManager() *TransactionManager {
	result := TransactionManager{
		utxoIndex:        map[string]*blockchain.Input{},
		transactionQueue: make([]*blockchain.Transaction, 0),
		lastBlock:        0,
		lastHash:         utils.Keccak256([]byte{}), // todo define genesis hash
		lastAccumulator:  []byte{3},                 //todo define genesis accumulator
		mutex:            sync.Mutex{},
	}
	return &result
}

// ValidateInputs checks that all inputs correspond to correct unspent outputs
func (p *TransactionManager) ValidateInputs(t *blockchain.Transaction) error {
	for _, in := range t.Inputs {
		utxo := p.utxoIndex[in.GetKey()]
		if utxo == nil {
			return fmt.Errorf("no such UTXO: %s", in.GetKey())
		}
		// todo deep equal instead of explicit, or make this resistant to modification of utxo model like adding assetId
		if bytes.Compare(utxo.Owner, in.Owner) != 0 || utxo.Slice.Begin != in.Slice.Begin || utxo.Slice.End != in.Slice.End {
			return fmt.Errorf("incorrect input data for UTXO: %s", in.GetKey())
		}
		if in.BlockIndex > p.lastBlock {
			return fmt.Errorf("multiple operations on a slice within the same block are forbidden")
		}
	}

	return nil
}

func (p *TransactionManager) SubmitTransaction(t *blockchain.Transaction) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// check that transaction is fully valid
	err := t.Validate()
	if err != nil {
		return err
	}
	err = p.ValidateInputs(t)
	if err != nil {
		return err
	}

	// spend inputs, add outputs to utxo index, queue transaction for the next block
	for _, in := range t.Inputs {
		delete(p.utxoIndex, in.GetKey())
	}
	for i, out := range t.Outputs {
		in := blockchain.Input{
			Output: blockchain.Output{
				Owner: out.Owner,
				Slice: out.Slice,
			},
			BlockIndex:  p.lastBlock + 1,
			TxIndex:     uint32(len(p.transactionQueue)),
			OutputIndex: uint8(i),
		}
		p.utxoIndex[in.GetKey()] = &in
	}
	p.transactionQueue = append(p.transactionQueue, t)

	return nil
}

// todo remove this
func dereference(t []*blockchain.Transaction) []blockchain.Transaction {
	result := make([]blockchain.Transaction, 0, len(t))
	for _, v := range t {
		result = append(result, *v)
	}
	return result
}

func (p *TransactionManager) AssembleBlock() (*blockchain.Block, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.assembleBlockFromTransactions(dereference(p.transactionQueue))
}

func (p *TransactionManager) AssembleDepositBlock(output blockchain.Output) (*blockchain.Block, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	t := blockchain.Transaction{
		UnsignedTransaction: blockchain.UnsignedTransaction{
			Outputs: []blockchain.Output{
				output,
			},
		},
	}

	// insert utxo
	// todo validate input slice
	in := blockchain.Input{
		Output:      output,
		BlockIndex:  p.lastBlock + 1,
		TxIndex:     0,
		OutputIndex: 0,
	}
	p.utxoIndex[in.GetKey()] = &in

	return p.assembleBlockFromTransactions([]blockchain.Transaction{t})
}

func (p *TransactionManager) assembleBlockFromTransactions(t []blockchain.Transaction) (*blockchain.Block, error) {
	block, err := blockchain.NewBlock(p.lastBlock+1, p.lastHash, p.lastAccumulator, dereference(p.transactionQueue))
	if err != nil {
		return nil, err
	}
	p.lastBlock++
	p.lastHash = block.GetHash()
	p.lastAccumulator = block.RSAAccumulator
	p.transactionQueue = make([]*blockchain.Transaction, 0)
	return block, nil
}

func (p *TransactionManager) GetUtxosForAddress(address string) ([]*blockchain.Input, error) {
	addr, err := hex.DecodeString(address[2:])
	if err != nil {
		return nil, err
	}
	result := make([]*blockchain.Input, 0)
	for _, out := range p.utxoIndex {
		if bytes.Compare(out.Owner, addr) == 0 {
			result = append(result, out)
		}
	}
	return result, nil
}

func (p *TransactionManager) GetLastBlockNumber() uint32 {
	return p.lastBlock
}

func (p *TransactionManager) GetUtxo(block, tx, output uint32) *blockchain.Input {
	return p.utxoIndex[fmt.Sprintf("%d:%d:%d", block, tx, output)]
}

// todo add utxo on deposit event, avoid double deposits
// todo spend utxo on withdraw event
