<script>
    import {wasAutoLoggedOut} from "$lib/states";
    import ErrorAlert from "$lib/components/ErrorAlert.svelte";
</script>
{#if $wasAutoLoggedOut}
    <ErrorAlert error={{
            title: "You were logged out.",
            description: "Sessions are automatically logged out after a minute of out-of-focus activity. You may login again to continue."
        }}/>
{/if}