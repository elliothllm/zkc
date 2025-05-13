package zkc

import (
	"testing"
)

func TestClient(t *testing.T) {
	erigonClient := CDKErigonClient("http://localhost:8124")

	res, err := erigonClient.GetL2BlockInfoTree(1)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	t.Log(res)
}
