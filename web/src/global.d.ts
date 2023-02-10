
export {};

declare global {
  export type Writable<T> = {
    -readonly [P in keyof T]: T[P];
  };
  interface Window {
    __POWERED_BY_QIANKUN__: boolean;
  }
}
