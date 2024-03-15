<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import Home from '@/components/home/Home.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';

	onMount(async () => {
		loading.finish();
	});
</script>

<MitochoPage title="Mitocho">
	<Loading>
		<Home />
	</Loading>
</MitochoPage>
