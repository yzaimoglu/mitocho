<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import type { SidebarOption } from '@/components/dashboard/Sidebar.svelte';

	import { page } from '$app/stores';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';
	import Profile2FaCreate from '@/components/profile/Profile2FACreate.svelte';
	import Profile2FaManage from '@/components/profile/Profile2FAManage.svelte';
	import Button from '@/components/ui/button/button.svelte';

	let sid = $page.url.searchParams.get('sid');
	let red = $page.url.searchParams.get('red');

	if (sid === null) sid = 'mitocho';
	if (red === null) red = 'mitocho';

	const sidebarOptions: SidebarOption[] = [
		{
			name: 'General',
			active: false,
			href: `/auth/profile?sid=${sid}&red=${red}`
		},
		{
			name: 'Password',
			active: false,
			href: `/auth/profile/password?sid=${sid}&red=${red}`
		},
		{
			name: '2FA',
			active: true,
			href: `/auth/profile/2fa?sid=${sid}&red=${red}`
		}
	];

	onMount(() => {
		console.log(`Site ID: ${sid}`);
		console.log(`Redirect URL: ${red}`);
		loading.finish();
	});
	let current = true;
</script>

<MitochoPage title="Profile" description="Manage your profile from here">
	<Loading>
		<FullPage>
			<SidebarLayoutPage
				current="/auth/profile"
				title="Profile"
				description="Change your username, email, password or enable two factor authentication"
				innerTitle="Two Factor Authentication"
				innerDescription="Manage your two factor authentication"
				{sidebarOptions}
				{red}
			>
				<Button class="mb-4" on:click={() => current = !current}>Change state</Button>
				{#if current}
					<Profile2FaCreate {sid} {red} />
				{:else}
					<Profile2FaManage {sid} {red} />
				{/if}
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
