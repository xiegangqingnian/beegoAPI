package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type getBlockJson struct {
	BlockNumOrID string `json:"block_num_or_id"`
}

type blockInfo struct {
	Timestamp         string        `json:"timestamp"`
	Producer          string        `json:"producer"`
	Confirmed         int           `json:"confirmed"`
	Previous          string        `json:"previous"`
	TransactionMroot  string        `json:"transaction_mroot"`
	ActionMroot       string        `json:"action_mroot"`
	ScheduleVersion   int           `json:"schedule_version"`
	NewProducers      interface{}   `json:"new_producers"`
	HeaderExtensions  []interface{} `json:"header_extensions"`
	ProducerSignature string        `json:"producer_signature"`
	Transactions      []interface{} `json:"transactions"`
	BlockExtensions   []interface{} `json:"block_extensions"`
	ID                string        `json:"id"`
	BlockNum          int           `json:"block_num"`
	RefBlockPrefix    int           `json:"ref_block_prefix"`
	Status            bool          `json:"stats"`
}

type resBlockInfo struct {
	HeadBlockNum   int    `json:"current_block_height"`
	HeadBlockID    string `json:"current_block_hash"`
	HeadBlockTime  string `json:"current_block_time"`
	RefBlockPrefix int    `json:"current_block_prefix"`
	Producer       string `json:"current_block_producer"`
	Transactions   []interface{} `json:"transactions"`
}

func getBlockInfo(blockNum string) interface{} {

	var res blockInfo
	var resBlockInfo resBlockInfo
	block_num_or_ld := getBlockJson{
		blockNum,
	}
	params, _ := json.Marshal(block_num_or_ld)
	body := HttpPost(string(params), "chain", "get_block")
	json.Unmarshal(body, &res)
	if res.ID == "" {
		err := &JSONStruct{201, "No this block_number_or_hash"}
		return err
	} else {
		resBlockInfo.HeadBlockNum = res.BlockNum
		resBlockInfo.HeadBlockID = res.ID
		resBlockInfo.HeadBlockTime = res.Timestamp
		resBlockInfo.RefBlockPrefix = res.RefBlockPrefix
		resBlockInfo.Producer = res.Producer
		return resBlockInfo
	}
}

func getBlocksInfo(params blocksParams) interface{} {

	var res blockInfo
	var resblockInfo []resBlockInfo
	number := params.End - params.Start
	fmt.Println("************start and  end*********")
	fmt.Println(params.Start)
	fmt.Println(params.End)
	for i:=0;i<=number;i++{
		block_num_or_ld := getBlockJson{
			strconv.Itoa(params.Start+i),
		}
		fmt.Println("params.Start+i")
		fmt.Println(params.Start+i)
		fmt.Println("strconv.Itoa(params.Start+i)")
		fmt.Println(strconv.Itoa(params.Start+i))
		par, _ := json.Marshal(block_num_or_ld)
		fmt.Println("**************")
		fmt.Println("*****par******")
		fmt.Println(string(par))
		body := HttpPost(string(par), "chain", "get_block")

		json.Unmarshal(body, &res)

		if res.ID == "" {
			err := &JSONStruct{201, "No this block_number_or_hash"}
			return err
		} else {
			tmp :=resBlockInfo{res.BlockNum,res.ID,res.Timestamp,
				res.RefBlockPrefix,res.Producer,res.Transactions}
			resblockInfo = append(resblockInfo,tmp)
		}
	}
	return resblockInfo
}