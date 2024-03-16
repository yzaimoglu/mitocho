<script lang="ts">
	import ProfileInputOtp from './ProfileInputOtp.svelte';

	export let sid = 'mitocho';
	export let red = 'mitocho';

	type Profile2FACreateForm = {
		totp: string;
	};

	let profile2faCreateForm: Profile2FACreateForm = {
		totp: ''
	};

	const formSubmit = async () => {
		console.log(sid);
		console.log(red);
		console.log(profile2faCreateForm);
	};
</script>

<form on:submit|preventDefault={formSubmit} class="flex flex-col gap-8 w-full xl:w-3/4 2xl:w-3/5">
	<p class="text-md text-muted-foreground text-justify">
		You have not activated two factor authentication yet. In order to activate it you have to scan
		the 2fa code with a one time password application of your choice.
	</p>
	<div class="flex justify-center">
		<img
			alt="Two Factor QR Code"
			class="max-w-60"
			src="https://cdn.ttgtmedia.com/rms/misc/qr_code_barcode.jpg"
		/>
	</div>

	<ProfileInputOtp {sid} {red} />
</form>
