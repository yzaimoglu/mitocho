<script lang="ts">
	import * as Card from '@/components/ui/card';
	import Separator from './ui/separator/separator.svelte';
	import Sidebar from './dashboard/Sidebar.svelte';
	import type { SidebarOption } from '@/components/dashboard/Sidebar.svelte';
	import Button from './ui/button/button.svelte';
	import { SquareArrowLeft } from 'lucide-svelte';

	export let title: string;
	export let description: string;
	export let sidebarOptions: SidebarOption[];
	export let innerTitle: string;
	export let innerDescription: string;
</script>

<Card.Root class="p-4">
	<Card.Header>
		<Button type="button" on:click={() => window.history.back()} variant="link" class="justify-start w-full pl-0">
			<div class="flex flex-row items-center gap-1">
				<SquareArrowLeft size="16px" />
				<p>Back</p>
			</div>
		</Button>
		<h1 class="text-2xl font-bold">{title}</h1>
		<h2 class="text-lg">
			{description}
		</h2>
		<Separator class="my-4" />
	</Card.Header>
	<Card.Content>
		<div class="flex flex-col md:flex-row">
			<Sidebar options={sidebarOptions} />
			<div class="md:w-3/4">
				<h1 class="mt-4 md:mt-0 text-lg font-bold">{innerTitle}</h1>
				<h2 class="text-md">{innerDescription}</h2>
				<Separator class="mt-2 mb-8" />
				<slot />
			</div>
		</div>
	</Card.Content>
</Card.Root>
