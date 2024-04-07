<script>
    import {isOffline} from "$lib/states";
    import ErrorAlert from "$lib/components/ErrorAlert.svelte";
</script>
{#if $isOffline}
    <ErrorAlert error={{
            title: "You are offline.",
            description: "Kuddles is built to interact with off-platform storage services. Without internet, Kuddles is useless and cannot be used."
        }}/>
{/if}