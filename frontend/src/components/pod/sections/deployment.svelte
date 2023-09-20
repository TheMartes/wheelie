<script>
  import { Circle } from "svelte-loading-spinners";
  import { GetDeployment } from "../../../../wailsjs/go/main/App";
  import hljs from "highlight.js/lib/core";
  import yaml from "highlight.js/lib/languages/yaml";
  import ContentCard from "../../utils/content-card.svelte";
  // import Push from "../../utils/icons/push.svelte";
  // import Refresh from "../../utils/icons/refresh.svelte";
  // import Trash from "../../utils/icons/trash.svelte";

  /** @type {string} */
    export let namespace;
  /** @type {string} */
    export let pod;
  let editMode = false;
  /** @type {string} */
  let originalContent;
  /** @type {HTMLTextAreaElement} */
  let ta;
  /** @type {boolean} */
  let contentModified = false;
  /** @type {Promise<string>}*/
  let deployment = GetDeployment(pod, namespace);

  hljs.registerLanguage("yaml", yaml);

  /** @param {string} output */
  const highlightOutput = (output) => {
      return hljs.highlight(output,{ language: "yaml" }).value;
  };

  // eslint-disable-next-line no-unused-vars
  const enableEditMode = (ogcontent) => {
      // originalContent = ogcontent
      // editMode = true
  };

    const handleKeydown = async (e) => {
        if (e.keyCode === 27) { // Escape
            editMode = false;
        }
        if (e.keyCode === 13) { // Escape
            if (ta.value !== originalContent) {
                contentModified = ta.value !== await deployment;
            }

            originalContent = ta.value;
            editMode = false;
        }
    };
</script>
<div class="deployment-container">
  <ContentCard cardTitle={editMode ? "Editing Deployment" : "Preview"} cardDescription="" width="100%">
    <div class="toolbar" slot="toolbar">
      <!--      <Refresh fill="#000000"></Refresh>-->
      <!--      <div class="divider"></div>-->
      <!--      <Trash></Trash>-->
      <!--      <div class="divider"></div>-->
      <!--      <button disabled={!editMode || !contentModified}><Push height="16px" width="16px" fill="#FFFFFF"></Push><span class="patch">Patch</span></button>-->
    </div>
    <div class="content" slot="content">
      {#await deployment}
      <div class="bg-message">
          <Circle size="60" color="#2E77D0" unit="px" duration="1s" />
      </div>
      {:then deployment}
        {#key editMode}
        <div>
          {#if !editMode}
          <!-- svelte-ignore a11y-click-events-have-key-events -->
          {#if contentModified}
            <pre on:click={() => enableEditMode(originalContent)} ><code class="language-yaml">{@html highlightOutput(originalContent)}</code></pre>
          {:else}
            <pre on:click={() => enableEditMode(deployment)} ><code class="language-yaml">{@html highlightOutput(deployment)}</code></pre>
          {/if}
          {:else}
            <textarea bind:this={ta} on:keydown={handleKeydown} rows={originalContent.split("\n").length}>{originalContent}</textarea>
          {/if}
        </div>
        {/key}
      {/await}
      </div>
  </ContentCard>
</div>
<style>
  .deployment-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: calc(100% - 40px);
    padding: 20px;
  }

  .content {
    padding: 0 20px
  }

  .toolbar {
    display: flex;
    align-items: center;
  }

  textarea {
    width: 100%;
    font-family: monospace;
    font-size: 1rem;
  }


  /*button {*/
  /*  background-color: var(--selectmenu-bg);*/
  /*  color: white;*/
  /*  border-radius: 5px;*/
  /*  cursor: pointer;*/
  /*  font-size: 16px;*/
  /*  padding: 5px 10px;*/
  /*  border: 1px solid var(--content-card-border);*/
  /*  vertical-align: middle;*/
  /*}*/

  /*button:disabled {*/
  /*  background-color: var(--content-card-border);*/
  /*}*/

  /*.patch {*/
  /*  margin-left: 5px;*/
  /*}*/
  /*.divider {*/
  /*  width: 10px;*/
  /*}*/
</style>
