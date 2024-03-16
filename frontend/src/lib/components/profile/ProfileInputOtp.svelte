<script lang="ts">
	import Input from '@/components/ui/input/input.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import { onMount } from 'svelte';
	import { toast } from 'svelte-sonner';
	import Label from '../ui/label/label.svelte';

	export let sid: string;
	export let red: string;

	type CreateOtpForm = {
		totp: string;
	};

	let createOtpForm: CreateOtpForm = {
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
		createOtpForm.totp = otpCode;
		console.log(sid);
		console.log(red);
		console.log(createOtpForm);
		toast.promise(test, {
			loading: 'Activating two factor authentication...',
			success: (data) => {
				return `Activated`;
			},
			error: 'Could not activate two factor authentication.'
		});
	};

	onMount(() => {
		document.getElementById('otp-0').focus();
	});
</script>

<form on:submit|preventDefault={handleSubmit}>
	<div class="flex flex-col gap-2">
		<Label for="otp-0">One Time Password</Label>
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
		<Button type="submit" class="mt-4">Activate 2FA</Button>
	</div>
</form>
