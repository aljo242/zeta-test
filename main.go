package main

import (
	"bytes"
	"context"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// create round a robin list of RPCs to access in case of failure or outage
	rpcEndpoint      = "wss://bsc-testnet.publicnode.com"
	headerBufferSize = 16
)

func main() {
	fmt.Println("hello world")

	ctx := context.Background()

	// create instance of ethclient and assign to cl
	cl, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		panic(err)
	}

	peerCount, err := cl.PeerCount(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(peerCount)

	heads := make(chan *ethtypes.Header, headerBufferSize)
	sub, err := cl.SubscribeNewHead(ctx, heads)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	// Start a goroutine to update the state from head notifications in the background
	update := make(chan *ethtypes.Header)

	go func() {
		var (
			prevHash    ethcommon.Hash
			currHash    ethcommon.Hash
			initialized bool
		)

		for head := range update {
			err := head.SanityCheck()
			if err != nil {
				panic(err)
			}

			prevHash = currHash
			currHash = head.Hash()
			if !initialized {
				initialized = true
				continue
			}

			if !bytes.Equal(prevHash.Bytes(), head.ParentHash.Bytes()) {
				panic("hash mismatch")
			} else {
				fmt.Println(currHash)
			}

		}
	}()

	// Wait for various events and passing to the appropriate background threads
	for {
		select {
		case head := <-heads:
			// New head arrived, send if for state update if there's none running
			select {
			case update <- head:
			default:
			}
		}
	}
}
