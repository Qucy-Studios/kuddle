<script lang="ts">
    import SetupEncryptionKey from "$lib/components/screens/setup/SetupEncryptionKey.svelte";
    import type {AlertError, Session} from "$lib/types";
    import SetupSession from "$lib/components/screens/setup/SetupSession.svelte";
    import {CreateSession, LocateSessions, Register} from "$lib/wailsjs/go/kuddle/App";
    import {currentSessionName, encryptionKey, hasRegistered, sessionNames} from "$lib/states";
    import ErrorAlert from "$lib/components/ErrorAlert.svelte";
    import * as Page from "$lib/components/page/components"

    let encryptKey = ""
    let session: Session

    let showSecrets = false

    // reset error
    let error: AlertError | null = null

    function toggleShowSecrets() {
        showSecrets = !showSecrets
    }

    async function nextStage() {
        error = null
        if (stage == 1) {
            await Register(encryptKey)
            await CreateSession(encryptKey, {
                name: session.name,
                access_key_id: session.accessKeyId,
                bucket: session.bucket,
                domain: session.domain,
                endpoint: session.endpoint,
                secret_access_key: session.secretAccessKey,
                token: session.token,
                use_ssl: true
            })
            let sessions = await LocateSessions(encryptKey)
            $encryptionKey = encryptKey
            $sessionNames = sessions
            $currentSessionName = sessions[0]
            $hasRegistered = true
            return
        }
        stage++
    }

    function accept(err: CustomEvent<AlertError>) {
        error = err.detail
    }

    let stage = 0
</script>

<Page.Container>
    <ErrorAlert bind:error/>
    <Page.Title>Setup</Page.Title>
    {#if stage === 0}
        <SetupEncryptionKey
                on:toggle={toggleShowSecrets}
                bind:showSecrets
                bind:encryptionKey={encryptKey}
                on:validation={accept}
                on:next={nextStage}
        />
    {:else if stage === 1}
        <SetupSession
                on:toggle={toggleShowSecrets}
                bind:showSecrets
                on:validation={accept}
                on:next={nextStage}
                bind:session
        />
    {/if}
</Page.Container>