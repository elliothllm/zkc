package zkc

// cdkErigonApiImpl implements the CdkErigonApi interface.
type cdkErigonApiImpl struct {
	client *Client
}

// BatchNumber calls method zkevm_batchNumber.
// It returns the current batch number in hexadecimal.
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

// BatchNumberByBlockNumber calls method zkevm_getBatchByNumber.
// It returns the batch number in hexadecimal for a given block number.
// It takes RpcBlockOrBatchNumber of type int64 where you can set the block number to negative values to represent string values.
// -1 is "latest", -2 is "pending", -3 is "safe", -4 is "finalized", -5 is "latestExecuted".
func (api *cdkErigonApiImpl) BatchNumberByBlockNumber(blockNumber RpcBlockOrBatchNumber) (string, error) {
	params := blockNumber.Params()

	req, err := newRequest(MethodZkevmBatchNumberByBlockNumber, params)
	if err != nil {
		return "", err
	}

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

// ConsolidatedBlockNumber calls method zkevm_consolidatedBlockNumber.
// It returns the latest consolidated block number in hexadecimal.
func (api *cdkErigonApiImpl) ConsolidatedBlockNumber() (string, error) {
	req := newRequestNoParams(MethodZkevmConsolidatedBlockNumber)

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

// EstimateCounters calls method zkevm_estimateCounters.
// It estimates the counters for a given transaction.
// It implements eth_estimateGas.
// It takes a slice of transactions. The transaction type must have fields "gas", "gasPrice", "nonce", "input", "to", "from", "value, "data".
func (api *cdkErigonApiImpl) EstimateCounters(transaction interface{}) (*EstimateCountersResponse, error) {
	req, err := newRequest(MethodZkevmEstimateCounters, transaction)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[*EstimateCountersResponse](resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetBatchByNumber calls method zkevm_getBatchByNumber.
// It returns the batch information for a given batch number.
// It takes RpcBlockOrBatchNumber of type int64 where you can set the batch number to negative values to represent string values.
// -1 is "latest", -2 is "pending", -3 is "safe", -4 is "finalized", -5 is "latestExecuted".
// FullTx bool indicates whether to return the full transaction details or not.
func (api *cdkErigonApiImpl) GetBatchByNumber(batchNumber RpcBlockOrBatchNumber, fullTx bool) (*Batch, error) {
	params := batchNumber.Params(fullTx)

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

// GetBatchCountersByNumber calls method zkevm_getBatchCountersByNumber.
// It returns the batch counters for a given batch number.
// It takes RpcBlockOrBatchNumber of type int64 where you can set the batch number to negative values to represent string values.
// -1 is "latest", -2 is "pending", -3 is "safe", -4 is "finalized", -5 is "latestExecuted".
func (api *cdkErigonApiImpl) GetBatchCountersByNumber(batchNumber RpcBlockOrBatchNumber) (*BatchCounters, error) {
	params := batchNumber.Params()

	req, err := newRequest(MethodZkevmGetBatchCountersByNumber, params)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[BatchCounters](resp)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetBatchWitness calls method zkevm_getBatchWitness.
// It returns the batch witness for a given batch number.
func (api *cdkErigonApiImpl) GetBatchWitness(batchNumber uint64) (string, error) {
	req, err := newRequest(MethodZkevmGetBatchWitness, []interface{}{batchNumber})
	if err != nil {
		return "", err
	}

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

// GetBlockRangeWitness calls method zkevm_getBlockRangeWitness.
// It returns the block range witness for a given start and end block number.
// It takes startBlock and endBlock as uint64 values.
// You may specify the WitnessMode which is a string type with values "full", or "trimmed".
// The debug parameter is a boolean value indicating whether to include debug information on the node.
func (api *cdkErigonApiImpl) GetBlockRangeWitness(startBlock, endBlock uint64, mode *WitnessMode, debug bool) (string, error) {
	params := []interface{}{startBlock, endBlock, mode, debug}

	req, err := newRequest(MethodZkevmGetBlockRangeWitness, params)
	if err != nil {
		return "", err
	}

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

// GetExitRootTable calls method zkevm_getExitRootTable.
// It returns the exit root table as a slice of ExitRoot structs.
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

// GetExitRootsByGER calls method zkevm_getExitRootsByGER.
func (api *cdkErigonApiImpl) GetExitRootsByGER(ger string) (*ExitRoots, error) {
	params := []interface{}{ger}

	req, err := newRequest(MethodZkevmGetExitRootsByGER, params)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[*ExitRoots](resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetForkByID calls method zkevm_getForkByID.
// It returns a ForkID struct for a given fork ID of type string.
// forkId must be a valid hexadecimal string value, e.g., "0x1", "0x2", etc.
func (api *cdkErigonApiImpl) GetForkByID(forkId string) (*ForkID, error) {
	params := []interface{}{forkId}

	req, err := newRequest(MethodZkevmGetForkById, params)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[*ForkID](resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetForkId calls method zkevm_getForkId.
// It returns the current fork ID as a HexString string type.
func (api *cdkErigonApiImpl) GetForkId() (*HexString, error) {
	req := newRequestNoParams(MethodZkevmGetForkId)

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[HexString](resp)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetForkIdByBatchNumber calls method zkevm_getForkIdByBatchNumber.
// It takes RpcBlockOrBatchNumber of type int64 where you can set the batch number to negative values to represent string values.
// -1 is "latest", -2 is "pending", -3 is "safe", -4 is "finalized", -5 is "latestExecuted".
func (api *cdkErigonApiImpl) GetForkIdByBatchNumber(batchNumber RpcBlockOrBatchNumber) (*HexString, error) {
	params := batchNumber.Params()

	req, err := newRequest(MethodZkevmGetForkIdByBatchNumber, params)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[HexString](resp)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetForks calls method zkevm_getForks.
// It returns a slice of ForkID structs.
func (api *cdkErigonApiImpl) GetForks() ([]ForkID, error) {
	req := newRequestNoParams(MethodZkevmGetForks)

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[[]ForkID](resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetFullBlockByHash calls method zkevm_getFullBlockByHash.
// It takes a block hash of type string. e.g., "0x1234...".
// FullTx bool indicates whether to return the full transaction details or not.
func (api *cdkErigonApiImpl) GetFullBlockByHash(blockHash string, fullTx bool) (*Block, error) {
	params := []interface{}{blockHash, fullTx}

	req, err := newRequest(MethodZkevmGetFullBlockByHash, params)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[Block](resp)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetFullBlockByNumber calls method zkevm_getFullBlockByNumber.
// It takes RpcBlockOrBatchNumber of type int64 where you can set the block number to negative values to represent string values.
// -1 is "latest", -2 is "pending", -3 is "safe", -4 is "finalized", -5 is "latestExecuted".
// FullTx bool indicates whether to return the full transaction details or not.
func (api *cdkErigonApiImpl) GetFullBlockByNumber(blockNumber RpcBlockOrBatchNumber, fullTx bool) (*Block, error) {
	params := blockNumber.Params(fullTx)

	req, err := newRequest(MethodZkevmGetFullBlockByNumber, params)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[Block](resp)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetL2BlockInfoTree calls method zkevm_getL2BlockInfoTree.
// It takes RpcBlockOrBatchNumber of type int64 where you can set the block number to negative values to represent string values.
// -1 is "latest", -2 is "pending", -3 is "safe", -4 is "finalized", -5 is "latestExecuted".
func (api *cdkErigonApiImpl) GetL2BlockInfoTree(blockNum RpcBlockOrBatchNumber) (*L2BlockInfoTree, error) {
	params := blockNum.Params()

	req, err := newRequest(MethodZkevmGetL2BlockInfoTree, params)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.handleRequest(req)
	if err != nil {
		return nil, err
	}

	result, err := getResult[L2BlockInfoTree](resp)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *cdkErigonApiImpl) GetLatestDataStreamBlock() (string, error) {
	req := newRequestNoParams(MethodZkevmGetLatestDataStreamBlock)

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

func (api *cdkErigonApiImpl) VerifiedBatchNumber() (string, error) {
	req := newRequestNoParams(MethodZkevmVerifiedBatchNumber)

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
