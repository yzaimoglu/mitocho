<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import { type SidebarOption } from '@/components/dashboard/Sidebar.svelte';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';
	import { page } from '$app/stores';
	import DashboardSiteUserCreate from '@/components/dashboard/DashboardSiteUserCreate.svelte';

	const sid = $page.params.sid;

	// Sidebar settings
	const sidebarOptions: SidebarOption[] = [
		{
			name: 'General',
			active: false,
			href: '/dashboard'
		},
		{
			name: 'Settings',
			active: false,
			href: '/dashboard/settings'
		},
		{
			name: 'Sites',
			active: true,
			href: '/dashboard/sites'
		}
	];

	onMount(() => {
		loading.finish();
	});
</script>

<MitochoPage title="Create a new Role" description="Create a new role from here">
	<Loading>
		<FullPage>
			<SidebarLayoutPage
				current="/dashboard"
				title="Dashboard"
				description="Monitor registered users and manage your sites"
				innerTitle="Create User"
				innerDescription={`${sid}`}
				{sidebarOptions}
			>
				<DashboardSiteUserCreate />
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
