<script>
	export let cardTitle;
	export let cardDescription;
	export let width;
</script>

<div class="content-card" style:width={width}>
	{#if cardTitle || cardDescription}
	<div class="card-header">
		<div class="title-section">
			{#if cardTitle}
				<h1>{cardTitle}</h1>
			{/if}
			{#if cardDescription}
				<span class="card-description">{cardDescription}</span>
			{/if}
		</div>
		<div class="toolbar">
			<slot name="toolbar"></slot>
		</div>
	</div>
	{/if}
	<slot name="content"></slot>
</div>

<style scoped>
	.content-card {
		background-color: var(--content-card-bg);
		border: 1px solid var(--content-card-border);
	}

	.card-header {
		display: flex;
		justify-content: space-between;
		padding: 20px;
		border-bottom: 1px solid var(--content-card-border);
	}

	h1 {
		font-weight: 700;
		margin: 0;
	}

	.card-description {
		color: var(--content-card-description)
	}

	.toolbar {
		display: flex;
		flex-direction: row;
	}
</style>