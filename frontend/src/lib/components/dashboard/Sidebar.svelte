<script context="module" lang="ts">
	export type SidebarOption = {
		name: string;
		active: boolean;
		level?: string;
		href: string;
	};

	export const levelSecondary = 'secondary';
	export const levelTertiary = 'tertiary';
</script>

<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	export let options: SidebarOption[];

	const getClass = (active: boolean, level?: string) => {
		let classes = [`justify-center md:justify-start md:w-2/3`];
		if (active) {
			classes.push('hover:bg-primary');
		}
		if (level === 'secondary') {
			classes.push('text-xs ml-4 w-1/2 h-8');
		} else if (level === 'tertiary') {
			classes.push('w-1/3');
		}
		return classes.join(' ');
	};
</script>

<nav class="flex flex-row flex-wrap justify-center items-center md:flex-nowrap md:justify-start md:items-start gap-3 md:gap-2 md:flex-col md:w-1/4">
	{#each options as option}
		<Button
			href={option.href}
			class={getClass(option.active, option.level)}
			variant={option.active ? 'default' : 'link'}>{option.name}</Button
		>
	{/each}
</nav>
