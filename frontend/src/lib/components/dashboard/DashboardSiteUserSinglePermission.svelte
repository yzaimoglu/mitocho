<script lang="ts">
	import { BookmarkX } from 'lucide-svelte';
	import MitochoTooltip from '../MitochoTooltip.svelte';
	import Badge from '../ui/badge/badge.svelte';
	import Button from '../ui/button/button.svelte';

	export let userPermissions: any;
	export let permission: string;
	export let permissionIndex: number;

	const removePermission = (index: number) => {
		let tmp = userPermissions;
		tmp.splice(index, 1);
		userPermissions = tmp;
	};
</script>

<div class="flex flex-row gap-2">
	<Badge variant="outline" class="text-xs w-5/6">{permission}</Badge>
	<MitochoTooltip tooltip="Remove permission">
		<Button class="h-6" on:click={() => removePermission(permissionIndex)} variant="destructive">
			<BookmarkX size="0.8rem" />
		</Button>
	</MitochoTooltip>
</div>
