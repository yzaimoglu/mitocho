<script lang="ts">
	import Button from '../ui/button/button.svelte';
	import ScrollArea from '../ui/scroll-area/scroll-area.svelte';
	import Input from '../ui/input/input.svelte';
	import Label from '../ui/label/label.svelte';
	import { toast } from 'svelte-sonner';
	import DashboardSiteUserSinglePermission from './DashboardSiteUserSinglePermission.svelte';

	export let role: any;

	let permissionInput = '';

	const addPermission = () => {
		if (permissionInput == '') {
			toast.error('Please enter a valid permission.');
			return;
		}
		let tmp = role.permissions;
		if (tmp.includes(permissionInput)) {
			toast.error('This role already has this permission.');
			return;
		}
		tmp.push(permissionInput);
		role.permissions = tmp;
		filterDuplicates(role.permissions);
		permissionInput = '';
	};

	const filterDuplicates = (array: any): any => {
		return [...new Set(array)];
	};

	const handleSubmit = () => {
		console.log(role.permissions);
		toast.success('Saved');
	};
</script>

<ScrollArea class="h-40 mb-6 pt-2">
	<div class="flex flex-col gap-2">
		{#if role.permissions.length <= 0}
			<p>The role has no permissions.</p>
		{/if}
		{#each role.permissions as permission, permissionIndex}
			<DashboardSiteUserSinglePermission
				{permission}
				{permissionIndex}
				bind:userPermissions={role.permissions}
			/>
		{/each}
		<!-- TODO: Maybe iterate over roles and get permissions? -->
	</div>
</ScrollArea>
<div class="flex flex-col gap-2 items-start">
	<Label for="permission_input">Add a new permission</Label>
	<form on:submit|preventDefault={addPermission} class="flex flex-row gap-1 w-full">
		<Input id="permission_input" bind:value={permissionInput} placeholder="mitocho.admin" />
		<Button type="submit">Add</Button>
	</form>
	<form on:submit|preventDefault={handleSubmit} class="w-full">
		<Button class="w-full" type="submit">Save Permissions</Button>
	</form>
</div>
