package main

import "fmt"



// data structure for header
type block struct {
	data string
	previousHash string
	timestamp string
	nonce string
	hash string
	
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



type blockchain struct {
	blocks []block
}












func main() {

// aes hasher	

var hashe = block{
	data:         "hash",
	previousHash: " ",
	timestamp:    "0000",
	nonce:        "ijksamd",
	hash:         "",
}




fmt.Println(hashe.hash,hashe.data,hashe.previousHash,hashe.timestamp,hashe.nonce)






fmt.Println(hashing("hello"))






}


func hashing(key string) string {
	
	
	key = key + "hashed"
	
	return key



}


/*
TODO:



FUNCTIONS:
INITIALIZE BLOCK
NONCE GENERATOR 
TIME STAMP




INITIALIZE BLOCKCHAIN

WALLET / TRANSACTION LOGIC
TRANSACTION WALLET BUG ?????


HASHING / VALIDATION LOGIC




*/