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
		value = value.replace(/\D/g, '');

		if (value) {
			otp[index] = value;
			if (index < otp.length - 1) {
				document.getElementById(`otp-${index + 1}`).focus();
			}
		} else {
			otp[index] = '';
		}
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

	const test = async () => {
		// Promise.reject("test"); --> if login fails
		return sid;
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
		toast.promise(test, {
			loading: 'Logging in...',
			success: (data) => {
				return `Logged in as ${data}`;
			},
			error: 'Could not login.'
		});
	};

	onMount(() => {
		document.getElementById('otp-0').focus();
	});
</script>

<form class="flex justify-center" on:submit|preventDefault={handleSubmit}>
	<div class="flex flex-col gap-4 w-full sm:w-3/4 lg:w-3/5 3xl:w-2/5">
		<div class="flex flex-row gap-2">
			{#each otp as _, index}
				<Input
					class="font-semibold text-center"
					id={`otp-${index}`}
					pattern="[0-9]*"
					inputmode="numeric"
					maxlength={1}
					on:input={(event) => updateOtp(event.target.value, index)}
					on:keydown={(event) => handleKeyDown(event, index)}
				/>
			{/each}
		</div>
		<Button type="submit">Login</Button>
	</div>
</form>
