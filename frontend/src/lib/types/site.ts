import type { SelectItem } from "@/components/MitochoSelect.svelte";

export type ColorSelect = SelectItem;

export const Colors: ColorSelect[] = [
  {
    value: 'green',
    text: 'Green'
  },
  {
    value: 'red',
    text: 'Red'
  },
  {
    value: 'blue',
    text: 'Blue'
  },
  {
    value: 'purple',
    text: 'Purple'
  }
];
