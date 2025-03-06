package zkc

import (
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	clientCfg := &ClientConfig{
		Url:     "http://localhost:8123",
		Timeout: 20 * time.Second,
	}

	erigonClient := NewClient(clientCfg).CDKErigon()

	res, err := erigonClient.GetExitRootTable()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	t.Log(res)
}
