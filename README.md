# ZKC

ZKC is a JSONRPC client built for interacting with the [CDK-Erigon](https://github.com/0xPolygonHermez/cdk-erigon) node.

_Currently a WIP..._

***

## Installation

To integrate ZKC into your project, use Go modules. Simply run:

```bash
go get github.com/elliothllm/zkc
```

Then, import the package in your project:

```go
import "github.com/elliothllm/zkc"
```

## Usage

Initialise the cdk-erigon client: _Currently only supports CDK Erigon_

```go
client := zkc.CDKErigonClient("localhost:8545")
```

You can configure the client with various options:

```go
client := zkc.CDKErigonClient("localhost:8545", zkc.WithTimeout(20*time.Second), zkc.WithMaxRetries(3, 5*time.Second))
```

Call an endpoint:

```go
batch, err := client.GetBatchByNumber(10, false)
if err != nil {
    return err
}
```

## Current Endpoints

- [x] zkevm_BatchNumber
- [x] zkevm_batchNumberByBlockNumber
- [x] zkevm_consolidatedBlockNumber
- [x] zkevm_estimateCounters 
- [x] zkevm_getBatchByNumber
- [x] zkevm_getBatchCountersByNumber 
- [x] zkevm_getBatchWitness 
- [x] zkevm_getBlockRangeWitness 
- [x] zkevm_getExitRootTable 
- [x] zkevm_getExitRootsByGER 
- [x] zkevm_getForkById 
- [x] zkevm_getForkId 
- [x] zkevm_getForkIdByBatchNumber 
- [x] zkevm_getForks 
- [x] zkevm_getFullBlockByHash 
- [x] zkevm_getFullBlockByNumber 
- [x] zkevm_getL2BlockInfoTree 
- [x] zkevm_getLatestDataStreamBlock 
- [ ] zkevm_getLatestGlobalExitRoot 
- [ ] zkevm_getProverInput 
- [ ] zkevm_getRollupAddress 
- [ ] zkevm_getRollupManagerAddress 
- [ ] zkevm_getVersionHistory 
- [ ] zkevm_getWitness 
- [ ] zkevm_isBlockConsolidated 
- [ ] zkevm_isBlockVirtualized 
- [x] zkevm_verifiedBatchNumber 
- [ ] zkevm_virtualBatchNumber

## Example Usage

```go
package main

import (
    "github.com/elliothllm/zkc"
    "time"
)

func main() {
	erigonClient := zkc.CDKErigonClient("", zkc.WithTimeout(20*time.Second), zkc.WithMaxRetries(3, 5*time.Second))

	res, err := erigonClient.GetBatchByNumber(10, true)
	if err != nil {
		return err
	}
}
```

## License

ZKC is released under the MIT License. See the [LICENSE](LICENSE) file for more details.

***

_This project is a work in progress. I appreciate your patience as ZKC gets developed._