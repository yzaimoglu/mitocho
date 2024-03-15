<script lang="ts">
	import Input from '@/components/ui/input/input.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import { onMount } from 'svelte';
	import { toast } from 'svelte-sonner';

	export let sid: string;
	export let red: string;
	export let email: string;
	export let password: string;

	type LoginOtpForm = {
		email: string;
		password: string;
		totp: string;
	};

	let loginForm: LoginOtpForm = {
		email: '',
		password: '',
		totp: ''
	};

	let otp = Array(6).fill('');

	const updateOtp = (value: any, index: number) => {
		otp[index] = value;
		if (value && index < otp.length - 1) {
			document.getElementById(`otp-${index + 1}`).focus();
		}
	};

	const handleKeyDown = (event: any, index: number) => {
		if ((event.key === 'Backspace' || event.key === 'Delete') && !otp[index] && index > 0) {
			document.getElementById(`otp-${index - 1}`).focus();
			if (event.key === 'Backspace') {
				otp[index - 1] = '';
			}
		}
	};

	const handleSubmit = async () => {
		const otpCode = otp.join('');
		loginForm = {
			email: email,
			password: password,
			totp: otpCode
		};
		console.log(sid);
		console.log(red);
		console.log(loginForm);
		toast("Does this work?");
	};

	onMount(() => {
		document.getElementById('otp-0').focus();
	});
</script>

<form on:submit|preventDefault={handleSubmit}>
	<div class="flex flex-row gap-2">
		{#each otp as _, index}
			<Input
				class="font-semibold text-center"
				id={`otp-${index}`}
				type="text"
				maxlength={1}
				on:input={(event) => updateOtp(event.target.value, index)}
				on:keydown={(event) => handleKeyDown(event, index)}
			/>
		{/each}
		<Button type="submit">Login</Button>
	</div>
</form>
