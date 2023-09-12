package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/trie"
)

// proveCmd represents the prove command
var proveCmd = &cobra.Command{
	Use:   "prove",
	Short: "prove some data",
	Long:  `Demo application to demonstrate cobra featues`,
	Args:  cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "tx":
			proveTx(args[1])
		case "rcpt":
			blockNo, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal(err)
			}

			proveReceipt(int64(blockNo), args[2])
		default:
			fmt.Println("invalid request")
		}

	},
}

func init() {
	rootCmd.AddCommand(proveCmd)
}

func proveTx(txHash string) {
	ctx := context.Background()
	// create instance of ethclient and assign to ethClient
	ethClient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := ethClient.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}

	block, err := ethClient.BlockByNumber(ctx, receipt.BlockNumber)
	if err != nil {
		log.Fatal(err)
	}

	hash := ethtypes.DeriveSha(block.Transactions(), trie.NewStackTrie(nil))

	// return final hash
	fmt.Println("txRoot Proof:")
	fmt.Println(hash.String())
	fmt.Println("block Number:")
	fmt.Println(receipt.BlockNumber)
}

func proveReceipt(blockNumber int64, receiptHash string) {
	ctx := context.Background()
	// create instance of ethclient and assign to ethClient
	ethClient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	block, err := ethClient.BlockByNumber(ctx, big.NewInt(blockNumber))
	if err != nil {
		log.Fatal(err)
	}

	found := false
	receipts := make(ethtypes.Receipts, len(block.Transactions()))
	for i, tx := range block.Transactions() {
		receipt, err := ethClient.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		receipts[i] = receipt

		// check if our receipt exists
		if receiptHash == receipt.TxHash.String() {
			found = true
		}
	}

	if found == false {
		fmt.Println("invalid receipt not in block")
		// return
	}

	hash := ethtypes.DeriveSha(receipts, trie.NewStackTrie(nil))

	// return final hash
	fmt.Println("receiptRoot Proof:")
	fmt.Println(hash.String())
}
