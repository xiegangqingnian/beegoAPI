基于beego的API开发

//1.trxjsontobin
		data1 :=transfer{
			Code:"zltio.token",
			Action:"transfer",
			Args:params,
		}
		fmt.Println(data1)
		binData := _trxjsontobin(data1)

		fmt.Println("****************************************")
		// 2.get_info
		data2 :=getInfo()
		fmt.Println(data2)
		blockId := data2.HeadBlockID
		//fmt.Println(block_id)
		// 3.get_block
		var block_info resBlockInfo
		info  := getBlockInfo(blockId)
		par,_ := json.Marshal(info)
		json.Unmarshal(par,&block_info)
	//	fmt.Println(block_info.RefBlockPrefix)
		resData:=sendtrx(block_info.HeadBlockNum,block_info.HeadBlockTime,block_info.RefBlockPrefix,binData)