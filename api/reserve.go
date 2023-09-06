package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

var FeishuUrl string

func Reserve(describe string) (string, error) {
	if len(describe) <= 10 {
		return "这么短的备注不太好吧", nil
	}

	jsonData := `{"msg_type":"text","content":{"text":"` + describe + `"}}`
	reader := bytes.NewReader([]byte(jsonData))
	req, err := http.NewRequest("POST", FeishuUrl, reader)
	if err != nil {
		return "请求错误", err
	}
	req.Header.Set("Content-Type", "application/json")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "请求错误", err
	}
	req.Close = true

	resbody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "请求错误", err
	}
	log.Println(string(resbody))

	return "老板的要求已经收到啦", nil

}
