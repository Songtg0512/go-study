### 安装
brew install solidity

### 执行命令
solc --abi --bin Counter.sol -o build

### 安装 abigen
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

### 生成代码
abigen \
  --abi build/Counter.abi \
  --bin build/Counter.bin \
  --pkg counter \
  --type Counter \
  --out counter.go
