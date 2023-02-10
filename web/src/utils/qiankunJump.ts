import { history as umiHistory } from 'umi';

export const qiankunJump = (url: string, name = '', params = null) => {
  // eslint-disable-next-line @typescript-eslint/no-unused-expressions
  window.__POWERED_BY_QIANKUN__ ? history.pushState(params, name, url) : umiHistory.push(url);
};
