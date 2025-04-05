package zkc

import (
	"encoding/json"
	"fmt"
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

type Batch struct {
	AccInputHash        string        `json:"accInputHash"`
	BatchL2Data         string        `json:"batchL2Data"`
	Blocks              []*Blocks     `json:"blocks"`
	Closed              bool          `json:"closed"`
	Coinbase            string        `json:"coinbase"`
	GlobalExitRoot      string        `json:"globalExitRoot"`
	LocalExitRoot       string        `json:"localExitRoot"`
	MainnetExitRoot     string        `json:"mainnetExitRoot"`
	Number              string        `json:"number"`
	RollupExitRoot      string        `json:"rollupExitRoot"`
	SendSequencesTxHash string        `json:"sendSequencesTxHash"`
	StateRoot           string        `json:"stateRoot"`
	Timestamp           string        `json:"timestamp"`
	Transactions        []Transaction `json:"transactions"`
	VerifyBatchTxHash   string        `json:"verifyBatchTxHash"`
}

type Blocks struct {
	ParentHash       string     `json:"parentHash"`
	Sha3Uncles       string     `json:"sha3Uncles"`
	Miner            string     `json:"miner"`
	StateRoot        string     `json:"stateRoot"`
	TransactionsRoot string     `json:"transactionsRoot"`
	ReceiptsRoot     string     `json:"receiptsRoot"`
	LogsBloom        string     `json:"logsBloom"`
	Difficulty       string     `json:"difficulty"`
	TotalDifficulty  string     `json:"totalDifficulty"`
	Size             string     `json:"size"`
	Number           string     `json:"number"`
	GasLimit         string     `json:"gasLimit"`
	GasUsed          string     `json:"gasUsed"`
	Timestamp        string     `json:"timestamp"`
	ExtraData        string     `json:"extraData"`
	MixHash          string     `json:"mixHash"`
	Nonce            string     `json:"nonce"`
	Hash             string     `json:"hash"`
	Transactions     []TxOrHash `json:"transactions"`
	Uncles           []string   `json:"uncles"`
	BlockInfoRoot    string     `json:"blockInfoRoot"`
	GlobalExitRoot   string     `json:"globalExitRoot"`
}

type TxOrHash struct {
	Hash        *string
	Transaction *Transaction
}

func (th *TxOrHash) UnmarshalJSON(data []byte) error {
	var hash string
	if err := json.Unmarshal(data, &hash); err == nil {
		th.Hash = &hash
		return nil
	}

	var tx Transaction
	if err := json.Unmarshal(data, &tx); err == nil {
		th.Transaction = &tx
		return nil
	}

	return fmt.Errorf("TxOrHash: unable to unmarshal JSON: %s", string(data))
}

type Transaction struct {
	Nonce       string   `json:"nonce"`
	GasPrice    string   `json:"gasPrice"`
	Gas         string   `json:"gas"`
	To          *string  `json:"to"`
	Value       string   `json:"value"`
	Input       string   `json:"input"`
	V           string   `json:"v"`
	R           string   `json:"r"`
	S           string   `json:"s"`
	Hash        string   `json:"hash"`
	From        string   `json:"from"`
	BlockHash   *string  `json:"blockHash"`
	BlockNumber *string  `json:"blockNumber"`
	TxIndex     *string  `json:"transactionIndex"`
	ChainID     *string  `json:"chainId,omitempty"`
	Type        string   `json:"type"`
	Receipt     *Receipt `json:"receipt,omitempty"`
	L2Hash      string   `json:"l2Hash,omitempty"`
}

type Receipt struct {
	CumulativeGasUsed string  `json:"cumulativeGasUsed"`
	LogsBloom         string  `json:"logsBloom"`
	Logs              []*Log  `json:"logs"`
	Status            string  `json:"status"`
	TxHash            string  `json:"transactionHash"`
	TxIndex           string  `json:"transactionIndex"`
	BlockHash         string  `json:"blockHash"`
	BlockNumber       string  `json:"blockNumber"`
	GasUsed           string  `json:"gasUsed"`
	FromAddr          string  `json:"from"`
	ToAddr            *string `json:"to"`
	ContractAddress   *string `json:"contractAddress"`
	Type              string  `json:"type"`
	EffectiveGasPrice *string `json:"effectiveGasPrice,omitempty"`
	TransactionL2Hash string  `json:"transactionL2Hash,omitempty"`
}

type Log struct {
	Address     string   `json:"address"`
	Topics      []string `json:"topics"`
	Data        string   `json:"data"`
	BlockNumber string   `json:"blockNumber"`
	TxHash      string   `json:"transactionHash"`
	TxIndex     string   `json:"transactionIndex"`
	BlockHash   string   `json:"blockHash"`
	Index       string   `json:"logIndex"`
	Removed     bool     `json:"removed"`
}
