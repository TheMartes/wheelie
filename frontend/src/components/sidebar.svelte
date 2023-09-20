<script>
  import {GetNamespaces, GetPodsFromNamespace} from "../../wailsjs/go/main/App.js";
  import MenuItem from "./utils/menu-item.svelte";
  import { BarLoader } from "svelte-loading-spinners";
  import { slide } from "svelte/transition";
  import { quintOut } from "svelte/easing";
  import { onDestroy } from "svelte";

  const getNamespacesFromBackend = async () => {
      const namespaces = await GetNamespaces();

      return namespaces.filter(namespace => namespace !== "");
  };

  /** @type {Object<string, MenuItem>} */
  let podMap = {};

  /**
   * @param {string} selectedNamespace
   */
  const getPodsFromBackend = async (selectedNamespace) => {
      let p = await GetPodsFromNamespace(selectedNamespace);
      p = p.filter(pod => pod.ShortName !== "" || pod.LongName !== "");

      p.forEach(pod => podMap[pod.LongName] = null);
      return p;
  };

  /** @param {string} pod */
  const selectActivePod = (pod, releaseName) => {
      if (selectedPod !== null) {
          podMap[selectedPod].$set({active: false});
      }

      selectedPod = pod;
      selectedReleaseName = releaseName;
      podMap[selectedPod].$set({active: true});
  };

  /**
   * @type {string|null}
   */
  export let selectedPod = null;
  export let selectedNamespace = "";
  export let selectedReleaseName = "";
  $: loadPods(selectedNamespace, true);
  $: {
      if (selectedNamespace === "") {
          pods = null;
      }
  }

  const loadPods = async (sn, resetPod = false) => {
      // Reset
      podMap = {};

      if (resetPod) {
        selectedPod = null;
      }

      if (sn !== "") {
          pods = await getPodsFromBackend(sn);
      }
  };

  let getNamespaces = getNamespacesFromBackend();
  let pods;

  let interval = setInterval(async () => {
    await loadPods(selectedNamespace)
  }, 5000);

  onDestroy(() => {
      clearInterval(interval);
  });
</script>

<div transition:slide={{ delay: 250, duration: 300, axis: "x", easing: quintOut  }} class="sidebar-container">
  <h2>Namespace</h2>
  {#await getNamespaces}
    <div class="bg-message">
      {#await pods}<BarLoader size="20" color="#7DF3E1" unit="px" duration="1s" />{/await}
    </div>
  {:then namespaces}
    <select bind:value={selectedNamespace} name="namespace">
        <option value="" selected>Please choose a namespace</option>
        {#each namespaces as namespace}
            <option value="{namespace.toLowerCase()}">{namespace}</option>
        {/each}
    </select>
  {/await}

  <div class="break-line"></div>
  <div class="pods-wrapper">
  <h3>Pods</h3>
  {#await pods then pods}
	{#if pods}
      {#each pods as pod}
		<MenuItem
			bind:this={podMap[pod.LongName]}
			active={false}
			text={pod.ShortName}
			podID={pod.LongName}
			status={pod.Status}
			on:switchActivePod={(e) => selectActivePod(e.detail.podID, pod.ReleaseName)}
		/>
      {/each}
	{:else}
		<div class="bg-message"><h2>No namespace.</h2></div>
	{/if}
  {/await}
  </div>
</div>

<style >
	.sidebar-container {
		position: fixed;
		height:100%;
		text-align: left;
		width: 20%;
		overflow: auto;
		background-color: var(--sidebar-bg);
		color: var(--main-bg-color);
	}

	.sidebar-container > h2 {
		margin: 10px 0 10px 10px;
		font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
		color: var(--sidebar-heading);
		font-weight: 200;
	}

	.pods-wrapper {
		overflow: auto;
	}

	.pods-wrapper::-webkit-scrollbar {
		display: block;
	}

	.sidebar-container > .pods-wrapper > h3 {
		margin: 10px 10px;
		font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
		color: var(--sidebar-heading);
		font-weight: 200;
	}

	.break-line {
		background-color: var(--sidebar-divider);
	}

	.bg-message {
		height: 200px;
	}

	.bg-message h2 {
		color: var(--sidebar-text);
	}

	select {
		-moz-appearance:none; /* Firefox */
		-webkit-appearance:none; /* Safari and Chrome */
		appearance:none;
		width: calc(100% - 20px);
		margin: 0 10px 10px 10px;
		border-radius: 0;
		background-color: var(--sidebar-bg);
		color: var(--sidebar-text);
		border: 1px solid var(--sidebar-divider);
		font-size: 0.75rem;
	}

	select:focus {
		outline: 1px solid var(--sideabr-heading);
	}
</style>
