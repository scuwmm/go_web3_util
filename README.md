# go_contract_util
    参考文档：
    https://mirror.xyz/daxiong.eth/AUa_o9LEH8N2YnIlPXcYNXneWU50hREzmNdP1d-eivk

### abigen abi生成go文件
    abigen --abi ./go_contract_util/router/TransitSwapRouterV4.abi --pkg router --out ./go_contract_util/router/TransitSwapRouterV4.go

#### go 通过Abi文件生成请求合约的calldata （abi.encodeWithSelector）
	routerAbi, err := abi.JSON(strings.NewReader(router.RouterMetaData.ABI))
	if err == nil {
		amountOutMin := new(big.Int)
		amountOutMin.SetString("5000000000000000", 10)
		data, err := routerAbi.Pack("swapTypeMode", uint8(3))
		if err == nil {
			fmt.Println("calldata2=", hex.EncodeToString(data))
		}
	}


#### go 合约参数encode （例如 solidity中 abi.encode("1", 1, "0x0000"...)）
    
    addressTy, _ := abi.NewType("address", "string", []abi.ArgumentMarshaling{})
	bytesTy, _ := abi.NewType("bytes", "string", nil)
	uint256Ty, _ := abi.NewType("uint256", "uint64", []abi.ArgumentMarshaling{})
	stringTy, _ := abi.NewType("string", "string", nil)

	args := abi.Arguments{
		{Type: addressTy},
		{Type: bytesTy},
		{Type: uint256Ty},
		{Type: stringTy},
	}
	_params2, _ := hex.DecodeString("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")
	packed, err := args.Pack(
		common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
		_params2,
		big.NewInt(1),
		"abc",
	)
	if err != nil {
		fmt.Println("abi.Pack error:", err)
	}
	fmt.Println("abi.Pack success:", hex.EncodeToString(packed))

#### go 结构体encode （abi.encode(结构体)）
    import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
    )
    ...
    
    type Transaction struct {
    To    common.Address `json:"to"`
    Value *big.Int       `json:"value"`
    Data  []byte         `json:"data"`
    }
    ...

    //golang数组的打包，在生成新类型的时候，中括号"[]"应该在类型名的后边
    structTy, _ := abi.NewType("tuple[]", "struct ty", []abi.ArgumentMarshaling{
    {Name: "params1", Type: "address"},
    {Name: "params2", Type: "uint256"},
    {Name: "params3", Type: "bytes"},
    })

	args := abi.Arguments{
		{Type: structTy},
	}

    _params3, _ := hex.DecodeString(params3)
    transactions := []Transaction{
        {
            common.HexToAddress(params1), 
            big.NewInt(params2), 
            params3
        }
    }

	packed, err := args.Pack(
		transactions,
	)

	argPacked := hexutil.Encode(packed)


