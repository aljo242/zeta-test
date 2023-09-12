package cmd

import (
	"bytes"
	"context"
	"fmt"
	synctypes "github.com/aljo242/sync/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	// create round a robin list of RPCs to access in case of failure or outage
	rpcEndpoint      = "wss://bsc-testnet.publicnode.com"
	syncAPIEndpoint  = "http://0.0.0.0:26657"
	headerBufferSize = 16
)

// watcherCmd represents the watcher command
var watcherCmd = &cobra.Command{
	Use:   "watcher",
	Short: "add values passed to function",
	Long:  `Demo application to demonstrate cobra featues`,
	Run: func(cmd *cobra.Command, args []string) {
		runWatcher()
	},
}

func init() {
	rootCmd.AddCommand(watcherCmd)
}

func runWatcher() {
	ctx := context.Background() // TODO add timeout option

	// build a client to connect to node
	cosmosClient, err := buildCosmosClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connecting to %s\n", rpcEndpoint)

	// create instance of ethclient and assign to ethClient
	ethClient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// subscribe to all new headers on chain
	heads := make(chan *ethtypes.Header, headerBufferSize)
	sub, err := ethClient.SubscribeNewHead(ctx, heads)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// Start a goroutine to update the state from head notifications in the background
	update := make(chan *ethtypes.Header)

	go func() {
		err := processHeaders(ctx, update, cosmosClient)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for events and pass to the appropriate background threads
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

func processHeaders(ctx context.Context, headers <-chan *ethtypes.Header, client cosmosclient.Client) error {
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
			log.Printf("hash mismatch")
		}

		log.Printf("recieved incoming block header %d with hash %s\n", header.Number.Uint64(), currHash.String())
		err = sendMsgAddHeader(ctx, header, client)
		if err != nil {
			return err
		}
	}

	return nil
}

func sendMsgAddHeader(ctx context.Context, header *ethtypes.Header, client cosmosclient.Client) error {
	var opts []cosmosaccount.Option
	opts = append(opts,
		cosmosaccount.WithKeyringBackend(cosmosaccount.KeyringTest),
		cosmosaccount.WithHome(os.Getenv("HOME")+"/.sync"),
	)

	reg, err := cosmosaccount.New(opts...)
	if err != nil {
		return fmt.Errorf("failed to create key registry: %w", err)
	}

	acct, err := reg.GetByName("alice")
	if err != nil {
		return fmt.Errorf("failed to get account: %w", err)
	}

	accAddr, _ := acct.Record.GetAddress()
	admin := sdk.MustBech32ifyAddressBytes("cosmos", accAddr)
	msg := ethHeaderToZetaHeader(admin, header)
	msgs := []sdk.Msg{sdk.Msg(msg)}

	_, err = client.BroadcastTx(ctx, acct, msgs...)
	if err != nil {
		return fmt.Errorf("failed to broadcast tx: %w", err)
	}

	log.Printf("successfully sent MsgAddHeader to sync chain\n")

	return nil
}

func ethHeaderToZetaHeader(admin string, header *ethtypes.Header) *synctypes.MsgCreateHeader {
	return &synctypes.MsgCreateHeader{
		Admin:       admin,
		ParentHash:  header.ParentHash.String(),
		UncleHash:   header.UncleHash.String(),
		RootHash:    header.Root.String(),
		TxHash:      header.TxHash.String(),
		ReceiptHash: header.ReceiptHash.String(),
		Hash:        header.Hash().String(),
		BlockNumber: header.Number.Uint64(),
	}
}

func buildCosmosClient(ctx context.Context) (cosmosclient.Client, error) {
	log.Printf("connecting to sync chain at %s\n", syncAPIEndpoint)

	var opts []cosmosclient.Option
	opts = append(opts,
		cosmosclient.WithKeyringBackend(cosmosaccount.KeyringTest),
		cosmosclient.WithNodeAddress(syncAPIEndpoint),
		cosmosclient.WithHome(os.Getenv("HOME")+"/.sync"),
		cosmosclient.WithFees("10stake"),
	)

	return cosmosclient.New(ctx, opts...)
}
