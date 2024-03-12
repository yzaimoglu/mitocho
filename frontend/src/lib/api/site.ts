import BaseService from './base';

export const postPublicKey = async (fetch: any, data: any) => {
  const base = new BaseService(fetch);
  return await base.post(`site/public_key`, data);
}
