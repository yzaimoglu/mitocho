import { verify } from 'paseto-ts/v4';
import { type FormattedPublicKey, type PasetoPayload } from './types';

const formatPublicKey = async (b64: string): Promise<FormattedPublicKey> => {
  return 'k4.public.' + b64;
};

const getPasetoPayload = async (publicKey: FormattedPublicKey, token: string): Promise<PasetoPayload> => {
  const formattedKey = await formatPublicKey(publicKey);
  let verified: any = { error: false, payload: undefined };

  try {
    verified = verify(formattedKey, token);
  } catch (error: any) {
    return { error: true, payload: undefined };
  }

  return { error: false, payload: verified.payload };
};

export { getPasetoPayload };
