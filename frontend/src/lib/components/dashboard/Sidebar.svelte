<script context="module" lang="ts">
	export type SidebarOption = {
		name: string;
		active: boolean;
		href: string;
	};
</script>

<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	export let options: SidebarOption[];
</script>

<nav class="flex flex-row gap-3 md:gap-2 md:flex-col md:w-1/4">
	{#each options as option}
		{#if option.active}
			<Button href={option.href} class="justify-center md:justify-start hover:bg-primary w-2/3">{option.name}</Button>
		{:else}
			<Button href={option.href} class="justify-center md:justify-start w-2/3" variant="link">{option.name}</Button>
		{/if}
	{/each}
</nav>
