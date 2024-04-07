<script lang="ts">
    import Copy from "$lib/components/Copy.svelte";
    import type {UploadResult} from "$lib/types";

    export let upload: UploadResult
</script>
<div class="flex flex-row gap-2 items-center">
    <img src="{upload.imageLink}" class="size-[6.5rem] object-cover border border-zinc-300 rounded text-transparent" alt="Upload Preview"/>
    <div class="border border-zinc-300 bg-zinc-300 bg-opacity-60 rounded p-4 w-full h-full">
        <div class="flex flex-row justify-between items-center">
            <ul>
                <li><b>Link:</b> {upload.imageLink}</li>
                <li><b>Session:</b> {upload.session}</li>
                <li><b>Date:</b> {upload.creationTime.toLocaleString()}</li>
            </ul>
            <Copy value={upload.imageLink}/>
        </div>
    </div>
</div>