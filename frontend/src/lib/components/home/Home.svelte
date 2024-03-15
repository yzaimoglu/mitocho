<script lang="ts">
	import Button from '../ui/button/button.svelte';
</script>

<div class="flex flex-col justify-center items-center h-screen gap-2">
	<a href="/">
		<img class="w-32" alt="Mitocho Logo" src="/assets/logo.png" />
	</a>
	<div class="text-center">
		<h1 class="text-2xl font-bold">Mitocho</h1>
		<p class="text-lg">The powerhouse of authentication</p>
	</div>
	<Button href="/auth/login">Login</Button>
</div>
