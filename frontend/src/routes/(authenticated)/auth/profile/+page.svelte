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
			<div class="flex flex-col items-center">
				<div class="flex flex-col gap-2 mt-2 w-4/5">
					<MainNavbar current="/auth/profile" />
					<SidebarLayout
						title="Profile"
						description="Change your email, username, password or enable two factor authentication"
						innerTitle="General"
						innerDescription="Change your email or username"
						{sidebarOptions}
					>
						<div>Form</div>
					</SidebarLayout>
				</div>
			</div>
		</FullPage>
	</Loading>
</MitochoPage>
