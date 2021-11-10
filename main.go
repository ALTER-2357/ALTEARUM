package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

/*data structure for header
type Block struct {
	data         string  `json:"Data"`
	previousHash string  `json:"PrevHash"`
	timestamp    string  `json:"timestamp"`
	nonce        string  `json:"nonce"`
	hash         string  `json:"hash"`
}

type Transaction struct {
	//	sender wallet.address   `json:"sender"`
	//	recipient wallet.address  `json:"recipent"`
	amount string `json:"amount"`
}

type Wallet struct {
	data    string  `json:"data"`
	address string  `json:"address"`	
	balance string  `json:"balance"`
}
*/



type BlockChain struct {
	blocks []*Block
}



type Block struct {
	Hash     []byte `json:"Hash"`
	Data     []byte `json:"Data"`
	PrevHash []byte `json:"PrevHash"`
}

//var blockchain []BlockChain = []BlockChain{}

func main() {

	for i := 0; i < 4; i++ {
	
		for _, block := range chain.blocks {

			data := jsonpaser()

			chain.AddBlock(data)

			fmt.Printf("Previous Hash: %x\n", block.PrevHash)
			fmt.Printf("Data in Block: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)

		}
		i++
	}
}





func jsonpaser() string {

	data, err := ioutil.ReadFile("userdata.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	var users Block

	json.Unmarshal(data, &users)

	error := ioutil.WriteFile("userdata.json", data, 0644)
	if error != nil {
		log.Fatal(error)
	}

	return string(data)

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
