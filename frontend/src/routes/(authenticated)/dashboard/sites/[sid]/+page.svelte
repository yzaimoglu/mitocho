<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import { levelSecondary, type SidebarOption } from '@/components/dashboard/Sidebar.svelte';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';
	import DashboardSiteSettings from '@/components/dashboard/DashboardSiteSettings.svelte';
	import { page } from '$app/stores';

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

<MitochoPage title="Site Settings" description="Manage your site settings">
	<Loading>
		<FullPage>
			<SidebarLayoutPage
				current="/dashboard"
				title="Dashboard"
				description="Monitor registered users and manage your sites"
				innerTitle="Site Settings"
				innerDescription={sid}
				{sidebarOptions}
			>
				<DashboardSiteSettings {sid} />
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
