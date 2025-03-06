package zkc

type cdkErigonApiImpl struct {
	client *Client
}

func (api *cdkErigonApiImpl) GetExitRootTable() ([]ExitRoot, error) {
	req := newRequestNoParams(CDKErigonZkevmGetExitRootTable)

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
	req := newRequestNoParams(CDKErigonZkevmGetForkId)

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
