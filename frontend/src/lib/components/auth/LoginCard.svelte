<script lang="ts">
	import * as Card from '@/components/ui/card';
	import Input from '@/components/ui/input/input.svelte';
	import MitochoLogo from '@/components/MitochoLogo.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import Label from '@/components/ui/label/label.svelte';
	import Alert from '../ui/alert/alert.svelte';

	export let sid = "mitocho";
	export let red = "mitocho";

	type LoginForm = {
		email: string;
		password: string;
	};

	let loginForm: LoginForm = {
		email: '',
		password: ''
	};

	let errors: string[] = [];

	const formSubmit = async () => {
		console.log(sid);
		console.log(red);	
		console.log(loginForm);
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
					<Card.Title>Login</Card.Title>
					<Card.Description>Please login with your account</Card.Description>
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
					bind:value={loginForm.email}
					id="email_input"
					type="email"
					placeholder="admin@mitocho.io"
					class="w-full"
				/>
			</div>
			<div class="flex flex-col gap-1">
				<Label for="password_input">Password</Label>
				<Input
					bind:value={loginForm.password}
					id="password_input"
					type="password"
					placeholder="************"
				/>
				<a href="/auth/forgot-password"><p class="text-xs">Forgot Password?</p></a>
			</div>
			<Button type="submit" class="mt-5">Login</Button>
		</form>
	</Card.Content>
</Card.Root>
