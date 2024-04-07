<script lang="ts">
    import {Cross1, MinusCircled} from "@steeze-ui/radix-icons";
    import type {AlertError} from "$lib/types";
    import {Icon} from "@steeze-ui/svelte-icon";

    export let error: AlertError | null
</script>

{#if error != null}
    <div class="border rounded-2xl bg-zinc-200 bg-opacity-60 px-8 py-4 mb-4">
        <div class="flex flex-row justify-between">
            <div class="flex flex-row gap-2 items-center">
                <Icon src={MinusCircled} size="18" class="text-red-500"/>
                <h3 class="text-lg font-bold text-red-500">{error.title}</h3>
            </div>
            <button on:click={() => error = null}>
                <Icon src={Cross1} size="16" class="text-red-500"/>
            </button>
        </div>
        <p>{error.description}</p>
    </div>
{/if}