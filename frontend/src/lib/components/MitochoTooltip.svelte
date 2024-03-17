<script setup lang="ts">
	import * as Tooltip from '@/components/ui/tooltip';
	export let tooltip: string;
</script>

<Tooltip.Root>
	<Tooltip.Trigger>
		<slot />
	</Tooltip.Trigger>
	<Tooltip.Content>
		<p>{tooltip}</p>
	</Tooltip.Content>
</Tooltip.Root>
