package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

// 实时交易监控器
type TransactionMonitor struct {
	client        *client.Client
	lastSignature string
	filters       []TransactionFilter
	callbacks     []TransactionCallback
}

type TransactionFilter func(tx *client.Transaction) bool
type TransactionCallback func(tx client.Transaction)

func NewTransactionMonitor(rpcEndpoint string) *TransactionMonitor {
	return &TransactionMonitor{
		client:    client.NewClient(rpcEndpoint),
		filters:   make([]TransactionFilter, 0),
		callbacks: make([]TransactionCallback, 0),
	}
}

// 添加过滤器
func (tm *TransactionMonitor) AddFilter(filter TransactionFilter) {
	tm.filters = append(tm.filters, filter)
}

// 添加回调
func (tm *TransactionMonitor) AddCallback(callback TransactionCallback) {
	tm.callbacks = append(tm.callbacks, callback)
}

func MonitorStart() {
	transactionMonitor := NewTransactionMonitor(rpc.DevnetRPCEndpoint)
	transactionMonitor.Start(context.Background(), time.Duration(6000))
}

// 开始监控
func (tm *TransactionMonitor) Start(ctx context.Context, interval time.Duration) {
	fmt.Println("🔍 开始监控交易...")

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🛑 停止监控")
			return
		case <-ticker.C:
			tm.pollRecentTransactions(ctx)
		}
	}
}

// 轮询最近交易
func (tm *TransactionMonitor) pollRecentTransactions(ctx context.Context) {
	// 获取最近的区块
	slot, err := tm.client.GetSlot(ctx)
	if err != nil {
		log.Printf("获取区块失败: %v", err)
		return
	}

	// 获取区块的交易签名
	blockSignatures, err := tm.client.GetBlock(ctx, slot)
	if err != nil {
		log.Printf("获取区块交易失败: %v", err)
		return
	}

	if blockSignatures == nil || len(blockSignatures.Transactions) == 0 {
		return
	}

	// 处理新交易
	for _, txSig := range blockSignatures.Transactions {
		tm.processTransaction(ctx, string(txSig.Transaction.Signatures[0]))
	}
}

// 处理单个交易
func (tm *TransactionMonitor) processTransaction(ctx context.Context, signature string) {
	// 如果是已经处理过的交易，跳过
	if signature == tm.lastSignature {
		return
	}

	// 获取交易详情
	tx, err := tm.client.GetTransaction(ctx, signature)
	if err != nil {
		log.Printf("获取交易详情失败: %v", err)
		return
	}

	if tx == nil {
		return
	}

	// 应用过滤器
	shouldProcess := true
	for _, filter := range tm.filters {
		if !filter(tx) {
			shouldProcess = false
			break
		}
	}

	if !shouldProcess {
		return
	}

	fmt.Printf("\n发现新交易: %s\n", signature)

	// 分析交易
	tm.analyzeTransaction(tx)

	// 调用回调
	for _, callback := range tm.callbacks {
		callback(*tx)
	}

	tm.lastSignature = signature
}

// 分析交易
func (tm *TransactionMonitor) analyzeTransaction(tx *client.Transaction) {
	fmt.Printf("📊 交易分析:\n")
	fmt.Printf("  区块: %d\n", tx.Slot)

	if tx.BlockTime != nil {
		timestamp := time.Unix(int64(*tx.BlockTime), 0)
		fmt.Printf("  时间: %s\n", timestamp.Format("2006-01-02 15:04:05"))
	}

	// 检查交易状态
	if tx.Meta != nil {
		if tx.Meta.Err != nil {
			fmt.Printf("  状态: ❌ 失败\n")
			fmt.Printf("  错误: %v\n", tx.Meta.Err)
		} else {
			fmt.Printf("  状态: ✅ 成功\n")
		}

		// 计算费用
		if tx.Meta.Fee != 0 {
			fmt.Printf("  费用: %d lamports (%.6f SOL)\n",
				tx.Meta.Fee, float64(tx.Meta.Fee)/1e9)
		}

		// 计算单元
		if tx.Meta.ComputeUnitsConsumed != nil {
			fmt.Printf("  计算单元: %d\n", *tx.Meta.ComputeUnitsConsumed)
		}
	}

	// 账户信息
	if tx.Transaction != nil && tx.Transaction.Message != nil {
		fmt.Printf("  签名者数量: %d\n", len(tx.Transaction.Signatures))
		fmt.Printf("  账户数量: %d\n", len(tx.Transaction.Message.AccountKeys))
	}
}
