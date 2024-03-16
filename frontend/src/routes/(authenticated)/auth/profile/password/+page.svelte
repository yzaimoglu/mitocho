<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import FullPage from '@/components/FullPage.svelte';
	import Loading from '@/loading/Loading.svelte';
	import { loading } from '@/loading/loading';
	import { onMount } from 'svelte';
	import type { SidebarOption } from '@/components/dashboard/Sidebar.svelte';

	import { page } from '$app/stores';
	import SidebarLayoutPage from '@/components/SidebarLayoutPage.svelte';
	import ProfilePasswordForm from '@/components/profile/ProfilePasswordForm.svelte';

	let sid = $page.url.searchParams.get('sid');
	let red = $page.url.searchParams.get('red');

	if (sid === null) sid = 'mitocho';
	if (red === null) red = 'mitocho';

	const sidebarOptions: SidebarOption[] = [
		{
			name: 'General',
			active: false,
			href: `/auth/profile?sid=${sid}&red=${red}`
		},
		{
			name: 'Password',
			active: true,
			href: `/auth/profile/password?sid=${sid}&red=${red}`
		},
		{
			name: '2FA',
			active: false,
			href: `/auth/profile/2fa?sid=${sid}&red=${red}`
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
				innerTitle="Password"
				innerDescription="Change your password"
				{sidebarOptions}
				{red}
			>
				<ProfilePasswordForm {sid} {red} />
			</SidebarLayoutPage>
		</FullPage>
	</Loading>
</MitochoPage>
