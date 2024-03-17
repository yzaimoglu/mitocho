<script lang="ts">
	import MitochoFooter from "./MitochoFooter.svelte";

	export let footer = true;
</script>

<div class={`min-h-screen`}>
	<slot />
	{#if footer}
		<MitochoFooter />
	{/if}
</div>
