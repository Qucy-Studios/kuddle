<script lang="ts">
    import {LockClosed} from "@steeze-ui/radix-icons";
    import {Key} from "@steeze-ui/lucide-icons";
    import SecretInputField from "$lib/components/SecretInputField.svelte";
    import {Button} from "$lib/components/ui/button";
    import type {AlertError} from "$lib/types";
    import * as Page from "$lib/components/page/components"
    import {useSetupDispatchEvents} from "$lib/components/screens/setup/utils";

    export let encryptionKey: string
    let confirmationKey: string

    export let showSecrets: boolean

    let {nextStage, toggleShowSecrets} = useSetupDispatchEvents(async () => {
        if (confirmationKey != encryptionKey) {
            return {
                title: "Encryption keys doesn't match",
                description: "It seems like you mistyped one of the inputs. Please check that both inputs have the same encryption key and is correct."
            } as AlertError
        }
        return undefined
    })
</script>

<Page.Container classNames="gap-2">
    <Page.Description>
        <p>
            To protect your storage platforms from being exposed to hackers, we encrypt all credential data with the
            same encryption protocol that applications such as Proton Mail uses. Before we can proceed with creating
            your session, we'll need to register your encryption key first.
        </p>
    </Page.Description>
    <SecretInputField
            icon={LockClosed}
            bind:value={encryptionKey}
            placeholder="Enter your encryption key."
            bind:showSecrets
            on:toggle={toggleShowSecrets}
    />
    <SecretInputField
            icon={Key}
            bind:value={confirmationKey}
            placeholder="Confirm your encryption key."
            bind:showSecrets
            on:toggle={toggleShowSecrets}
    />
    <div class="fixed bottom-0 right-0 my-9 mx-8">
        <div class="flex flex-row gap-3">
            <Button on:click={nextStage} class="w-40">
                Next
            </Button>
        </div>
    </div>
</Page.Container>