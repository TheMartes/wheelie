<script>
  // @ts-nocheck
  import Bug from "./utils/icons/bug.svelte";
  import ContextSelect from "./utils/overlay/context-select.svelte";
  import Close from "./utils/window/controls/close.svelte";
  import Maximise from "./utils/window/controls/maximise.svelte";
  import Minimise from "./utils/window/controls/minimise.svelte";


  export let isMacOS;
  export let selectedContext;
</script>

<header>
	<div class="topbar" style="--wails-draggable:drag">
		<div class="control-buttons-macos">
		</div>
		<div class="middle-sector">
			<ContextSelect on:switchContext bind:selectedContext></ContextSelect>
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<div class="bug-report" on:click={() => window.runtime.BrowserOpenURL("https://github.com/TheMartes/wheelie/issues/new?assignees=TheMartes&labels=&projects=&template=bug_report.md&title=%5BBUG%5D")}>
				<Bug></Bug>
			</div>
		</div>
		<div class="control-buttons-windows">
			{#if !isMacOS}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<div class="minimise" on:click={window.runtime.WindowMinimise()}>
					<Minimise color="white"} />
				</div>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<div class="maximise" on:click={window.runtime.WindowMaximise()}>
					<Maximise color="white" />
				</div>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<div class="close" on:click={window.runtime.Quit()}>
					<Close color="white" />
				</div>
			{/if}
		</div>
	</div>
</header>

<style>
	header {
		position: fixed;
		width: 100%;
		z-index: 9;
	}
	.topbar {
		display: flex;
		justify-content: space-between;
		align-items: center;
		background-color: var(--sidebar-bg);
		height: 36px;
	}

	.control-buttons-macos, .control-buttons-windows {
		height: 36px;
		display: flex;
	}

	.minimise, .maximise, .close {
		cursor: pointer;
		padding: 2px;
	}

	.minimise:hover, .maximise:hover {
		background-color: rgba(255, 255, 255, 0.1);
	}

	.minimise, .maximise {
		height: 40px;
		width: 50px;
	}

	.close {
		height: 32px;
		width: 50px;
	}

	.close:hover {
		background-color: rgb(231, 22, 22);
	}

	.middle-sector {
		display: flex;
		align-items: center;
	}

	.bug-report {
		margin-left: 10px;
		cursor: pointer;
		margin-top: 2px; /** Svg is kinda broken */
	}
</style>
