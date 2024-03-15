<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import { getInstallFinish } from '@/api/install';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { loading } from '@/loading/loading';
	import Loading from '@/loading/Loading.svelte';
	import InstallSetup from '@/components/install/InstallSetup.svelte';
	import CenterPage from '@/components/CenterPage.svelte';

	onMount(async () => {
		const res = await getInstallFinish(fetch);
		const finished_state = res.response.finished;
		if (finished_state) {
			goto('/auth/login');
		} else {
			loading.finish();
		}
	});
</script>

<MitochoPage title="Setup" description="Setup your Mitocho instance">
	<Loading>
		<CenterPage>
			<InstallSetup />
		</CenterPage>
	</Loading>
</MitochoPage>
