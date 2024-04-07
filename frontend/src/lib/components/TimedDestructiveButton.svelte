<script lang="ts">
    import {createEventDispatcher, onDestroy, onMount} from "svelte";
    import {Button} from "$lib/components/ui/button";

    let timer = 5
    let isTimerOver = false
    let countdownTimer: number | undefined;

    $: {
        if (timer === 0) {
            isTimerOver = true
            clearInterval(countdownTimer)
        }
    }

    const dispatch = createEventDispatcher()
    export let classNames = "w-40"

    onMount(() => {
        countdownTimer = setInterval(() => {
            if (timer > 0) {
                timer--
            }
        }, 1_000)
    })

    onDestroy(() => {
        clearInterval(countdownTimer)
    })

    function onClick() {
        dispatch('click')
    }
</script>
<Button variant="destructive" disabled={!isTimerOver} on:click={onClick} class={classNames}>
    {#if !isTimerOver}
        {timer}
    {:else}
        <slot/>
    {/if}
</Button>