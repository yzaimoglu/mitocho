<script lang="ts" context="module">
	export type Role = {
		id: string;
		name: string;
		readable_name: string;
		permissions: string[];
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
	import DashboardSiteRolePermissions from './DashboardSiteRolePermissions.svelte';

	export let role;
	export let sid;
</script>

<Card.Root>
	<Card.Header>
		<div class="flex flex-row justify-between">
			<div class="flex flex-col gap-1">
				<h1 class="text-sm font-bold">{role.readable_name}</h1>
				<Badge
					class="text-xs justify-center items-center font-semibold text-slate-800 bg-green-100 hover:bg-green-100"
					>{role.id}</Badge
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
					title="Role permissions"
					description={`Change the permissions of ${role.username}`}
				>
					<div slot="trigger">
						<Button class="md:justify-start w-1/3 md:w-full" size="sm" variant="link"
							>Permissions</Button
						>
					</div>
					<div slot="content">
						<DashboardSiteRolePermissions {role} />
					</div>
				</MitochoDialog>
			</div>
			<MitochoTooltip tooltip="Remove this role">
				<Button size="sm" variant="destructive">Remove</Button>
			</MitochoTooltip>
		</div>
	</Card.Content>
</Card.Root>
