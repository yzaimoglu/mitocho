<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import { type SidebarOption } from '@/components/dashboard/Sidebar.svelte';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';

	// Sidebar settings
	const sidebarOptions: SidebarOption[] = [
		{
			name: 'General',
			active: false,
			href: '/dashboard'
		},
		{
			name: 'Settings',
			active: true,
			href: '/dashboard/settings'
		},
		{
			name: 'Sites',
			active: false,
			href: '/dashboard/sites'
		},
	];

	onMount(() => {
		loading.finish();
	});
</script>

<MitochoPage
	title="Dashboard"
	description="Manage your sites connected to your Mitocho instance from here"
>
	<Loading>
		<FullPage>
			<SidebarLayoutPage
				current="/dashboard"
				title="Dashboard"
				description="Monitor registered users and manage your sites"
				innerTitle="Settings"
				innerDescription="Manage your settings from here"
				{sidebarOptions}
			>
			Manage Settings
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
