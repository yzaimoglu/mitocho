<script lang="ts">
	import LoadingSpinner from '@/components/LoadingSpinner.svelte';
	import { loading } from '@/loading/loading';

	loading.start();
</script>

{#if $loading}
	<LoadingSpinner />
{:else}
	<slot />
{/if}
