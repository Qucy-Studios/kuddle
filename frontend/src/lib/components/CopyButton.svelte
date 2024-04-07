<script lang="ts">
    import {ClipboardSetText} from "$lib/wailsjs/runtime";
    import {Copy} from "@steeze-ui/radix-icons";
    import {Icon} from "@steeze-ui/svelte-icon";
    import {CopyCheck} from "@steeze-ui/lucide-icons";

    export let value: string
    let copied: boolean = false

    function copy() {
        if (copied) return
        ClipboardSetText(value)
        copied = true

        setTimeout(() => copied = false, 500)
    }
</script>
<button on:click={copy} class="flex flex-row gap-4 hover:brightness-90 transition ease-in-out duration-500 items-center bg-zinc-300 bg-opacity-30 py-1 px-4 rounded border border-zinc-300">
    <b>
        {#if copied}
            Copied
        {:else}
            <slot/>
        {/if}
    </b>
    {#if copied}
        <Icon src={CopyCheck} size="18"/>
    {:else}
        <Icon src={Copy} size="18"/>
    {/if}
</button>