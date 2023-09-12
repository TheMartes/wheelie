<script>
  import Sidebar from './components/sidebar.svelte'
  import Overview from './components/pod/sections/overview.svelte';
  import Topbar from './components/topbar.svelte';
  import Navigation from './components/pod/navigation.svelte';
  import Container from './components/pod/container.svelte';
  import Logs from './components/pod/sections/logs.svelte';
  import GuidingArrow from './components/utils/window/guiding-arrow.svelte';
  import { fade } from 'svelte/transition';
  import { quintOut } from 'svelte/easing';

  // Global settings
  export let selectedPod;
  export let selectedNamespace;
  export let selectedContext = null;
  export let currentTab;
  const views = [Overview, Logs]
  let viewportComponent = null
	let currentView = 0

  let component = Overview;

  $: switchComponent(currentTab)

  /** @param {string} currentTab */
  const switchComponent = (currentTab) => {
    switch (currentTab) {
      case "overview":
        component = Overview
      case "logs":
        component = Logs
      default:
        component = Overview
    }
  }

	function updateViewportComponent() {
		viewportComponent = views[currentView]
	}

  updateViewportComponent()

  const onSwitchContext = (e) => {
    selectedContext = e.detail.context;
    selectedNamespace = ""
    selectedPod = null;
  }

  const isMacOS = navigator.userAgent.indexOf("Mac")!= -1;
</script>

<main class="root-container">
  <Topbar on:switchContext={(e) => onSwitchContext(e)} bind:selectedContext isMacOS={isMacOS} />
  {#if selectedContext}
  <div class="content-area">
    <Sidebar bind:selectedNamespace bind:selectedPod />

    <Container>
      {#if !selectedPod}
        <div class="sidebar-guide">
          <div class="guiding-arrow-left"><GuidingArrow rotate={"-60deg"}></GuidingArrow></div>
          <div class="guiding-text-left">Select namespace & pod!</div>
        </div>
        {:else}
        <Navigation bind:currentView />
        {#if viewportComponent == views[currentView]}
          <div id="viewport" on:outroend={updateViewportComponent} transition:fade={{ duration: 150 }}>
            <svelte:component this={viewportComponent} pod={selectedPod} namespace={selectedNamespace}></svelte:component>
          </div>
        {/if}
      {/if}
    </Container>
  </div>
  {:else}
    <div class="context-guide">
      <div class="guiding-arrow-up"><GuidingArrow rotate={"15.216deg"}></GuidingArrow></div>
      <div class="guiding-text-up">First, let's select a context!</div>
    </div>
  {/if}
</main>

<style scoped>
  .root-container {
    display: flex;
    flex-direction: column;
  }

  .content-area {
    margin-top: 36px;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    width: 100%;
    z-index: 1;
  }

  .context-guide {
    width: 100%;
    display: flex;
    z-index: 0;
    flex-direction: column;
    align-items: center;
    margin-top: var(--header-size);

  }

  .guiding-arrow-up {
    margin-left: 100px
  }

  .guiding-arrow-left {
    margin-top: -20px;
    margin-left: 100px
  }

  .guiding-text-up, .guiding-text-left {
    color: #D0D0D0;
    font-feature-settings: 'clig' off, 'liga' off;
    font-family: Indie Flower;
    transform: rotate(-2.663deg);
    font-size: 36px;
    font-style: normal;
    font-weight: 400;
    line-height: 32px; /* 88.889% */
    letter-spacing: -0.25px;
    text-transform: uppercase;
  }

  .guiding-text-left {
    margin-top: -40px;
    margin-left: 100px;
  }
</style>