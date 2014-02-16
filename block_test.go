package blockchain_test

import (
	"github.com/qedus/blockchain"
	"net/http"
	"testing"
	"time"
)

const (
	blockHash = "0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103"
)

func TestRequestBlock(t *testing.T) {
	bc := blockchain.New(http.DefaultClient)
	block := &blockchain.Block{Hash: blockHash}
	if err := bc.Request(block); err != nil {
		t.Fatal(err)
	}

	if block.Fee != 200000 {
		t.Fatal("fee not 200000")
	}

	if block.TransactionCount != 22 {
		t.Fatal("transaction count not 22")
	}

	if len(block.Transactions) != 22 {
		t.Fatal("transactions length not 22")
	}

	if block.Transactions[0].Hash != "5b09bbb8d3cb2f8d4edbcf30664419fb7c9deaeeb1f62cb432e7741c80dbe5ba" {
		t.Fatal("first transaction hash incorrect")
	}
}

func TestRequestLatestBlock(t *testing.T) {

	bc := blockchain.New(http.DefaultClient)
	block := &blockchain.LatestBlock{}
	if err := bc.Request(block); err != nil {
		t.Fatal(err)
	}

	if time.Unix(block.Time, 0).Before(time.Now().Add(-30 * time.Minute)) {
		t.Fatal("latest block too old")
	}

	if len(block.TransactionIndexes) < 1 {
		t.Fatal("no transactions in latest block")
	}
}

func TestRequestBlockHeight(t *testing.T) {
	bc := blockchain.New(http.DefaultClient)
	bh := &blockchain.BlockHeight{Height: 285180}
	if err := bc.Request(bh); err != nil {
		t.Fatal(err)
	}

	if len(bh.Blocks) != 2 {
		t.Fatal("should be two blocks")
	}

	if !bh.Blocks[0].MainChain {
		t.Fatal("this block should be on the main chain")
	}

	if bh.Blocks[1].MainChain {
		t.Fatal("this block should not on the main chain")
	}
}