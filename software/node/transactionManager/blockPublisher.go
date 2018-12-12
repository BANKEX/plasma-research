package transactionManager

import (
	"../blockchain"
	"log"
	"time"
)

const (
	blockInterval = 10 * time.Second
)

var Blockchain []*blockchain.Block

type BlockPublisher struct {
	transactionManager *TransactionManager
}

func NewBlockPublisher(m *TransactionManager) *BlockPublisher {
	result := BlockPublisher{
		transactionManager: m,
	}

	go func() {
		for {
			time.Sleep(blockInterval)
			result.AssembleBlock()
		}
	}()

	return &result
}

func (p *BlockPublisher) AssembleBlock() {
	// after any error in this function, blockchain data becomes corrupted and we should terminate execution
	block, err := p.transactionManager.AssembleBlock()
	if err != nil {
		log.Fatalf("Failed to assemble block: %s", err)
	}

	// upload to a durable storage (S3/IPFS) or write to a local file system
	Blockchain = append(Blockchain, block)
	//data, err := block.Serialize()
	//if err != nil {
	//	log.Fatalf("Failed to write block: %s", err)
	//}
	//err = ioutil.WriteFile(fmt.Sprintf("./blockchain/%d.bin", block.BlockNumber), data, 0666)
	//if err != nil {
	//	log.Fatalf("Failed to write block: %s", err)
	//}

	// todo
	//ethereum.PushHashBlock(block.BlockNumber, block.GetHash())
}
