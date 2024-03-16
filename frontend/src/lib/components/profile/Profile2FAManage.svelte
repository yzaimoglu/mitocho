<script lang="ts">
	import Button from '../ui/button/button.svelte';
	import * as AlertDialog from '@/components/ui/alert-dialog';

	export let sid = 'mitocho';
	export let red = 'mitocho';

	const formSubmit = async () => {
		console.log(sid);
		console.log(red);
	};
</script>

<form on:submit|preventDefault={formSubmit} class="flex flex-col gap-8 w-full xl:w-3/4 2xl:w-3/5">
	<p class="text-md text-muted-foreground text-justify">
		You have already activated two factor authentication. We highly recommend using two factor
		authentication and strongly advise against removing it. If you still have the need to do so you
		can do it here.
	</p>

	<AlertDialog.Root>
		<AlertDialog.Trigger>
			<Button class="w-full" type="submit" variant="destructive">Delete 2FA</Button>
		</AlertDialog.Trigger>
		<AlertDialog.Content>
			<AlertDialog.Header>
				<AlertDialog.Title>Are you absolutely sure?</AlertDialog.Title>
				<AlertDialog.Description>
					This will remove your current two factor authentication and thus your secondary account protection.
				</AlertDialog.Description>
			</AlertDialog.Header>
			<AlertDialog.Footer>
				<AlertDialog.Cancel>Cancel</AlertDialog.Cancel>
				<AlertDialog.Action
					on:click={formSubmit}
					class="inline-flex items-center justify-center rounded-md text-sm font-medium whitespace-nowrap transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-destructive text-destructive-foreground shadow-sm hover:bg-destructive/90 h-9 px-4 py-2"
					>Remove</AlertDialog.Action
				>
			</AlertDialog.Footer>
		</AlertDialog.Content>
	</AlertDialog.Root>
</form>
