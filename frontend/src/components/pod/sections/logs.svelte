<script>
  import { GetLogsFromPod } from "../../../../wailsjs/go/main/App";
  import { Circle } from "svelte-loading-spinners";
  import { splitLogsByNewline } from "../../../services/pod.service";
  import ContentCard from "../../utils/content-card.svelte";

  export let pod;
  export let namespace;
  $: {
    podLogsOriginal = GetLogsFromPod(pod, namespace, Number(podTailLines));
    podLogsCurrent = GetLogsFromPod(pod, namespace, Number(podTailLines));
  }

  let searchNeedle
  $: filterLogs(searchNeedle)

  const changeTitle = (count) => {
    if (count === 0) {
      title = "Logs" 
    } else if (count === 1) {
      title = `1 Result`
    } else {
      title = `${count} Results`
    }
  }

  const filterLogs = async (searchNeedle) => {
    if (searchNeedle === undefined || searchNeedle === "") {
      podLogsCurrent = podLogsOriginal
      filterCount = 0
    } else {
      let splitLogs = splitLogsByNewline(await podLogsCurrent)
      let filteredLogs = splitLogs.filter((line) => line.includes(searchNeedle)) 
      filterCount = filteredLogs.length
      podLogsCurrent = filteredLogs.join("\n")
    }
  }

  $: changeTitle(filterCount)

  let title = "Logs"
  let filterCount = 0
  let podLogsCurrent
  let podLogsOriginal
  let podTailLines = "100"
</script>
<div class="logs-container">
  <ContentCard cardTitle={title} cardDescription="" width="100%">
    <div slot="toolbar">
      <input spellcheck="false" type="text" bind:value={searchNeedle} class="search-logs" placeholder="Search logs.." />
      <select bind:value={podTailLines} name="tailLines">
        <option value="100" selected>100 lines</option>
        <option value="200">200 lines</option>
        <option value="500">500 lines</option>
        <option value="1000">1000 lines</option>
        <option value="0">Infinity</option>
      </select>
      <select name="logFormat">
        <option value="plaintext" selected>Plain Text</option>
      </select>
    </div>
    <div slot="content">
      {#await podLogsCurrent}
        <div class="bg-message">
        <Circle size="60" color="#2E77D0" unit="px" duration="1s" />
        </div>
      {:then logs} 
        {#if logs === undefined || logs === ''}
        <div class="bg-message">
          <h2>No logs available.</h2>
        </div>
        {:else}
        <div class="logs">
          <pre>{logs}</pre>
        </div>
        {/if}
      {/await}
    </div>
  </ContentCard>
</div>
<style scoped>
  .logs-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: calc(100% - 40px);
    padding: 20px;
  }

  .logs {
    display: flex;
    background-color: black;
    color:white;
    width: calc(100% - 10px);
    height: 50vh;
    padding: 5px;
    text-align: left;
    font-family: "monospace";
    word-wrap: normal;
    overflow: auto;
  }

  .search-logs {
    padding: 10px;
    border-radius: 5px;
    border: 1px solid var(--content-card-border);
    width: 150px;
  }

  .search-logs:focus {
    outline: 0;
    border: 1px solid var(--selectmenu-bg);
  }

  .bg-message {
    margin: 20px 0;
  }

  select {
    background-color: var(--selectmenu-bg);
    color: white;
    border-radius: 5px;
    border: 1px solid var(--content-card-border);
  }
</style>