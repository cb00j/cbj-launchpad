# .PHONY 告诉 make 这是命令，不是文件
.PHONY: help deploy-local deploy-sepolia deploy-all test clean

# 默认目标（运行 make 不带参数时执行）
help:
	@echo "Available commands:"
	@echo "  make deploy-local    - Deploy to local Anvil"
	@echo "  make deploy-sepolia  - Deploy to Sepolia testnet"
	@echo "  make test            - Run tests"
	@echo "  make clean           - Clean build artifacts"

# 变量定义
NETWORK ?= local
PRIVATE_KEY := $(shell grep PRIVATE_KEY .env | cut -d '=' -f2)

# 本地部署
deploy-local:
	@echo "Deploying to local..."
	@forge script script/AirDropScript.s.sol --rpc-url local --broadcast
	@forge script script/LPTokenScirpt.s.sol --rpc-url local --broadcast
	@forge script script/FarmingCBJScript.s.sol --rpc-url local --broadcast
	@echo "Local deployment complete"

# 测试网部署
deploy-sepolia:
	@echo "Deploying to Sepolia..."
	@forge script script/AirDropScript.s.sol \
		--rpc-url sepolia \
		--broadcast \
		--verify
	@forge script script/LPTokenScirpt.s.sol \
		--rpc-url sepolia \
		--broadcast \
		--verify
	@forge script script/FarmingCBJScript.s.sol \
		--rpc-url sepolia \
		--broadcast \
		--verify
	@echo "Sepolia deployment complete"

# 一键部署所有合约
deploy-all: deploy-lp
	@echo "All contracts deployed"

deploy-lp:
	@forge script script/LPTokenScirpt.s.sol --rpc-url $(NETWORK) --broadcast


# 测试
test:
	@forge test -vv

# 清理
clean:
	@forge clean
	@rm -rf deployments/*.json