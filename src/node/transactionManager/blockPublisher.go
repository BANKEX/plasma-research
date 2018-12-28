package transactionManager

import (
	"github.com/BANKEX/plasma-research/src/node/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"

	"github.com/BANKEX/plasma-research/src/node/blockchain"
)

const (
	blockInterval = 10 * time.Second
)

var Blockchain []*blockchain.Block

type BlockPublisher struct {
	transactionManager *TransactionManager
	client             *ethclient.Client
}

func NewBlockPublisher(m *TransactionManager) (*BlockPublisher, error) {
	c, err := ethclient.Dial(config.GetOperator().GethHost)
	if err != nil {
		return nil, err
	}

	result := BlockPublisher{
		transactionManager: m,
		client:             c,
	}

	go func() {
		for {
			time.Sleep(blockInterval)
			result.AssembleBlock()
		}
	}()

	return &result, nil
}

func (p *BlockPublisher) AssembleBlock() {
	// after any error in this function, blockchain data becomes corrupted and we should terminate execution
	block, err := p.transactionManager.AssembleBlock()
	if err != nil {
		log.Fatalf("Failed to assemble block: %s", err)
	}

	err = p.PublishBlock(block)
	if err != nil {
		log.Fatalf("Failed to publish block: %s", err)
	}
}

func (p *BlockPublisher) PublishBlock(block *blockchain.Block) error {
	// upload to a durable storage (S3/IPFS) or write to a local file system
	Blockchain = append(Blockchain, block)
	// data, err := block.Serialize()
	// if err != nil {
	// 	log.Fatalf("Failed to write block: %s", err)
	// }
	// err = ioutil.WriteFile(fmt.Sprintf("./blockchain/%d.bin", block.BlockNumber), data, 0666)
	// if err != nil {
	// 	log.Fatalf("Failed to write block: %s", err)
	// }

	// todo
	// ethereum.PushHashBlock(block.BlockNumber, block.GetHash())
	return nil
}
