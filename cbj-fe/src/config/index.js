// boba token
export const STAKED_TOKEN_ADDRESS =
    process.env.NEXT_PUBLIC_STAKED_TOKEN_ADDRESS;

// bre token
export const EARNED_TOKEN_ADDRESS =
    process.env.NEXT_PUBLIC_EARNED_TOKEN_ADDRESS;

// staking address
export const stakingPoolAddresses = [
    {
        chainId: 11155111,
        stakingAddress: "0x6C336a43bC47648Dac96b1419958B8a4e78E05C1",
        depositTokenAddress: "0x4E71E941878CE2afEB1039A0FE16f5eb557571C8",
        earnedTokenAddress: "0x4E71E941878CE2afEB1039A0FE16f5eb557571C8",
    },
    {
        chainId: 31337,
        stakingAddress: "0xD1eeE1b97d4E082252f15Eda10d870b2B7145491", // 填AllocationStakingProxy的地址
        depositTokenAddress: "0x31Afb76742cd07DAc1C92Cd567b045cA63653515", // 填LP-Token的地址
        earnedTokenAddress: "0x3831A9994e23C444b3BAdFc0fc803814b6f88Edf", // 填CBJ-Token的地址
    },
];

export const API_DOMAIN = process.env.NEXT_PUBLIC_SERVER_DOMAIN;

export const VALID_CHAIN_IDS = [
    // Boba Network
    288,
    // Boba Rinkeby test
    28,
    // bsc main network
    56,
    // bsc test network
    97,
    31337,
];

export * from "./valid_chains";

// 0: bre pool 1: boba pool
export const STAKING_POOL_ID = 0;

export const APPROVE_STAKING_AMOUNT_ETHER = 1000000;
 
export const TELEGRAM_BOT_ID = process.env.NEXT_PUBLIC_TG_BOT_ID;

export const BASE_URL = "https://pancakeswap.finance";
export const BASE_BSC_SCAN_URL = "https://bscscan.com";

export const tokenAbi = [
    // Read-Only Functions
    "function deposited(uint256 pid, address to) view returns (uint256)",
    "function balanceOf(address owner) view returns (uint256)",
    "function decimals() view returns (uint8)",
    "function symbol() view returns (string)",
    "function allowance(address owner, address spender) view returns (uint256)",
    "function userInfo(uint pid, address spender) view returns (uint256)",
    "function poolInfo(uint pid) view returns (uint256)",

    // Authenticated Functions
    "function deposit(uint256 pid, uint256 amount) returns (bool)",
    "function withdraw(uint256 pid, uint256 amount) returns (bool)",
    "function approve(address spender, uint256 amount) returns (bool)",
    "function transfer(address to, uint amount) returns (bool)",

    // Events
];

export const tokenImage =
    "http://bobabrewery.oss-ap-southeast-1.aliyuncs.com/brewery_logo.jpg";


export const TOKEN_ADDRESS_MAP = {
    11155111: "0x4E71E941878CE2afEB1039A0FE16f5eb557571C8", // 测试链sepolia
    31337: "0x3831A9994e23C444b3BAdFc0fc803814b6f88Edf", // 本地链 填CBJTOKEN的地址
}

// LP token地址
export const LP_TOKEN_ADDRESS_MAP = {
    11155111: "0x4E71E941878CE2afEB1039A0FE16f5eb557571C8", // 测试链sepolia
    31337: "0x31Afb76742cd07DAc1C92Cd567b045cA63653515", // 本地链 填LPTOKEN的地址
}

export const AIRDROP_ADDRESS_MAP = {
    11155111: {

    },
    31337: {
        'CBJ':"0x16477603B3A9bBa51151B336A7C2057294Ca56B2", //Airdrop-CBJ
        'LP-CBJ':"0x029a8268152296430739a562aD0776fE8BDa1958", //Airdrop-LP-CBJ
    }, 
}

export const tokenSymbols = [
    { chainId: 11155111, symbol: 'CBJ', address: TOKEN_ADDRESS_MAP[11155111] },
    { chainId: 31337, symbol: 'CBJ', address: TOKEN_ADDRESS_MAP[31337] },
]

export const tokenInfos = [
    { chainId: 11155111, symbol: 'CBJ', address: TOKEN_ADDRESS_MAP[11155111]},
    { chainId: 31337, symbol: 'CBJ', address: TOKEN_ADDRESS_MAP[31337]},
]

export const lpTokenInfos = [
    { chainId: 11155111, symbol: 'LP-CBJ', address: LP_TOKEN_ADDRESS_MAP[11155111]},
    { chainId: 31337, symbol: 'LP-CBJ', address: LP_TOKEN_ADDRESS_MAP[31337]},
]
