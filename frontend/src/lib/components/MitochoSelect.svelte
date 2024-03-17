<script lang="ts" context="module">
	export type SelectItem = {
		text: string;
		value: string;
	};
</script>

<script lang="ts">
	import * as Select from '@/components/ui/select';
	import type { Selected } from 'bits-ui';

	export let selected: any;
	export let className: string = 'w-full';
	export let label: string = null;
	export let placeholder: string;
	export let items: SelectItem[];

	const selectedChange = (v: Selected<unknown>) => {
		selected = v.value;
	};
</script>

<Select.Root onSelectedChange={selectedChange}>
	<Select.Trigger class={className}>
		<Select.Value {placeholder} />
	</Select.Trigger>
	<Select.Content>
		{#if label !== null}
			<Select.Label>{label}</Select.Label>
		{/if}
		{#each items as item}
			<Select.Item value={item.value}>{item.text}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>
