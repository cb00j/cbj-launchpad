import { useMemo } from 'react';
import { useAppSelector } from "@src/redux/hooks";
import { Contract, providers } from "ethers";
import AirdropAbi from '@src/util/abi/Airdrop.json'

export const useAirdropContract = (address) => {
  const chain = useAppSelector(state => state.wallet.chain);
  const signer = useAppSelector(state => state.contract.signer);
  const walletAddress = useAppSelector(state => state.contract.walletAddress);

  // 使用 useMemo 缓存合约实例，只有当 address 或 signer 变化时才重新实例化
  return useMemo(() => {
    // 核心修复 1：如果没传地址，或者取到了 undefined，直接返回 null 拦截掉
    if (!address) {
      return null;
    }

    // 核心修复 2：优雅降级。如果有 signer 就用 signer（可读写），没有就用 provider（只读）
    let providerOrSigner = signer;
    
    if (!signer || !walletAddress) {
      // 注意：确保 chain?.rpc[0] 在没连钱包时也有默认值，否则这里可能拿不到 RPC
      if (chain?.rpc?.[0]) {
        providerOrSigner = new providers.JsonRpcProvider(chain.rpc[0]);
      } else {
        // 如果连默认 RPC 都没有，也只能放弃实例化
        return null; 
      }
    }

    try {
      const actualAbi = (AirdropAbi as any).default || (AirdropAbi as any).abi || AirdropAbi;
      return new Contract(address, actualAbi, providerOrSigner);
    } catch (error) {
      console.error("实例化空投合约失败:", error);
      return null;
    }
  }, [address, signer, walletAddress, chain]);
}