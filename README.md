# CBJ Launchpad

> 🌐 演示地址：[https://cbj-launchpad.cb00j.com/](https://cbj-launchpad.cb00j.com/)

CBJ Launchpad 是一个去中心化的 IDO（Initial DEX Offering）募资平台，让项目方可以发起代币销售，用户可以注册、认购、并按 vesting 计划分批提取代币。系统由智能合约、Go 后端、Next.js 前端三部分组成，并配套了一套基于 6 天周期的自动化部署运维脚本。

---

## 目录

- [系统架构](#系统架构)
- [模块介绍](#模块介绍)
- [数据库表](#数据库表)
- [核心功能](#核心功能)
- [业务流程](#业务流程)
- [技术栈](#技术栈)
- [环境要求](#环境要求)
- [部署指南](#部署指南)

---

## 系统架构

```
┌─────────────────────────────────────────────────────────────┐
│                         用户浏览器                            │
│              (MetaMask 钱包 + HTTPS 前端页面)                 │
└───────────────┬─────────────────────────┬───────────────────┘
                │ HTTPS                     │ HTTPS (RPC/API)
                ▼                           ▼
┌─────────────────────────────────────────────────────────────┐
│                    Nginx (反向代理 + HTTPS)                   │
│   /  → 前端    /api → 后端    /rpc → 区块链节点               │
└───────┬─────────────────┬──────────────────┬─────────────────┘
        ▼                 ▼                  ▼
┌──────────────┐  ┌───────────────┐  ┌──────────────────┐
│  cbj-fe      │  │   cbj-be      │  │   anvil (链)     │
│  Next.js     │  │   Go/Gin      │  │   本地测试链      │
│  :3000       │  │   :8080       │  │   :8545          │
└──────────────┘  └───────┬───────┘  └────────┬─────────┘
                          │                    │
                          │  监听链上事件        │
                          ▼                    │
                  ┌───────────────┐            │
                  │    MySQL      │◀───────────┘
                  │  cbj-launchpad│   (合约部署后同步地址)
                  └───────────────┘
```

数据流向：用户通过浏览器（MetaMask）与前端交互，前端通过 Nginx 反向代理访问后端 API 和区块链 RPC。后端监听链上事件，将销售、注册、认购等数据同步写入 MySQL，供前端展示。

---

## 模块介绍

### 1. 智能合约（cbj-contracts）

基于 Foundry 框架开发的 Solidity 合约，是整个系统的链上核心。

| 合约                | 职责                                            |
| ------------------- | ----------------------------------------------- |
| `CBJToken`          | 平台 ERC20 代币                                 |
| `CBJSale`           | 单场 IDO 销售合约，处理注册、认购、vesting 提取 |
| `SalesFactory`      | 销售工厂，通过它部署每一场 CBJSale              |
| `AllocationStaking` | 质押分配合约，管理用户质押与额度分配            |
| `FarmingCBJ`        | 流动性挖矿（farming）合约，质押 LP 获取奖励     |
| `LP-CBJ`            | 流动性代币                                      |

`CBJSale` 是业务最复杂的合约，关键方法：

- `setSaleParams(...)`：设置销售参数（代币、价格、销售量、销售结束时间、解锁时间、单人上限、vesting 精度），**注意不含 saleStart 参数**
- `setSaleStart(uint256)`：单独设置销售开始时间（要求 `> now` 且 `< saleEnd`）
- `setRegistrationTime(start, end)`：设置注册时间窗口
- `setVestingParams(unlockTimes[], percentPerPortion[], precision)`：设置分批解锁的时间和比例
- `registerForSale(sign, poolId)`：用户携带后端签名进行注册
- `participate(...)`：用户认购代币
- `withdrawMultiplePortions(...)`：按 vesting 计划分批提取已购代币

### 2. 后端（cbj-be）

基于 Go + Gin + GORM 的服务端，连接 go-ethereum 监听链上事件。

主要职责：

- **REST API**：向前端提供产品信息、注册签名、用户认购记录等接口
- **事件监听**：通过 WebSocket 连接链节点，监听 `UserRegistered`、`TokensSold` 等事件，实时把链上数据同步到数据库
- **签名服务**：为用户注册生成后端签名（防止任意地址注册）

事件监听采用"折中"架构：一个共享的 `listener`（负责连接、订阅、补漏、重连、分发），加上多个 `EventHandler`（如 `RegisterHandler`、`ParticipateHandler`）分别处理各类事件。

### 3. 前端（cbj-fe）

基于 Next.js + React + ethers.js v5 的 Web 应用。

主要页面：

- **首页 / 项目列表**：展示所有 IDO 项目及状态
- **项目详情**：展示单个 IDO 的时间线（注册/销售/解锁）、价格、进度，提供注册、认购、提取入口
- **Farm**：流动性挖矿页面
- **Staking**：质押页面

前端通过 MetaMask 与链交互，通过后端 API 获取展示数据。时间统一以 **UTC 为基准**：后端返回毫秒时间戳，前端用 `new Date(timestamp)` 按用户本地时区显示，实现全球化。

---

## 数据库表

数据库 `cbj-launchpad` 共 4 张表（建表语句见 `cbj-be/conf/init.sql`）：

| 表名                    | 用途                                                                                                                                                                                                                                                                  |
| ----------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `product_contract`      | **产品/销售主表**。存每个 IDO 项目的核心信息：合约地址（销售合约、代币地址）、各阶段时间（注册开始/结束、销售开始/结束、解锁时间、TGE）、销售状态、代币价格、销售总量、已售数量、已募集金额、vesting 分批解锁时间、链 ID 等。是前端展示项目详情和时间线的主要数据源。 |
| `product_registration`  | **用户注册记录表**。存用户注册参与资格的记录（账户地址、产品 ID、交易信息等）。由后端监听链上 `UserRegistered` 事件写入。是用户认购前的资格门槛记录。                                                                                                                 |
| `product_participation` | **用户认购记录表**。存用户在销售期认购代币的记录（账户地址、产品 ID、认购数量、支付金额、交易哈希、区块号、认购时间等）。由后端监听链上 `TokensSold` 事件写入。                                                                                                       |
| `sync_state`            | **事件同步状态表**。记录事件监听器的同步进度（已处理到的区块号）。后端断线重连或重启后，从该区块继续补漏（pollMissed），避免漏处理或重复处理事件。                                                                                                                    |

> 数据同步方向：链上事件（注册、认购）由后端实时监听并写入对应表；`product_contract` 的合约地址在每次重新部署后由运维脚本同步更新。

---

## 核心功能

- **IDO 项目展示**：项目列表、详情、时间线、销售进度、ROI 等
- **用户注册**：社交认证 + 后端签名 + 链上 `registerForSale`
- **代币认购**：在销售期内按设定价格认购代币，受单人上限约束
- **Vesting 分批提取**：销售结束后，按 vesting 计划分批解锁、提取已购代币
- **流动性挖矿**：质押 LP 代币获取 CBJ 奖励
- **链上事件同步**：后端自动监听并同步链上数据到数据库
- **全球化时间**：UTC 基准存储与传输，前端按用户本地时区展示

---

## 业务流程

一个完整的 IDO 周期分为三个阶段：

```
   注册期            销售期            提现期
 ┌────────┐       ┌────────┐       ┌──────────────────┐
 │ 第 1 天 │  →   │ 第 2 天 │  →   │  第 3 ~ 6 天       │
 │ 用户注册 │       │ 用户认购 │       │  分批提取代币       │
 │status=1│       │status=3│       │     status=4      │
 └────────┘       └────────┘       └──────────────────┘
```

1. **注册期**：用户完成社交认证后，向后端获取签名，调用合约 `registerForSale` 注册参与资格
2. **销售期**：注册成功的用户可在销售时间窗口内认购代币
3. **提现期**：销售结束后，代币按 vesting 计划分批解锁，用户调用 `withdrawMultiplePortions` 分批提取

---

## 技术栈

| 层       | 技术                                             |
| -------- | ------------------------------------------------ |
| 智能合约 | Solidity ^0.8.23, Foundry (forge / anvil / cast) |
| 后端     | Go, Gin, GORM, go-ethereum, MySQL                |
| 前端     | Next.js, React, ethers.js v5, antd, Redux        |
| 部署     | Nginx, Bash 脚本, cron                           |
| 链       | anvil 本地链（chainId 31337）                    |

---

## 环境要求

服务器需安装：

- **Foundry**（anvil / forge / cast）
- **Go**（建议 1.20+）
- **Node.js**（建议 18+，含 npm）
- **MySQL**
- **Nginx**（HTTPS + 反向代理）
- **辅助工具**：`jq`、`lsof`、`curl`、`mysql` 客户端

---

## 部署指南

### 1. 准备数据库

创建数据库并导入表结构：

```bash
mysql -u root -p -e "CREATE DATABASE \`cbj-launchpad\` DEFAULT CHARSET utf8mb4;"
mysql -u root -p cbj-launchpad < cbj-be/conf/init.sql
```

确保 `product_contract` 表中有对应的产品记录（本系统默认产品 id = 4）。

### 2. 部署智能合约

```bash
cd cbj-contracts

# 启动本地链
anvil --host 0.0.0.0 --port 8545 --block-time 2 --chain-id 31337

# 另开终端，清理旧状态并部署
rm -rf broadcast/ cache/ deployments/
make deploy-local
```

部署完成后，合约地址会写入 `deployments/contract_addresses.json`。**每次重新部署 CBJSale 地址都会变，必须把新地址同步到数据库**（自动化脚本会处理）。

### 3. 启动后端

```bash
cd cbj-be
go run .
```

后端配置（数据库连接、监听的合约地址、链 RPC 等）在配置文件中设置。数据库连接的时区基准为 UTC。

### 4. 启动前端

前端有多套环境文件：

- `.env.development`：本地开发（默认连 localhost）
- `.env.production`：生产/远程部署（连服务器地址，通过 HTTPS 代理）

**生产部署使用生产模式**（读 `.env.production`）：

```bash
cd cbj-fe
NODE_ENV=production npx next build           # 构建（读 .env.production）
NODE_ENV=production npx next start -p 3000    # 启动
```

> 注意：`npm run dev` 读 `.env.development`（localhost），不适合远程部署。远程部署务必用生产模式读 `.env.production`。
