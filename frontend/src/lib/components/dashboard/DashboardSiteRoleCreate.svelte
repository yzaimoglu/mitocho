<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	import Input from '@/components/ui/input/input.svelte';
	import Label from '@/components/ui/label/label.svelte';
	import MitochoTooltip from '../MitochoTooltip.svelte';
	import MitochoDialog from '../MitochoDialog.svelte';
	import DashboardSiteRolePermissions from './DashboardSiteRolePermissions.svelte';

	type SiteRoleCreateForm = {
		name: string;
		readable_name: string;
		permissions: string[];
	};

	let roleCreateForm: SiteRoleCreateForm = {
		name: '',
		readable_name: '',
		permissions: []
	};

	const formSubmit = async () => {
		console.log(roleCreateForm);
	};
</script>

<form on:submit|preventDefault={formSubmit} class="flex flex-col gap-8 w-full xl:w-3/4 2xl:w-3/5">
	<!-- Name -->
	<div class="flex flex-col gap-1">
		<Label for="name_input">Role ID Name</Label>
		<Input bind:value={roleCreateForm.name} id="name_input" type="text" placeholder="admin" />
		<p class="text-sm text-muted-foreground">
			This is the name of the role with which it is identified.
		</p>
	</div>

	<!-- Readable Name -->
	<div class="flex flex-col gap-1">
		<Label for="readable_name_input">Role Name</Label>
		<Input
			bind:value={roleCreateForm.name}
			id="readable_name_input"
			type="text"
			placeholder="Admin"
		/>
		<p class="text-sm text-muted-foreground">
			This is the name of the role that is displayed on your Mitocho instance.
		</p>
	</div>

	<div class="flex flex-col gap-1">
		<MitochoDialog
			title="Manage role permissions"
			description="You can manage the permissions of the role you want to create."
		>
			<div slot="trigger">
				<Button type="button" class="w-full" variant="outline">Manage permissions</Button>
			</div>
			<div slot="content">
				<DashboardSiteRolePermissions bind:role={roleCreateForm} />
			</div>
		</MitochoDialog>
	</div>

	<MitochoTooltip tooltip="Create the role with these parameters">
		<Button class="w-full" type="submit">Create Role</Button>
	</MitochoTooltip>
</form>
