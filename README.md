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
		aggregateDescriptionTt, _ = abi.NewType("tuple", "aggregateDesc", []abi.ArgumentMarshaling{
			{Name: "field_one", Type: "address"},
			{Name: "field_two", Type: "address"},
			{Name: "field_three", Type: "uint64[]"},
			{Name: "field_four", Type: "uint64[]"},
			{Name: "field_five", Type: "address[]"},
			{Name: "field_six", Type: "address[]"},
			{Name: "field_seven", Type: "bytes[]"},
		})
		args = abi.Arguments{
			{Type: aggregateDescriptionTt},
		}
	)

	record := struct {
		FieldOne   common.Address
		FieldTwo   common.Address
		FieldThree []uint64
		FieldFour  []uint64
		FieldFive  []common.Address
		FieldSix   []common.Address
		FieldSeven [][]byte
	}{
		common.HexToAddress("0x302BaE587Ab9E1667a2d2b0FD67730FEfDD1AB2d"),
		common.HexToAddress("0x03Ca6DEfffD0ed6d9540d770ee8EC33D0EC57563"),
		[]uint64{997000000000000}, //[]big.Int{*big.NewInt(997000000000000)},
		[]uint64{0},               //[]big.Int{*big.NewInt(0)},
		[]common.Address{common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")}, //swap path
		[]common.Address{common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")}, //swap path
		[][]byte{swapExactETHForTokensData},
	}

	packed, err := args.Pack(&record)
	if err != nil {
		fmt.Println("Pack1 error:", err)
	}

	fmt.Println("Pack1 data:", hex.EncodeToString(packed))


