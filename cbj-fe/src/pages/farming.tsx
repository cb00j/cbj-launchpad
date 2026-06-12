import type { AppProps } from 'next/app'
import { Row, Col } from 'antd'

import styles from './farming.module.scss'
import FarmingForm from '@src/containers/FarmingForm/FarmingForm'
import { useResponsive } from '@src/hooks/useResponsive';
import farmConfigs from '@src/config/farms'
import { QuestionCircleOutlined } from '@ant-design/icons';

import {
  EARNED_TOKEN_ADDRESS,
} from '@src/config'

/**
 * Stake form page
 */
export default function Pools({ Component, pageProps }: AppProps) {
  const {
    isDesktopOrLaptop,
    isTabletOrMobile,
  } = useResponsive();

  return (
    <main className={styles['container'] + " container"}>
      <section className={styles['intro'] + ' main-content'}>
        <h2 className={styles['stake-title']}>
          <Row justify="space-between">
            <Col>
              <span>Yield Farms</span>
            </Col>
            <Col>
              <span style={{ fontSize: '16px', verticalAlign: 'middle' }}>
                <QuestionCircleOutlined style={{ fontSize: '36px', verticalAlign: 'middle', marginRight: '.2em' }}></QuestionCircleOutlined>
                See Tutorial: &nbsp;
                <span
                  className={styles['link']}
                  onClick={() => {
                  }}>{isDesktopOrLaptop ? 'CBJ Farm Tutorial' : 'Tutorial'} </span>
              </span>
            </Col>
          </Row>
        </h2>
        <h3 className={styles['stake-subtitle']}>
          Yield Farms allow users to earn Reward token while supporting CBJ by staking LP Tokens.
        </h3>
      </section>
      <section className={styles['staking']}>
        <div className="main-content">
          <Row gutter={32}>
            {
              farmConfigs.map((item, index) => {
                return (
                  <Col span={isDesktopOrLaptop ? 8 : 24} key={index}>
                    
                    {/* 添加一个相对定位的外壳 */}
                    <div style={{ position: 'relative', height: '100%' }}>
                      
                      {/* 如果不 available，则显示遮罩层 */}
                      {!item.available && (
                        <div
                          style={{
                            position: 'absolute',
                            top: 0,
                            left: 0,
                            width: '100%',
                            height: '100%',
                            backgroundColor: 'rgba(0, 0, 0, 0.4)', // 半透明黑底
                            backdropFilter: 'blur(3px)', // 毛玻璃效果
                            zIndex: 10, // 确保蒙层在最上层
                            display: 'flex',
                            justifyContent: 'center',
                            alignItems: 'center',
                            borderRadius: '16px', // 请根据 FarmingForm 实际的 UI 圆角调整这里
                            cursor: 'not-allowed', // 鼠标变成禁止符号
                          }}
                        >
                          <span style={{
                            color: '#fff',
                            fontSize: '28px',
                            fontWeight: 'bold',
                            letterSpacing: '2px',
                            textShadow: '0 2px 8px rgba(0,0,0,0.6)' // 增加文字阴影让其更清晰
                          }}>
                            In Coming!
                          </span>
                        </div>
                      )}

                      {/* 底部的表单。加了 grayscale 让未开放的池子变灰，视觉提示更明显 */}
                      <div style={{ 
                        height: '100%', 
                        filter: !item.available ? 'grayscale(80%)' : 'none',
                        pointerEvents: !item.available ? 'none' : 'auto' // 防止未开放时还能被点击
                      }}>
                        <FarmingForm
                          chainId={item.chainId}
                          depositTokenAddress={item.depositTokenAddress}
                          earnedTokenAddress={item.earnedTokenAddress}
                          stakingAddress={item.stakingAddress}
                          poolId={item.poolId}
                          available={item.available}
                          depositSymbol={item.depositSymbol}
                          earnedSymbol={item.earnedSymbol}
                          title={item.title}
                          depositLogo={item.depositLogo}
                          earnedLogo={item.earnedLogo}
                          getLptHref={item.getLptHref}
                          aprRate={item.aprRate}
                          aprUrl={item.aprUrl}
                        />
                      </div>

                    </div>
                  </Col>
                )
              })
            }
          </Row>
        </div>
      </section>
    </main>
  )
}