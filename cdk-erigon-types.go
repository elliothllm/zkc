package zkc

import (
	"strconv"
	"strings"
)

type ForkId string

func (id ForkId) String() string {
	return string(id)
}

func (id ForkId) Uint64() uint64 {
	trimmed := strings.TrimPrefix(string(id), "0x")
	value, err := strconv.ParseUint(trimmed, 16, 64)
	if err != nil {
		return 0
	}
	return value
}

type ExitRoot struct {
	Index           uint64 `json:"index"`
	Ger             string `json:"ger"`
	InfoRoot        string `json:"info_root"`
	MainnetExitRoot string `json:"mainnet_exit_root"`
	RollupExitRoot  string `json:"rollup_exit_root"`
	ParentHash      string `json:"parent_hash"`
	MinTimestamp    uint64 `json:"min_timestamp"`
	BlockNumber     uint64 `json:"block_number"`
}
