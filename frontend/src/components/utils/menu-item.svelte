<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  export let text;
  export let status;
  export let podID;
  function setActivePod() {
      dispatch("switchActivePod", {
          podID
      });
  }

  export let active;
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="pod" class:active={active} id={podID} on:click={setActivePod}>
	{#if status === "Running"} <div class="green-indicator"></div> {/if}
	<div class="pod-text">{text}</div>
</div>

<style >
	.pod-text {
		font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
		color: var(--sidebar-text);
	}

	.pod {
		padding: 10px 10px 10px 12px;
		margin: 10px 0;
		display: flex;
		align-items: center;
		text-overflow: ellipsis;
		white-space: nowrap;
		overflow: hidden;
		cursor: pointer;
		border-left: 3px solid rgb(0, 0, 0, 0);
	}
	
	.pod-text:hover {
		color: var(--sidebar-text-active);
	}

	.pod.active {
		color: var(--sidebar-text-active);
		border-left: 3px solid var(--sidebar-row-active);
	}
</style>