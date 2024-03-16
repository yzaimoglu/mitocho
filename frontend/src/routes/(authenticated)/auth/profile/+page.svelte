<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import MainNavbar from '@/components/MainNavbar.svelte';
	import SidebarLayout from '@/components/SidebarLayout.svelte';
	import type { SidebarOption } from '@/components/dashboard/Sidebar.svelte';

	import { page } from '$app/stores';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';

	const sid = $page.url.searchParams.get('sid');
	const red = $page.url.searchParams.get('red');

	const sidebarOptions: SidebarOption[] = [
		{
			name: 'General',
			active: true,
			href: '/auth/profile/general'
		},
		{
			name: 'Password',
			active: false,
			href: '/auth/profile/password'
		},
		{
			name: '2FA',
			active: false,
			href: '/auth/profile/2fa'
		}
	];

	onMount(() => {
		console.log(`Site ID: ${sid}`);
		console.log(`Redirect URL: ${red}`);
		loading.finish();
	});
</script>

<MitochoPage title="Profile" description="Manage your profile from here">
	<Loading>
		<FullPage>
			<SidebarLayoutPage
				current="/auth/profile"
				title="Profile"
				description="Change your username, email, password or enable two factor authentication"
				innerTitle="General"
				innerDescription="Manage your general user settings"
				{sidebarOptions}
			>
			<div>Test</div>
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
