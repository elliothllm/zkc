package zkc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Request represents a JSON-RPC request.
type Request struct {
	JsonRpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

func newRequest(method string, params interface{}) (*Request, error) {
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return &Request{
		JsonRpc: "2.0",
		Id:      1,
		Method:  method,
		Params:  paramsBytes,
	}, nil
}

func newRequestNoParams(method string) *Request {
	return &Request{
		JsonRpc: "2.0",
		Id:      1,
		Method:  method,
	}
}

// Response represents a JSON-RPC response.
type Response struct {
	JsonRpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *Error          `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func getResult[T any](resp *Response) (T, error) {
	var result T
	err := json.Unmarshal(resp.Result, &result)
	return result, err
}

// RpcBlockOrBatchNumber represents a block or batch number in the Erigon RPC API.
// It implements the ethereum RPC BlockNumber interface.
// It can be used to specify special block numbers like "latest", "earliest", etc.
//
// 0 is the earliest block number.
// -1 is the latest block number.
// -2 is the pending block number.
// -3 is the safe block number.
// -4 is the finalized block number.
// -5 is the latest executed block number.
//
// Any positive number is a block number.
type RpcBlockOrBatchNumber int64

const (
	LatestExecuted = RpcBlockOrBatchNumber(-5)
	Finalized      = RpcBlockOrBatchNumber(-4)
	Safe           = RpcBlockOrBatchNumber(-3)
	Pending        = RpcBlockOrBatchNumber(-2)
	Latest         = RpcBlockOrBatchNumber(-1)
	Earliest       = RpcBlockOrBatchNumber(0)
)

func (b *RpcBlockOrBatchNumber) Resolve() (uint64, string, error) {
	switch *b {
	case LatestExecuted:
		return 0, "latestExecuted", nil
	case Finalized:
		return 0, "finalized", nil
	case Safe:
		return 0, "safe", nil
	case Pending:
		return 0, "pending", nil
	case Latest:
		return 0, "latest", nil
	case Earliest:
		return 0, "earliest", nil
	default:
		if *b < 0 {
			return 0, "", fmt.Errorf("invalid block number: %d", *b)
		}
		return uint64(*b), "", nil
	}
}

func (b *RpcBlockOrBatchNumber) Params(p ...interface{}) []interface{} {
	num, str, err := b.Resolve()
	if err != nil {
		return nil
	}
	if num == 0 {
		return append([]interface{}{str}, p...)
	}
	return append([]interface{}{num}, p...)
}

// WitnessMode represents a string value of either "full" or "trimmed".
type WitnessMode string

const (
	WitnessModeFull    WitnessMode = "full"
	WitnessModeTrimmed WitnessMode = "trimmed"
)

// HexString represents a hexadecimal string.
// It provides a method to convert it to a uint64 value.
type HexString string

func (s HexString) String() string {
	return string(s)
}

func (s HexString) Uint64() uint64 {
	trimmed := strings.TrimPrefix(string(s), "0x")
	value, err := strconv.ParseUint(trimmed, 16, 64)
	if err != nil {
		return 0
	}
	return value
}
