<script lang="ts">
	import MitochoFooter from './MitochoFooter.svelte';

	export let footer = true;
	export let gap = 4;
</script>

<div class={`flex flex-col justify-center items-center min-h-screen gap-${gap}`}>
	<slot />
	{#if footer}
		<MitochoFooter />
	{/if}
</div>
