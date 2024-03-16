<script lang="ts">
	import { onMount } from 'svelte';
	import LoadingSpinner from '../LoadingSpinner.svelte';
	import { goto } from '$app/navigation';

	onMount(() => {
		setTimeout(() => {
			goto('/auth/login');
		}, 1000);
	});
</script>

<div class="flex flex-col justify-center items-center h-screen gap-4">
	<a href="/">
		<img class="w-32" alt="Mitocho Logo" src="/assets/logo.png" />
	</a>
	<div class="text-center">
		<h1 class="text-2xl font-bold">Mitocho</h1>
		<p class="text-lg">The powerhouse of authentication</p>
	</div>
	<LoadingSpinner />
</div>
