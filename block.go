package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	headerPrefix = []byte("h") // headerPrefix + num (uint64 big endian) + hash -> header
	numSuffix    = []byte("n") // headerPrefix + num (uint64 big endian) + numSuffix -> hash
)

const (
	targetBlock = 3
)

func main() {

	// Connection to leveldb
	db, _ := leveldb.OpenFile("/Users/Aaron/Documents/Workspace/solidity/erc20/eth/geth/chaindata", nil)

	// 40 to bytes (Big endian)
	blockNumber := make([]byte, 8)
	binary.BigEndian.PutUint64(blockNumber, uint64(targetBlock))

	fmt.Printf("Details of Blocknumber:- \nHex: %x \nBytes: %d\n\n\n", blockNumber, blockNumber)

	// create key to get hash (headerPrefix + num (uint64 big endian) + numSuffix)
	hashKey := append(headerPrefix, blockNumber...) // adding prefix
	hashKey = append(hashKey, numSuffix...)         // adding suffix

	fmt.Printf("Details of leveldb key for Block Hash:- \nType: %T  \nHex: %x \nbytes: %v \nLength:  %d\n\n\n", hashKey, hashKey, hashKey, len(hashKey))

	// Getting hash using hashKey
	blockHash, _ := db.Get(hashKey, nil)
	fmt.Printf("Details of Block hash:- \nType: %T \nHex: %x \nBytes: %v\n\n\n", blockHash, blockHash, blockHash)

	//Create key to get header (headerPrefix + num (uint64 big endian) + hash)
	headerKey := append(headerPrefix, blockNumber...) // adding prefix
	headerKey = append(headerKey, blockHash...)       // adding suffix

	fmt.Printf("Details of leveldb key for Block Header:- \nType: %T  \nHex: %x \nVytes: %v \nLength:  %d\n\n\n", headerKey, headerKey, headerKey, len(headerKey))

	//get Block Header data from db
	blockHeaderData, _ := db.Get(headerKey, nil)

	fmt.Printf("Details of Raw Block Header:- \nType: %T  \nHex: %x \nBytes: %v \nLength:  %d\n\n\n", blockHeaderData, blockHeaderData, blockHeaderData, len(blockHeaderData))

	//new Blockheader type
	blockHeader := new(types.Header)
	fmt.Printf("Details of new Header Type:- \nType: %T  \nHex: %x \nValue: %v\n\n\n", blockHeader, blockHeader, blockHeader)

	// Read blockHeaderData in a tmp variable
	tmpByteData := bytes.NewReader(blockHeaderData)
	fmt.Printf("Details of tmpByteData:- \nType: %T  \nHex: %x \nValue: %v\n\n\n", tmpByteData, tmpByteData, tmpByteData)

	//Decode tmpByteData to new blockHeader
	rlp.Decode(tmpByteData, blockHeader)
	fmt.Printf("Details of Header for block number %d:- \nType: %T  \nHex: %x \nValue: %v\n\n\n", targetBlock, blockHeader, blockHeader, blockHeader)

}
