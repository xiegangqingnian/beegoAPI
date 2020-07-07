package controllers

import (
	"encoding/json"
)

type AddrResult struct {
	AccountName            string        `json:"account_name"`
	CoreLiquidBalance      string        `json:"core_liquid_balance"`
	Created                string        `json:"created"`
	HeadBlockNum           int           `json:"head_block_num"`
	HeadBlockTime          string        `json:"head_block_time"`
	LastCodeUpdate         string        `json:"last_code_update"`
	Permissions            []Permissions `json:"permissions"`
	Privileged             bool          `json:"privileged"`
	RefundRequest          interface{}   `json:"refund_request"`
	SelfDelegatedBandwidth interface{}   `json:"self_delegated_bandwidth"`
	VoterInfo              interface{}   `json:"voter_info"`
}
type Keys struct {
	Key    string `json:"key"`
	Weight int    `json:"weight"`
}
type RequiredAuth struct {
	Accounts  []interface{} `json:"accounts"`
	Keys      []Keys        `json:"keys"`
	Threshold int           `json:"threshold"`
	Waits     []interface{} `json:"waits"`
}
type Permissions struct {
	Parent       string       `json:"parent"`
	PermName     string       `json:"perm_name"`
	RequiredAuth RequiredAuth `json:"required_auth"`
}

type Addr struct {
	AccountName string `json:"account_name"`
}

type resGetAddr struct {
	AccountName     string `json:"account_name"`
	Current_Balance string `json:"current_balance"`
	Created         string `json:"created"`
}

func getAddr(addr string) interface{} {

	var addrResult AddrResult
	var resGetAddr resGetAddr

	name := Addr{addr}
	params, _ := json.Marshal(name)

	body := HttpPost(string(params), "chain", "get_account")
	json.Unmarshal(body, &addrResult)

	if addrResult.AccountName == "" {
		errRes := &JSONStruct{201, "No this address"}
		return errRes
	} else {
		resGetAddr.AccountName = addrResult.AccountName
		resGetAddr.Current_Balance = addrResult.CoreLiquidBalance
		resGetAddr.Created = addrResult.Created
		return resGetAddr
	}
}

func checkAddrExist(addr string) bool {
	var res interface{}
	res = getAddr(addr)

	switch res.(type) {
	case JSONStruct:
		return false
	case resGetAddr:
		return true
	default:
		return false
	}
}
