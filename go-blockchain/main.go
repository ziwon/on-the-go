package main

import (
)

func main() {
//	bc := NewBlockchain()

//	bc.AddBlock("Received 100 BTC from Jeff Bezos")
//	bc.AddBlock("Received 200 BTC from Elon Musk")
//	bc.AddBlock("Received 200 BTC from Mark Zuckerberg")
//	bc.AddBlock("Received 500 ETH from Vitalik Buterin")
//	bc.AddBlock("Received 100 BTC from Satoshi Nakamoto")

//	for _, block := range bc.blocks {
//		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
//		fmt.Printf("Data: %s\n", block.Data)
//		fmt.Printf("Hash: %x\n", block.Hash)
//		pow := NewProofOfWork(block)
//		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
//		fmt.Println()
//	}

	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
