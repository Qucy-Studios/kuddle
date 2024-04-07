<script lang="ts">
    import {Cube, Link1, Link2, LockClosed, Person} from "@steeze-ui/radix-icons";
    import {Key} from "@steeze-ui/lucide-icons";
    import SecretInputField from "$lib/components/SecretInputField.svelte";
    import {Button} from "$lib/components/ui/button";
    import * as Tabs from "$lib/components/ui/tabs"
    import InputField from "$lib/components/InputField.svelte";
    import * as Page from "$lib/components/page/components"
    import {useSetupDispatchEvents, validateNotEmpty} from "$lib/components/screens/setup/utils";
    import type {Session} from "$lib/types";
    import {httpUrlRegex} from "$lib/regex";
    import {sessionNames} from "$lib/states";
    import {Test} from "$lib/wailsjs/go/kuddle/App";
    import {nonNullOrDefault} from "$lib/utils";

    export let session: Session
    export let defaultSession: Session | null

    let sessionName: string = nonNullOrDefault<string>(defaultSession?.name, "")
    let domain: string = nonNullOrDefault<string>(defaultSession?.domain, "")

    let storageEndpoint: string = nonNullOrDefault<string>(defaultSession?.endpoint, "")
    let bucketName: string = nonNullOrDefault<string>(defaultSession?.bucket, "")
    let accessKeyId: string = nonNullOrDefault<string>(defaultSession?.accessKeyId, "")
    let secretAccessKey: string = nonNullOrDefault<string>(defaultSession?.secretAccessKey, "")
    let secretToken: string =  nonNullOrDefault<string>(defaultSession?.token, "")

    let isLoading = false
    export let showSecrets: boolean

    let {nextStage, toggleShowSecrets} = useSetupDispatchEvents(async () => {
        isLoading = true
        try {
            if (!validateNotEmpty(sessionName, domain, storageEndpoint, bucketName) ||
                (currentTab == "access_key" && secretAccessKey == "" || accessKeyId == "") ||
                (currentTab == "access_token"  && secretToken == "")) {
                console.info('err')
                return {
                    title: "You have fields that are not filled in.",
                    description: "All the fields must be filled in to create your session. If there are fields that you do not understand, please refer to our documentations."
                }
            }

            if (!httpUrlRegex.test(domain)) {
                return {
                    title: "Invalid domain.",
                    description: "Domain must be a valid URL that points to the bucket, this will be used to generate the links to the images."
                }
            }

            if (!httpUrlRegex.test(storageEndpoint)) {
                return {
                    title: "Invalid storage endpoint.",
                    description: "Storage endpoint must be a valid URL that points to the S3-compatible API of the storage, such as AWS S3 or Cloudflare R2."
                }
            }

            if (currentTab == "access_token") {
                secretAccessKey = ""
                accessKeyId = ""
            }

            if (currentTab == "access_key") {
                secretToken = ""
            }

            let useSsl = !storageEndpoint.startsWith("http://") // automatically assume HTTPS if without any prefix.
            session = {
                accessKeyId,
                bucket: bucketName,
                domain,
                endpoint: storageEndpoint,
                name: sessionName,
                secretAccessKey,
                token: secretToken,
                useSsl
            }

            try {
                const isValidSession = await Test( {
                    name: session.name,
                    access_key_id: session.accessKeyId,
                    bucket: session.bucket,
                    domain: session.domain,
                    endpoint: session.endpoint,
                    secret_access_key: session.secretAccessKey,
                    token: session.token,
                    use_ssl: session.useSsl
                })

                if (!isValidSession) {
                    return {
                        title: "Invalid session credentials.",
                        description: "It seems like you've typed invalid credentials to the storage platform. Please try double-checking and trying again."
                    }
                }
            } catch (e) {
                return {
                    title: "Failed to test session.",
                    description: "We encountered an error while trying to test session: " + e
                }
            }
            return undefined
        } finally {
            isLoading = false
        }
    })
    let currentTab = "access_key"

    function select(tab: string) {
        currentTab = tab
    }
</script>

<Page.Container classNames="gap-2">
    <Page.Description>
        {#if $sessionNames != null}
            <p>
                Kuddle uploads all the images directly to your storage platform, such as AWS S3, or Cloudflare R2 without
                passing through any middleman as it runs directly through your computer. As such, we require essential
                and secret information, such as your <b>bucket endpoint</b>, <b>bucket name</b>, <b>access token</b> and <b>secret key</b>.
            </p>
        {:else}
            <p>
                You can add up to 5 sessions maximum per computer. Sessions are encrypted, and no information about them
                is sent to any telemetry, or anywhere online other than the cloud provider, or the storage platform that
                needs them to upload the images.
            </p>
        {/if}
    </Page.Description>
    <InputField icon={Person} bind:value={sessionName} placeholder="Session Name"/>
    <InputField icon={Link2} bind:value={domain} placeholder="Public Domain"/>
    <InputField icon={Link1} bind:value={storageEndpoint} placeholder="Storage Endpoint"/>
    <InputField icon={Cube} bind:value={bucketName} placeholder="Bucket Name"/>
    <Tabs.Root value={currentTab} onValueChange={(value) => select(value ?? 'access_key')}>
        <Tabs.List>
            <Tabs.Trigger value="access_key">Access Keys</Tabs.Trigger>
            <Tabs.Trigger value="access_token">Tokens</Tabs.Trigger>
        </Tabs.List>
        <Tabs.Content value="access_key">
            <div class="flex flex-col gap-2">
                <SecretInputField
                        icon={LockClosed}
                        bind:value={accessKeyId}
                        placeholder="Access Key Id"
                        bind:showSecrets
                        on:toggle={toggleShowSecrets}
                />
                <SecretInputField
                        icon={Key}
                        bind:value={secretAccessKey}
                        placeholder="Secret Access Key"
                        bind:showSecrets
                        on:toggle={toggleShowSecrets}
                />
            </div>
        </Tabs.Content>
        <Tabs.Content value="access_token">
            <div class="flex flex-col gap-2">
                <SecretInputField
                        icon={Key}
                        bind:value={secretToken}
                        placeholder="Secret Token"
                        bind:showSecrets
                        on:toggle={toggleShowSecrets}
                />
            </div>
        </Tabs.Content>
    </Tabs.Root>
    <div class="justify-end my-4 flex flex-row gap-3">
        <Button disabled={isLoading} on:click={nextStage} class="w-40">
            {#if !isLoading}
                Next
            {:else}
                Loading...
            {/if}
        </Button>
    </div>
</Page.Container>