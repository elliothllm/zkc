package zkc

type CdkErigonApi interface {
	GetExitRootTable() ([]ExitRoot, error)
	GetForkId() (*ForkId, error)
	BatchNumber() (string, error)
	GetBatchByNumber(number uint64, fullTx bool) (*Batch, error)
}
