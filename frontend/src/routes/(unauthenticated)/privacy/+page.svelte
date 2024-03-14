<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';

	onMount(() => {
		loading.finish();
	});
</script>

<MitochoPage>
	<Loading>
		<div>Privacy Policy</div>
	</Loading>
</MitochoPage>
