<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import CenterPage from '@/components/CenterPage.svelte';
	import LoginOtpCard from '@/components/auth/LoginOtpCard.svelte';
  import { loginStore } from '@/stores/loginStore';
	import { goto } from '$app/navigation';

	const sid = $page.url.searchParams.get('sid');
	const red = $page.url.searchParams.get('red');

	onMount(() => {
    if(loginStore.isEmpty()) {
      goto("/");
    }
		console.log(`Site ID: ${sid}`);
		console.log(`Redirect URL: ${red}`);
		loading.finish();
	});
</script>

<MitochoPage>
	<Loading>
		<CenterPage>
			<LoginOtpCard email={$loginStore.email} password={$loginStore.password} {sid} {red} />
		</CenterPage>
	</Loading>
</MitochoPage>
