package zkc

import "encoding/json"

type request struct {
	JsonRpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

func newRequest(method string, params interface{}) (*request, error) {
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return &request{
		JsonRpc: "2.0",
		Id:      1,
		Method:  method,
		Params:  paramsBytes,
	}, nil
}

func newRequestNoParams(method string) *request {
	return &request{
		JsonRpc: "2.0",
		Id:      1,
		Method:  method,
	}
}

type response struct {
	JsonRpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
}

func getResult[T any](resp *response) (T, error) {
	var result T
	err := json.Unmarshal(resp.Result, &result)
	return result, err
}
