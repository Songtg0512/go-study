package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

func SearchAccount() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	ctx := context.Background()

	resp, err := c.GetVersion(ctx)
	if err != nil {
		log.Fatalf("GetVersion: %v", err)
	}
	log.Println("GetVersion", resp.SolanaCore)

	address := os.Getenv("address")
	balance, err := c.GetBalance(ctx, address)
	fmt.Println(balance)

	info, err := c.GetAccountInfo(ctx, address)
	if err != nil {
		log.Fatalf("GetAccountInfo: %v", err)
	}
	log.Printf("GetAccountInfo%v", info)

	balance1, err1 := c.GetBalanceWithConfig(context.TODO(), address, client.GetBalanceConfig{
		Commitment: rpc.CommitmentFinalized,
	})
	if err1 != nil {
		log.Fatalf("GetBalanceWithConfig: %v", err1)
	}
	log.Printf("GetBalanceWithConfig %v", balance1)

	// 获取最新的区块高度
	slot, err := client.GetSlot(ctx)
	if err != nil {
		log.Fatal("获取最新slot失败:", err)
	}
	fmt.Printf("最新slot: %d\n", slot)
	// 获取最新区块
	recentBlock, err := client.GetBlock(ctx, slot)
	if err != nil {
		panic("查询失败: " + err.Error())
	}

	fmt.Printf("区块高度: %d\n", recentBlock.BlockHeight)
	fmt.Printf("交易数量: %d\n", len(recentBlock.Transactions))
}
