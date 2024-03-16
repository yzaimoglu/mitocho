<script lang="ts">
	export let title = 'Mitocho';
	export let description = 'The powerhouse of authentication';
	export let color = 'normal';

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

<div class={getBackgroundColor()}>
	<slot />
</div>
