<script>
	export let currentView = 0;

	let active = {
		overview: true,
		logs: false,
		deployment: false,
	}

	const setCurrentView = (i) => {
		currentView = i
		Object.keys(active).forEach((k) => active[k] = false)
		active[Object.keys(active)[i]] = true
	}

</script>
<div class="pod-navigation">
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<div class:active={active.overview} on:click={() => setCurrentView(0)} class="page active">
		Overview
	</div>
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<div class:active={active.logs} on:click={() => setCurrentView(1)} class="page">
		Logs
	</div>
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<div class:active={active.deployment} on:click={() => setCurrentView(2)} class="page">
		Deployment
	</div>
</div>
<style scoped>
	.pod-navigation {
		background-color: var(--pod-navigation-bg);
		display: flex;
		flex-direction: row;
		width: 100%;
	}

	.page {
		color: var(--pod-navigation-text);
		padding: 18px 18px 15px 18px;
		cursor: pointer;
	}

	.page:hover {
		color: var(--pod-navigation-text-selected);
	}
	.page:first-child {
		margin-left: 20px;
	}

	.active {
		color: var(--pod-navigation-text-selected);
		border-bottom: 3px solid var(--pod-navigation-column-active);
	}
</style>
