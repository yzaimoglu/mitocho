import BaseService from './base';
import BaseRes from './base';

type postPublicKeyReq = {
  access_token: string;
};

type postPublicKeyRes = BaseRes & {
  response: {
    public_key?: string;
    error?: string;
  }
};

export const postPublicKey = async (fetch: any, data: postPublicKeyReq) => {
  const base = new BaseService(fetch);
  return await base.post<postPublicKeyReq, postPublicKeyRes>(`site/public_key`, data);
};
