import BaseService from './base';
import type { BaseRes } from './base';

type PostInstallFinishRes = BaseRes & {
  response?: {
    error?: string;
  };
};

export type PostInstallFinishReq = {
  username: string;
  email: string;
  password: string;
  password_repeat: string;
  domain: string;
};

type GetInstallFinishRes = BaseRes & {
  response: {
    finished: boolean;
  };
};

export const postInstallFinish = async (fetch: any, data: any): Promise<PostInstallFinishRes> => {
  const base = new BaseService(fetch);
  return await base.post<PostInstallFinishReq, PostInstallFinishRes>(`install/finish`, data);
};

export const getInstallFinish = async (fetch: any): Promise<GetInstallFinishRes> => {
  const base = new BaseService(fetch);
  return await base.get<GetInstallFinishRes>('install/finish');
};
