<script>
    import { ChangeContext, GetKubernetesContext } from "../../../../wailsjs/go/main/App";
    import { createEventDispatcher } from "svelte";
    import { shortcut } from "../../../services/utils/shortcut.service";
    const dispatch = createEventDispatcher();

    /** @type {string|null} */
    export let selectedContext = null;
    /** @type {boolean} */
    export let expand = false;
    /** @type {HTMLInputElement} */
    let input;
    /** @type {Promise<Array<string>>} */
    let kubeContexts = GetKubernetesContext();
    /** @type {Array<string>} */
    let filteredContexts = [];
    /** @type {HTMLElement} */
    let selectedElement;
    /** @type {{[key: string]: HTMLElement}} */
    let filteredElements = {};

    filteredContexts.forEach(ctx => filteredElements[ctx] = null);

    const toggleContextSelect = () => expand = !expand;

    /**
     * @param {Array<string>} contexts
     */
    const filter = async (contexts) => {
        filteredContexts = contexts.filter(ctx => ctx !== "" && ctx);
        filteredContexts = filteredContexts.filter(ctx => ctx.startsWith(input.value));
        filteredContexts.forEach(ctx => filteredElements[ctx] = null);

        selectedElement = filteredElements[Object.keys(filteredElements)[0]];
        filteredElements[Object.keys(filteredElements)[0]].classList.add("selected");
    };

    /** @param {string} context */
    const loadContext = (context) => {
        selectedContext = context;
        dispatch("switchContext", {
            context
        });

        expand = false;
    };

    const toggleExpand = () => {
        if (selectedContext) {
            if (selectedElement) {
                loadContext(selectedElement.innerText);
            }
            expand = false;
        }

        if (selectedElement) {
            loadContext(selectedElement.innerText);
            expand = false;
        }
    };

    const handleKeydown = (e) => {
        if (e.keyCode == 27) { // Escape
            input.blur();
            filteredContexts = []; // reset
        }
    };

    const refreshFilter = (contexts) => {
        filteredContexts = contexts.filter(ctx => ctx != "" && ctx);
        filteredContexts.forEach(ctx => filteredElements[ctx] = null);
    };

    const removeSelectedElement = () => {
        selectedElement = null;
    };

    const addSelectedElement = (elements) => {
        selectedElement = elements[Object.keys(elements)[0]];
    };

    const init = (el) => {
        el.focus();
    };

    $: ChangeContext(selectedContext);
</script>
<!-- svelte-ignore a11y-click-events-have-key-events -->
{#if !expand}
<div use:shortcut={{control: true, code: "KeyK"}} class="selected-context" on:click={() => toggleContextSelect()}>
    {#if selectedContext === null}
        <span class="context-info-message">Missing Context</span>
    {:else}
        {selectedContext}
    {/if}
</div>
{:else}
    {#await kubeContexts}
        Loading...
    {:then contexts}
    <div class="context-selector"
        on:mouseenter={() => removeSelectedElement()}
        on:mouseleave={() => addSelectedElement(filteredElements)}
    >
        <input
            spellcheck="false"
            bind:this={input}
            on:keydown={handleKeydown}
            on:input={() => filter(contexts)}
            on:focus={() => refreshFilter(contexts)}
            placeholder="Search contexts..."
            type="text"
            on:blur={() => toggleExpand()}
            use:init
        />
        <div class="context-autocomplete">
            {#if filteredContexts.length > 0}
            {#each filteredContexts as context}
                <!-- svelte-ignore a11y-mouse-events-have-key-events -->
                <div
                    class="autocomplete-option"
                    on:mouseleave={() => selectedElement = null}
                    on:mouseover={() => selectedElement = filteredElements[context]}
                    bind:this={filteredElements[context]}
                    on:click={() => loadContext(context)}
                    >
                    {context}
                </div>
            {/each}
            {:else}
            {#each contexts.filter(ctx => ctx != "") as context}
                <!-- svelte-ignore a11y-mouse-events-have-key-events -->
                <div
                    class="autocomplete-option"
                    class:selected={selectedElement === filteredElements[context]}
                    on:mouseleave={() => selectedElement = null}
                    on:mouseover={() => selectedElement = filteredElements[context]}
                    on:click={() => loadContext(context)}
                    bind:this={filteredElements[context]}>
                    {context}
                </div>
            {/each}
            {/if}
        </div>
    </div>
    {/await}
{/if}
<style>
    .selected-context {
        background-color: var(--sidebar-bg);
        padding: 3px 30px;
        border: 1px solid var(--sidebar-divider);
        color: var(--sidebar-text);
        min-width: 200px;
        text-align: center;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 0.7rem;
    }

    .context-selector {
        z-index: 10;
    }

    .context-selector > input {
        background-color: #FFFFFF;
        width: 400px;
        padding: 5px;
        font-size: 1rem;
        margin: 0;
        border-radius: 5px 5px 0 0;
        border: 1px solid #FFFFFF;
    }

    .context-selector > input:focus {
        outline: 0
    }

    .context-autocomplete {
        background-color: #FFFFFF;
        position: fixed;
        width: 412px;
        margin: 0 0px;
        z-index: 10;
        margin-top: -1px;
        box-shadow: 0 4px 2px -2px gray;
        border-radius: 0px 0px 5px 5px;
    }

    .context-autocomplete > div {
        padding: 10px 11px;
    }

    .context-autocomplete > div:last-child {
        border-radius: 0px 0px 5px 5px;
    }

    .context-autocomplete > div:hover {
        background-color: var(--selectmenu-bg);
        color: #FFFFFF;
        cursor: pointer;
    }

    .autocomplete-option.selected {
        background-color: var(--selectmenu-bg);
        color: #FFFFFF;
    }

    .context-info-message {
        font-weight: bold;
        text-transform: uppercase;
    }
</style>