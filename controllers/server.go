package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func HttpPost(PostData string,class string,method string) ([]byte){

	url_class:=class
	url_method:=method
	ip_chain:=chain_url+"/v1/"
	//获取配置里host属性的value
	url:=ip_chain+url_class+"/"+url_method
	fmt.Println(url)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(PostData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	//查看通信是否正常
	//fmt.Println("returnJson:", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}