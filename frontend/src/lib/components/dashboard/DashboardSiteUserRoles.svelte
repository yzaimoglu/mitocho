<script lang="ts">
	import Button from '../ui/button/button.svelte';
	import ScrollArea from '../ui/scroll-area/scroll-area.svelte';
	import Label from '../ui/label/label.svelte';
	import { toast } from 'svelte-sonner';
	import DashboardSiteUserSinglePermission from './DashboardSiteUserSinglePermission.svelte';
	import MitochoSelect, { type SelectItem } from '../MitochoSelect.svelte';
	import DashboardSiteUserSingleRole from './DashboardSiteUserSingleRole.svelte';

	export let user: any;

	const roles: SelectItem[] = [
		{
			text: 'Admin',
			value: 'admin'
		},
		{
			text: 'Moderator',
			value: 'mod'
		},
		{
			text: 'User',
			value: 'user'
		}
	];
	let roleInput = '';

	const addRole = () => {
		if (roleInput == '') {
			toast.error('Please choose a valid role.');
			return;
		}
		let tmp = user.roles;
		if (tmp.includes(roleInput)) {
			toast.error('This user already has this role.');
			return;
		}
		tmp.push(roleInput);
		user.roles = tmp;
		filterDuplicates(user.roles);
	};

	const filterDuplicates = (array: any): any => {
		return [...new Set(array)];
	};

	const handleSubmit = () => {
		console.log(user.roles);
		toast.success('Saved');
	};
</script>

<ScrollArea class="h-40 mb-6 pt-2">
	<div class="flex flex-col gap-2">
		{#if user.roles.length <= 0}
			<p>The user has no roles.</p>
		{/if}
		{#each user.roles as role, roleIndex}
			<DashboardSiteUserSingleRole bind:userRoles={user.roles} {role} {roleIndex} />
		{/each}
		<!-- TODO: Maybe iterate over roles and get permissions? -->
	</div>
</ScrollArea>
<div class="flex flex-col gap-2 items-start">
	<Label>Add a new role</Label>
	<form on:submit|preventDefault={addRole} class="flex flex-row gap-1 w-full">
		<MitochoSelect bind:selected={roleInput} placeholder="Roles" label="Roles" items={roles} />
		<Button type="submit">Add</Button>
	</form>
	<form on:submit|preventDefault={handleSubmit} class="w-full">
		<Button class="w-full" type="submit">Save Roles</Button>
	</form>
</div>
