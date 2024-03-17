<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import type { SidebarOption } from '@/components/dashboard/Sidebar.svelte';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';
	import DashboardSiteCreate from '@/components/dashboard/DashboardSiteCreate.svelte';

	// Sidebar settings
	const sidebarOptions: SidebarOption[] = [
		{
			name: 'General',
			active: false,
			href: '/dashboard'
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

<MitochoPage
	title="Create a new Site"
	description="Create a new site from here"
>
	<Loading>
		<FullPage>
			<SidebarLayoutPage
				current="/dashboard"
				title="Dashboard"
				description="Monitor registered users and manage your sites"
				innerTitle="Create Site"
				innerDescription="Create a new site you want to connect with your mitocho instance"
				{sidebarOptions}
			>
				<DashboardSiteCreate />
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
