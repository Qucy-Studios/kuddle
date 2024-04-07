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

        setTimeout(() => copied = false, 1_500)
    }
</script>
{#if copied}
    <Icon src={CopyCheck} size="18"/>
{:else}
    <button class="group-hover:brightness-110 transition duration-500" on:click={copy}>
        <Icon src={Copy} size="18"/>
    </button>
{/if}