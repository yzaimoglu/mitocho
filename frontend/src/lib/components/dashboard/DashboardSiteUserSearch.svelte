<script lang="ts">
	import { ShieldPlus } from 'lucide-svelte';
	import MitochoTooltip from '../MitochoTooltip.svelte';
	import Button from '../ui/button/button.svelte';
	import Input from '../ui/input/input.svelte';
	import Label from '../ui/label/label.svelte';

	export let sid;
	export let search = '';

	const handleSearch = async () => {
		console.log(`searching for ${search}...`);
	};
</script>

<form on:submit|preventDefault={handleSearch} class="flex flex-col gap-2">
	<Label for="site_search_input">Search for a specific user or create one</Label>
	<div class="flex flex-row items-end mb-4 gap-2">
		<div class="flex flex-col w-full gap-2">
			<Input
				id="site_search_input"
				class="w-full"
				bind:value={search}
				type="text"
				placeholder="admin"
			/>
		</div>
		<MitochoTooltip tooltip="Search for an existing user">
			<Button type="submit">Search</Button>
		</MitochoTooltip>
		<MitochoTooltip tooltip="Create a new user">
			<Button href={`/dashboard/sites/${sid}/users/create`}>
				<ShieldPlus />
			</Button>
		</MitochoTooltip>
	</div>
</form>
