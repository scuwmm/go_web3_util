package main

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"scumm/go_contract_util/floki"
	"scumm/go_contract_util/pancake"
	"scumm/go_contract_util/pancakefactory"
	"scumm/go_contract_util/router"
	"strings"
)

var (
	//网络
	netUrl = "https://goerli.infura.io/v3/49fb3b593a324bffa7bcd9653f73f7c3"
	//测试私钥
	privateKey, _ = crypto.HexToECDSA("c9dc95a37d3e5954a176ffb375840f6d6aaefc2dc038abb99b56a7bb74d3a9d6")
	//发起者
	fromAddress = common.HexToAddress("0x03Ca6DEfffD0ed6d9540d770ee8EC33D0EC57563")
	//transit router
	routerAddress = common.HexToAddress("0x1fb8547525Ce41Ec1CeCb2763b8E5DfF7A0B46c3")
	//pancake router合约地址
	pancakeAddress = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
)

func swap(
	token0 string,
	token1 string,
	//decimal0 int,
	//decimal1 int,
	//impact int,
	//part int,
	amountIn int64,
	amountOutMin int64,
	to string,
	//chain string,
	//issuer string,
	channel string,
	dealline int64) {

	//链接ETH网络
	conn, err := GetEthConn()
	if err != nil {
		fmt.Print("Dial err", err)
		return
	}

	defer conn.Close()

	//pancake, err := pancake.NewPancake(pancakeAddress, conn)
	//if err != nil {
	//	fmt.Print("NewPancake err", err)
	//	return
	//}
	//

	swapAbi, _ := abi.JSON(strings.NewReader(router.RouterMetaData.ABI))
	fmt.Println(swapAbi)
	pancakeAbi, _ := abi.JSON(strings.NewReader(pancake.PancakeMetaData.ABI))
	fmt.Println(pancakeAbi)

	desc := router.TransitStructsTransitSwapDescription{
		SwapType:        0, //aggregatePreMode, aggregatePostMode, swap, cross， 默认0
		SrcToken:        common.HexToAddress(token0),
		DstToken:        common.HexToAddress(token1),
		SrcReceiver:     common.HexToAddress("0x470F30D2E7a2618c3cc374b3382CCFC64F820b48"), //TransitSwap的地址
		DstReceiver:     common.HexToAddress(to),                                           //token1的接收地址
		Amount:          big.NewInt(amountIn),                                              //token0的数量
		MinReturnAmount: big.NewInt(amountOutMin),                                          //最小返回token1的数量 （滑点）
		Channel:         channel,                                                           // android、ios、web
		ToChainID:       big.NewInt(0),                                                     //
		WrappedNative:   common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"), //WBNB的合约地址
	}
	//fmt.Println(desc)

	swapExactETHForTokensData, err := GenSwapExactETHForTokensData(
		pancakeAbi,
		big.NewInt(amountOutMin),
		[]common.Address{common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"), common.HexToAddress("0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d")},
		common.HexToAddress(to),
		big.NewInt(dealline),
	)
	if err != nil {
		fmt.Println("swapExactETHForTokens err:", err)
	}
	fmt.Println("swapExactETHForTokensData:", hex.EncodeToString(swapExactETHForTokensData))

	var (
		aggregateDescriptionTy, _ = abi.NewType("tuple", "struct TransitStructs.AggregateDescription", []abi.ArgumentMarshaling{
			{Name: "DstToken", Type: "address", InternalType: "address"},
			{Name: "Receiver", Type: "address", InternalType: "address"},
			{Name: "Amounts", Type: "uint256[]", InternalType: "uint256[]"},
			{Name: "NeedTransfer", Type: "uint256[]", InternalType: "uint256[]"},
			{Name: "Callers", Type: "address[]", InternalType: "address"},
			{Name: "ApproveProxy", Type: "address[]", InternalType: "address"},
			{Name: "Calls", Type: "bytes[]", InternalType: "bytes[]"},
		})
		args = abi.Arguments{
			{Type: aggregateDescriptionTy},
		}
	)

	amount := big.NewInt(int64(997000000000000))
	record := struct {
		DstToken     common.Address
		Receiver     common.Address
		Amounts      []*big.Int
		NeedTransfer []*big.Int
		Callers      []common.Address
		ApproveProxy []common.Address
		Calls        [][]byte
	}{
		DstToken:     common.HexToAddress("0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d"),
		Receiver:     common.HexToAddress("0x03Ca6DEfffD0ed6d9540d770ee8EC33D0EC57563"),
		Amounts:      []*big.Int{amount},                                                                  //[]big.Int{*big.NewInt(997000000000000)},
		NeedTransfer: []*big.Int{big.NewInt(0)},                                                           //[]big.Int{*big.NewInt(0)},
		Callers:      []common.Address{common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")}, //swap path
		ApproveProxy: []common.Address{common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")}, //swap path
		Calls:        [][]byte{swapExactETHForTokensData},
	}

	packed, err := args.Pack(&record)
	if err != nil {
		fmt.Println("Pack1 error:", err)
	}

	fmt.Println("Pack1 data:", hex.EncodeToString(packed))

	call := router.TransitStructsCallbytesDescription{
		Flag:      0,
		SrcToken:  common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Calldatas: packed,
	}
	//fmt.Println(call)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))
	//设置GasPrice
	gasPrice, _ := conn.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))
	//设置Nonce
	nonce, _ := conn.PendingNonceAt(context.Background(), fromAddress)
	fmt.Println("nonce=", nonce)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000)
	auth.Value = big.NewInt(1000000000000000)

	swap, err := router.NewRouter(routerAddress, conn)
	if err != nil {
		fmt.Println("NewContract error", err)
	}

	fmt.Println("address2string:", common.HexToAddress("0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d").String())

	//swap交易
	trx, err := swap.Swap(auth, desc, call)
	if err == nil {
		fmt.Println("Swap hash:", trx.Hash())
	}

	//生成wap的encode data
	{
		swapData, err := swapAbi.Pack("swap", desc, call)
		if err != nil {
			fmt.Println("swapData error:", err)
		}
		fmt.Println("swapData:", hex.EncodeToString(swapData))

		//swap.
		fmt.Print("swap END")
	}

}

func GenSwapExactETHForTokensData(pancakeAbi abi.ABI, amountOutMin *big.Int, path []common.Address, toAddress common.Address, deadline *big.Int) ([]byte, error) {
	return pancakeAbi.Pack(
		"swapExactETHForTokens",
		amountOutMin,
		path,
		toAddress,
		deadline,
	)
}

func GetEthConn() (*ethclient.Client, error) {
	conn, err := ethclient.Dial(netUrl)
	if err != nil {
		fmt.Print("Dial err", err)
		return nil, errors.New("GetEthConn failed!")
	}
	return conn, nil
}

func main() {

	conn, err := GetEthConn()
	if err != nil {
		fmt.Print("Dial err", err)
		return
	}

	defer conn.Close()

	//Approve(conn, "0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d", "0xD4C95beDC1ef456B75a8BFeB255C3e7468F2840d")
	//BalanceOf(conn, "0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d", "0xAac4310416dcc69643876c0143cf897DE9db7073")
	//BalanceOf(conn, "0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d", "0x03Ca6DEfffD0ed6d9540d770ee8EC33D0EC57563")
	//Allowance(conn, "0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d", "0xAac4310416dcc69643876c0143cf897DE9db7073", "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	GetPair(conn, "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f", "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", "0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d")

	//swap(
	//	"0x0000000000000000000000000000000000000000",
	//	"0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d",
	//	1000000000000000,
	//	2,
	//	"0x03Ca6DEfffD0ed6d9540d770ee8EC33D0EC57563",
	//	"web",
	//	1678950866)

}

func GetPair(conn *ethclient.Client, factory string, token0 string, token1 string) {
	//fmt.Println(time.Now().UnixMilli())
	pancakeFactory, err := pancakefactory.NewPancakefactory(common.HexToAddress(factory), conn)
	if err != nil {
		return
	}
	pair, err := pancakeFactory.GetPair(&bind.CallOpts{}, common.HexToAddress(token0), common.HexToAddress(token1))
	//fmt.Println(time.Now().UnixMilli())
	if err != nil {
		fmt.Println("GetPair failed")
	}
	fmt.Print("pair: ", pair.String())
}

func BalanceOf(conn *ethclient.Client, contract string, addr string) {
	token, err := floki.NewFloki(common.HexToAddress(contract), conn)
	if err != nil {
		return
	}

	balance, err := token.BalanceOf(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, common.HexToAddress(addr))

	if err != nil {
		log.Error("BalanceOf failed.")
	}

	fmt.Println("balance= ", balance)
}

func Allowance(conn *ethclient.Client, contract string, owner string, spender string) {
	token, err := floki.NewFloki(common.HexToAddress(contract), conn)
	if err != nil {
		return
	}
	r, err := token.Allowance(&bind.CallOpts{}, common.HexToAddress(owner), common.HexToAddress(spender))
	if err != nil {
		fmt.Println("Allowance error")
	}
	fmt.Println("Allowance r=", r)
}

func Approve(conn *ethclient.Client, contract string, spender string) {
	token, err := floki.NewFloki(common.HexToAddress(contract), conn)
	if err != nil {
		return
	}

	auth, _ := BuildTransactOpts(conn, big.NewInt(0))

	r, err := token.Approve(auth, common.HexToAddress(spender), math.MaxBig256)
	if err != nil {
		log.Error("Approve failed.")
	}
	log.Info("Approve tx= ", r.Hash())
}

func BuildTransactOpts(conn *ethclient.Client, value *big.Int) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))
	if err != nil {
		return nil, errors.New("NewKeyedTransactorWithChainID err")
	}
	//设置GasPrice
	gasPrice, _ := conn.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))
	//设置Nonce
	nonce, _ := conn.PendingNonceAt(context.Background(), fromAddress)
	fmt.Println("nonce=", nonce)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000)
	auth.Value = value

	return auth, nil
}
