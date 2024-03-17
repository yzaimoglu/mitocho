import { writable } from 'svelte/store';

export interface Role {
  name: string;
  readable_name: string;
  permissions: string[];
};

export interface User {
  sid: string;
  username: string;
  email: string;
  roles: Role[];
  permissions: string[];
}

const createAuthStore = () => {
  const { subscribe, set } = writable<User>(
    {
      sid: '',
      username: '',
      email: '',
      roles: [],
      permissions: []
    }
  );

  return {
    subscribe,
    setUser: (user: User) => set(user),
    isEmpty: () => {
      let isEmpty = true;
      const unsubscribe = subscribe(currentValue => {
        isEmpty = !currentValue.sid;
      });
      unsubscribe();
      return isEmpty;
    },

    reset: () => set(
      {
        sid: '',
        username: '',
        email: '',
        roles: [],
        permissions: []
      }
    )
  };
}

export const authStore = createAuthStore();
