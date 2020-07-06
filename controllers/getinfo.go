package controllers

import "encoding/json"

type resInfo struct {
	HeadBlockNum             int    `json:"current_block_height"`
	HeadBlockID              string `json:"current_block_hash"`
	HeadBlockTime            string `json:"current_block_time"`
	HeadBlockProducer        string `json:"current_block_producer"`
}


func getInfo() (resInfo) {


	type chainInfo struct {
		ServerVersion            string `json:"server_version"`
		ChainID                  string `json:"chain_id"`
		HeadBlockNum             int    `json:"head_block_num"`
		LastIrreversibleBlockNum int    `json:"last_irreversible_block_num"`
		LastIrreversibleBlockID  string `json:"last_irreversible_block_id"`
		HeadBlockID              string `json:"head_block_id"`
		HeadBlockTime            string `json:"head_block_time"`
		HeadBlockProducer        string `json:"head_block_producer"`
	}

	var blockInfo chainInfo
	var resInfo resInfo

	body := HttpPost("", "chain", "get_info")
	json.Unmarshal(body, &blockInfo)

	resInfo.HeadBlockNum 	  = blockInfo.HeadBlockNum
	resInfo.HeadBlockID   	  = blockInfo.HeadBlockID
	resInfo.HeadBlockTime 	  = blockInfo.HeadBlockTime
	resInfo.HeadBlockProducer = blockInfo.HeadBlockProducer

	return resInfo
}