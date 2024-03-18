<script lang="ts">
	import { BookmarkX } from 'lucide-svelte';
	import MitochoTooltip from '../MitochoTooltip.svelte';
	import Badge from '../ui/badge/badge.svelte';
	import Button from '../ui/button/button.svelte';

	export let userRoles: any;
	export let role: string;
	export let roleIndex: number;

	const removeRole = (index: number) => {
		let tmp = userRoles;
		tmp.splice(index, 1);
		userRoles = tmp;
	};
</script>

<div class="flex flex-row gap-2">
	<Badge variant="outline" class="text-xs w-5/6">{role}</Badge>
	<MitochoTooltip tooltip="Remove role">
		<Button class="h-6" on:click={() => removeRole(roleIndex)} variant="destructive">
			<BookmarkX size="0.8rem" />
		</Button>
	</MitochoTooltip>
</div>
