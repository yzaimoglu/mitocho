<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import LoginCard from '@/components/auth/LoginCard.svelte';
	import CenterPage from '@/components/CenterPage.svelte';

	let sid = $page.url.searchParams.get('sid');
	let red = $page.url.searchParams.get('red');

	onMount(() => {
		console.log(`Site ID: ${sid}`);
		console.log(`Redirect URL: ${red}`);
		if (sid === null) sid = 'mitocho';
		if (red === null) red = 'mitocho';
		loading.finish();
	});
</script>

<MitochoPage title="Login">
	<Loading>
		<CenterPage>
			<LoginCard {sid} {red} />
		</CenterPage>
	</Loading>
</MitochoPage>
