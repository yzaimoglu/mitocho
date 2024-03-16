<script lang="ts">
	import * as Card from '@/components/ui/card';
	import MitochoLogo from './MitochoLogo.svelte';
	import Button from './ui/button/button.svelte';
	import { onMount } from 'svelte';

	type NavbarOption = {
		name: string;
		active: boolean;
		admin: boolean;
		href: string;
	};

	export let current: string;
	export let red: string = "mitocho";
	export let options: NavbarOption[] = [
		{
			name: 'Dashboard',
			admin: true,
			active: true,
			href: '/dashboard'
		},
		{
			name: 'Profile',
			admin: true,
			active: false,
			href: '/auth/profile'
		},
		{
			name: 'Logout',
			admin: true,
			active: false,
			href: '/auth/logout'
		}
	];

	onMount(() => {
		for (let i = 0; i < options.length; i++) {
			if (current === options[i].href) {
				options[i].active = true;
			} else {
				options[i].active = false;
			}
		}
	});
</script>

<Card.Root class="p-3">
	<Card.Content class="p-0">
		<nav class="flex flex-row justify-between">
			<div class="flex flex-row gap-2">
				{#if red !== 'mitocho'}
					<MitochoLogo href={red} width="w-8" />
					<Button href={red} variant="link">Back to App</Button>
				{:else}
					<MitochoLogo href="/dashboard" width="w-8" />
				{/if}
			</div>
			<div class="flex flex-row gap-2">
				{#each options as option}
					{#if option.active}
						<Button href={option.href} class="bg-primary">{option.name}</Button>
					{:else}
						<Button href={option.href} variant="link">{option.name}</Button>
					{/if}
				{/each}
			</div>
		</nav>
	</Card.Content>
</Card.Root>
