import { verify } from 'paseto-ts/v4';

const formatPublicKey = async (b64: string): Promise<string> => {
  return 'k4.public.' + b64;
};

const getPayload = async (publicKey: string, token: string): Promise<object> => {
  const formattedKey = await formatPublicKey(publicKey);
  let verified: any = { error: false, payload: undefined };
  
  try {
    verified = verify(formattedKey, token);
  } catch (error: any) {
    return { error: true, payload: undefined };  
  }

  return { error: false, payload: verified.payload };
};

export { getPayload };
