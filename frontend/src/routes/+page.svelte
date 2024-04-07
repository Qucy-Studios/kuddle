<script lang="ts">
    import {Icon} from "@steeze-ui/svelte-icon";
    import {Symbol} from "@steeze-ui/radix-icons";
    import {encryptionKey, hasRegistered} from "$lib/states";
    import DecryptionScreen from "$lib/components/screens/DecryptionScreen.svelte";
    import {onMount} from "svelte";
    import {HasRegistered} from "$lib/wailsjs/go/kuddle/App";
    import Setup from "$lib/components/screens/Setup.svelte";
    import SessionScreen from "$lib/components/screens/SessionScreen.svelte";

    let isLoading = true
    onMount(async () => {
        $hasRegistered = await HasRegistered()
        isLoading = false
    })
</script>

{#if isLoading}
    <div class="h-screen w-full flex flex-col justify-center align-middle items-center">
        <div class="mb-2 flex flex-col items-center animate-pulse">
            <Icon src={Symbol} size="18" class="animate-spin mb-2"></Icon>
            <p class="text-xs text-zinc-400 sofia">Loading</p>
        </div>
    </div>
{:else}
    {#if $hasRegistered}
        {#if $encryptionKey == null}
            <DecryptionScreen/>
        {:else}
            <SessionScreen/>
        {/if}
    {:else}
        <Setup/>
    {/if}
{/if}