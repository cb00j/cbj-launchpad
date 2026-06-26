import React, { useCallback, useEffect, useMemo, useState } from "react";
import { ProjectData } from "@src/types/ProjectData";
import { formatDate, formatNumber } from "@src/util/index";

import { Row, Col, Statistic } from 'antd';
const { Countdown } = Statistic;

import AppPopover from '@src/components/elements/AppPopover'

import styleNames from './PoolCard.module.scss';
import { formatEther, seperateNumWithComma } from "@src/util/index";
import { useResponsive } from '@src/hooks/useResponsive';


/**
 * Basic Button
 */

type PoolCardProds = {
  info: ProjectData,
  className?: any;
  styleNames?: any;
  children?: any;
};

const judgeClassName = (className, value) => {
  return value !== undefined ? className : (className + " loading-element");
};

export default function PoolCard(props: PoolCardProds) {
  const {
    isDesktopOrLaptop,
    isTabletOrMobile,
  } = useResponsive();

  const [status, setStatus] = useState<number>(-1)
  // 解锁阶段(status=4)某批倒计时结束后,强制重算"下一批"
  const [unlockRefresh, setUnlockRefresh] = useState<number>(0)

  const info: any = useMemo(() => {
    return props.info || {};
  }, [props])

  useEffect(() => {
    info && setStatus(info.status);
  }, [info]);

  const styles = props.styleNames && Object.assign(styleNames, props.styleNames) || styleNames;

  const basicElement = useCallback((key, className, mapper?) => {
    const _className = (className || key || '').replace(/([A-Z])/g, '-$1').toLowerCase();
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

  const progress = useMemo(() => {
    let p = formatEther(info.totalTokensSold) * 100 / formatEther(info.amountOfTokensToSell || 1);
    p = p > 100 ? 100 : p < 0 ? 0 : p;
    p = parseFloat(p.toFixed(2));
    return p;
  }, [info]);

  // 解锁阶段:从 vestingPortionsUnlockTime 找下一个还没到的解锁批次时间(毫秒)
  // 返回 null 表示全部已解锁
  const getNextUnlockTime = useCallback((): number | null => {
    const now = Date.now();
    let vesting: any = info.vestingPortionsUnlockTime;
    if (typeof vesting === 'string') {
      try { vesting = JSON.parse(vesting); } catch { vesting = []; }
    }
    if (!Array.isArray(vesting) || vesting.length === 0) {
      return info.unlockTime ? Number(info.unlockTime) : null;
    }
    const next = vesting
      .map((t: any) => Number(t) * 1000)   // 秒 → 毫秒
      .filter((t: number) => !isNaN(t))
      .find((t: number) => t > now);
    return next ?? null;
  }, [info, unlockRefresh]);

  const timer = useMemo(() => {
    const countdownValue = status === 4
      ? getNextUnlockTime()
      : [
          info.registrationTimeStarts,
          info.registrationTimeEnds,
          info.saleStart,
          info.saleEnd,
        ][status];

    return (
      <div className={styles['timer']}>
        <span style={{ marginRight: '4px' }}>
          {
            status === 4 && !countdownValue
              ? 'All unlocked'
              : ([
                  'Register starts in:',
                  'Register ends in:',
                  'Sale starts in:',
                  'Sale ends in:',
                  'Token unlocks in:',
                  'Sale ended',
                ][status] || 'Coming soon')
          }
        </span>
        {
          status > -1 && status < 5 && countdownValue
            ?
            <Countdown
              className={styles['counter']}
              valueStyle={{ fontSize: '16px', color: '#55BC7E' }}
              key={status + '-' + countdownValue}
              value={countdownValue}
              format="HH:mm:ss"
              onFinish={() => {
                if (status === 4) {
                  setUnlockRefresh(v => v + 1);
                } else {
                  setStatus(status + 1);
                }
              }} />
            : ''
        }
      </div>
    );
  }, [info, status, unlockRefresh, getNextUnlockTime]);

  const totalTokensSoldInEther = useMemo(() => {
    return formatEther(info?.totalTokensSold || 0);
  }, [info]);

  const amountOfTokensToSellInEther = useMemo(() => {
    return formatEther(info?.amountOfTokensToSell || 0);
  }, [info]);

  return (
    <div
      className={styles['pool-card'] + ' ' + (props.className || '')}
    >
      {isDesktopOrLaptop ? timer : <></>}
      <Row justify="start" align={isDesktopOrLaptop ? 'middle' : 'top'} style={{ marginTop: '0' }} gutter={16}>
        {/* icon */}
        <Col span={isDesktopOrLaptop ? 3 : 8} style={{ textAlign: 'left' }}>
          <i
            className={judgeClassName(styles['icon-logo'], info.img)}
            style={info.img ? { backgroundImage: `url(${info.img})` } : {}}
          >
          </i>
        </Col>
        {/* product name / title */}
        <Col span={16}>
          {basicElement('name', 'productName')}
          {isTabletOrMobile ? timer : <></>}
          {basicElement('description', 'describe')}
        </Col>
      </Row>
      <Row justify="space-between" align="middle" className={styles['bottom-info']} style={{ marginTop: '30px' }} gutter={[16, 16]}>
        <Col span={isDesktopOrLaptop ? 6 : 12}>
          <div className={styles['total-raise-wrapper']}>
            <div className={styles['total-raise-label']}>Total raised</div>
            <AppPopover content={<>{seperateNumWithComma(info && (info.totalRaised / 1e36).toFixed(2))}</>}>
              <div className={styles['total-raise']}>
                {
                  status > 2
                    ? <>$ {info?.totalRaised && formatNumber((info?.totalRaised / 1).toFixed(2)) || '0.00'}</>
                    : <span style={{ fontSize: '.8em', whiteSpace: 'nowrap' }}>Starts on {formatDate(info.registrationTimeEnds, 'Month DD, YYYY')}</span>
                }
              </div>
            </AppPopover>
          </div>
        </Col>
        <Col span={isDesktopOrLaptop ? 4 : 12}>
          <div className={styles['bottom-info-item']}>
            <div className={styles['row']}>
              <i className={[styles['bottom-icon'], 'icon', 'icon-project-card-1'].join(' ')}></i>
              <label>Followers</label>
            </div>
            <div className={styles['row']}>
              {basicElement('follower', 'followers', v => v || '0')}
            </div>
          </div>
        </Col>
        <Col span={isDesktopOrLaptop ? 4 : 12}>
          <div className={styles['bottom-info-item']}>
            <div className={styles['row']}>
              <i className={[styles['bottom-icon'], 'icon', 'icon-project-card-2'].join(' ')}></i>
              <label>Start Date</label>
            </div>
            <div className={styles['row']}>
              {basicElement('saleStart', 'startDate', v => formatDate(v, 'Month DD, YYYY'))}
            </div>
          </div>
        </Col>
        <Col span={isDesktopOrLaptop ? 4 : 12}>
          <div className={styles['bottom-info-item']}>
            <div className={styles['row']}>
              <i className={[styles['bottom-icon'], 'icon', 'icon-project-card-3'].join(' ')}></i>
              <label>Token Price</label>
            </div>
            <div className={styles['row']}>
              {/* <AppPopover content={info && info.tokenPriceInUsd}> */}
              <div className={styles['token-price']}>
                {(info?.tokenPriceInETH)} ETH
                <div className={styles['extra']}>
                  ~ ${(info?.tokenPriceInUsd)?.toFixed(4) || '0'}
                </div>
              </div>
              {/* </AppPopover> */}
            </div>
          </div>
        </Col>
      </Row>
      <Row style={{ marginTop: '30px' }}>
        <Col span={24} offset={0}>
          <AppPopover content={
            <>Sale:{
              seperateNumWithComma(totalTokensSoldInEther > amountOfTokensToSellInEther ? amountOfTokensToSellInEther : totalTokensSoldInEther)
            } / {
                seperateNumWithComma(amountOfTokensToSellInEther / 1.25)
              }
            </>}>
            <div className={styles['card-progress']}
              onClick={() => {
                console.log('sale')
              }}
            >
              <div className={styles['progress-background']}></div>
              <div className={styles['progress-colored']} style={{ width: progress + '%' }}></div>
              <div className={styles['progress-text']} style={{ color: progress > 70 ? '#ffffff' : '#000000' }}>
                Sale: {progress || '00.00'}%
              </div>
            </div>
          </AppPopover>
        </Col>
      </Row>
    </div>
  );
}