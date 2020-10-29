package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DiceRequestParams struct {
	ApiKey string `json:"apiKey"`
	Max    int32  `json:"max"`
	Min    int32  `json:"min"`
	N      int32  `json:"n"`
}

type DiceRequest struct {
	Id      int32             `json:"id"`
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  DiceRequestParams `json:"params"`
}

func newDiceRequest(amount int32, size int32) DiceRequest {
	return DiceRequest{
		Id:      1,
		Jsonrpc: "2.0",
		Method:  "generateIntegers",
		Params: DiceRequestParams{
			ApiKey: randomOrgApiKey,
			Max:    size,
			Min:    1,
			N:      amount,
		},
	}
}

type DiceResponse struct {
	Id      int32               `json:"id"`
	Jsonrpc string              `json:"jsonrpc"`
	Result  DiceResponseResults `json:"result"`
}

type DiceResponseResults struct {
	Random        DiceResponseRandom `json:"random"`
	BitsUsed      int32              `json:"bitsUsed"`
	BitsLeft      int32              `json:"bitsLeft"`
	RequestsLeft  int32              `json:"requestsLeft"`
	AdvisoryDelay int32              `json:"advisoryDelay"`
}

type DiceResponseRandom struct {
	Data           []int32 `json:"data"`
	CompletionTime string  `json:"completionTime"`
}

func rollDice(amount int32, size int32) (result int32, err error) {
	request_body, err := json.Marshal(newDiceRequest(amount, size))
	if err != nil {
		return
	}

	response, err := http.Post("https://api.random.org/json-rpc/2/invoke", "application/json", bytes.NewReader(request_body))
	if err != nil {
		return
	}

	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	response_data := DiceResponse{}
	err = json.Unmarshal(response_body, &response_data)
	if err != nil {
		return
	}

	for _, roll := range response_data.Result.Random.Data {
		result += roll
	}

	return
}
