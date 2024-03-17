<script lang="ts">
	import CenterPage from '@/components/CenterPage.svelte';
	import LoadingSpinner from '@/components/LoadingSpinner.svelte';
	import { loading } from '@/loading/loading';

	loading.start();
</script>

{#if $loading}
	<CenterPage footer={false}>
		<LoadingSpinner />
	</CenterPage>
{:else}
	<slot />
{/if}
