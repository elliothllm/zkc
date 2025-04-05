package zkc

import (
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	erigonClient := CDKErigonClient("", WithTimeout(20*time.Second), WithMaxRetries(3, 5*time.Second))

	res, err := erigonClient.GetBatchByNumber(10, true)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	t.Log(res)
}
