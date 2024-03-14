import { writable } from 'svelte/store';

function createLoadingStore() {
  const { subscribe, set } = writable(false);

  return {
    subscribe,
    start: () => set(true),
    finish: () => set(false)
  };
}

export const loading = createLoadingStore();
