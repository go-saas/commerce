import Footer from '@/components/Footer';
import RightContent from '@/components/RightContent';
import { LinkOutlined } from '@ant-design/icons';
import type { Settings as LayoutSettings } from '@ant-design/pro-components';
import { SettingDrawer } from '@ant-design/pro-components';
import type { RunTimeLayoutConfig } from '@umijs/max';
import { history, Link } from '@umijs/max';
import { getRequestInstance } from '@@/plugin-request/request';
import { setDefaultAxiosFactory } from '@kit/platform-api';
import defaultSettings from '../config/defaultSettings';
import type { AxiosResponse } from 'umi';
import {
  authRequestInterceptor,
  csrfRequestInterceptor,
  csrfRespInterceptor,
  saasRequestInterceptor,
  authRespInterceptor,
  bizErrorInterceptor,
  tenantErrorInterceptor,
  ErrorShowType,
  FriendlyError,
} from '@kit/core';
import type { RequestConfig } from 'umi';
import { message, notification } from 'antd';

import 'echarts-wordcloud';

const isDev = process.env.NODE_ENV === 'development';
const loginPath = '/user/login';

/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * */
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  currentUser?: API.CurrentUser;
  loading?: boolean;
  fetchUserInfo?: () => Promise<API.CurrentUser | undefined>;
}> {
  setDefaultAxiosFactory(getRequestInstance);
  return {
    settings: defaultSettings,
  };
}

// ProLayout 支持的api https://procomponents.ant.design/components/layout
export const layout: RunTimeLayoutConfig = ({ initialState, setInitialState }) => {
  return {
    rightContentRender: () => <RightContent />,
    disableContentMargin: false,
    waterMarkProps: {
      content: initialState?.currentUser?.name,
    },
    footerRender: () => <Footer />,
    onPageChange: () => {
      // const { location } = history;
      console.log(history);
      // // 如果没有登录，重定向到 login
      // if (!initialState?.currentUser && location.pathname !== loginPath) {
      //   history.push(loginPath);
      // }
    },
    links:[],
    menuHeaderRender: undefined,
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    // 增加一个 loading 的状态
    childrenRender: (children, props) => {
      // if (initialState?.loading) return <PageLoading />;
      return (
        <>
          {children}
          {!props.location?.pathname?.includes('/login') && (
            <SettingDrawer
              disableUrlParams
              enableDarkTheme
              settings={initialState?.settings}
              onSettingChange={(settings) => {
                setInitialState((preInitialState) => ({
                  ...preInitialState,
                  settings,
                }));
              }}
            />
          )}
        </>
      );
    },
    ...initialState?.settings,
  };
};

function errorInterceptor() {
  return [
    (resp: AxiosResponse) => {
      return resp;
    },
    (error: any) => {
      console.log(error.wrap || error);
      const { config, code, response, request } = error.wrap || error || {};
      let showType = ErrorShowType.ERROR_MESSAGE;
      let errorMessage = '';
      let errorCode = code;

      if ('showType' in config) {
        showType = config.showType;
      }

      if (error instanceof FriendlyError) {
        errorCode = error.reason;
        errorMessage = error.message || errorCode;
      }
      if (response) {
        if (!errorMessage) {
          const status = response.status;
          errorMessage = status.toString();
        }
        if (!errorCode) {
          errorCode = errorMessage;
        }
      } else if (request) {
        // 请求已经成功发起，但没有收到响应
        // \`error.request\` 在浏览器中是 XMLHttpRequest 的实例，
        // 而在node.js中是 http.ClientRequest 的实例
        errorMessage = 'None response! Please retry.';
        showType = ErrorShowType.ERROR_MESSAGE;
      } else {
        // 发送请求时出了点问题
        errorMessage = 'Request error, please retry.';
        showType = ErrorShowType.ERROR_MESSAGE;
      }
      switch (showType) {
        case ErrorShowType.SILENT:
          // do nothing
          break;
        case ErrorShowType.WARN_MESSAGE:
          message.warn(errorMessage);
          break;
        case ErrorShowType.ERROR_MESSAGE:
          message.error(errorMessage);
          break;
        case ErrorShowType.NOTIFICATION:
          notification.open({
            description: errorMessage,
            message: errorCode,
          });
          break;
        case ErrorShowType.REDIRECT:
          // TODO: redirect
          break;
        default:
          message.error(errorMessage);
      }
      return Promise.reject(new FriendlyError(code, errorCode, errorMessage));
    },
  ];
}

export const request: RequestConfig = {
  baseURL: BASE_URL,
  withCredentials: true,
  requestInterceptors: [
    authRequestInterceptor(),
    csrfRequestInterceptor(),
    saasRequestInterceptor(),
  ],
  responseInterceptors: [
    csrfRespInterceptor(),
    bizErrorInterceptor(),
    errorInterceptor() as any,
    authRespInterceptor(() => {
      //redirect to login
      history.push(loginPath);
    }),
    tenantErrorInterceptor(),
  ],
};
