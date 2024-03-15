<script lang="ts">
	import CenterPage from '@/components/CenterPage.svelte';
	import MitochoPage from '@/components/MitochoPage.svelte';
	import ForgotPasswordCard from '@/components/auth/ForgotPasswordCard.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';

	onMount(() => {
		loading.finish()
	});
</script>

<MitochoPage
	title="Password Reset"
	description="Reset your password to gain access to your account"
>
	<Loading>
		<CenterPage>
			<ForgotPasswordCard />
		</CenterPage>
	</Loading>
</MitochoPage>
