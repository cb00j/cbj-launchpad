import type { AppProps } from 'next/app'
import Router from 'next/router'
import { Row, Col, Statistic, message, Timeline, Modal, Input, Spin } from 'antd';
import { QuestionCircleOutlined, InfoCircleOutlined, CopyOutlined } from '@ant-design/icons';
const { Countdown } = Statistic;

import {
  ProjectData,
  RegistrationData,
  SaleData,
} from '@src/types'
import {
  STAKING_POOL_ID, tokenAbi
} from '@src/config'
import styles from './project.module.scss'
import TransactionButton from "@src/components/elements/TransactionButton";
import PoolCard from '@src/components/elements/PoolCard'
import ParticipateModal from '@src/containers/ParticipateModal/ParticipateModal';
import InviteModal from '@src/containers/InviteModal/InviteModal';
import { useEffect, useState, useMemo } from 'react'
import { formatDate, hexToBytes, parseEther, formatEther, parseWei, seperateNumWithComma } from "@src/util/index"
import axios from '@src/api/axios'
import { useWallet } from '@src/hooks/useWallet'
import { usePageLoading } from '@src/hooks/usePageLoading'
import { useErrorHandler } from '@src/hooks/useErrorHandler'
import { useResponsive } from '@src/hooks/useResponsive';

import {
  BigNumber, Contract, ethers
} from 'ethers'

import { useMessage } from '../hooks/useMessage'
import { useLogin } from '../hooks/useLoginModal'
import { useStake } from '../hooks/useStake'
import { useThirdParty } from '../hooks/useThirdParty';
import AppPopover from '@src/components/elements/AppPopover';
import { WarningOutlined, CheckCircleTwoTone } from '@ant-design/icons';
import { parseUnits } from 'ethers/lib/utils';

type OtherPoolInfoProps =
  {
    registration: RegistrationData,
    sale: SaleData,
  };

/**
 * Pool detail page
 */
export default function Pool({ Component, pageProps }: AppProps) {
  let search, params, pId;
  if (typeof window !== 'undefined') {
    search = window.location.search;
    params = new URLSearchParams(search);
    pId = params.get('id');
  }

  const {
    walletAddress,
    addToken,
    saleContract,
    saleAddress,
    setSaleAddress,
    setLoading,
    signer,
    chain,
  } = useWallet();

  const {
    setSuccessMessage,
    setErrorMessage
  } = useMessage();

  const {
    LoginModal,
    isUserRegister,
    showLoginModal,
    setLoginConfig,
    setLoginState,
  } = useLogin({ pId: pId });

  const {
    setDepositTokenAddress,
    depositedAmount,
    setStakingAddress,
    approve,
    stakingContract,
    depositDecimals,
    depositSymbol,
    setAllowanceAddress,
    allowance,
    globalPoolStakingAddress,
  } = useStake();

  useEffect(() => {
    setStakingAddress(globalPoolStakingAddress);
  }, [globalPoolStakingAddress])

  const {
    PageLoader,
    setPageLoading,
  } = usePageLoading();

  const {
    getErrorMessage,
  } = useErrorHandler();

  const {
    isDesktopOrLaptop,
    isTabletOrMobile,
  } = useResponsive();

  const [projectInfo, setProjectInfo] = useState<ProjectData>({} as ProjectData);
  const [tab, setTab] = useState('sale_info');
  // TODO: fix pool Id
  const [poolId, setPoolId] = useState<number>(STAKING_POOL_ID);
  const [statesReady, setStatesReady] = useState<boolean>(false);
  const [isRegistered, setIsRegistered] = useState<boolean>(false);
  const [canRegister, setCanRegister] = useState<boolean>(true);
  const [isParticipated, setIsParticipated] = useState<boolean>(false);
  const [participateAmount, setParticipateAmount] = useState();
  const [allocationTop, setAllocationTop] = useState('1000000000000000000000');
  const [participateInfo, setParticipateInfo] = useState<any>();
  const [onchainSale, setOnchainSale] = useState<any>(null);

  const [otherPoolInfo, setOtherPoolInfo] = useState<OtherPoolInfoProps>();
  const [mileStones, setMileStones] = useState<any>({})

  const [paymentTokenAddress, setPaymentTokenAddress] = useState<string>('');
  const [paymentTokenDecimal, setPaymentTokenDecimal] = useState<number>(18);

  // status: -1:not ready, 0: not started, 1: in registration, 
  // 2: after registration and before participation,
  // 3: in sale,
  // 4: sale ended,
  const [status, setStatus] = useState(-1);

  const [showParticipateModal, setShowParticipateModal] = useState(false);

  const [referralCode, setReferralCode] = useState(null);
  const [referralModalVisible, setReferralModalVisible] = useState(false);

  const { ethToUsd, getEthToUsd } = useThirdParty();
  const unImplementedInterfaceLog = (text) => {
    console.log('未实现的接口：', text);
  }

  const [chainNow, setChainNow] = useState(0);

  useEffect(() => {
    if (!saleContract) return;
    const fetchChainTime = () => {
      saleContract.provider.getBlock('latest')
        .then((block) => {
          setChainNow(block.timestamp * 1000);  // 链上时间是秒,转毫秒
        })
        .catch((e) => console.error('get chain time failed', e));
    };
    fetchChainTime();
    // 每 5 秒刷新一次链上时间(本地推进时间后能及时反映)
    const timer = setInterval(fetchChainTime, 5000);
    return () => clearInterval(timer);
  }, [saleContract]);


  useEffect(() => {
    setPageLoading(true);

    // get projects info
    axios.get('/api/product/base_info', { params: { productId: pId } })
      .then((response) => {
        let data = response.data || {};
        data.allocationTop = 1;
        data.chainId = chain.chainId;
        setProjectInfo(data);
        setLoginConfig({
          tweetId: data.tweetId,
          tgGroupLink: data.medias,
          projectName: data.name,
        })
      })
      .catch((e) => {
        setErrorMessage('Network error, please check your network and refresh.')
      })
      .finally(() => {
        setPageLoading(false);
      })

    refreshStates();

    getEthToUsd();

    return (() => {
    })
  }, []);

  useEffect(() => {
    if (!walletAddress) {
      return;
    }
    // // check whether has registered other projects
    // axios.get('/boba/register/can_register', {
    //   params: { accountId: walletAddress }
    // })
    //   .then((response) => {
    //     setCanRegister(true);
    //   })
    //   .catch((e) => {
    //     setCanRegister(false);
    //   })

    const referralCode = localStorage?.getItem('referral');
    if (referralCode) {
      bindReferral(referralCode);
    }
  }, [walletAddress])

  useEffect(() => {
    refreshStates();
  }, [saleContract, walletAddress])

  useEffect(() => {
    if (!saleContract) return;
    saleContract.getSaleInfo()
      .then(([totalTokensSold, totalETHRaised, amountOfTokensToSell,tokenPriceInETH]) => {
        setOnchainSale({ totalTokensSold, totalETHRaised, amountOfTokensToSell,tokenPriceInETH });
      })
      .catch(e => console.error('getSaleInfo failed', e));
  }, [saleContract]);

  useEffect(() => {
    if (isRegistered && depositedAmount) {
      getParticipateAmount();
    }
  }, [isRegistered, depositedAmount]);

  const amountBought: BigNumber = useMemo(() => {
    return participateInfo && BigNumber.from(participateInfo[0] as string) || BigNumber.from(0);
  }, [participateInfo])

  const amountETHPaid: BigNumber = useMemo(() => {
    return participateInfo && BigNumber.from(participateInfo[1] as string) || BigNumber.from(0);
  }, [participateInfo])

  const timeParticipated: number = useMemo(() => {
    return participateInfo && parseWei(participateInfo[2]) || 0;
  }, [participateInfo])

  const isPortionWithdrawn: Array<boolean> = useMemo(() => {
    return participateInfo && participateInfo[3] || 0;
  }, [participateInfo])

  const vestingInfo = useMemo(() => {
    if (projectInfo && projectInfo.vestingPercentPerPortion && projectInfo.vestingPortionsUnlockTime) {
      return [projectInfo.vestingPercentPerPortion, projectInfo.vestingPortionsUnlockTime];
    } else {
      return [[], []];
    }
  }, [projectInfo]);

  const vestingPercentPerPortion: Array<number> = useMemo(() => {
    return vestingInfo && vestingInfo[0] || [];
  }, [vestingInfo]);

  const vestingPortionsUnlockTime: Array<number> = useMemo(() => {
    return vestingInfo && vestingInfo[1] || [];
  }, [vestingInfo]);

  const canWithdrawArr: Array<boolean> = useMemo(() => {
    // 没有 vesting 时间或没有提取记录时的兜底
    if (!isPortionWithdrawn || !Array.isArray(isPortionWithdrawn)) {
      return [true];
    }

    // 有解锁时间数组:判断每个 portion 是否「已解锁 且 未提取」
    if (vestingPortionsUnlockTime && vestingPortionsUnlockTime.length > 0) {
      return isPortionWithdrawn.map((withdrawn, index) => {
        // 该 portion 的解锁时间(链上是秒,转毫秒)
        const unlockTimeMs = Number(vestingPortionsUnlockTime[index] || 0) * 1000;
        // 可提条件:未提取(!withdrawn) 且 已到解锁时间(用链上时间比较)
        return !withdrawn && unlockTimeMs <= chainNow;
      });
    }

    // 没有解锁时间数组,只按是否已提取判断
    return isPortionWithdrawn.map((withdrawn) => !withdrawn);
  }, [isPortionWithdrawn, vestingPortionsUnlockTime, chainNow]);

  useEffect(() => {
    const bigNumSec2Milsec = (bigSec) => {
      return (bigSec && (BigNumber.from(bigSec).toString())) || '';
    }
    if (projectInfo) {
      // init sale contract
      setSaleAddress(projectInfo.saleContractAddress);

      const mileStones = {
        registrationTimeStarts: bigNumSec2Milsec(projectInfo.registrationTimeStarts),
        registrationTimeEnds: bigNumSec2Milsec(projectInfo.registrationTimeEnds),
        saleStart: bigNumSec2Milsec(projectInfo.saleStart),
        saleEnd: bigNumSec2Milsec(projectInfo.saleEnd),
        unlock: bigNumSec2Milsec(projectInfo.unlockTime),
      }
      setMileStones(mileStones);
      setStatus(projectInfo.status);
      // setPaymentTokenAddress(projectInfo.paymentToken);
      setPaymentTokenAddress('0x1Dd5dcB05E451EfC6d6D4Fb2B905b02Dc3679aB4');
    }
  }, [projectInfo, otherPoolInfo]);

  useEffect(() => {
    const timer = setInterval(() => {
      judgePoolStatus();
    }, 1000)
    return () => {
      clearInterval(timer);
    }
  }, [mileStones])

  /**
   * payment token contract
   */
  useEffect(() => {
    setDepositTokenAddress(paymentTokenAddress);
  }, [paymentTokenAddress])

  useEffect(() => {
    setAllowanceAddress(saleAddress);
  }, [saleAddress])

  function judgePoolStatus() {
    const now = Date.now();
    if (!mileStones) {
      return;
    }

    if (now < mileStones.registrationTimeStarts) {
      // not started
      setStatus(0);
    } else if (now < mileStones.registrationTimeEnds) {
      // in registration
      setStatus(1);
    } else if (now < mileStones.saleStart) {
      // before sale
      setStatus(2);
    } else if (now < mileStones.saleEnd) {
      // in sale
      setStatus(3);
    } else if (now < mileStones.unlock) {
      // sale ends
      setStatus(4);
    } else if (now >= mileStones.unlock) {
      // sale ends and unlock(means can withdraw token)
      setStatus(5);
    }
  }

  function refreshStates() {
    if (!saleContract || !walletAddress) {
      return;
    }
    let option = { gasLimit: 100000 }
    let promiseA = saleContract.isRegistered(walletAddress, option)
      .then((data) => {
        setIsRegistered(data);
      })
      .catch(e => {
        console.error(e);
      });
    let promiseB = saleContract.isParticipated(walletAddress, option)
      .then((data) => {
        setIsParticipated(data);
        if (data) {
          getParticipateInfo();
        }
      })
      .catch(e => {
        console.error(e);
      });
    Promise.all([promiseA, promiseB])
      .then(() => {
        setStatesReady(true);
      })
  }

  function refreshOnchainSale() {
    if (!saleContract) return;
    saleContract.sale()
      .then(setOnchainSale)
      .catch(e => console.error(e));
  }

  /**
   * 
   * @returns Promise<string> - registration sign
   */
  function getRegistrationSign() {
    const f = new FormData();
    f.append('userAddress', walletAddress || '');// 钱包地址
    f.append('contractAddress', saleAddress);// 这个 CBJSale 合约地址

    return axios.post('/api/encode/sign_registration', f)
      .then((response) => {
         // 返回一段签名 hex字符串，这个签名是后端用一个特定的私钥对钱包地址和合约地址做的签名，合约会验证这个签名，目的是证明这个用户确实有权限注册这个合约的这个项目
        let data = response.data;
        return data;
      })
      .catch(e => {
        // handle request error
        console.error(e);
        setErrorMessage('Register fail');
      })
  }
  /**
   * 
   * @returns Promise<string> - participate sign
   */
  function getParticipateSign() {
    const NUMBER_1E18 = "1000000000000000000";

    const f = new FormData();
    f.append('userAddress', walletAddress || '');
    f.append('contractAddress', saleAddress);
    // FIXME: get Participate amount
    f.append('amount', allocationTop);
    return axios.post('/api/encode/sign_participation', f)
      .then((response) => {
        let data = response.data;
        return data;
      })
      .catch(e => {
        // handle request error
        console.error(e);
        setErrorMessage('Purchase fail');
      })
  }
  /**
   * bind referral codes
   * @param referralcode 
   * @returns 
   */
  function bindReferral(referralcode) {
    const f = new FormData();
    f.append('referralCode', referralcode);
    f.append('participant', walletAddress);
    return axios.post('/boba/referral/bind', f)
  }

  /**
   * Register 
   * @returns Promise
   */
  function registerForSale() {
    // 第 1 步:检查用户是否完成了社交登录/认证(推特、TG 等)
    if (!isUserRegister) {
      showLoginModal();// 没认证 → 弹登录框,流程中止
      return;
    }
    // 第 2 步:合约实例必须就绪
    if (!saleContract) {
      return Promise.reject();
    }
    // 第 3 步:向后端要注册签名
    return getRegistrationSign()
      .then((registrationSign) => {
        console.log(registrationSign, 'sign')
        // 把 hex 签名转成字节
        //const signBuffer = hexToBytes(registrationSign);
        // 第 4 步:带着签名调合约的 registerForSale
        return saleContract.registerForSale(registrationSign, poolId)
          .then(transaction => {
            // 等上链确认
            return transaction.wait();
          })
          .then((receipt) => {
            const txHash = receipt.transactionHash;
            setSuccessMessage('Register success');
            // 刷新链上状态(isRegistered 等)
            refreshStates();
            return txHash;
          })
          .then((txHash) => {
            // on register success
            const sendRegisterSuccess = async () => {
              const referralCode = localStorage?.getItem('referral');
              const f = new FormData();
              f.append('accountId', walletAddress);
              f.append('productId', pId);
              f.append('referralCode', referralCode)
              f.append('txHash', txHash)
              return axios.post('/api/product/register/user_register', f).then(() => {
                if (typeof window !== 'undefined') {
                  window.localStorage.removeItem('referral');
                }
              })
            }
            let retryCount = 0;
            const MAX_RETRY = 3; 
            function loop() {
              setTimeout(() => {
                sendRegisterSuccess()
                  .catch((e) => {
                    console.error('sendRegisterSuccess failed:', e);
                    retryCount++;
                    if (retryCount < MAX_RETRY) {
                      loop(); // retry
                    } else {
                      console.error('register sync failed after max retries');
                    }
                  })
              }, 3000)
            }
             loop();
          })
          .then(() => {
            function loop() {
              setTimeout(() => {
                addRegisterAmount()
                  .then(() => {
                    getParticipateAmount();
                  })
                  .catch(() => {
                    loop();
                  })
              }, 3000)
            }
            loop();
          })
      })
      .catch(e => {
        setErrorMessage('Participate failed. ' + (getErrorMessage(e) || ''));
        console.error(e);
      })
  }

  function addRegisterAmount(amount?, productId?, accountId?) {
    unImplementedInterfaceLog('/boba/amount/register/add');
    return Promise.resolve();
    // const f = new FormData();
    // Object.entries({ stakeAmount: amount || depositedAmount || 0, productId: productId || pId, accountId: accountId || walletAddress })
    //   .forEach(([key, value]) => {
    //     f.append(key, value as string);
    //   })
    // return axios.post('/boba/amount/register/add', f)
    //   .then((response) => {
    //   })
    //   .catch((e) => {
    //     console.error(e);
    //   })
  }

  function getParticipateAmount() {
    let search = window.location.search;
    let params = new URLSearchParams(search);
    let pId = params.get('id');
    const f = new FormData();
    Object.entries({ accountId: walletAddress, productId: pId })
      .forEach(([key, value]) => {
        f.append(key, value as string);
      })
    return axios.post('/api/product/allocation/calc', f)
      .then((response: any) => {
        setAllocationTop(String(response.amount));
        return response;
      })
  }
  /**
   * Participate project
   * @returns Promise
   */
  async function participate(value) {
    if (!saleContract) {
      return Promise.reject();
    }
    setLoading(true);
    // const decimals = await depositTokenContract.decimals();
    // const paymentAmount = BigNumber.from(projectInfo.tokenPriceInPT).mul(~~value.value).div(Math.pow(10, 18-decimals));
    const paymentAmount = BigNumber.from(projectInfo.tokenPriceInPT).mul(~~value.value);
    /**
     * 购买代币
     * 设置足够的代币授权，后面在withdraw时，会自动扣除相应的代币
     */
    return approve(saleAddress, Number(ethers.utils.formatUnits(paymentAmount, depositDecimals)), depositDecimals)
      .then(getParticipateSign)
      .then(async (participateSign) => {
        //const signBuffer = hexToBytes(participateSign);
        // get participation value
        const options = {
          value: paymentAmount
        };
        return saleContract.participate(
          participateSign,
          BigNumber.from(allocationTop + ''),
          options
        )
          .then((transaction) => {
            return transaction.wait();
          })
          .then((transaction) => {
            setSuccessMessage('Purchase success');
            refreshStates();
            refreshOnchainSale();
            setShowParticipateModal(false);
            return Promise.resolve();
          })
      })
      .catch((e) => {
        console.error(e);
        // if(e && e.code === -32603) {
        //   setErrorMessage('Purchase fail, it seems you don\'t have enough ETH tokens to pay, please check your balance.')
        //   return;
        // }
        let msg = getErrorMessage(e);
        setErrorMessage('Purchase fail. ' + (msg || ''));
        return;
      })
      .finally(() => {
        setLoading(false);
      })
  }

  function withdrawTokens() {
    const paymentAmount = BigNumber.from(projectInfo.tokenPriceInPT).mul(~~'30');
    if (!saleContract) {
      return Promise.reject();
    }
    // set protionId
    let portionIds = canWithdrawArr.map((canWithdraw, index) => {
      return canWithdraw ? index : -1;
    })
      .filter(v => v > -1);
    return saleContract.withdrawMultiplePortions(portionIds)
      .then(transaction => {
        return transaction.wait();
      })
      .then((transaction) => {
        setSuccessMessage('withdraw success');
        refreshStates();
        return Promise.resolve();
      })
      .catch((e) => {
        let msg = getErrorMessage(e);
        setErrorMessage('Withdraw fail. ' + (msg || ''));
      })
  }

  function getParticipateInfo() {
    if (!saleContract) {
      return Promise.reject;
    }
    //let option = { gasLimit: 100000 }
    return saleContract.getParticipation(walletAddress)
      .then(data => {
        return setParticipateInfo(data);
      })
      .catch((e) => {
        setErrorMessage('Get participation info fail, please check your network and refresh!');
        console.error(e);
      })
  }
  const SaleInfo = (
    <div className={styles['sale-info'] || ''}>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> Project Website </Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>
          <a href={projectInfo.projectWebsite} target="_blank" rel="noreferrer">{projectInfo.projectWebsite}</a>
        </Col>
      </Row>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> Number of Registrations </Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>{projectInfo.numberOfRegistrants}</Col>
      </Row>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> Vesting </Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>{projectInfo.vesting}</Col>
      </Row>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> TGE </Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>{formatDate(projectInfo.tge, 'HH:mm, Month DD, YYYY')}</Col>
      </Row>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> Sale Contract Address </Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>
          {projectInfo?.saleContractAddress || ''}
        </Col>
      </Row>
    </div>
  )
  const TokenInfo = (
    <div className={styles['token-info'] || ''}>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> Token Name </Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>{projectInfo.tokenName}</Col>
      </Row>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> Token Symbol</Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>{projectInfo.symbol}</Col>
      </Row>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> Token Decimals </Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>{projectInfo.decimals}</Col>
      </Row>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> IDO Total Supply </Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>{projectInfo && seperateNumWithComma(formatEther(projectInfo.amountOfTokensToSell) / 1.25)}</Col>
      </Row>
      <Row className={styles['row']}>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> Token Address </Col>
        <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>
          {projectInfo?.tokenAddress || ''}
        </Col>
      </Row>
    </div>
  )

  const alloTable = <>
    {isRegistered ? (
      <div className={styles['allocation-info']}>
        {/* visible after registration */}
        <Row className={styles['row']}>
          <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}>
            My Allocation
            <AppPopover
              content={<div style={{ width: '300px' }}>Your allocations could be changed during the registration period,
                and the final allocations will be fixed when the registration ends.
                You can always add staking tokens to increase your allocations during the registration time.</div>}>
              <QuestionCircleOutlined
                style={{ marginLeft: '10px', color: '#FFB852', cursor: 'pointer' }}
              />
            </AppPopover>
          </Col>
          <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>
            {
              !allocationTop || allocationTop == 0
                ? <><Spin style={{ verticalAlign: 'text-bottom' }}></Spin> <span>Calculating...</span></>
                : <>{seperateNumWithComma(formatEther(allocationTop, 2))} {projectInfo.symbol}</>
            }
          </Col>
        </Row>
        {/* visible after participation */}
        {
          isParticipated && !!participateInfo ? (<>
            <Row className={styles['row']}>
              <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}>
                My Participation
              </Col>
              <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>
                <AppPopover content={<>{formatEther(amountETHPaid).toFixed(2)} ETH</>}>
                  {seperateNumWithComma(formatEther(amountBought, 2))} {projectInfo.symbol}
                </AppPopover>
              </Col>
            </Row>
            <Row className={styles['row']}>
              <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['label']}> Purchase Time </Col>
              <Col span={isDesktopOrLaptop ? 12 : 24} className={styles['value']}>
                {formatDate(timeParticipated + '000', 'YYYY-MM-DD HH:mm:ss')}
              </Col>
            </Row>
            {/* <Row className={styles['row']}>
          <Col span={12} className={styles['label']}> Withdraw Status </Col>
          <Col span={12} className={styles['value']}>
            A
            Locked,
            Withdrawn, 
          </Col>
        </Row> */}
          </>)
            : <></>
        }
      </div>) : (<><Row className={styles['row']} justify="center" align="middle">
        <WarningOutlined
          style={{ color: 'red', fontSize: '32px', marginRight: '15px' }}
        />
        <div style={{ 'whiteSpace': 'nowrap', 'lineHeight': '50px' }}>
          You need to participate to check your allocation.
        </div>
      </Row></>)}
  </>

  const getTabInfo = () => {
    switch (tab) {
      case 'token_info':
        return TokenInfo;
      case 'about':
        return <>
          <div className={styles['about']} dangerouslySetInnerHTML={{ __html: projectInfo.aboutHtml || '' }}></div>
        </>
      case 'allocations':
        return alloTable;
      case 'sale_info':
      default:
        return SaleInfo;
    }
  }
  const getActionButton = () => {
    if (projectInfo.type == 1) {

    }
    // no contract in this chain  
    if (chain && chain.chainId != projectInfo.chainId) {
      return <TransactionButton
        noConnectText={'Connect wallet to stake'}
        className={'button ' + styles['register-button']}
        requiredChainId={projectInfo.chainId}
        switchNetworkText={'Switch Network'}
      >
        Switch Network
      </TransactionButton>
    }
    // not started
    if (statesReady && [0].includes(status)) {
      if (depositedAmount == 0) {
        return <TransactionButton
          noConnectText={'Connect wallet to stake'}
          className={'button ' + styles['register-button']}
          onClick={() => Router.push('/stake')}
        >
          Staking Now
        </TransactionButton>
      } else {
        return <AppPopover content={'Not started!'} wrap={true}>
          {
            <TransactionButton
              disabled={true}
              noConnectText={'Connect wallet to register'}
              className={'button ' + styles['register-button'] + ' ' + (isRegistered ? 'disabled' : '')}
              onClick={registerForSale}
            >
              {
                // isUserRegister ? 'Participate': 
                'Register'}
            </TransactionButton>
          }
        </AppPopover>
      }
    }
    // in registration
    if (statesReady && [1].includes(status)) {
      if (!isRegistered && !canRegister) {
        return <AppPopover content={'You have participated another project.'} wrap={true}>
          <TransactionButton
            disabled={true}
            disabledText={'Register'}
            className={'button ' + styles['register-button'] + ' ' + ('disabled')}
          >
            {
              // isUserRegister ? 'Participate' :
                'Register'}
          </TransactionButton>
        </AppPopover>
      } else if (!isUserRegister) {
        return <TransactionButton
          disabled={isRegistered}
          disabledText={'Participated'}
          noConnectText={'Connect wallet to register'}
          className={'button ' + styles['register-button'] + ' ' + (isRegistered ? 'disabled' : '')}
          onClick={showLoginModal}
          style={{ backgroundColor: isUserRegister ? '#D7FF1E' : '' }}
        >
          {'Register'}
        </TransactionButton>
      } else if (depositedAmount == 0) {
        return <TransactionButton
          noConnectText={'Connect wallet to stake'}
          className={'button ' + styles['register-button'] + ' ' + (isRegistered ? 'disabled' : '')}
          onClick={() => Router.push('/stake')}
        >
          Staking Now
        </TransactionButton>
      } else {
        return <TransactionButton
          disabled={isRegistered}
          disabledText={'Registered'}
          noConnectText={'Connect wallet to register'}
          className={'button ' + styles['register-button'] + ' ' + (isRegistered ? 'disabled' : '')}
          onClick={registerForSale}
          style={{ backgroundColor: isUserRegister ? '#D7FF1E' : '' }}
        >
          {
          // isUserRegister ? 'Participate' : 
          'Register'}
        </TransactionButton>
      }
    }
    // in sale
    if (statesReady && [3].includes(status) && isRegistered) {
      // if (isRegistered) {
      return <TransactionButton
        disabled={isParticipated}
        disabledText={'You have purchased'}
        noConnectText={'Connect wallet to purchase'}
        className={'button ' + styles['purchase-button'] + ' ' + (isParticipated ? 'disabled' : '')}
        onClick={() => setShowParticipateModal(true)}
      >
        Purchase
      </TransactionButton>
    }
    // sale end
    if (statesReady && isParticipated && [4].includes(status)) {
      return <AppPopover content={'Tokens not unlocked yet'}>
        <TransactionButton
          disabled={true}
          className={'button ' + styles['purchase-button']}
          noConnectText={'Connect wallet to withdraw'}
        >
          Withdraw
        </TransactionButton>
      </AppPopover>
    }
    if (statesReady && isParticipated && [5].includes(status)) {
      return <AppPopover
        wrap={true}
        content={withdrawTimeline()}
        placement={vestingPercentPerPortion.length >= 5 ? 'right' : 'top'}
      >
        <TransactionButton
          disabled={!canWithdrawArr.includes(true)}
          className={'button ' + styles['purchase-button']}
          noConnectText={'Connect wallet to withdraw'}
          onClick={withdrawTokens}
        >
          Withdraw
        </TransactionButton>
      </AppPopover>
    }
    if (statesReady && isRegistered && !isParticipated && [4, 5].includes(status)) {
      return <TransactionButton
        disabled={true}
        className={'button ' + styles['purchase-button']}
        noConnectText={'Connect wallet to withdraw'}
      >
        Sale ends
      </TransactionButton>
    }
  }

  function getReferralCode() {
    const f = new FormData;
    const params = {
      walletAddress,
      pid: pId,
    }
    for (let key in params) {
      f.append(key, params[key]);
    }
    return axios.post('/boba/referral/code', f)
      .then((res) => {
        setReferralCode(res.data);
        setReferralModalVisible(true);
      })
  }

  const referralLink = useMemo(() => {
    if (typeof window !== 'undefined') {
      return window.location.href.replace(/aff=.*/, '') + `&aff=${referralCode}`;
    } else {
      return '';
    }
  }, [referralCode])

  const getReferralButton = () => {
    if (status == 1 && isRegistered) {
      return <TransactionButton
        className={'button ' + styles['register-button']}
        onClick={() => { getReferralCode() }}
      >
        Invite
      </TransactionButton>
    } else {
      return <></>
    }
  }
  const cardInfo = useMemo(() => {
    const paymentTokenDecimas = projectInfo.paymentTokenDecimals;
    //let tokenPriceInUsd = projectInfo.tokenPriceInPT ? Number(projectInfo.tokenPriceInPT) / Math.pow(10, paymentTokenDecimas) : 0;
    let tokenPriceInUsd = projectInfo.tokenPriceInPT 
    ? (Number(projectInfo.tokenPriceInPT) / Math.pow(10, paymentTokenDecimas)) * (ethToUsd || 0) 
    : 0;
    // 链上实时值(读到了就用链上的,没读到回退到后端值)
    const totalTokensSold = onchainSale?.totalTokensSold ?? projectInfo.totalTokensSold;
    const totalETHRaised = onchainSale?.totalETHRaised;
    return {
      ...projectInfo,
      tokenPriceInUsd: tokenPriceInUsd,
      totalTokensSold,
      amountOfTokensToSell: onchainSale?.amountOfTokensToSell ?? projectInfo.amountOfTokensToSell,
      tokenPriceInETH: formatEther(onchainSale?.tokenPriceInETH || '0'),
      // totalRaised 以 ETH 计;如需 USD,再乘 ethToUsd。这里先给 ETH 数值
      totalRaised: totalETHRaised
        ? Number(formatEther(totalETHRaised))* (ethToUsd || 0)  
        : projectInfo.totalRaised * (ethToUsd || 0) ,
    }
  }, [projectInfo,onchainSale,ethToUsd]);

  const participateModalData = useMemo(() => {
    return { ...cardInfo, allocationTop, paymentTokenDecimal: depositDecimals || 18, paymentTokenSymbol: depositSymbol };
  }, [cardInfo, allocationTop, depositDecimals, depositSymbol])

  const withdrawTimeline = () => {
    return <Timeline
      mode="left"
      style={{ color: '#ffffff', minWidth: '400px' }}
    >
      {
        vestingPortionsUnlockTime && vestingPercentPerPortion ?
          vestingPortionsUnlockTime.map((sec, index) => {
            const unlockTimeMs = Number(sec) * 1000;  // 解锁时间:秒→毫秒
            // status: 0-未解锁, 1-可提, 2-已提
            let status: number = unlockTimeMs > chainNow ? 0 : 1;  // ← 用 chainNow(链上时间)数值比较
            if (isPortionWithdrawn && isPortionWithdrawn[index] && status === 1) {
              status = 2;
            }
            return <Timeline.Item key={index}
              color={status === 0 ? 'gray' : '#55BC7E'}
              label={formatDate(sec + '000', 'hh:mm:ss, Month DD, YYYY')}
            >
              {vestingPercentPerPortion[index] / 10}% &nbsp;
              {status === 0 ? <span>(Locked)</span> : <></>}
              {status === 1 ? <span>(Available now)</span> : <></>}
              {status === 2 ? <span style={{ color: '#55BC7E' }}>(Withdrawn)</span> : <></>}
            </Timeline.Item>
          })
          : <></>
      }
    </Timeline>
  };

  const trickers = useMemo(() => {
    let ret;
    try {
      ret = JSON.parse(projectInfo?.tricker)
    } catch (e) {
      // do nothing
    }
    return ret;
  }, [projectInfo])
  return (
    <main className={styles['container'] + " container"}
      style={trickers
        ? { backgroundImage: `url(${trickers?.background})`, backgroundSize: '100%', backgroundRepeat: 'repeat' }
        : {}
      }
    >
      {LoginModal}
      <InviteModal
        visible={referralModalVisible}
        onCancel={() => { setReferralModalVisible(false) }}
        url={referralLink}
        projectName={projectInfo.name}
      ></InviteModal>
      <PageLoader>
        <section className={styles['intro'] + ' main-content'}>
          <Row justify="space-between" align="middle" style={{ marginTop: '42px' }}>
            <Col span={24}>
              <h2 className={styles['pool-title']}>
                {projectInfo.name || ''}
              </h2>
            </Col>
          </Row>
          <Row style={{ marginTop: '42px' }}>
            <Col span={24}>
              <div className={styles['pool-desc']}>
                {projectInfo.description || ''}
              </div>
            </Col>
          </Row>
          {/* buttons */}
          <Row justify="start" align="middle" style={{ marginTop: '42px' }} gutter={[16, 16]}>
            <Col span={isDesktopOrLaptop ? 6 : 24}>
              <TransactionButton
                className={'button ' + styles['add-token-button']}
                onClick={() => { addToken(projectInfo.tokenAddress, projectInfo.symbol) }}
              >
                Add {projectInfo.symbol} to Wallet
              </TransactionButton>
            </Col>
            <Col span={isDesktopOrLaptop ? 6 : 24}>
              {getActionButton()}
            </Col>
            {/* <Col span={isDesktopOrLaptop ? 6 : 24}>
              {getReferralButton()}
            </Col> */}
            {/* <Col span={isDesktopOrLaptop ? 6 : 24}>
              <div className={styles['medias']}>
                <i className="icon"></i>
              </div>
            </Col> */}
          </Row>
          <Row style={{ marginTop: '55px' }}>
            <Col span={24}>
              <PoolCard
                className={styles['pool-info']}
                info={{ paymentTokenDecimals: depositDecimals, ...cardInfo, status }}
                styleNames={styles}
              >
              </PoolCard>
            </Col>
          </Row>
          <Row className={styles['progress']} gutter={[0, 16]}>
            {
              isDesktopOrLaptop
                ? <div className={styles['divider']}></div>
                : <></>
            }
            {
              [
                { stage: 'Registration Opens', time: mileStones.registrationTimeStarts },
                { stage: 'Registration Closes', time: mileStones.registrationTimeEnds },
                { stage: 'Sale Starts', time: mileStones.saleStart },
                { stage: 'Sale Ends', time: mileStones.saleEnd },
              ].map((info, index) => {
                return (
                  <Col span={isDesktopOrLaptop ? 6 : 24} className={styles['item'] + ' ' + (index < status ? styles['active'] : '')} key={index}>
                    <Row justify="center" align="middle" gutter={[0, 16]}>
                      <Col span={isDesktopOrLaptop ? 24 : 12}>
                        <i className={styles['icon'] + ' ' + styles['icon-pool-' + (index + 1)]}></i>
                      </Col>
                      <Col span={isDesktopOrLaptop ? 24 : 12}>
                        <div className={styles['label']}>
                          {info.stage}
                          <br />
                          {formatDate(info.time, 'HH:mm')},
                          <br />
                          {formatDate(info.time, 'Month DD, YYYY')}
                        </div>
                      </Col>
                    </Row>
                  </Col>
                )
              })
            }
          </Row>

        </section>
        <section className={styles['detail']}>
          <div className={styles['tab-info']}>
            <div className={styles['menu']}>
              <Row className="main-content" justify="start" align="middle">
                {
                  [
                    ['Sale Info', 'sale_info'],
                    ['Token Info', 'token_info'],
                    [isDesktopOrLaptop ? 'About the Project' : 'About', 'about'],
                    [isDesktopOrLaptop ? 'Your Allocations' : 'Allocations', 'allocations'],
                  ].map(([title, key], index) => (
                    <Col span={isDesktopOrLaptop ? 3 : 6}
                      key={index}
                      className={styles['item'] + ' ' + (tab == key && styles['active'] || '')}
                      onClick={() => setTab(key)}
                    >
                      {title}
                    </Col>
                  ))
                }
              </Row>
            </div>
            <div className="main-content">
              <div className={styles['tabs']}>
                {getTabInfo()}
              </div>
            </div>
          </div>

          {/* <div className={styles['alerts']}>
            <b>C2N. People First.</b>
            <div className={styles['entry']}>
              <p>Get Alerts For New Pools</p>
              <div className={styles['button']}>
                Sign up
              </div>
            </div>
          </div> */}
        </section>
      </PageLoader>
      <ParticipateModal
        visible={showParticipateModal}
        data={participateModalData}
        handleOk={participate}
        handleCancel={() => setShowParticipateModal(false)}
        tokenPriceInPT={projectInfo.tokenPriceInPT}
      />
    </main>
  )
}