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

***

## Usage

Configure the client by creating a new config instance with your desired configuration:

```go
clientCfg := &zkc.ClientConfig{
    Url:     "http://localhost:8123",
    Timeout: 20 * time.Second,
}
```

Initialise the client and call the desired client: _Currently only supports CDK Erigon_

```go
erigonClient := zkc.NewClient(clientCfg).CDKErigon()
```

***

## Current Endpoints

- [x] zkevm_getForkId
- [x] zkevm_getExitRootTable

***

## Example Usage

```go
package main

import (
    "github.com/elliothllm/zkc"
    "time"
    "log"
    "fmt"
)

func main() {
	clientCfg := &zkc.ClientConfig{
		Url:     "http://localhost:8123",
		Timeout: 20 * time.Second,
	}

	erigonClient := zkc.NewClient(clientCfg).CDKErigon()

	res, err := erigonClient.GetForkId()
	if err != nil {
		log.Fatalf("Error getting ForkId: %v", err)
	}

	forkId := res.Uint64()
	
	fmt.Printf("ForkId: %d\n", forkId)
}
```

***

## License

ZKC is released under the MIT License. See the [LICENSE](LICENSE) file for more details.

***

_This project is a work in progress. I appreciate your patience as ZCK gets developed._