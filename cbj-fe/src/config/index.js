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
        stakingAddress: "0xA1F663Bd77C53A3c2a4725F12741E85216357BB8", // 填AllocationStakingProxy的地址
        // TODO
        depositTokenAddress: "0x9C645aE1079817ac0F415b39f7800038702C4af2", // 填C2N-Token的地址
        earnedTokenAddress: "0x8F93a6D0C4DD211C737F28CfB74bCdaD2b3804E2", // 填C2N-Token的地址
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
    31337: "0x8F93a6D0C4DD211C737F28CfB74bCdaD2b3804E2", // 本地链 填C2N-TOKEN的地址
}

// LP token地址
export const LP_TOKEN_ADDRESS_MAP = {
    11155111: "0x4E71E941878CE2afEB1039A0FE16f5eb557571C8", // 测试链sepolia
    31337: "0x82550B351160A5FeFb42a900247A0c321be23350", // 本地链 填LPTOKEN的地址
}

export const tokenSymbols = [
    { chainId: 11155111, symbol: 'CBJ', address: TOKEN_ADDRESS_MAP[11155111] },
    { chainId: 31337, symbol: 'CBJ', address: TOKEN_ADDRESS_MAP[31337] },
]

export const tokenInfos = [
    { chainId: 11155111, symbol: 'CBJ', address: TOKEN_ADDRESS_MAP[11155111] },
    { chainId: 31337, symbol: 'CBJ', address: TOKEN_ADDRESS_MAP[31337] },
]

export const lpTokenInfos = [
    { chainId: 11155111, symbol: 'CBJ', address: LP_TOKEN_ADDRESS_MAP[11155111] },
    { chainId: 31337, symbol: 'CBJ', address: LP_TOKEN_ADDRESS_MAP[31337] },
]

export const AIRDROP_CONTRACT = "0x37c0166D96702b3446854448189b7bc4Bb6c475D" // AIRDROP_TOKEN的地址：Airdrop-C2N
