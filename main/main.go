package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"scumm/go_contract_util/router"
	"strings"
)

var (
	//初始化网络、私钥、地址等
	netUrl        = "https://goerli.infura.io/v3/49fb3b593a324bffa7bcd9653f73f7c3"
	privateKey, _ = crypto.HexToECDSA("c9dc95a37d3e5954a176ffb375840f6d6aaefc2dc038abb99b56a7bb74d3a9d6")
	fromAddress   = common.HexToAddress("0x03Ca6DEfffD0ed6d9540d770ee8EC33D0EC57563")
	swapAddress   = common.HexToAddress("0x1fb8547525Ce41Ec1CeCb2763b8E5DfF7A0B46c3")
)

type Test struct {
	Num  uint8
	Name string
}

func main() {
	//fmt.Println("abc")
	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "Hello World")
	//})
	//router.Run(":8000")

	//链接ETH网络
	conn, err := ethclient.Dial(netUrl)
	if err != nil {
		fmt.Print("Dial err", err)
		return
	}

	defer conn.Close()

	//new合约对象
	swap, err := router.NewRouter(swapAddress, conn)
	if err != nil {
		fmt.Println("NewContract error", err)
	}

	//调用合约函数
	owner, err := swap.Owner(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	})
	if err != nil {
		fmt.Println("Owner error", err)
	}
	fmt.Println("Owner=", owner)

	//调用合约函数
	executor, err := swap.Executor(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	})
	if err != nil {
		fmt.Println("Owner error", err)
	}
	fmt.Println("Executor=", executor)

	//fmt.Println("CallData=", common.Hex2BytesFixed(calldataStr, len(calldataStr)))

	desc := router.TransitStructsTransitSwapDescription{
		SwapType:        0,
		SrcToken:        common.HexToAddress("0x0000000000000000000000000000000000000000"),
		DstToken:        common.HexToAddress("0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d"),
		SrcReceiver:     common.HexToAddress("0x470F30D2E7a2618c3cc374b3382CCFC64F820b48"),
		DstReceiver:     common.HexToAddress("0x03Ca6DEfffD0ed6d9540d770ee8EC33D0EC57563"),
		Amount:          big.NewInt(1000000000000000),
		MinReturnAmount: big.NewInt(1),
		Channel:         "web",
		ToChainID:       big.NewInt(0),
		WrappedNative:   common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"),
	}
	fmt.Println("desc=", desc)

	calldataStr := "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000302bae587ab9e1667a2d2b0fd67730fefdd1ab2d00000000000000000000000003ca6defffd0ed6d9540d770ee8ec33d0ec5756300000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000038ac426d750000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d00000000000000000000000000000000000000000000000000000000000000010000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000e47ff36ab50000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000008000000000000000000000000003ca6defffd0ed6d9540d770ee8ec33d0ec57563000000000000000000000000000000000000000000000000000000006412c1d20000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b4fbf271143f4fbf7b91a5ded31805e42b2208d6000000000000000000000000302bae587ab9e1667a2d2b0fd67730fefdd1ab2d00000000000000000000000000000000000000000000000000000000"
	calldata, err := hex.DecodeString(calldataStr)
	if err != nil {
		fmt.Println("calldata=", calldata)
	}

	//通过Abi文件生成请求合约的calldata （abi.encodeWithSelector）
	routerAbi, err := abi.JSON(strings.NewReader(router.RouterMetaData.ABI))
	if err == nil {
		amountOutMin := new(big.Int)
		amountOutMin.SetString("5000000000000000", 10)
		data, err := routerAbi.Pack("swapTypeMode", uint8(3))
		if err == nil {
			fmt.Println("calldata2=", hex.EncodeToString(data))
		}
	}

	//

	//t := Test{
	//	Num:  9,
	//	Name: "name",
	//}

	//arg := abi.Arguments{}.Pack()

	//
	//abi.NewType("tuple", "Aggre", []abi.ArgumentMarshaling{
	//	{Name: ""},
	//})
	//
	//abi.Arguments{
	//	{Type: string, Name: ""},
	//}.Pack()

	call := router.TransitStructsCallbytesDescription{
		Flag:      0,
		SrcToken:  common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Calldatas: calldata,
	}
	fmt.Println("call=", call)
	//
	//auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))
	////设置GasPrice
	//gasPrice, _ := conn.SuggestGasPrice(context.Background())
	//auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))
	////设置Nonce
	//nonce, _ := conn.PendingNonceAt(context.Background(), fromAddress)
	//fmt.Println("nonce=", nonce)
	//auth.Nonce = big.NewInt(77) //big.NewInt(int64(nonce))
	//auth.GasLimit = uint64(300000)
	//auth.Value = big.NewInt(1000000000000000)
	//
	////调用合约函数 swap
	//trx, err := swap.Swap(auth, desc, call)
	//if err == nil {
	//	fmt.Println("Swap:", trx.Hash())
	//}

	//hex.Decode

	//hex.EncodeToString()

	//routerAbi :=
	//abi.Arguments{}

}

//func encodeAbi()  {
//
//}

func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
