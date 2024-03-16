<script lang="ts">
	import { pushState } from '$app/navigation';
	import * as Pagination from '@/components/ui/pagination';
	import { ChevronLeft, ChevronRight } from 'svelte-radix';

	export let currentPage = '1';
	export let count = 100;
	export let perPage = 4;
	export let siblingCount = 0;
	export let changePageFunction = async (page: Number) => {
		console.log(`Changed to page ${page}`);
	};
	if (Number(currentPage) > count / perPage) {
		currentPage = (count / perPage).toString();
	} else if (Number(currentPage) <= 0) {
		currentPage = '1';
	}
	const initialCurrentPage = Number(currentPage);

	const changeUrlBar = () => {
		if (typeof window !== 'undefined') {
			const newUrl = new URL(window.location); // Create a new URL object based on the current URL
			newUrl.searchParams.set('page', currentPage); // Set the 'page' query parameter
			pushState(newUrl.toString(), {
				showModal: false
			});
		}
	};

	const changePage = async (pageNum: number) => {
		currentPage = pageNum.toString();
		changeUrlBar();
		await changePageFunction(pageNum);
	};
</script>

<Pagination.Root
	{count}
	{perPage}
	{siblingCount}
	page={initialCurrentPage}
	let:pages
	let:currentPage
>
	<Pagination.Content>
		<Pagination.Item>
			<Pagination.PrevButton on:click={() => changePage(currentPage - 1)}>
				<ChevronLeft class="h-4 w-4" />
				<span class="hidden sm:block">Previous</span>
			</Pagination.PrevButton>
		</Pagination.Item>
		{#each pages as page (page.key)}
			{#if page.type === 'ellipsis'}
				<Pagination.Item>
					<Pagination.Ellipsis />
				</Pagination.Item>
			{:else}
				<Pagination.Item>
					<Pagination.Link
						on:click={() => changePage(page.value)}
						{page}
						isActive={currentPage == page.value}
					>
						{page.value}
					</Pagination.Link>
				</Pagination.Item>
			{/if}
		{/each}
		<Pagination.Item>
			<Pagination.NextButton on:click={() => changePage(currentPage + 1)}>
				<span class="hidden sm:block">Next</span>
				<ChevronRight class="h-4 w-4" />
			</Pagination.NextButton>
		</Pagination.Item>
	</Pagination.Content>
</Pagination.Root>
