<script lang="ts">
    import type {AlertError, Session} from "$lib/types";
    import SetupSession from "$lib/components/screens/setup/SetupSession.svelte";
    import {EditSession, GetSession, LocateSessions} from "$lib/wailsjs/go/kuddle/App";
    import {currentSessionName, encryptionKey, sessionNames} from "$lib/states";
    import ErrorAlert from "$lib/components/ErrorAlert.svelte";
    import {onMount} from "svelte";
    import * as Page from "$lib/components/page/components";
    import ProgrammaticableRedirect from "$lib/components/ProgrammaticableRedirect.svelte";
    import Loading from "$lib/components/screens/Loading.svelte";

    let originalSession: Session
    let session: Session
    let originalSessionName: string
    let showSecrets = false

    $: {
        if ($currentSessionName !== originalSessionName) {
            reloadSession()
        }
    }

    onMount(async () => {
        if ($encryptionKey == null || $currentSessionName == null) {
            goHome()
            return
        }

        await reloadSession()
    })

    async function reloadSession() {
        let selectedSession = await GetSession($currentSessionName!)
        originalSession = {
            name: selectedSession.name,
            secretAccessKey: selectedSession.secret_access_key,
            useSsl: selectedSession.use_ssl,
            token: selectedSession.token,
            endpoint: selectedSession.endpoint,
            domain: selectedSession.domain,
            accessKeyId: selectedSession.access_key_id,
            bucket: selectedSession.bucket
        }
        originalSessionName = originalSession.name
        session = originalSession
    }

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
        await EditSession($encryptionKey, originalSessionName, {
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
        $currentSessionName = session.name
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

{#if originalSession == null}
    <Loading/>
{:else}
    <Page.Container>
        <ProgrammaticableRedirect id="_home" href="/"/>
        <ErrorAlert bind:error/>
        <Page.Title>Editing {originalSession?.name ?? ""}</Page.Title>
        {#key originalSession}
            <SetupSession
                    defaultSession={originalSession}
                    on:toggle={toggleShowSecrets}
                    bind:showSecrets
                    on:validation={accept}
                    on:next={nextStage}
                    bind:session
            />
        {/key}
    </Page.Container>
{/if}
