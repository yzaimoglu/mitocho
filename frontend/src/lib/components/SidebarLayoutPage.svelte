<script lang="ts">
	import MainNavbar from "./MainNavbar.svelte";
	import SidebarLayout from "./SidebarLayout.svelte";

	import type { SidebarOption } from '@/components/dashboard/Sidebar.svelte';

	export let current: string;
	export let red: string = "mitocho";
	export let title: string;
	export let description: string;
	export let sidebarOptions: SidebarOption[];
	export let innerTitle: string;
	export let innerDescription: string;
</script>

<div class="flex flex-col items-center">
	<div class="flex flex-col gap-2 my-3 w-5/6 2xl:w-3/5">
		<MainNavbar {current} {red} />
		<SidebarLayout
			{title}
			{description}
			{sidebarOptions}
			{innerTitle}
			{innerDescription}
		>
			<slot />
		</SidebarLayout>
	</div>
</div>
