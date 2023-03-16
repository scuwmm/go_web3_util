package keccka256

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

// 1. keccka256 算法
// 2. 取后20位
// 3. 与ETH地址进行gf256算法  （ETH地址20 byte）
const lengthOfBytesSub = 16

func Encrypt(source string) []byte {
	byteData, _ := hex.DecodeString(source)

	result := crypto.Keccak256(byteData)
	startIdx := len(result) - lengthOfBytesSub
	//fmt.Println("keccka256 result: ", hex.EncodeToString(result))
	//fmt.Println("keccka256 resultX: ", result[startIdx:])
	//fmt.Println("keccka256 sizeX: ", len(result[startIdx:]))

	return result[startIdx:]
}
