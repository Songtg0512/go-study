package counter

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"counter-demo/counter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. 连接 Sepolia
	rpcURL := "https://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMbZ1AkVvIRu9N8PyUX0Z"
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	// 2. 加载私钥
	privateKey, err := crypto.HexToECDSA("key")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 3. 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 4. 构造交易参数
	chainID := big.NewInt(11155111) // Sepolia Chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // 不转 ETH
	auth.GasLimit = uint64(3000000) // 视情况调整
	auth.GasPrice = big.NewInt(20e9)

	// 5. 部署合约
	address, tx, counterInstance, err :=
		counter.DeployCounter(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract deployed at:", address.Hex())
	fmt.Println("Deploy tx:", tx.Hash().Hex())

	// 6. 调用 increment()
	tx, err = counterInstance.Increment(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Increment tx:", tx.Hash().Hex())

	// 7. 读取 count
	count, err := counterInstance.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current count:", count)
}
