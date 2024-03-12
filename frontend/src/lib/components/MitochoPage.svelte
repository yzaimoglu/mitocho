<script lang="ts">
	import { onMount } from 'svelte';

	export let loaded = false;
	export let title = 'Mitocho';
	export let description = 'The powerhouse of authentication';
	export let color = 'normal';

	onMount(async () => {
		loaded = true;
	});

	const getBackgroundColor = (): string => {
		switch (color) {
			case 'error':
				return 'bg-red-50';
			default:
				return 'bg-green-50';
		}
	};
</script>

<svelte:head>
	<title>{title}</title>
	<meta name="description" content={description} />
</svelte:head>

{#if loaded}
	<div class={getBackgroundColor()}>
		<slot />
	</div>
{/if}
