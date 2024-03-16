<script lang="ts">
	import * as Card from '@/components/ui/card';
	import MitochoLogo from '../MitochoLogo.svelte';
	import Button from '../ui/button/button.svelte';
	import Separator from '../ui/separator/separator.svelte';

	export let name = 'Mitocho';
	export let domains = ['mitocho.io', 'auth.mitocho.io'];
	export let sid = '12345-12345-fefefefe-123244';
</script>

<Card.Root>
	<Card.Header>
		<div class="flex flex-row justify-between">
			<div class="flex flex-col">
				<h1 class="text-sm font-bold">{name}</h1>
				<p class="text-xs italic font-semibold">{sid}</p>
			</div>
			<MitochoLogo width="w-12" />
		</div>
		<Separator class="my-2" />
	</Card.Header>
	<Card.Content>
		<div class="flex flex-col md:flex-row justify-between gap-1">
			<div class="flex flex-row justify-between gap-1">
				<Button class="w-1/3 md:w-full" size="sm">Users</Button>
				<Button class="w-1/3 md:w-full" size="sm">Roles</Button>
				<Button class="w-1/3 md:w-full" size="sm">Settings</Button>
			</div>
			<Button size="sm" variant="destructive">Remove</Button>
		</div>
	</Card.Content>
</Card.Root>
