<script>
    import { getPodMetadata } from "../../../services/pod.service";
    import { Circle } from "svelte-loading-spinners";
    import { calculateUptime } from "../../../services/time.service";
    import ContentCard from "../../utils/content-card.svelte";
    import Lock from "../../utils/icons/lock.svelte";
    import Refresh from "../../utils/icons/refresh.svelte";

    export let namespace;
    export let pod;

    /** @type {HTMLInputElement} */
    let copyImageNameButton;
    let refreshRate;
    let podInfo;
    $: {
        podInfo = getPodMetadata(pod, namespace);
    }

    /**
     * @param {string} value
     */
	// eslint-disable-next-line no-unused-vars
    const copyToClipboard = (value) => {
        navigator.clipboard.writeText(value);
        copyImageNameButton.value = "Copied! ❤️";

        setTimeout(() => {
            copyImageNameButton.value = "copy image name";
        }, 5000);
    };
</script>
<div class="pod-info-container">
    {#await podInfo}
    <div class="bg-message">
        <Circle size="60" color="#2E77D0" unit="px" duration="1s" />
    </div>
    {:then podInfo}
        {#if !podInfo.Retrieved}
        <div class="bg-message">
            <h2>Whoops... something went wrong.</h2>
        </div>
        {:else}
        <ContentCard width="100%" cardTitle="Pod Meta" cardDescription="Basic pod metadata.">
            <div slot="toolbar" class="pod-info">
                <div class="refresh-rate-container">
                    <Refresh fill="#000000"></Refresh>
                    <select bind:this={refreshRate} name="refreshRate">
                        <option value="5000" selected>5 seconds</option>
                        <option value="30000">30 seconds</option>
                        <option value="60000">1 minute</option>
                        <option value="120000">2 minutes</option>
                        <option value="300000">5 minutes</option>
                    </select>
                </div>
            </div>
            <div slot="content" class="pod-info">
                    <div class="pod-props">
                        <div class="prop">
                            <div><strong>Full Name</strong></div>
                            <div>{podInfo.Name}</div>
                        </div>
                        <div class="prop">
                            <div><strong>Namespace</strong></div>
                            <div>{podInfo.Namespace}</div>
                        </div>
                        <div class="prop">
                            <div><strong>Deployment Name</strong></div>
                            <div>{podInfo.ReleaseName}</div>
                        </div>
                        <div class="prop">
                            <div><strong>Status</strong></div>
                            <div>
                                {#if podInfo.Status === "Running"}
                                    <div class="green-indicator animate-green-indicator"></div>
                                {/if}
                                {podInfo.Status}
                            </div>
                        </div>
                        <div class="prop">
                            <div><strong>Uptime</strong></div>
                            <div>{calculateUptime(podInfo.StartTime)}</div>
                        </div>
                    </div>
            </div>
        </ContentCard>
        <div class="split">
            {#if podInfo.EnvVariables}
            <ContentCard width="calc(50% - 10px)" cardTitle="Environment" cardDescription="Container environment variables.">
                <div slot="content" class="env-table">
                    <div class="table-heading">
                        <h3>Key</h3>
                        <h3>Value</h3>
                    </div>
                    {#each podInfo.EnvVariables as variable}
                    <div class="table-row">
                        <div class="env-key">{variable.name}</div>
                        {#if variable.valueFrom}
                            <div class="env-value sealed"><Lock /><span>Sealed</span></div>
                        {:else}
                            <div class="env-value">{variable.value}</div>
                        {/if}
                    </div>
                    {/each}
                </div>
            </ContentCard>
            {/if}
            <ContentCard width="calc(50% - 10px)" cardTitle="Port Forwarding" cardDescription="">
                <div slot="content" class="content">
                    <div class="wip-message">
                        Work in progress.
                    </div>
                </div>
            </ContentCard>
        </div>
        {/if}
    {/await}
</div>
<style scoped>
  .pod-info-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: calc(100% - 40px);
    padding: 20px;
  }

  .split {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    width: calc(100% + 2px);
    margin-top: 20px;
    flex-direction: row;
  }

  .pod-info {
    padding: 20px;
    padding-bottom: 0px;
    width: calc(100% - 40px);
  }

  .pod-props {
    display: flex;
    flex-wrap: wrap;
  }

  .prop {
    margin-bottom: 20px;
    margin-right: 80px;
  }

  .prop div:nth-child(1) {
    text-transform: uppercase;
    font-size: 0.7rem;
    color: var(--content-card-prop-heading)
  }

  .prop div:nth-child(2) {
    font-weight: 700;
    font-size: 1.2rem;
    display: flex;
    align-items: center;
    color: var(--content-card-prop-value)
  }

  .env-table {
    display: flex;
    flex-direction: column;
    padding: 20px;
  }

  .table-heading, .table-row {
    display: flex;
    flex-direction: row;
  }

  .table-row:nth-child(2n) {
    background-color: rgba(0, 0, 0, 0.05);
  }

  .table-heading > h3, .table-row > div {
    width: 50%;
    padding: 20px;
    margin: 0
  }

  .table-row > div {
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
  }

  .sealed {
    display: flex;
    align-items: center;
  }

  .sealed > span {
    font-weight: 700;
    text-transform: uppercase;
    margin-left: 5px;
    font-size: 0.8rem;
  }

  .refresh-rate-container {
    display: flex;
    align-items: center;
  }

  .refresh-rate-container > select {
    margin-left: 10px;
  }

  .wip-message {
      display: flex;
      justify-content: center;
      padding: 20px 0;
      font-size: 2rem;
      font-weight: 100;
      color: rgba(0, 0, 0, 0.5)
  }
</style>