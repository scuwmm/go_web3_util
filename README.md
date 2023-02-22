# go_contract_util

### abigen abi生成go文件
    abigen --abi ./go_contract_util/router/TransitSwapRouterV4.abi --pkg router --out ./go_contract_util/router/TransitSwapRouterV4.go

#### 通过Abi文件生成请求合约的calldata （abi.encodeWithSelector）
	routerAbi, err := abi.JSON(strings.NewReader(router.RouterMetaData.ABI))
	if err == nil {
		amountOutMin := new(big.Int)
		amountOutMin.SetString("5000000000000000", 10)
		data, err := routerAbi.Pack("swapTypeMode", uint8(3))
		if err == nil {
			fmt.Println("calldata2=", hex.EncodeToString(data))
		}
	}


