package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

/*data structure for header
type Block struct {
	data         string
	previousHash string
	timestamp    string
	nonce        string
	hash         string
}

type transaction struct {
	//	sender wallet.address
	//	recipient wallet.address
	amount string
}

type wallet struct {
	address string
	balance string
}



*/

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}



func main() {

	for i := 0; i < 10; i++ {
		for _, block := range chain.blocks {

			data := RandStringRunes(10)
			chain.AddBlock(data)

			fmt.Printf("Previous Hash: %x\n", block.PrevHash)
			fmt.Printf("Data in Block: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)

		}
		i++
	}
}










func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}






var chain = InitBlockChain()

func InitBlockChain() *BlockChain {
	x := []*Block{Genesis()}
	return &BlockChain{x}
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new2 := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new2)

}
