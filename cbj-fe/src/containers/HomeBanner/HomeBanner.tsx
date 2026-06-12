import { useMemo,useState } from 'react';
import { useRouter } from 'next/router';
import { Layout, Row, Col, Menu, message } from 'antd';
import { CopyOutlined } from '@ant-design/icons'

import { useResponsive } from '@src/hooks/useResponsive';
import { useWallet } from '@src/hooks/useWallet';
import { useMessage } from '@src/hooks/useMessage';

import styles from './HomeBanner.module.scss'
import { IconCBJ,IconLpCBJ } from '@src/components/icons';
import Image from 'next/image';
import { useAirdropContract } from '@src/hooks/useContract';
import { tokenInfos,lpTokenInfos, AIRDROP_ADDRESS_MAP } from '@src/config';

export default function Header() {
  const {
    isDesktopOrLaptop,
    isTabletOrMobile,
  } = useResponsive();

  const {
    chain,
    walletAddress
  } = useWallet();
  
  const {
    setSuccessMessage
  } = useMessage();

  const token = useMemo(() => {
    return tokenInfos.find(item => item.chainId == chain?.chainId) || tokenInfos[1];
  }, [chain]);

  const lptoken = useMemo(() => {
    return lpTokenInfos.find(item => item.chainId == chain?.chainId) || lpTokenInfos[1];
  }, [chain]);

  const [claimingSymbol, setClaimingSymbol] = useState(null);

  // 1. 获取当前链下对应的空投合约地址 (使用可选链 ?. 防止未配置的网络报错)
  const cbjAirdropAddress = AIRDROP_ADDRESS_MAP[chain?.chainId]?.[token.symbol];
  const lpAirdropAddress = AIRDROP_ADDRESS_MAP[chain?.chainId]?.[lptoken.symbol];

  // 2. 顶层调用 Hook，分别获取两个不同的合约实例
  const cbjAirdropContract = useAirdropContract(cbjAirdropAddress);
  const lpAirdropContract = useAirdropContract(lpAirdropAddress);

  function copy(text) {
    navigator.clipboard.writeText(text).then(function () {
      setSuccessMessage('Copied')
    }, function (err) {
    });
  }
  async function addToken(tokenAddress, symbolName) {
    if (!chain) {
      message.error('connect wallet and try again !')
      return
    }
    if (chain.chainId !== token.chainId) {
      message.error('switch network and try again !')
      return
    }
    await window.ethereum && window.ethereum.request({
      method: "wallet_watchAsset",
      params: {
        type: "ERC20", // Initially only supports ERC20, but eventually more!
        options: {
          address: tokenAddress, // The address that the token is at.
          symbol: symbolName, // A ticker symbol or shorthand, up to 5 chars.
          decimals: 18, // The number of decimals in the token
          image: '',
        }
      }
    });
  }
  const handleClaim = async (contractInstance, symbol) => {
    if (!walletAddress) {
      message.warn('Please connect wallet first');
      return;
    }
    if (!contractInstance) {
      message.error(`Airdrop contract for ${symbol} is not available on this network`);
      return;
    }

    // 防抖：如果当前有任何 Claim 正在进行中，直接 return 阻止执行
    if (claimingSymbol) return; 

    try {
      // 开启 Loading 状态
      setClaimingSymbol(symbol);
      
      // 1. 唤起 MetaMask，等待用户确认并发送交易
      const tx = await contractInstance.claim();
      message.loading({ content: `Claiming ${symbol}, waiting for confirmation...`, key: 'claim' });
      
      // 2. 【核心修复】等待区块链打包确认这笔交易
      const receipt = await tx.wait(); 
      
      // 3. 只有真正打包成功了，才提示成功
      if (receipt.status === 1) {
        message.success({ content: `Successfully claimed ${symbol}!`, key: 'claim', duration: 3 });
      } else {
        message.error({ content: `Transaction failed on chain`, key: 'claim' });
      }

    } catch (error) {
      console.error(error);
      // 如果用户在 MetaMask 里点了取消，这里会捕获到错误
      message.error({ content: error.reason || error?.data?.message || error?.message || 'Claim failed or canceled', key: 'claim' });
    } finally {
      // 无论成功还是失败，最后都要解除 Loading 状态
      setClaimingSymbol(null);
    }
  }

  return <div className={styles['home-banner']}>
    <Row justify="space-between" align="middle" className={styles['main']}>
      <Col span={isDesktopOrLaptop ? 15 : 24}>
        <Row gutter={16}>
          <Col span={isDesktopOrLaptop ? 4 : 24}>
            <Row justify="center" align="middle">
              <IconCBJ className={styles.icon} />
            </Row>
          </Col>
          <Col span={isDesktopOrLaptop ? 20 : 24}>
            <Row>
              <Col span={24} className={styles['text1']}>
                {token.symbol} Tokens Online Now!
              </Col>
              <Col className={styles['text2']}>
                Contract Address: &nbsp;
                {
                  isDesktopOrLaptop ? <></> : <br />
                }
                {token.address}
                &nbsp;
                {/* <CopyOutlined className={styles['copy']} onClick={()=>{copy(token.address)}}></CopyOutlined> */}
              </Col>
            </Row>
          </Col>
        </Row>
      </Col>
      <Col span={isDesktopOrLaptop ? 4 : 12}>
        <div 
          className={styles['button']}
          // 如果正在 claim 这个代币，改变样式表现为禁用
          style={{ opacity: claimingSymbol === token.symbol ? 0.5 : 1, cursor: claimingSymbol ? 'not-allowed' : 'pointer' }}
          onClick={() => handleClaim(cbjAirdropContract, token.symbol)}
        >
          {claimingSymbol === token.symbol ? 'Claiming...' : `Claim ${token.symbol}`}
        </div>
      </Col>
      <Col span={isDesktopOrLaptop ? 4 : 12}>
        <div className={styles['button']}
          onClick={() => {
            addToken(token.address, token.symbol)
          }}
        >
          Add {token.symbol} to Wallet
        </div>
      </Col>
    </Row>
    <Row justify="space-between" align="middle" className={styles['main']}>
      <Col span={isDesktopOrLaptop ? 15 : 24}>
        <Row gutter={16}>
          <Col span={isDesktopOrLaptop ? 4 : 24}>
            <Row justify="center" align="middle">
              <IconLpCBJ className={styles.icon} />
            </Row>
          </Col>
          <Col span={isDesktopOrLaptop ? 20 : 24}>
            <Row>
              <Col span={24} className={styles['text1']}>
                {lptoken.symbol} Tokens Online Now!
              </Col>
              <Col className={styles['text2']}>
                Contract Address: &nbsp;
                {
                  isDesktopOrLaptop ? <></> : <br />
                }
                {lptoken.address}
                &nbsp;
                {/* <CopyOutlined className={styles['copy']} onClick={()=>{copy(token.address)}}></CopyOutlined> */}
              </Col>
            </Row>
          </Col>
        </Row>
      </Col>
      <Col span={isDesktopOrLaptop ? 4 : 12}>
        <div 
        className={styles['button']}
        style={{ opacity: claimingSymbol === lptoken.symbol ? 0.5 : 1, cursor: claimingSymbol ? 'not-allowed' : 'pointer' }}
        onClick={() => handleClaim(lpAirdropContract, lptoken.symbol)}
      >
        {claimingSymbol === lptoken.symbol ? 'Claiming...' : `Claim ${lptoken.symbol}`}
      </div>
      </Col>
      <Col span={isDesktopOrLaptop ? 4 : 12}>
        <div className={styles['button']}
          onClick={() => {
            addToken(lptoken.address, lptoken.symbol)
          }}
        >
          Add {lptoken.symbol} to Wallet
        </div>
      </Col>
    </Row>
  </div>
}
