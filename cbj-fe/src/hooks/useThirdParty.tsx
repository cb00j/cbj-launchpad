import axios from '@src/api/axios'
import { useAppSelector, useAppDispatch } from "@src/redux/hooks";
import { setBobaToUsd, setEthToUsd, setBreToUsd } from '@src/redux/modules/third-party';

export const useThirdParty = () => {
  const dispatch = useAppDispatch();
  const bobaToUsd = useAppSelector(state => state.thirdParty.bobaToUsd);
  const ethToUsd = useAppSelector(state => state.thirdParty.ethToUsd);
  const breToUsd:number = useAppSelector(state => state.thirdParty.breToUsd);
  
  async function getBobaToUsd() {
    const ret:any = await axios.get('https://price-api.crypto.com/price/v1/exchange/boba-network')
    dispatch(setBobaToUsd(ret.fiat.usd));
  }

  function getEthToUsd() {
    return axios.get('https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=BTC,USD,EUR')
    .then((response:any)=>{
      dispatch(setEthToUsd(response.USD))
    })
  }

  function getBreToUsd() {
    // return axios.get('/boba/apr/bre_price')
    // .then((response:any)=>{
    //   dispatch(setBreToUsd(Number(response.data)))
    // })
    // 1. 发送 dispatch 通知 Redux 仓库更新价格为 0.5
    dispatch(setBreToUsd(0.5));
    
    // 2. 返回一个立刻成功的 Promise，防止外面的 .then() 报错
    return Promise.resolve();
  }

  return {
    bobaToUsd,
    ethToUsd,
    breToUsd,
    stakedTokenToUsd: breToUsd,
    earnedTokenToUsd: breToUsd,

    getBobaToUsd,
    getEthToUsd,
    getBreToUsd
  }
}