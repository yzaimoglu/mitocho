<script lang="ts">
	import * as Card from '@/components/ui/card';
	import MitochoLogo from '@/components/MitochoLogo.svelte';
	import { loginStore } from '@/stores/loginStore';
	import InputOtpVert from './InputOtpVert.svelte';
	import InputOtpHorizont from './InputOtpHorizont.svelte';

	export let sid = "mitocho";
	export let red = "mitocho";
  export let email = $loginStore.email;
  export let password = $loginStore.password;

  
</script>

<Card.Root class="w-3/4 md:w-1/2 xl:w-1/3">
	<Card.Header>
		<div class="flex flex-col gap-4">
			<div class="flex flex-row justify-between gap-2">
				<div class="flex flex-col justify-center gap-1.5">
					<Card.Title>Two Factor Authentication</Card.Title>
					<Card.Description>Please enter your one time password</Card.Description>
				</div>
				<MitochoLogo width="w-16" />
			</div>
		</div>
	</Card.Header>
	<Card.Content>
		<div class="flex flex-col gap-4">
			<div class="flex flex-col gap-1">
			  <InputOtpHorizont {sid} {red} {email} {password} />
      </div>
		</div>
	</Card.Content>
</Card.Root>
