<script lang="ts">
	export let href: string = undefined;
	export let width: string = "w-32";
</script>

{#if href !== undefined}
<a href={href}>
	<img class={width} alt="Mitocho Logo" src="/assets/logo.png" />
</a>
{:else}
	<img class={width} alt="Mitocho Logo" src="/assets/logo.png" />
{/if}
