<script lang="ts">
	import * as Dialog from '@/components/ui/dialog';
	export let title: string;
	export let description = null;
</script>

<Dialog.Root>
	<Dialog.Trigger>
		<slot name="trigger">Open</slot>
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>{title}</Dialog.Title>
			{#if description}
				<Dialog.Description>
					{description}
				</Dialog.Description>
			{/if}
			<slot name="content" />
		</Dialog.Header>
	</Dialog.Content>
</Dialog.Root>
