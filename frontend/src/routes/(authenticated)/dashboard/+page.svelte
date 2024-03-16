<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import MainNavbar from '@/components/MainNavbar.svelte';
	import SidebarLayout from '@/components/SidebarLayout.svelte';
	import type { SidebarOption } from '@/components/dashboard/Sidebar.svelte';

	const sidebarOptions: SidebarOption[] = [
		{
			name: 'General',
			active: true,
			href: '/auth/profile/general'
		},
		{
			name: 'Sites',
			active: false,
			href: '/auth/profile/password'
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
			<div class="flex flex-col items-center">
				<div class="flex flex-col gap-2 mt-2 w-4/5">
					<MainNavbar current="/dashboard" />
					<SidebarLayout
						title="Dashboard"
						description="Monitor registered users and manage your sites"
						innerTitle="General"
						innerDescription="Usage information regarding your Mitocho instance"
						{sidebarOptions}
					>
						<div>Test</div>
					</SidebarLayout>
				</div>
			</div>
		</FullPage>
	</Loading>
</MitochoPage>
