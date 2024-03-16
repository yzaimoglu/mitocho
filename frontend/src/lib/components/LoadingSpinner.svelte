<div class="spinner"></div>

<style>
	.spinner {
		border: 4px solid rgba(0, 0, 0, 0.1);
		border-left-color: hsl(222.2 47.4% 11.2%);
		border-radius: 50%;
		width: 40px;
		height: 40px;
		animation: spin 1s linear infinite;
	}

	@keyframes spin {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}
</style>
