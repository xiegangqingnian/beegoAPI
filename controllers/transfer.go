package controllers

import "encoding/json"

type SendTransaction struct {
	Compression string      `json:"compression"`
	Transaction interface{} `json:"transaction"`
	Signatures  []string    `json:"signatures"`
}

type Transaction struct {
	RefBlockNum           int           `json:"ref_block_num"`
	RefBlockPrefix        int           `json:"ref_block_prefix"`
	Expiration            string        `json:"expiration"`
	Actions               []interface{} `json:"actions"`
	ContextFreeActions    []interface{} `json:"context_free_actions"`
	TransactionExtensions []interface{} `json:"transaction_extensions"`
}

type transfer struct {
	Code   string `json:"code"`
	Action string `json:"action"`
	Args   struct {
		From     string `json:"from"`
		To       string `json:"to"`
		Quantity string `json:"quantity"`
		Memo     string `json:"memo"`
	} `json:"args"`
}

type args_0 struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity string `json:"quantity"`
	Memo     string `json:"memo"`
}

type args_1 struct {
	From           string `json:"from"`
	To             string `json:"to"`
	Quantity       string `json:"quantity"`
	Memo           string `json:"memo"`
	RefBlockNum    int    `json:"ref_block_num"`
	RefBlockPrefix int    `json:"ref_block_prefix"`
	Expiration     string `json:"expiration"`
	Data           string `json:"data"`
	Signatures     string `json:"signatures"`
}

type binargs struct {
	Binargs string `json:"binargs"`
}

func trxjsontobin(data transfer) interface{} {

	params, _ := json.Marshal(data)
	body := HttpPost(string(params), "chain", "abi_json_to_bin")
	var res binargs
	json.Unmarshal(body, &res)

	if res.Binargs == "" {
		err := JSONStruct{500, "try it again"}
		return err
	} else {
		return res
	}
}
func _trxjsontobin(data transfer) string {

	params, _ := json.Marshal(data)
	body := HttpPost(string(params), "chain", "abi_json_to_bin")
	var res binargs
	json.Unmarshal(body, &res)
	return res.Binargs
}

func sendtrx(params args_1) {

}
