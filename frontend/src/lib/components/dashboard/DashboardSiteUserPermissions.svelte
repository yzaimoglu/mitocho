<script lang="ts">
	import { BookmarkX } from 'lucide-svelte';
	import MitochoTooltip from '../MitochoTooltip.svelte';
	import Badge from '../ui/badge/badge.svelte';
	import Button from '../ui/button/button.svelte';
	import ScrollArea from '../ui/scroll-area/scroll-area.svelte';
	import Input from '../ui/input/input.svelte';
	import Label from '../ui/label/label.svelte';

	export let user;

	let permission = '';

	const removePermission = (index: number) => {
		let tmp = user.permissions;
		tmp.splice(index, 1);
		user.permissions = tmp;
		filterDuplicates(user.permissions);
	};

	const addPermission = () => {
		let tmp = user.permissions;
		tmp.push(permission);
		user.permissions = tmp;
		filterDuplicates(user.permissions);
		permission = "";
	};

	const filterDuplicates = (array: any): any => {
		return [...new Set(array)];
	};

	const handleSubmit = () => {
		console.log(user.permissions);
	};
</script>

<ScrollArea class="h-40 mb-6 pt-2">
	<div class="flex flex-col gap-2">
		{#if user.permissions.length <= 0}
			<p>The user has no permissions.</p>
		{/if}
		{#each user.permissions as permission, permissionIndex}
			<div class="flex flex-row gap-2">
				<Badge class="text-xs bg-green-100 hover:bg-green-100 text-slate-700 w-5/6"
					>{permission}</Badge
				>
				<MitochoTooltip tooltip="Remove permission">
					<Button
						class="h-6"
						on:click={() => removePermission(permissionIndex)}
						variant="destructive"
					>
						<BookmarkX size="0.8rem" />
					</Button>
				</MitochoTooltip>
			</div>
		{/each}
	</div>
</ScrollArea>
<div class="flex flex-col gap-2 items-start">
	<Label for="permission_input">Add a new permission</Label>
	<form on:submit|preventDefault={addPermission} class="flex flex-row gap-1 w-full">
		<Input id="permission_input" bind:value={permission} />
		<Button type="submit">Add</Button>
	</form>
	<form on:submit|preventDefault={handleSubmit} class="w-full">
		<Button class="w-full" type="submit">Save Permissions</Button>
	</form>
</div>
