<script lang="ts">
	import { onMount } from "svelte";
	import { authStore } from "@/stores/authStore";
	onMount(() => {
		if(authStore.isEmpty()) {
			// paseto check cookie
			// set auth token and user
			// if not valid, check refresh token
			// get new access token
			// if refresh token not valid, redirect to login
		}
		authStore.setUser({
			email: "yz@yagi.sh",
			permissions: ["*"],
			roles: [],
			sid: "123",
			username: "yz",
		})
		console.log("check for authentication and save user to user store");
	});
</script>

<slot />
