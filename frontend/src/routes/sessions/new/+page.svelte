<script lang="ts">
    import type {AlertError, Session} from "$lib/types";
    import SetupSession from "$lib/components/screens/setup/SetupSession.svelte";
    import {CreateSession, LocateSessions} from "$lib/wailsjs/go/kuddle/App";
    import {encryptionKey, sessionNames} from "$lib/states";
    import ErrorAlert from "$lib/components/ErrorAlert.svelte";
    import {onMount} from "svelte";
    import * as Page from "$lib/components/page/components";
    import ProgrammaticableRedirect from "$lib/components/ProgrammaticableRedirect.svelte";

    let session: Session
    let showSecrets = false

    onMount(async () => {
        if ($encryptionKey == null) {
            goHome()
            return
        }
    })

    // reset error
    let error: AlertError | null = null

    function toggleShowSecrets() {
        showSecrets = !showSecrets
    }

    async function nextStage() {
        if ($encryptionKey == null) {
            goHome()
            return
        }

        error = null
        await CreateSession($encryptionKey, {
            name: session.name,
            access_key_id: session.accessKeyId,
            bucket: session.bucket,
            domain: session.domain,
            endpoint: session.endpoint,
            secret_access_key: session.secretAccessKey,
            token: session.token,
            use_ssl: true
        })

        $sessionNames = await LocateSessions($encryptionKey)
        goHome()
        return
    }

    function goHome() {
        document.getElementById("_home")?.click()
    }

    function accept(err: CustomEvent<AlertError>) {
        error = err.detail
    }
</script>

<Page.Container>
    <ProgrammaticableRedirect id="_home" href="/"/>
    <ErrorAlert bind:error/>
    <Page.Title>Add New Session</Page.Title>
    <SetupSession
            defaultSession={null}
            on:toggle={toggleShowSecrets}
            bind:showSecrets
            on:validation={accept}
            on:next={nextStage}
            bind:session
    />
</Page.Container>