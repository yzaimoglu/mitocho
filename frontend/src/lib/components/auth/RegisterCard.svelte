<script lang="ts">
	import * as Card from '@/components/ui/card';
	import Input from '@/components/ui/input/input.svelte';
	import MitochoLogo from '@/components/MitochoLogo.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import Label from '@/components/ui/label/label.svelte';
	import Alert from '../ui/alert/alert.svelte';
	import Checkbox from '../ui/checkbox/checkbox.svelte';

	export let sid = 'mitocho';
	export let red = 'mitocho';

	type RegisterForm = {
		email: string;
		username: string;
		password: string;
		password_repeat: string;
		accept_tos: boolean;
	};

	let registerForm: RegisterForm = {
		email: '',
		username: '',
		password: '',
		password_repeat: '',
		accept_tos: false
	};

	let errors: string[] = [];

	const formSubmit = async () => {
		console.log(sid);
		console.log(red);
		console.log(registerForm);
	};

	const getLoginUrl = () => {
		return `/auth/login?sid=${sid}&red=${red}`;
	};
</script>

<Card.Root class="w-3/4 md:w-1/2 xl:w-1/3">
	<Card.Header>
		<div class="flex flex-col gap-4">
			{#if errors.length !== 0}
				<Alert class="bg-red-50" variant="destructive">
					<h1 class="font-semibold">The following changes have to be made:</h1>
					<ul class="list-disc list-inside">
						{#each errors as error}
							<li>
								{error}
							</li>
						{/each}
					</ul>
				</Alert>
			{/if}
			<div class="flex flex-row justify-between gap-2">
				<div class="flex flex-col justify-center gap-1.5">
					<Card.Title>Register</Card.Title>
					<Card.Description
						>Please register an account by entering your information</Card.Description
					>
				</div>
				<MitochoLogo width="w-16" />
			</div>
		</div>
	</Card.Header>
	<Card.Content>
		<form on:submit|preventDefault={formSubmit} class="flex flex-col gap-4">
			<div class="flex flex-col gap-1">
				<Label for="email_input">Email</Label>
				<Input
					bind:value={registerForm.email}
					id="email_input"
					type="email"
					placeholder="admin@mitocho.io"
					class="w-full"
				/>
			</div>
			<div class="flex flex-col gap-1">
				<Label for="username_input">Username</Label>
				<Input
					bind:value={registerForm.username}
					id="username_input"
					type="text"
					placeholder="admin"
					class="w-full"
				/>
			</div>

			<div class="flex flex-col gap-1">
				<Label for="password_input">Password</Label>
				<Input
					bind:value={registerForm.password}
					id="password_input"
					type="password"
					placeholder="************"
				/>
			</div>
			<div class="flex flex-col gap-1">
				<Label for="password_repeat_input">Password Repeat</Label>
				<Input
					bind:value={registerForm.password_repeat}
					id="repeatpassword_input"
					type="password"
					placeholder="************"
				/>
			</div>
			<div class="flex flex-row gap-1">
				<div class="flex items-center space-x-2">
					<Checkbox id="terms" bind:checked={registerForm.accept_tos} />
					<Label
						for="terms"
						class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
					>
						Accept terms and conditions
					</Label>
				</div>
			</div>
			<Button type="submit" class="mt-5">Register</Button>
			<Button href={getLoginUrl()} class="text-xs" type="button" variant="link"
				>Login instead</Button
			>
		</form>
	</Card.Content>
</Card.Root>
