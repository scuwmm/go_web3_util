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


#### GO生成 android SDK 
    1.安装android环境： android studio ; android SDK; android NDK   
    2.安装gomobile  go get xxxx
    3.gomobile init
    4.gomobile bind -target=android package_name
    生成 aar 和 jar包，android使用 aar包


