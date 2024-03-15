<script lang="ts">
	import * as Card from '@/components/ui/card';
	import Input from '@/components/ui/input/input.svelte';
	import MitochoLogo from '@/components/MitochoLogo.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import Label from '@/components/ui/label/label.svelte';

	type InstallationForm = {
		email: string;
		username: string;
		password: string;
		password_repeat: string;
		domain: string;
	};

	let installationForm: InstallationForm = {
		domain: '',
		email: '',
		password: '',
		password_repeat: '',
		username: ''
	};

	const formSubmit = async () => {
		console.log(installationForm);
	};
</script>

<MitochoLogo width="w-24" />
<Card.Root class="w-3/4 md:w-3/5 xl:w-1/2">
	<Card.Header>
		<Card.Title>Mitocho Installation</Card.Title>
		<Card.Description>Please create your admin acccount</Card.Description>
	</Card.Header>
	<Card.Content>
		<form on:submit|preventDefault={formSubmit} class="flex flex-col gap-6">
			<div class="flex flex-col gap-2">
				<Label for="email_input">Email</Label>
				<Input
					bind:value={installationForm.email}
					id="email_input"
					type="email"
					placeholder="admin@mitocho.io"
					class="w-full"
				/>
			</div>
			<div class="flex flex-col gap-2">
				<Label for="username_input">Username</Label>
				<Input
					bind:value={installationForm.username}
					id="username_input"
					type="text"
					placeholder="admin"
					class="w-full"
				/>
			</div>
			<div class="flex flex-row w-full gap-2">
				<div class="flex flex-col gap-2 w-1/2">
					<Label for="password_input">Password</Label>
					<Input
						bind:value={installationForm.password}
						id="password_input"
						type="password"
						placeholder="************"
					/>
				</div>
				<div class="flex flex-col gap-2 w-1/2">
					<Label for="password_input">Password Repeat</Label>
					<Input
						bind:value={installationForm.password_repeat}
						id="password_repeat_input"
						type="password"
						placeholder="************"
					/>
				</div>
			</div>
			<div class="flex flex-col gap-2">
				<Label for="domain_input">Domain</Label>
				<Input
					bind:value={installationForm.domain}
					id="domain_input"
					type="text"
					placeholder="auth.mitocho.io"
					class="w-full"
				/>
			</div>
			<Button type="submit" class="mt-5">Finish Setup</Button>
		</form>
	</Card.Content>
</Card.Root>
