<script lang="ts">
	import { ShieldPlus } from "lucide-svelte";
	import MitochoTooltip from "../MitochoTooltip.svelte";
	import Button from "../ui/button/button.svelte";
	import Input from "../ui/input/input.svelte";
	import Label from "../ui/label/label.svelte";

	export let search = "";

	const handleSearch = async () => {
		console.log(`searching for ${search}...`);
	};
</script>

<form on:submit|preventDefault={handleSearch} class="flex flex-row gap-2 items-end mb-4">
	<div class="flex flex-col w-full gap-2">
		<Label for="site_search_input">Search for a specific site or create one</Label>
		<Input
			id="site_search_input"
			class="w-full"
			bind:value={search}
			type="text"
			placeholder="mitocho"
		/>
	</div>
	<MitochoTooltip tooltip="Search for an existing site">
		<Button type="submit">Search</Button>
	</MitochoTooltip>
	<MitochoTooltip tooltip="Create a new site">
		<Button href="/dashboard/sites/create">
			<ShieldPlus />
		</Button>
	</MitochoTooltip>
</form>
