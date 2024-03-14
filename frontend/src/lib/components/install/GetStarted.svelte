<script lang="ts">
	// import { onMount } from 'svelte';
	import Button from '../ui/button/button.svelte';
	// import { postPublicKey } from '@/api/site';

	// onMount(async () => {
	// 	const test = await postPublicKey(fetch, {
	// 		"access_token": "822a6512f869505f9348266f737197a22e0679ae118f0275a290543ce71ce69762f28c7671ad0230ef9a1d5edaaf2a798d98bbd2e57751890b0ddf06ec6113d7",
	// 	});
	// })
</script>

<div class="flex flex-col justify-center items-center h-screen gap-2">
	<a href="/">
		<img class="w-32" alt="Mitocho Logo" src="/assets/logo.png" />
	</a>
	<div class="text-center">
		<h1 class="text-2xl font-bold">Mitocho</h1>
		<p class="text-lg">The powerhouse of authentication</p>
	</div>
	<Button href="/install/setup">Get Started</Button>
</div>
