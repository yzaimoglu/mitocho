<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	import Input from '@/components/ui/input/input.svelte';
	import Label from '@/components/ui/label/label.svelte';
	import MitochoTooltip from '../MitochoTooltip.svelte';
	import MitochoSelect from '@/components/MitochoSelect.svelte';
	import type { SelectItem } from '@/components/MitochoSelect.svelte';
	import { BookmarkPlus, BookmarkX } from 'lucide-svelte';
	import { toast } from 'svelte-sonner';

	type SiteCreateForm = {
		name: string;
		description?: string;
		domains: string[];
		callbacks: string[];
		privacy_policy?: string;
		imprint?: string;
		color: string;
		logo: string;
		terms_of_service: string;
	};

	let siteCreateForm: SiteCreateForm = {
		name: '',
		description: '',
		domains: [''],
		callbacks: [''],
		privacy_policy: '',
		imprint: '',
		color: '',
		logo: '',
		terms_of_service: ''
	};

	const addNewDomain = () => {
		if (siteCreateForm.domains.length > 2) {
			toast.error('Please do not add more than 3 domains.');
		} else {
			siteCreateForm.domains[siteCreateForm.domains.length] = '';
		}
	};

	const addNewCallback = () => {
		if (siteCreateForm.callbacks.length > 2) {
			toast.error('Please do not add more than 3 callbacks.');
		} else {
			siteCreateForm.callbacks[siteCreateForm.callbacks.length] = '';
		}
	};

	const removeDomain = (index: number) => {
		let tmp = siteCreateForm.domains;
		tmp.splice(index, 1);
		siteCreateForm.domains = tmp;
	};

	const removeCallback = (index: number) => {
		let tmp = siteCreateForm.callbacks;
		tmp.splice(index, 1);
		siteCreateForm.callbacks = tmp;
	};

	const filterDuplicates = (array: any): any => {
		return [...new Set(array)];
	};

	const colors: SelectItem[] = [
		{
			value: 'green',
			text: 'Green'
		},
		{
			value: 'red',
			text: 'Red'
		},
		{
			value: 'blue',
			text: 'Blue'
		},
		{
			value: 'purple',
			text: 'Purple'
		}
	];

	const uploadFile = async (event: Event) => {
		const input = event.target as HTMLInputElement;
		const file = input.files ? input.files[0] : null;
		if (file) {
			// Check file type
			const validTypes = ['image/png', 'image/jpeg', 'image/jpg', 'image/webp'];
			if (!validTypes.includes(file.type)) {
				input.value = '';
				toast.error(
					'The uploaded file is not an accepted image. Please only upload .png, .jpeg, .jpg or .webp files.'
				);
				return;
			}

			// 1MB Check
			if (file.size > 1 * 1024 * 1024) {
				input.value = '';
				toast.error('The uploaded image is too big.');
				return;
			}
			const reader = new FileReader();

			// Read bytes
			reader.onload = (e: ProgressEvent<FileReader>) => {
				const dataUrl = e.target.result as string;
				siteCreateForm.logo = dataUrl;
			};

			// Read error
			reader.onerror = () => {
				toast.error('An error occured during your file upload.');
				return;
			};

			reader.readAsDataURL(file);
		} else {
			toast.error('An error occured during your file upload.');
			return;
		}
	};

	const formSubmit = async () => {
		siteCreateForm.callbacks = filterDuplicates(siteCreateForm.callbacks);
		siteCreateForm.domains = filterDuplicates(siteCreateForm.domains);
		console.log(siteCreateForm);
	};
</script>

<form on:submit|preventDefault={formSubmit} class="flex flex-col gap-8 w-full xl:w-3/4 2xl:w-3/5">
	<!-- Name -->
	<div class="flex flex-col gap-1">
		<Label for="name_input">Site Name</Label>
		<Input bind:value={siteCreateForm.name} id="name_input" type="text" placeholder="mitocho" />
		<p class="text-sm text-muted-foreground">
			This is the name of this site. You will see this on the dashboard.
		</p>
	</div>

	<!-- Description -->
	<div class="flex flex-col gap-1">
		<Label for="description_input">Site Description</Label>
		<Input
			bind:value={siteCreateForm.description}
			id="description_input"
			type="text"
			placeholder="Mitocho instance"
		/>
		<p class="text-sm text-muted-foreground">
			This is the description of this site. You will see this on the dashboard.
		</p>
	</div>

	<!-- Imprint -->
	<div class="flex flex-col gap-1">
		<Label for="imprint_input">Imprint</Label>
		<Input
			bind:value={siteCreateForm.imprint}
			id="imprint_input"
			type="text"
			placeholder="https://example.com/imprint"
		/>
		<p class="text-sm text-muted-foreground">
			This is your imprint link which will be displayed in the footer of all Mitocho pages
			when users are coming from this site.
		</p>
	</div>

	<!-- Privacy Policy & Tos -->
	<div class="flex flex-col gap-8 md:flex-row md:gap-4">
		<div class="flex flex-col gap-1 md:w-1/2">
			<Label for="privacy_input">Privacy Policy</Label>
			<Input
				bind:value={siteCreateForm.privacy_policy}
				id="privacy_input"
				type="text"
				placeholder="https://example.com/privacy"
			/>
			<p class="text-sm text-muted-foreground">
				This is your privacy policy link which will be displayed in the footer of all Mitocho pages
				when users are coming from this site.
			</p>
		</div>
		<div class="flex flex-col gap-1 md:w-1/2">
			<Label for="tos_input">Terms of Service</Label>
			<Input
				bind:value={siteCreateForm.terms_of_service}
				id="tos_input"
				type="text"
				placeholder="https://example.com/tos"
			/>
			<p class="text-sm text-muted-foreground">
				This is your terms of service link which will be displayed in the footer of all Mitocho
				pages when users are coming from this site.
			</p>
		</div>
	</div>

	<!-- Color and Logo -->
	<div class="flex flex-col gap-8 md:flex-row md:gap-4">
		<div class="flex flex-col gap-1 md:w-1/2">
			<Label for="color_input">Color</Label>
			<MitochoSelect
				className="w-full"
				bind:selected={siteCreateForm.color}
				label="Color"
				placeholder="Color"
				items={colors}
			/>
			<p class="text-sm text-muted-foreground">
				This is the color your Mitocho instance will have in the background when your users are
				coming from this site.
			</p>
		</div>
		<div class="flex flex-col gap-1 md:w-1/2">
			<Label for="logo_input">Logo</Label>
			<Input id="logo_input" type="file" accept=".png, .jpeg, .jpg, .webp" on:change={uploadFile} />
			<p class="text-sm text-muted-foreground">
				This is the logo that will be displayed on all the pages for users coming from this site.
				Maximum upload size is 1MB.
			</p>
		</div>
	</div>

	<!-- Domains -->
	<div class="flex flex-col gap-8 md:flex-row md:gap-4">
		<div class="flex flex-col gap-1 md:w-1/2">
			<Label class="flex flex-row gap-1" for="domains_input">
				Domains
				<MitochoTooltip tooltip="Add a new domain">
					<button type="button" on:click={addNewDomain}>
						<BookmarkPlus size="1rem" />
					</button>
				</MitochoTooltip>
			</Label>
			{#each siteCreateForm.domains as _, domainIndex}
				<div class="flex flex-row gap-1">
					<Input
						bind:value={siteCreateForm.domains[domainIndex]}
						id={`domains_${domainIndex}_input`}
						type="text"
						placeholder="example.com"
					/>
					{#if domainIndex !== 0}
						<MitochoTooltip tooltip="Remove domain">
							<Button on:click={() => removeDomain(domainIndex)} variant="destructive">
								<BookmarkX size="1rem" />
							</Button>
						</MitochoTooltip>
					{/if}
				</div>
			{/each}
			<p class="text-sm text-muted-foreground">
				These are the domains that are linked to this site.
			</p>
		</div>

		<!-- Callbacks -->
		<div class="flex flex-col gap-1 md:w-1/2">
			<Label class="flex flex-row gap-1" for="callbacks_input">
				Callbacks
				<MitochoTooltip tooltip="Add a new callback">
					<button type="button" on:click={addNewCallback}>
						<BookmarkPlus size="1rem" />
					</button>
				</MitochoTooltip>
			</Label>
			{#each siteCreateForm.callbacks as _, callbackIndex}
				<div class="flex flex-row gap-1">
					<Input
						bind:value={siteCreateForm.callbacks[callbackIndex]}
						id={`callbacks_${callbackIndex}_input`}
						type="text"
						placeholder="https://example.com/callback"
					/>
					{#if callbackIndex !== 0}
						<MitochoTooltip tooltip="Remove callback">
							<Button on:click={() => removeCallback(callbackIndex)} variant="destructive">
								<BookmarkX size="1rem" />
							</Button>
						</MitochoTooltip>
					{/if}
				</div>
			{/each}
			<p class="text-sm text-muted-foreground">
				These are the callbacks that will be called when a user registers or logins.
			</p>
		</div>
	</div>

	<MitochoTooltip tooltip="Create the site with these parameters">
		<Button class="w-full" type="submit">Create Site</Button>
	</MitochoTooltip>
</form>
