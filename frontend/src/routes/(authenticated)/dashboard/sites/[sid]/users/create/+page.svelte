<script lang="ts">
	import { page } from '$app/stores';
	import MitochoPage from '@/components/MitochoPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
  
  const sid = $page.params.sid;
	
	onMount(() => {
		console.log(`Site ID: ${sid}`)
		loading.finish();
	});
</script>

<MitochoPage>
	<Loading>
		<div>Dashboard Sites Users Create</div>
	</Loading>
</MitochoPage>
