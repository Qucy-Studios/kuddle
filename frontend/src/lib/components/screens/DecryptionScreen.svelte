<script lang="ts">
    import {LockClosed} from "@steeze-ui/radix-icons";
    import {Button} from "$lib/components/ui/button";
    import SecretInputField from "$lib/components/SecretInputField.svelte";
    import {Authenticate, LocateSessions} from "$lib/wailsjs/go/kuddle/App";
    import type {AlertError} from "$lib/types";
    import {
        currentSessionName,
        encryptionKey,
        hasRegistered,
        isOffline,
        sessionNames,
        wasAutoLoggedOut
    } from "$lib/states";
    import {WindowMinimise} from "$lib/wailsjs/runtime";
    import ErrorAlert from "$lib/components/ErrorAlert.svelte";
    import AutoLoggedOutAlert from "$lib/components/alerts/AutoLoggedOutAlert.svelte";
    import OfflineAlert from "$lib/components/alerts/OfflineAlert.svelte";
    import * as Page from "$lib/components/page/components"

    let error: AlertError | null = null

    let showSecrets: boolean = false
    let key: string = ""

    let isLoading = false

    async function decrypt() {
        if (key == "") {
            error = {
                title: "No encryption key given.",
                description: "You did not fill any encryption key in the field."
            }
            return
        }

        isLoading = true
        try {
            try {
                let isAuthenticated = await Authenticate(key)
                if (!isAuthenticated) {
                    error = {
                        title: "Invalid encryption key.",
                        description: "It seems like you've typed the wrong encryption key. Please try again."
                    }
                    return
                }
            } catch (e) {
                console.error(e)
                return
            }

            let sessions = await LocateSessions(key)
            $hasRegistered = true
            $sessionNames = sessions
            $currentSessionName = sessions.at(0)
            $encryptionKey = key
            $wasAutoLoggedOut = false
        } finally {
            isLoading = false
        }
    }

    function toggleShowSecrets() {
        showSecrets = !showSecrets
    }

    function onKeyPress(evt: KeyboardEvent) {
        if (evt.key === "Enter") {
            decrypt()
        }
    }
</script>

<svelte:window on:keypress={onKeyPress}/>
<Page.Container>
    <ErrorAlert bind:error/>
    <OfflineAlert/>
    <AutoLoggedOutAlert/>
    <Page.Header>
        <Page.Title>Single-factor Authentication</Page.Title>
        <Page.Description>
            <p>
                Kuddle keeps session information encrypted to protect your storage keys from being exposed which can lead to
                various security risks such as allowing unauthorized people with access to your personal computer's storage knowledge
                to your storage keys, granting them access to it and the ability to delete, upload and know files inside the storage.
            </p>
        </Page.Description>
    </Page.Header>
    <div class="mt-2">
        <SecretInputField
                on:toggle={toggleShowSecrets}
                bind:showSecrets
                placeholder="Enter your encryption key"
                icon={LockClosed}
                bind:disabled={$isOffline}
                bind:value={key}
        />
    </div>
    <div class="fixed bottom-0 right-0 my-9 mx-8">
        <div class="flex flex-row gap-3">
            <Button on:click={WindowMinimise} variant="secondary" class="w-40 bg-zinc-200 hover:bg-zinc-300">
                Cancel
            </Button>
            <Button disabled={isLoading} on:click={decrypt} class="w-40">
                {#if isLoading}
                    Loading...
                {:else}
                    Decrypt
                {/if}
            </Button>
        </div>
    </div>
</Page.Container>