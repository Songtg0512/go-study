package task1

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 这个 API Key 暂时用的别人的
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMbZ1AkVvIRu9N8PyUX0Z")

	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)

	// 来返回有关一个区块的头信息。若传入 nil，它将返回最新的区块头
	header, err := client.HeaderByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.Uint64())
	//fmt.Println(header.Time)
	//fmt.Println(header.Difficulty.Uint64())
	//fmt.Println(header.Hash())

	// 方法来获得完整区块。您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	//fmt.Println(block.Number().Uint64())     // 5671744
	//fmt.Println(block.Time())                // 1712798400
	//fmt.Println(block.Difficulty().Uint64()) // 0
	//fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	//fmt.Println(len(block.Transactions()))   // 70

	// 先从客户端拿到链 ID。
	// 这里拿到的 chainId ，是执行 client.Dial 链接的地址的 chainId
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 查询交易的详细信息
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		// 如果调用的是合约的话，这里会是方法签名
		fmt.Println(tx.Data())
		// 这里指的是这笔交易的接收方
		fmt.Println(tx.To().Hex())

		// 为了读取发送方的地址，我们需要在事务上调用 AsMessage，它返回一个 Message 类型，其中包含一个返回 sender（from）地址的函数。 AsMessage 方法需要 EIP155 签名者
		// 先从客户端拿到链 ID。, 在使用 EIP155Signer 还原出 sender 地址：
		if sender, err := types.Sender(types.NewEIP155Signer(chainId), tx); err == nil {
			fmt.Println("sender", sender.Hex())
		} else {
			log.Fatal(err)
		}
		// 每个交易都有一个收据，其中包含执行交易的结果，例如所有的返回值和日志，以及“1”（成功）或“0”（失败）的交易结果状态。
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		fmt.Println(receipt.ContractAddress)

		// 这里，如果我们提前知道交易的 hash 的话，那么我们可以直接通过交易 hash 去获取交易数据，同时还可以获取到这笔交易是否是 isPending 状态
		txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")

		tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(isPending)       // false

	}

}
