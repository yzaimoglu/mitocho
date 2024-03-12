<script lang="ts">
	import MitochoPage from '@/components/MitochoPage.svelte';
	import { page } from '$app/stores';
	import Button from '@/components/ui/button/button.svelte';
</script>

<MitochoPage color="error" title="{$page.error.message}" description="An error occured, please contact an administrator.">
	<div class="flex flex-col justify-center items-center h-screen gap-2">
		<a href="/">
			<img class="w-32" alt="Mitocho Logo" src="/assets/logo.png" />
		</a>
		<div class="text-center">
			<h1 class="text-2xl font-bold">Error {$page.status}</h1>
			<p class="text-lg">{$page.error.message}</p>
		</div>
		<Button href="/">Home</Button>
	</div>
</MitochoPage>
