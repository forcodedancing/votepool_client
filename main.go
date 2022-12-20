package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/rpc/jsonrpc/client"
	"github.com/tendermint/tendermint/votepool"
)

func main() {
	bs := common.Hex2Bytes("38ebe1ea64da024ac6ac7a9b5f484293c6f391b65837c9b5a8b27b47bef96f42")
	secretKey, err := blst.SecretKeyFromBytes(bs)
	if err != nil {
		panic(err)
	}
	pubKey := secretKey.PublicKey()

	eventHash := common.HexToHash("0xeefacfed87736ae1d8e8640f6fd7951862997782e5e79842557923e2779d5d5a").Bytes()
	//secKey, _ := blst.SecretKeyFromBytes(privKey.Marshal())
	sign := secretKey.Sign(eventHash).Marshal()

	fmt.Printf("pubkey: %s \n", base64.StdEncoding.EncodeToString(pubKey.Marshal()))
	fmt.Printf("eventHash: %s \n", base64.StdEncoding.EncodeToString(eventHash))
	fmt.Printf("sign: %s \n", base64.StdEncoding.EncodeToString(sign))

	vote := votepool.Vote{
		PubKey:    pubKey.Marshal(),
		Signature: sign,
		EventType: 1,
		EventHash: eventHash,
	}

	c1, err := client.New("http://127.0.0.1:26657")
	if err != nil {
		panic(err)
	}

	c2, err := client.New("http://127.0.0.1:27657")
	if err != nil {
		panic(err)
	}

	broadcastMap := make(map[string]interface{})
	broadcastMap["vote"] = vote
	var broadcastVote coretypes.ResultBroadcastVote

	queryMap1 := make(map[string]interface{})
	queryMap1["event_type"] = 1

	queryMap2 := make(map[string]interface{})
	queryMap2["event_type"] = 1
	queryMap2["event_hash"] = eventHash
	var queryVote coretypes.ResultQueryVote

	// send from c1
	_, err = c1.Call(context.Background(), "broadcast_vote", broadcastMap, &broadcastVote)
	if err != nil {
		panic(err)
	}

	fmt.Println("sleep 3 seconds")
	time.Sleep(3 * time.Second)

	// query from c1
	_, err = c1.Call(context.Background(), "query_vote", queryMap1, &queryVote)
	if err != nil {
		panic(err)
	}
	if 1 != len(queryVote.Votes) {
		panic("not found")
	}
	fmt.Printf("votes from c1: %x \n", queryVote.Votes[0].EventHash)

	queryVote = coretypes.ResultQueryVote{}
	_, err = c1.Call(context.Background(), "query_vote", queryMap2, &queryVote)
	if err != nil {
		panic(err)
	}
	if 1 != len(queryVote.Votes) {
		panic("not found")
	}
	fmt.Printf("votes from c1: %x \n", queryVote.Votes[0].EventHash)

	// query from c2
	queryVote = coretypes.ResultQueryVote{}
	_, err = c2.Call(context.Background(), "query_vote", queryMap1, &queryVote)
	if err != nil {
		panic(err)
	}
	if 1 != len(queryVote.Votes) {
		panic("not found")
	}
	fmt.Printf("votes from c2: %x \n", queryVote.Votes[0].EventHash)

	queryVote = coretypes.ResultQueryVote{}
	_, err = c2.Call(context.Background(), "query_vote", queryMap2, &queryVote)
	if err != nil {
		panic(err)
	}
	if 1 != len(queryVote.Votes) {
		panic("not found")
	}
	fmt.Printf("votes from c2: %x \n", queryVote.Votes[0].EventHash)

	fmt.Println("sleep 30 seconds")
	time.Sleep(30 * time.Second)

	// query from c1
	queryVote = coretypes.ResultQueryVote{}
	_, err = c1.Call(context.Background(), "query_vote", queryMap1, &queryVote)
	if err != nil {
		panic(err)
	}
	if 0 != len(queryVote.Votes) {
		panic("still found")
	}

	_, err = c1.Call(context.Background(), "query_vote", queryMap2, &queryVote)
	if err != nil {
		panic(err)
	}
	if 0 != len(queryVote.Votes) {
		panic("still found")
	}

	// query from c2
	queryVote = coretypes.ResultQueryVote{}
	_, err = c2.Call(context.Background(), "query_vote", queryMap1, &queryVote)
	if err != nil {
		panic(err)
	}
	if 0 != len(queryVote.Votes) {
		panic("still found")
	}

	_, err = c2.Call(context.Background(), "query_vote", queryMap2, &queryVote)
	if err != nil {
		panic(err)
	}
	if 0 != len(queryVote.Votes) {
		panic("still found")
	}

	// send from c1 again
	_, err = c1.Call(context.Background(), "broadcast_vote", broadcastMap, &broadcastVote)
	if err != nil {
		panic(err)
	}

	fmt.Println("sleep 1 seconds")
	time.Sleep(1 * time.Second)

	// query from c1 again
	queryVote = coretypes.ResultQueryVote{}
	_, err = c1.Call(context.Background(), "query_vote", queryMap1, &queryVote)
	if err != nil {
		panic(err)
	}
	if 1 != len(queryVote.Votes) {
		panic("not found")
	}
	fmt.Printf("votes from c1: %x \n", queryVote.Votes[0].EventHash)

	// query from c2 again
	queryVote = coretypes.ResultQueryVote{}
	_, err = c2.Call(context.Background(), "query_vote", queryMap1, &queryVote)
	if err != nil {
		panic(err)
	}
	if 1 != len(queryVote.Votes) {
		panic("not found")
	}
	fmt.Printf("votes from c2: %x \n", queryVote.Votes[0].EventHash)
}
