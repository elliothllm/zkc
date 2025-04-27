package zkc

import (
	"testing"
)

func TestClient(t *testing.T) {
	erigonClient := CDKErigonClient("http://localhost:8123")

	res, err := erigonClient.GetForkByID("0x9")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	t.Log(res)
}
