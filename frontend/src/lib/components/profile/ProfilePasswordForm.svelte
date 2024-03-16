<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	import Input from '@/components/ui/input/input.svelte';
	import Label from '@/components/ui/label/label.svelte';

	export let sid = 'mitocho';
	export let red = 'mitocho';

	type ProfilePasswordForm = {
		current_password: string;
		password: string;
		password_repeat: string;
	};

	let profilePasswordForm: ProfilePasswordForm = {
		current_password: '',
		password: '',
		password_repeat: ''
	};

	const formSubmit = async () => {
		console.log(sid);
		console.log(red);
		console.log(profilePasswordForm);
	};
</script>

<form on:submit|preventDefault={formSubmit} class="flex flex-col gap-8 w-full xl:w-3/4 2xl:w-3/5">
	<div class="flex flex-col gap-1">
		<Label for="current_password_input">Current Password</Label>
		<Input
			bind:value={profilePasswordForm.current_password}
			id="current_password_input"
			type="password"
			placeholder="*****************"
		/>
		<p class="text-sm text-muted-foreground">
			This is the password you currently use during login.
		</p>
	</div>
	<div class="flex flex-col gap-1">
		<Label for="password_input">New Password</Label>
		<Input
			bind:value={profilePasswordForm.password}
			id="password_input"
			type="password"
			placeholder="*****************"
		/>
	</div>
	<div class="flex flex-col gap-1">
		<Label for="password_repeat_input">New Password Repeat</Label>
		<Input
			bind:value={profilePasswordForm.password_repeat}
			id="password_repeat_input"
			type="password"
			placeholder="*****************"
		/>
		<p class="text-sm text-muted-foreground">
			Please enter your new password two times to make sure it is typed in correctly.
		</p>
	</div>
	<Button type="submit">Change Password</Button>
</form>
