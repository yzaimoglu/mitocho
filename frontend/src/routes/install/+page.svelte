<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import GetStarted from '@/components/install/GetStarted.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';

	onMount(async () => {
		loading.finish();
	});
</script>

<MitochoPage title="Installation" description="Get started with your Mitocho installation">
	<Loading>
		<GetStarted />
	</Loading>
</MitochoPage>
