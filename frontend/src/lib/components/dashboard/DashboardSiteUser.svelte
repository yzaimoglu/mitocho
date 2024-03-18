<script lang="ts" context="module">
	export type User = {
		id: string;
		username: string;
		email: string;
		permissions: string[];
		roles: string[];
	};
</script>

<script lang="ts">
	import * as Card from '@/components/ui/card';
	import MitochoLogo from '../MitochoLogo.svelte';
	import Button from '../ui/button/button.svelte';
	import Separator from '../ui/separator/separator.svelte';
	import Badge from '../ui/badge/badge.svelte';
	import MitochoTooltip from '../MitochoTooltip.svelte';
	import MitochoDialog from '@/components/MitochoDialog.svelte';
	import DashboardSiteUserPermissions from './DashboardSiteUserPermissions.svelte';
	import DashboardSiteUserRoles from './DashboardSiteUserRoles.svelte';

	export let user;
	export let sid;
</script>

<Card.Root>
	<Card.Header>
		<div class="flex flex-row justify-between">
			<div class="flex flex-col gap-1">
				<h1 class="text-sm font-bold">{user.username}</h1>
				<Badge
					class="text-xs justify-center items-center font-semibold text-slate-800 bg-green-100 hover:bg-green-100"
					>{user.id}</Badge
				>
			</div>
			<MitochoLogo width="w-12" />
		</div>
		<Separator class="my-2" />
	</Card.Header>
	<Card.Content>
		<div class="flex flex-col md:flex-row justify-between gap-1">
			<div class="flex flex-row justify-center md:justify-between gap-6 md:gap-1">
				<MitochoDialog
					title="User permissions"
					description={`Change the permissions of ${user.username}`}
				>
					<div slot="trigger">
						<Button class="md:justify-start w-1/3 md:w-full" size="sm" variant="link"
							>Permissions</Button
						>
					</div>
					<div slot="content">
						<DashboardSiteUserPermissions {user} />
					</div>
				</MitochoDialog>
				<MitochoDialog
					title="User roles"
					description={`Change the roles of ${user.username}`}
				>
					<div slot="trigger">
						<Button class="md:justify-start w-1/3 md:w-full" size="sm" variant="link"
							>Roles</Button
						>
					</div>
					<div slot="content">
						<DashboardSiteUserRoles {user} />
					</div>
				</MitochoDialog>
			</div>
			<MitochoTooltip tooltip="Remove this user">
				<Button size="sm" variant="destructive">Remove</Button>
			</MitochoTooltip>
		</div>
	</Card.Content>
</Card.Root>
