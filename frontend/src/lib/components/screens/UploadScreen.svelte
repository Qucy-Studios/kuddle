<script lang="ts">
    import {Gear} from "@steeze-ui/radix-icons";
    import {Icon} from "@steeze-ui/svelte-icon";
    import SuccessfulUploadAlert from "$lib/components/alerts/SuccessfulUploadAlert.svelte";
    import {Button} from "$lib/components/ui/button";
    import type {UploadResult} from "$lib/types";
    import {createEventDispatcher} from "svelte";
    import * as Page from "$lib/components/page/components"

    export let uploadResult: UploadResult | null = null

    export let imageSrc: string | null = null
    export let isLoading: boolean = false

    const dispatch = createEventDispatcher()

    function goSettingsScreen() {
        dispatch('navigate', 'settings')
    }

    function upload() {
        dispatch('upload')
    }

    function discard() {
        dispatch('discard')
    }
</script>

<SuccessfulUploadAlert bind:result={uploadResult}/>
<Page.Header>
    <Page.Title>Upload an Image</Page.Title>
    <Page.Description>
        <p>
            Use the command <b>CTRL+C</b> and <b>CTRL+V</b>, otherwise known as Copy-Paste, to upload a file into here.
            We will generate a random link for that image on your bucket and give you the link.
        </p>
    </Page.Description>
</Page.Header>
{#if imageSrc == null}
    <div class="w-full p-4 h-96 bg-zinc-300 object-cover">
    </div>
{:else}
    <img alt="Preview" src={imageSrc} class="w-full h-96 object-cover"/>
{/if}
<div class="justify-end my-4 flex flex-row gap-3">
    <Button variant="secondary" on:click={goSettingsScreen} class="w-fit">
        <Icon src={Gear} size="18"/>
    </Button>
    <Button disabled={isLoading || imageSrc == null} variant="secondary" on:click={discard} class="w-40">
        Discard
    </Button>
    <Button disabled={isLoading || imageSrc == null} on:click={upload} class="w-40">
        {#if isLoading}
            Loading...
        {:else}
            Upload
        {/if}
    </Button>
</div>