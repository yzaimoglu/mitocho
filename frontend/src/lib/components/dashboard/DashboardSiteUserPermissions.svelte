<script lang="ts">
	import Button from '../ui/button/button.svelte';
	import ScrollArea from '../ui/scroll-area/scroll-area.svelte';
	import Input from '../ui/input/input.svelte';
	import Label from '../ui/label/label.svelte';
	import { toast } from 'svelte-sonner';
	import DashboardSiteUserSinglePermission from './DashboardSiteUserSinglePermission.svelte';

	export let user: any;

	let permissionInput = '';

	const addPermission = () => {
		if (permissionInput == '') {
			toast.error('Please enter a valid permission.');
			return;
		}
		let tmp = user.permissions;
		if (tmp.includes(permissionInput)) {
			toast.error('This user already has this permission.');
			return;
		}
		tmp.push(permissionInput);
		user.permissions = tmp;
		filterDuplicates(user.permissions);
		permissionInput = '';
	};

	const filterDuplicates = (array: any): any => {
		return [...new Set(array)];
	};

	const handleSubmit = () => {
		console.log(user.permissions);
		toast.success('Saved');
	};
</script>

<ScrollArea class="h-40 mb-6 pt-2">
	<div class="flex flex-col gap-2">
		{#if user.permissions.length <= 0}
			<p>The user has no permissions.</p>
		{/if}
		{#each user.permissions as permission, permissionIndex}
			<DashboardSiteUserSinglePermission
				{permission}
				{permissionIndex}
				bind:userPermissions={user.permissions}
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
