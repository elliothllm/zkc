package zkc

import (
	"encoding/json"
	"fmt"
)

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
	AccInputHash        string         `json:"accInputHash"`
	BatchL2Data         string         `json:"batchL2Data"`
	Blocks              Blocks         `json:"blocks"`
	Closed              bool           `json:"closed"`
	Coinbase            string         `json:"coinbase"`
	GlobalExitRoot      string         `json:"globalExitRoot"`
	LocalExitRoot       string         `json:"localExitRoot"`
	MainnetExitRoot     string         `json:"mainnetExitRoot"`
	Number              string         `json:"number"`
	RollupExitRoot      string         `json:"rollupExitRoot"`
	SendSequencesTxHash string         `json:"sendSequencesTxHash"`
	StateRoot           string         `json:"stateRoot"`
	Timestamp           string         `json:"timestamp"`
	Transactions        []*Transaction `json:"transactions"`
	VerifyBatchTxHash   string         `json:"verifyBatchTxHash"`
}

type Blocks struct {
	BlockHashes []string
	FullBlocks  []Block
}

func (b *Blocks) UnmarshalJSON(data []byte) error {
	var ss []string
	if err := json.Unmarshal(data, &ss); err == nil {
		b.BlockHashes = ss
		return nil
	}
	var ff []Block
	if err := json.Unmarshal(data, &ff); err == nil {
		b.FullBlocks = ff
		return nil
	}
	return fmt.Errorf("blocks: cannot unmarshal %s into []string or []BlockFull", string(data))
}

func (b Blocks) MarshalJSON() ([]byte, error) {
	if b.BlockHashes != nil {
		return json.Marshal(b.BlockHashes)
	}
	return json.Marshal(b.FullBlocks)
}

type Block struct {
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

type BatchCounters struct {
	SmtDepth       uint64   `json:"smtDepth"`
	BatchNumber    uint64   `json:"batchNumber"`
	BlockFrom      uint64   `json:"blockFrom"`
	BlockTo        uint64   `json:"blockTo"`
	CountersUsed   Counters `json:"countersUsed"`
	CountersLimits Counters `json:"countersLimits"`
}

type Counters struct {
	Gas              uint64 `json:"gas"`
	KeccakHashes     uint64 `json:"keccakHashes"`
	PoseidonHashes   uint64 `json:"poseidonHashes"`
	PoseidonPaddings uint64 `json:"poseidonPaddings"`
	MemAligns        uint64 `json:"memAligns"`
	Arithmetics      uint64 `json:"arithmetics"`
	Binaries         uint64 `json:"binaries"`
	Steps            uint64 `json:"steps"`
	SHA256Hashes     uint64 `json:"SHA256Hashes"`
}

// EstimateCountersRequest represents the request for estimating counters.
// All fields must be hexadecimal strings. Both to, and from are 40 byte long addresses.
type EstimateCountersRequest struct {
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Nonce    string `json:"nonce"`
	Input    string `json:"input"`
	To       string `json:"to"`
	From     string `json:"from"`
	Value    string `json:"value"`
	Data     string `json:"data"`
}

type EstimateCountersResponse struct {
	CountersUsed   Counters `json:"countersUsed"`
	CountersLimits Counters `json:"countersLimits"`
	RevertInfo     Revert   `json:"revertInfo"`
	OocError       string   `json:"oocError"`
}

type Revert struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ExitRoots struct {
	BlockNumber     HexString `json:"blockNumber"`
	Timestamp       HexString `json:"timestamp"`
	MainnetExitRoot string    `json:"mainnetExitRoot"`
	RollupExitRoot  string    `json:"rollupExitRoot"`
}

type ForkID struct {
	ForkId          HexString `json:"forkId"`
	FromBatchNumber HexString `json:"fromBatchNumber"`
	ToBatchNumber   HexString `json:"toBatchNumber"`
	Version         string    `json:"version"`
	BlockNumber     HexString `json:"blockNumber"`
}
