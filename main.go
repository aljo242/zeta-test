package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// create round a robin list of RPCs to access in case of failure or outage
	rpcEndpoint      = "wss://bsc-testnet.publicnode.com"
	syncAPIEndpoint  = "http://0.0.0.0:1317"
	headerBufferSize = 16
)

func main() {
	fmt.Println("hello world")

	ctx := context.Background()

	log.Printf("connecting to %s\n", rpcEndpoint)

	// create instance of ethclient and assign to cl
	cl, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	heads := make(chan *ethtypes.Header, headerBufferSize)
	sub, err := cl.SubscribeNewHead(ctx, heads)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// Start a goroutine to update the state from head notifications in the background
	update := make(chan *ethtypes.Header)

	go func() {
		err := processHeaders(ctx, update)
		if err != nil {
			log.Fatal(err)
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

func processHeaders(ctx context.Context, headers <-chan *ethtypes.Header) error {
	var (
		prevHash    ethcommon.Hash
		currHash    ethcommon.Hash
		initialized bool
	)

	for header := range headers {
		err := header.SanityCheck()
		if err != nil {
			return err
		}

		prevHash = currHash
		currHash = header.Hash()
		if !initialized {
			initialized = true
			continue
		}

		if !bytes.Equal(prevHash.Bytes(), header.ParentHash.Bytes()) {
			return errors.New("hash mismatch")
		}

		log.Printf("recieved incoming block header with hash %s\n", currHash.String())
		err = sendMsgAddHeader(ctx, header)
		if err != nil {
			return err
		}
	}

	return nil
}

func sendMsgAddHeader(ctx context.Context, header *ethtypes.Header) error {
	log.Printf("successfully sent MsgAddHeader to sync chain\n")
	return nil
}
