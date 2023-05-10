package main

import (
	"encoding/json"
	"fmt"
)

const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}
`

type (
	Response struct {
		// поля с тегами
		Header ResponseHeader `json:"header"`
		Data   ResponseData   `json:"data"`
	}

	ResponseHeader struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	ResponseData []ResponseDataItem

	ResponseDataItem struct {
		Type       string               `json:"type"`
		Id         int                  `json:"id"`
		Attributes ResponseDataItemAttr `json:"attributes"`
	}

	ResponseDataItemAttr struct {
		Email       string `json:"email"`
		Article_ids []int  `json:"article_isd"`
	}
)

func ReadResponse(rawResp string) (Response, error) {
	// код десериализации
	resp := Response{}
	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		return Response{}, fmt.Errorf("JSONL: unmarshal: %w", err)
	}
	return resp, nil
}

func main() {
	resp, err := ReadResponse(rawResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v/n", resp)
}
