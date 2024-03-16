<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	import Input from '@/components/ui/input/input.svelte';
	import Label from '@/components/ui/label/label.svelte';

	export let sid = 'mitocho';
	export let red = 'mitocho';

	type ProfileGeneralForm = {
		email: string;
		username: string;
	};

	let profileGeneralForm: ProfileGeneralForm = {
		email: '',
		username: ''
	};

	const formSubmit = async () => {
		console.log(sid);
		console.log(red);
		console.log(profileGeneralForm);
	};
</script>

<form on:submit|preventDefault={formSubmit} class="flex flex-col gap-8 w-full xl:w-3/4 2xl:w-3/5">
	<div class="flex flex-col gap-1">
		<Label for="email_input">Email</Label>
		<Input
			bind:value={profileGeneralForm.email}
			id="email_input"
			type="email"
			placeholder="admin@mitocho.io"
		/>
		<p class="text-sm text-muted-foreground">This is the email address you login with. You will need to verify the email change.</p>
	</div>
	<div class="flex flex-col gap-1">
		<Label for="username_input">Username</Label>
		<Input
			bind:value={profileGeneralForm.username}
			id="username_input"
			type="text"
			placeholder="admin"
		/>
		<p class="text-sm text-muted-foreground">This is the unique username you have, you can change your username without verification as long as the username is not taken.</p>
	</div>
	<Button type="submit">Save</Button>
</form>
