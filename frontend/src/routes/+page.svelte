<script lang="ts">
	// import { onMount } from 'svelte';
	// import { getPayload } from '@/paseto';
	import MitochoPage from '@/components/MitochoPage.svelte';
	import GetStarted from '@/components/install/GetStarted.svelte';

	// onMount(async () => {
	// 	const access_token =
	// 		'v4.public.eyJhdWQiOiJtaXRvY2hvLmlvIiwiZXhwIjoiMjAyNC0wMy0xMVQyMjoxNToyMiswMTowMCIsImlhdCI6IjIwMjQtMDMtMTFUMjE6MTU6MjIrMDE6MDAiLCJpc3MiOiJtaXRvY2hvIiwibmJmIjoiMjAyNC0wMy0xMVQyMToxNToyMiswMTowMCIsInBheWxvYWQiOnsiVGVzdFJpZ2h0IjoxMjMsIkFiYyI6NDU2LCJEZmUiOjc4OX0sInN1YiI6ImFjY2VzcyJ9v1xZTcycjfZOvwadY1L4hZelMPom3d_PoQ222FBx4msubGzayENFF8OO4hpdhgptIPRk10KvM3c1Tw47HkiGDA';
	// 	const public_key = 'RfDzruLNgVMy0SdbW0tgYE81Upe7ixrMxDfmHTvyTg8=';

	// 	const verified = await getPayload(public_key, access_token);
	// 	console.log(verified);
	// });
</script>

<MitochoPage title="Mitocho">
	<GetStarted />
</MitochoPage>
