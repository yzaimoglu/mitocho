import { writable } from 'svelte/store';

export interface LoginCredentials {
  email: string;
  password: string;
}

const createLoginStore = () => {
  const { subscribe, set } = writable<LoginCredentials>({ email: '', password: '' });

  return {
    subscribe,
    setCredentials: (credentials: LoginCredentials) => set(credentials),
    isEmpty: () => {
      let isEmpty = true;
      const unsubscribe = subscribe(currentValue => {
        isEmpty = !currentValue.email && !currentValue.password;
      });
      unsubscribe();
      return isEmpty;
    },
    reset: () => set({ email: '', password: '' })
  };
}

export const loginStore = createLoginStore();
