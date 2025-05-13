package zkc

type CdkErigonApi interface {
	BatchNumber() (string, error)
	BatchNumberByBlockNumber(blockNumber RpcBlockOrBatchNumber) (string, error)
	ConsolidatedBlockNumber() (string, error)
	EstimateCounters(transaction interface{}) (*EstimateCountersResponse, error)
	GetBatchByNumber(batchNumber RpcBlockOrBatchNumber, fullTx bool) (*Batch, error)
	GetBatchCountersByNumber(batchNumber RpcBlockOrBatchNumber) (*BatchCounters, error)
	GetBatchWitness(batchNumber uint64) (string, error)
	GetBlockRangeWitness(startBlock, endBlock uint64, mode *WitnessMode, debug bool) (string, error)
	GetExitRootTable() ([]ExitRoot, error)
	GetExitRootsByGER(ger string) (*ExitRoots, error)
	GetForkByID(forkId string) (*ForkID, error)
	GetForkId() (*HexString, error)
	GetForkIdByBatchNumber(batchNumber RpcBlockOrBatchNumber) (*HexString, error)
	GetForks() ([]ForkID, error)
	GetFullBlockByHash(blockHash string, fullTx bool) (*Block, error)
	GetFullBlockByNumber(blockNumber RpcBlockOrBatchNumber, fullTx bool) (*Block, error)
	GetL2BlockInfoTree(blockNum RpcBlockOrBatchNumber) (*L2BlockInfoTree, error)
	GetLatestDataStreamBlock() (string, error)
	VerifiedBatchNumber() (string, error)
}
