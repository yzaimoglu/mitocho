<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import { type SidebarOption } from '@/components/dashboard/Sidebar.svelte';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';
	import { page as storePage } from '$app/stores';
	import DashboardSiteRole from '@/components/dashboard/DashboardSiteRole.svelte';
	import { type Role } from '@/components/dashboard/DashboardSiteRole.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import MitochoTooltip from '@/components/MitochoTooltip.svelte';

	// Site ID
	const sid = $storePage.params.sid;

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
	let roles: Role[] = [
		{
			id: '12345-12345-fefefefe-12345',
			name: 'admin',
			readable_name: 'Admin',
			permissions: ['*', '1', '2', '3', '4', '5', '6', '7']
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
				innerTitle="Site Roles"
				innerDescription={sid}
				{sidebarOptions}
			>
				<div class="flex flex-col gap-2">
					<MitochoTooltip tooltip="Create a new role">
						<Button href="roles/create" class="w-full">Create a new role</Button>
					</MitochoTooltip>
					{#if roles.length === 0}
						<p>This site has no roles.</p>
					{/if}
					<div class="grid grid-cols-1 xl:grid-cols-2 gap-2 mb-4">
						{#each roles as role}
							<DashboardSiteRole {sid} {role} />
						{/each}
					</div>
				</div>
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
