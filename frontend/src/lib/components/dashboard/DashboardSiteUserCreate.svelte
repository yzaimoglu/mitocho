<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	import Input from '@/components/ui/input/input.svelte';
	import Label from '@/components/ui/label/label.svelte';
	import MitochoTooltip from '../MitochoTooltip.svelte';
	import MitochoDialog from '../MitochoDialog.svelte';
	import DashboardSiteUserRoles from './DashboardSiteUserRoles.svelte';
	import DashboardSiteUserPermissions from './DashboardSiteUserPermissions.svelte';

	type SiteUserCreateForm = {
		username: string;
		email: string;
		password: string;
		password_repeat: string;
		roles: string[];
		permissions: string[];
	};

	let userCreateForm: SiteUserCreateForm = {
		username: '',
		email: '',
		password: '',
		password_repeat: '',
		roles: [],
		permissions: []
	};

	const formSubmit = async () => {
		console.log(userCreateForm);
	};
</script>

<form on:submit|preventDefault={formSubmit} class="flex flex-col gap-8 w-full xl:w-3/4 2xl:w-3/5">
	<div class="flex flex-col gap-1">
		<Label for="email_input">Email</Label>
		<Input
			bind:value={userCreateForm.email}
			id="email_input"
			type="text"
			placeholder="admin@mitocho.io"
		/>
		<p class="text-sm text-muted-foreground">This is the email of the user.</p>
	</div>

	<div class="flex flex-col gap-1">
		<Label for="username_input">Username</Label>
		<Input
			bind:value={userCreateForm.username}
			id="username_input"
			type="text"
			placeholder="admin"
		/>
		<p class="text-sm text-muted-foreground">This is the username of the user.</p>
	</div>

	<div class="flex flex-col gap-1">
		<div class="flex flex-col gap-8 md:flex-row md:gap-2">
			<div class="flex flex-col gap-1 w-full">
				<Label for="username_input">Password</Label>
				<Input
					bind:value={userCreateForm.username}
					id="password_input"
					type="password"
					placeholder="***************"
				/>
			</div>
			<div class="flex flex-col gap-1 w-full">
				<Label for="password_repeat_input">Password Repeat</Label>
				<Input
					bind:value={userCreateForm.username}
					id="password_repeat_input"
					type="password"
					placeholder="***************"
				/>
			</div>
		</div>
		<p class="text-sm text-muted-foreground">This is the password of the user.</p>
	</div>

	<div class="flex flex-row gap-2 w-full">
		<MitochoDialog
			title="Manage user permissions"
			description="You can manage the permissions of the user you want to create."
		>
			<Button slot="trigger" class="w-full" type="button" variant="outline">Manage permissions</Button>
			<div slot="content">
				<DashboardSiteUserPermissions bind:user={userCreateForm} />
			</div>
		</MitochoDialog>
		<MitochoDialog
			title="Manage user roles"
			description="You can manage the roles of the user you want to create."
		>
			<Button slot="trigger" class="w-full" type="button" variant="outline">Manage roles</Button>
			<div slot="content">
				<DashboardSiteUserRoles bind:user={userCreateForm} />
			</div>
		</MitochoDialog>
	</div>

	<MitochoTooltip tooltip="Create the user with these parameters">
		<Button class="w-full" type="submit">Create User</Button>
	</MitochoTooltip>
</form>
