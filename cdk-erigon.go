package zkc

import "strconv"

type cdkErigonApiImpl struct {
	client *Client
}

func (api *cdkErigonApiImpl) GetExitRootTable() ([]ExitRoot, error) {
	req := newRequestNoParams(MethodZkevmGetExitRootTable)

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[[]ExitRoot](resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (api *cdkErigonApiImpl) GetForkId() (*ForkId, error) {
	req := newRequestNoParams(MethodZkevmGetForkId)

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[ForkId](resp)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *cdkErigonApiImpl) BatchNumber() (string, error) {
	req := newRequestNoParams(MethodZkevmBatchNumber)

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return "", err
	}

	result, err := getResult[string](resp)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (api *cdkErigonApiImpl) GetBatchByNumber(number uint64, fullTx bool) (*Batch, error) {
	params := []interface{}{strconv.FormatUint(number, 10), fullTx}

	req, err := newRequest(MethodZkevmGetBatchByNumber, params)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[Batch](resp)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
