<script lang="ts">
	import * as Card from '@/components/ui/card';
	import MitochoLogo from '../MitochoLogo.svelte';
	import Button from '../ui/button/button.svelte';
	import Separator from '../ui/separator/separator.svelte';
	import Badge from '../ui/badge/badge.svelte';

	export let name = 'Mitocho';
	export let sid = '12345-12345-fefefefe-123244';
</script>

<Card.Root>
	<Card.Header>
		<div class="flex flex-row justify-between">
			<div class="flex flex-col gap-1">
				<h1 class="text-sm font-bold">{name}</h1>
				<Badge
					class="text-xs justify-center items-center font-semibold text-slate-800 bg-green-100 hover:bg-green-100"
					>{sid}</Badge
				>
			</div>
			<MitochoLogo width="w-12" />
		</div>
		<Separator class="my-2" />
	</Card.Header>
	<Card.Content>
		<div class="flex flex-col md:flex-row justify-between gap-1">
			<div class="flex flex-row justify-between gap-1">
				<Button href={`/dashboard/sites/${sid}`} class="w-1/3 md:w-full" size="sm">Settings</Button>

				<Button href={`/dashboard/sites/${sid}/users`} class="w-1/3 md:w-full" size="sm"
					>Users</Button
				>
				<Button href={`/dashboard/sites/${sid}/roles`} class="w-1/3 md:w-full" size="sm"
					>Roles</Button
				>
			</div>
			<Button size="sm" variant="destructive">Remove</Button>
		</div>
	</Card.Content>
</Card.Root>
