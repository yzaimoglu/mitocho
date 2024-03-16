<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import type { SidebarOption } from '@/components/dashboard/Sidebar.svelte';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';
	import DashboardSite from '@/components/dashboard/DashboardSite.svelte';
	import Input from '@/components/ui/input/input.svelte';
	import Label from '@/components/ui/label/label.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import { page as storePage } from '$app/stores';
	import MitochoPagination from '@/components/MitochoPagination.svelte';

	let currentPage = $storePage.url.searchParams.get('page');
	if (currentPage === null) currentPage = '1';
	let x = 0;

	const changePage = async (page: Number) => {
		console.log(`Changed to page ${page}`);
	};

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

	let dashboardSites = [
		{
			name: 'Object1',
			domains: ['domain1.com', 'domain2.com'],
			sid: '12345-fewfefeff-12312341-fefef'
		},
		{ name: 'Object2', domains: ['domain3.com', 'domain4.com'], sid: 'SID2' },
		{ name: 'Object3', domains: ['domain5.com', 'domain6.com'], sid: 'SID3' },
		{ name: 'Object4', domains: ['domain7.com', 'domain8.com'], sid: 'SID4' },
		{ name: 'Object5', domains: ['domain9.com', 'domain10.com'], sid: 'SID5' }
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
				innerTitle="Sites"
				innerDescription="Manage your sites from here"
				{sidebarOptions}
			>
				<form class="flex flex-row gap-2 items-end mb-4">
					<div class="flex flex-col w-full gap-2">
						<Label for="site_search_input">Search for a specific site</Label>
						<Input id="site_search_input" class="w-full" type="text" placeholder="mitocho" />
					</div>
					<Button>Search</Button>
				</form>

				<div class="flex flex-col gap-2 mb-4">
					{#each dashboardSites as site}
						<DashboardSite name={site.name} sid={site.sid} />
					{/each}
				</div>
				<MitochoPagination bind:currentPage changePageFunction={changePage} count={200} siblingCount={1} perPage={10} />
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
