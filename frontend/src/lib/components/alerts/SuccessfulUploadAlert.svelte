<script lang="ts">
    import {Cross1, Rocket} from "@steeze-ui/radix-icons";
    import type {UploadResult} from "$lib/types";
    import {Icon} from "@steeze-ui/svelte-icon";
    import Copy from "$lib/components/Copy.svelte";
    import CopyButton from "$lib/components/CopyButton.svelte";

    export let result: UploadResult | null

    function html() {
        return `<img src="${result?.imageLink ?? ""}" alt=""/>`
    }

    function markdown() {
        return `![](${result?.imageLink ?? ""})`
    }
</script>

{#if result != null}
    <div class="border rounded-2xl bg-zinc-200 bg-opacity-60 px-8 py-4 mb-4">
        <div class="flex flex-row justify-between">
            <div class="flex flex-row gap-2 items-center">
                <Icon src={Rocket} size="18" class="text-green-500"/>
                <h3 class="text-lg font-bold text-green-500">Image successfully uploaded</h3>
            </div>
            <button on:click={() => result = null}>
                <Icon src={Cross1} size="16" class="text-red-500"/>
            </button>
        </div>
        <p>We've uploaded the image to your storage bucket and copied the link to your clipboard.</p>
        <div class="flex flex-row gap-2 items-center py-2">
            <CopyButton value={result.imageLink}>Copy image link</CopyButton>
            <CopyButton value={markdown()}>Copy as Markdown</CopyButton>
            <CopyButton value={html()}>Copy as HTML</CopyButton>
        </div>
    </div>
{/if}