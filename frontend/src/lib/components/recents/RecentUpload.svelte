<script lang="ts">
    import Copy from "$lib/components/Copy.svelte";
    import type {UploadResult} from "$lib/types";
    import {htmlImage, markdownImage} from "$lib/utils";
    import CopyButton from "$lib/components/CopyButton.svelte";

    export let upload: UploadResult
</script>
<div class="flex flex-row items-center">
    <img src="{upload.imageLink}" class="w-[100px] h-[140px] object-cover border border-zinc-300 rounded rounded-r-none text-transparent" alt="Upload Preview"/>
    <div class="border border-zinc-300 bg-zinc-300 bg-opacity-60 rounded rounded-l-none p-4 w-full h-full">
        <div class="flex flex-row justify-between items-center">
            <ul>
                <li><b>Link:</b> {upload.imageLink}</li>
                <li><b>Session:</b> {upload.session}</li>
                <li><b>Date:</b> {upload.creationTime.toLocaleString()}</li>
            </ul>
        </div>
        <div class="flex flex-row items-center gap-4">
            <CopyButton value={upload.imageLink}>Copy link</CopyButton>
            <CopyButton value={markdownImage(upload.imageLink)}>Copy Markdown</CopyButton>
            <CopyButton value={htmlImage(upload.imageLink)}>Copy HTML</CopyButton>
        </div>
    </div>
</div>