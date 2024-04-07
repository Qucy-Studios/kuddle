<script lang="ts">
    import {Cross1, MinusCircled, Rocket} from "@steeze-ui/radix-icons";
    import type {AlertError} from "$lib/types";
    import {Icon} from "@steeze-ui/svelte-icon";

    export let alert: AlertError | null
</script>

{#if alert != null}
    <div class="border rounded-2xl bg-zinc-200 bg-opacity-60 px-8 py-4 mb-4">
        <div class="flex flex-row justify-between">
            <div class="flex flex-row gap-2 items-center">
                <Icon src={Rocket} size="18" class="text-green-500"/>
                <h3 class="text-lg font-bold text-green-500">{alert.title}</h3>
            </div>
            <button on:click={() => alert = null}>
                <Icon src={Cross1} size="16" class="text-red-500"/>
            </button>
        </div>
        <p>{alert.description}</p>
    </div>
{/if}