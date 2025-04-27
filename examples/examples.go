package examples

import "github.com/elliothllm/zkc"

func sendRpcRequest() error {
	req := &zkc.Request{
		JsonRpc: "2.0",
		Id:      1,
		Method:  "zkevm_getForkId",
		Params:  nil,
	}

	res, err := zkc.SendRpcRequest("http://localhost:8123", req)
	if err != nil {
		return err
	}

	_ = res

	return nil
}

func estimateCounters() error {
	erigonClient := zkc.CDKErigonClient("http://localhost:8123")

	tx := []zkc.EstimateCountersRequest{
		{
			Gas:      "0x135168",
			GasPrice: "0x3b9aca00",
			Nonce:    "0x0",
			Input:    "0x",
			To:       "0x0",
			From:     "0x0",
			Value:    "0x0",
			Data:     "0x",
		},
	}

	res, err := erigonClient.EstimateCounters(tx)
	if err != nil {
		return err
	}

	_ = res

	return nil
}
