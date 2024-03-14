import { verify } from 'paseto-ts/v4';

type FormattedPayload = string;

const formatPublicKey = async (b64: string): Promise<FormattedPayload> => {
  return 'k4.public.' + b64;
};

type Payload = {
  error: boolean;
  payload: FormattedPayload | undefined;
};

const getPayload = async (publicKey: string, token: string): Promise<Payload> => {
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
