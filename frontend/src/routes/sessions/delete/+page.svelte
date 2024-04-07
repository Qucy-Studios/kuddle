<script lang="ts">
    import type {AlertError, Session} from "$lib/types";
    import SetupSession from "$lib/components/screens/setup/SetupSession.svelte";
    import {DeleteSession, EditSession, GetSession, LocateSessions} from "$lib/wailsjs/go/kuddle/App";
    import {currentSessionName, encryptionKey, sessionNames} from "$lib/states";
    import ErrorAlert from "$lib/components/ErrorAlert.svelte";
    import {onDestroy, onMount} from "svelte";
    import * as Page from "$lib/components/page/components";
    import ProgrammaticableRedirect from "$lib/components/ProgrammaticableRedirect.svelte";
    import Loading from "$lib/components/screens/Loading.svelte";
    import {Button} from "$lib/components/ui/button";
    import TimedDestructiveButton from "$lib/components/TimedDestructiveButton.svelte";

    onMount(async () => {
        if ($encryptionKey == null || $currentSessionName == null) {
            goHome()
            return
        }

        // you cannot delete your only session
        if (($sessionNames?.length ?? 0)< 2) {
            goHome()
            return
        }
    })

    // reset error
    let error: AlertError | null = null

    async function del() {
        if ($encryptionKey == null) {
            goHome()
            return
        }

        error = null

        await DeleteSession($currentSessionName!)
        $sessionNames = await LocateSessions($encryptionKey)
        $currentSessionName = $sessionNames?.at(0) ?? null

        goHome()
        return
    }

    function goHome() {
        document.getElementById("_home")?.click()
    }

    function discard() {
        goHome()
    }
</script>

<Page.Container>
    <ProgrammaticableRedirect id="_home" href="/"/>
    <ErrorAlert bind:error/>
    <Page.Header>
        <Page.Title>Deleting {$currentSessionName ?? "Unknown"}</Page.Title>
        <Page.Description>
            <p>Are you sure you want to delete this session? Deleting this session means that
                all data including the tokens, and credentials will be erased. Recent upload histories will not
                be touched.
                Please note that this is a <b>irreversible action</b>.
            </p>
        </Page.Description>
    </Page.Header>
    <div class="justify-end my-4 flex flex-row gap-3">
        <Button variant="secondary" on:click={discard} class="w-40">
            Cancel
        </Button>
        <TimedDestructiveButton on:click={del}>
            Delete
        </TimedDestructiveButton>
    </div>
</Page.Container>
