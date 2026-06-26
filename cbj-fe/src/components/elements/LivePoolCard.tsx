import React, { useCallback, useMemo, useState, useEffect } from "react";
import { ProjectData } from "@src/types/ProjectData";
import { formatDate } from "@src/util/index"
import { Row, Col, Popover, Statistic } from 'antd';
import styleNames from './LivePoolCard.module.scss';
import AppPopover from '@src/components/elements/AppPopover'
import { formatEther, seperateNumWithComma, formatNumber } from "@src/util/index";
import { useThirdParty } from '@src/hooks/useThirdParty';
import {
  BigNumber
} from 'ethers'
import { useResponsive } from "@src/hooks/useResponsive";

const { Countdown } = Statistic;
/**
 * Basic Button
 */

interface LivePoolCardProds {
  info: ProjectData,
  className?: any;
  styleNames?: any;
  onClick?: any;
};

const judgeClassName = (className, value) => {
  return value !== undefined ? className : ( className + " loading-element" );
};

export default function LivePoolCard(props: LivePoolCardProds) {
  const info = props.info || ({} as ProjectData);
  const styles = props.styleNames && Object.assign(styleNames, props.styleNames) || styleNames;

  const {
    ethToUsd,
  } = useThirdParty();

  const {
    isDesktopOrLaptop
  } = useResponsive()

  const [status, setStatus] = useState<number>(-1)
  // 用于在解锁阶段(status=4)某批倒计时结束后,强制重新计算"下一批"的时间
  const [unlockRefresh, setUnlockRefresh] = useState<number>(0)

  useEffect(()=> {
    info && setStatus(info.status);
  }, [info]);

  const basicElement = useCallback((key, className, mapper?) => {
    const _className = (className || key || '').replace(/([A-Z])/g,'-$1').toLowerCase();
    const value = mapper && mapper(info[key]) || info[key];
    return (
      <div
        className={judgeClassName(
          styles[_className],
          value
        )}
      >
        {value || '\u00a0'}
      </div>
    );
  }, [info]);

  const progress = useMemo(()=>{
    let p = formatEther(info.totalTokensSold) * 100 / formatEther(info.amountOfTokensToSell||1);
    p = p > 100 ? 100 : p < 0 ? 0 : p;
    p = parseFloat(p.toFixed(2));
    return p;
  }, [info]);

  const totalRaisedUsd = useMemo(() => {
    if (!info.totalRaised) return 0;
    try {
      const ethAmount = Number(formatEther(info.totalRaised));  // wei → ETH
      return ethAmount * (ethToUsd || 0);                        // ETH → USD
    } catch (e) {
      console.error('totalRaisedUsd calc failed:', e);
      return 0;
    }
  }, [info.totalRaised, ethToUsd]);


  const tokenPriceInUsd:number = useMemo(()=>{
    const paymentTokenDecimas = props?.info?.paymentTokenDecimals ? Number(props.info.paymentTokenDecimals) : 0;
    return props?.info?.tokenPriceInPT ? Number(props.info.tokenPriceInPT)/Math.pow(10, paymentTokenDecimas) * (ethToUsd || 0) : 0;
  }, [props, ethToUsd])

  // 解锁阶段:从 vestingPortionsUnlockTime 找下一个还没到的解锁批次时间(毫秒)
  // 返回 null 表示所有批次都已解锁
  const getNextUnlockTime = useCallback((): number | null => {
    const now = Date.now();
    let vesting: any = info.vestingPortionsUnlockTime;
    // 兼容:万一是字符串则解析
    if (typeof vesting === 'string') {
      try { vesting = JSON.parse(vesting); } catch { vesting = []; }
    }
    if (!Array.isArray(vesting) || vesting.length === 0) {
      // 没有 vesting 数组,退回总解锁时间
      return info.unlockTime ? Number(info.unlockTime) : null;
    }
    // vesting 是秒级时间戳,转毫秒后找第一个 > now 的批次
    const next = vesting
      .map((t: any) => Number(t) * 1000)
      .filter((t: number) => !isNaN(t))
      .find((t: number) => t > now);
    return next ?? null;  // 全部解锁则 null
  }, [info, unlockRefresh]);

  const timer = useMemo(() =>{
    // status=4(解锁阶段):倒计时指向下一个未解锁批次;其它阶段用对应单点时间
    const countdownValue = status === 4
      ? getNextUnlockTime()
      : [
          info.registrationTimeStarts,
          info.registrationTimeEnds,
          info.saleStart,
          info.saleEnd,
        ][status];

    return (<div className={styles['timer']}>
      <span style={{marginRight:'4px'}}>
        {
          // 解锁阶段且已全部解锁 → 显示"已全部解锁";否则正常文案
          status === 4 && !countdownValue
            ? 'All unlocked'
            : ([
                'Register starts in:',
                'Register ends in:',
                'Sale starts in:',
                'Sale ends in:',
                'Token unlocks in:',
                'Sale ended',
              ][status] || 'To start')
        }
      </span>
      {
        status > -1 && status < 5 && countdownValue
        ? 
        <Countdown 
          className={styles['counter']}
          valueStyle={{fontSize:'16px',color:'#D7FF1E'}}
          key={status + '-' + countdownValue}
          value={countdownValue} 
          format="HH:mm:ss" 
          onFinish={()=>{
            if (status === 4) {
              // 解锁阶段:某批倒计时结束,触发重算下一批(不推进 status)
              setUnlockRefresh(v => v + 1);
            } else {
              setStatus(status + 1);
            }
          }} />
          : ''
      }
    </div>)
  }, [info, status, unlockRefresh, getNextUnlockTime])

  
  const trickers = useMemo(()=>{
    let ret;
    try{
      ret = JSON.parse(info?.tricker)
    } catch(e) {
      // do nothing
    }
    return ret;
  }, [info])

  if(info.status == -1) {
    return (
      <div 
      className={`${styles['live-pool-card']} ${(props.className || '')} ${styles['not-start']}`} 
      onClick={props.onClick}
      style={trickers?.cardBackground ? {
        backgroundImage: `url(${trickers.cardBackground})`,
      }: {}}
      >
        {isDesktopOrLaptop ? timer : <></>}
        <Row justify="start" align={isDesktopOrLaptop ? "middle" : "top"} gutter={16} style={{marginTop:'30px'}}>
          <Col span={isDesktopOrLaptop ? 6 : 8} style={{textAlign:'center'}}>
            {/* icon */}
            <i
              className={judgeClassName(styles['icon-logo'], info.img)}
              style={info.img ? { backgroundImage: `url(${info.img})` } : {}}
            >
            </i>
          </Col>
          <Col span={isDesktopOrLaptop ? 18 : 16}>
            {/* product name / title */}
            {basicElement('name', 'productName')}
            {basicElement('description', 'describe')}
            {isDesktopOrLaptop ? <></> : timer}
          </Col>
        </Row>
        <Row 
        justify="start" 
        align="middle" 
        style={{marginTop:'20px'}}
        className={styles['total-raise-wrapper']}>
          <div className={styles['total-raise-label']}>Total raised</div>
          <AppPopover content={<>{seperateNumWithComma(info&&(info.totalRaised/1).toFixed(2))}</>}>
            {basicElement('totalRaised', 'totalRaise', v=>`$ --`)}
          </AppPopover>
        </Row>
        <Row 
        justify={'center'}
        style={{marginTop:'20px'}}
        className={styles['bottom-info']} 
        >
          <div className={styles['coming-soon']}>
            ~ Coming Soon ~
          </div>
        </Row>
        <Row>
          <Col span={24} offset={0}>
            <div className={styles['card-progress']}>
              <div className={styles['progress-background']}></div>
            </div>
          </Col>
        </Row>
      </div>
    );
  }

  return (
    <div 
     className={styles['live-pool-card']+' '+(props.className || '')} 
     // 如果 status 是 0，可以选择不绑定 onClick 事件
     onClick={status === 0 ? undefined : props.onClick} 
     style={{ position: 'relative' }} // 【关键修改】：添加相对定位作为蒙层参照
    >
      {status === 0 && (
        <div
          style={{
            position: 'absolute',
            top: 0,
            left: 0,
            width: '100%',
            height: '100%',
            backgroundColor: 'rgba(0, 0, 0, 0.5)', // 半透明黑色蒙层
            backdropFilter: 'blur(4px)', // 毛玻璃效果
            zIndex: 10, // 确保蒙层盖住卡片内所有元素
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            borderRadius: 'inherit', // 自动继承你在 css 里给卡片写的 border-radius
            cursor: 'not-allowed'
          }}
          onClick={(e) => e.stopPropagation()} // 阻止点击事件穿透到底层
        >
          <span style={{
            color: '#FFFFFF',
            fontSize: '32px',
            fontWeight: '900',
            letterSpacing: '2px',
            textShadow: '0 4px 8px rgba(0, 0, 0, 0.5)' // 增加阴影让文字在复杂背景下更清晰
          }}>
            In Coming!
          </span>
        </div>
      )}
      
      {isDesktopOrLaptop ? timer : <></>}
      <Row justify="start" align={isDesktopOrLaptop ? "middle" : "top"} gutter={16} style={{marginTop:'30px'}}>
        <Col span={isDesktopOrLaptop ? 6 : 8} style={{textAlign:'center'}}>
          {/* icon */}
          <i
            className={judgeClassName(styles['icon-logo'], info.img)}
            style={info.img ? { backgroundImage: `url(${info.img})` } : {}}
          >
          </i>
        </Col>
        <Col span={isDesktopOrLaptop ? 18 : 16}>
          {/* product name / title */}
          {basicElement('name', 'productName')}
          {basicElement('description', 'describe')}
          {isDesktopOrLaptop ? <></> : timer}
        </Col>
      </Row>
      <Row 
       justify="start" 
       align="middle" 
       style={{marginTop:'20px'}}
       className={styles['total-raise-wrapper']}>
        <div className={styles['total-raise-label']}>
          Total raised
        </div>
        <AppPopover content={<>$ {seperateNumWithComma(totalRaisedUsd.toFixed(2))}</>}>
        <div className={styles['total-raise']}>
          {
            status > 2
              ? <>$ {totalRaisedUsd ? formatNumber(totalRaisedUsd.toFixed(2)) : '0.00'}</>
              : <span style={{fontSize: '.8em'}}>Starts on {formatDate(info.registrationTimeEnds, 'Month DD, YYYY')}</span>
          }
        </div>
      </AppPopover>
      </Row>
      <Row 
       justify={'space-between'}
       style={{marginTop:'20px'}}
       className={styles['bottom-info']} 
       >
        {/* 1 */}
        <Col span={8} className={styles['bottom-info-item']}>
          <div className={styles['row']}>
            <i className="bottom-icon icon icon-project-card-1"></i>
            <label className={styles['label']}>Followers</label>
          </div>
          <div className={[styles['row'], styles['value']].join(' ')}>
            {basicElement('follower', 'followers', v=>v||'0')}
          </div>
        </Col>
        {/* 2 */}
        <Col span={8} className={styles['bottom-info-item']}>
          <div className={styles['row']}>
            <i className="bottom-icon icon icon-project-card-2"></i>
            <label className={styles['label']}>Start Date</label>
          </div>
          <div className={[styles['row'], styles['value']].join(' ')}>
            {basicElement('saleStart', 'startDate', v=>formatDate(v, 'YYYY-MM-DD'))}
          </div>
        </Col>
        {/* 3 */}
        <Col span={8} className={styles['bottom-info-item']}>
          <div className={styles['row']}>
            <i className="bottom-icon icon icon-project-card-3"></i>
            <label className={styles['label']}>Token Price</label>
          </div>
          <div className={[styles['row'], styles['value']].join(' ')}>
            {/* <AppPopover content={tokenPriceInUsd}> */}
              <div
                className={styles['token-price']}
              >
                $ {tokenPriceInUsd && tokenPriceInUsd.toFixed(4)}
              </div>
            {/* </AppPopover> */}
          </div>
        </Col>
      </Row>
      <Row>
        <Col span={24} offset={0}>
        <AppPopover content={`Sale: ${progress||'00.00'}%`}>
          <div className={styles['card-progress']}>
            <div className={styles['progress-background']}></div>
            <div className={styles['progress-colored']} style={{width: progress+'%'}}></div>
            {/* <div className={styles['progress-text']} style={{color: progress>0.6?'#ffffff':'#000000'}}>
              Sale: {progress||'-'}%
            </div> */}
          </div>
        </AppPopover>
        </Col>
      </Row>
    </div>
  );
}