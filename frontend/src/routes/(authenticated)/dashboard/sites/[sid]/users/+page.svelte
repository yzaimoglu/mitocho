<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import { levelSecondary, type SidebarOption } from '@/components/dashboard/Sidebar.svelte';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';
	import DashboardSite from '@/components/dashboard/DashboardSite.svelte';
	import { page as storePage } from '$app/stores';
	import MitochoPagination from '@/components/MitochoPagination.svelte';
	import DashboardSiteSearch from '@/components/dashboard/DashboardSiteSearch.svelte';
	import DashboardSiteUserSearch from '@/components/dashboard/DashboardSiteUserSearch.svelte';
	import DashboardSiteUser, { type User } from '@/components/dashboard/DashboardSiteUser.svelte';

	// Site ID
	const sid = $storePage.params.sid;

	// Pagination
	let currentPage = $storePage.url.searchParams.get('page');
	if (currentPage === null) currentPage = '1';
	const changePage = async (page: Number) => {
		console.log(`Changed to page ${page}`);
	};

	// Search
	let search = '';

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

	// Users
	let users: User[] = [
		{
			email: 'yz@yagi.sh',
			id: '12345-12345-fefefefe-12345',
			permissions: ['*', '1', '2', '3', '4', '5', '6', '7'],
			roles: ['admin'],
			username: 'yz'
		}
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
				<DashboardSiteUserSearch {sid} bind:search />
				<div class="grid grid-cols-1 xl:grid-cols-2 gap-2 mb-4">
					{#each users as user}
						<DashboardSiteUser {sid} {user} />
					{/each}
				</div>
				{#if users.length > 0}
					<MitochoPagination
						bind:currentPage
						changePageFunction={changePage}
						count={100}
						siblingCount={1}
						perPage={users.length}
					/>
				{/if}
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
