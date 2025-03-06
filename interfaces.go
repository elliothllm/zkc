package zkc

type CdkErigonApi interface {
	GetExitRootTable() ([]ExitRoot, error)
	GetForkId() (*ForkId, error)
}
