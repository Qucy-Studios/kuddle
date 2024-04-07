<script lang="ts">
    import {RulerHorizontal} from "@steeze-ui/radix-icons";
    import {Folder} from "@steeze-ui/lucide-icons";
    import NumericInputField from "$lib/components/NumericInputField.svelte";
    import InputField from "$lib/components/InputField.svelte";
    import {createEventDispatcher} from "svelte";
    import {Button} from "$lib/components/ui/button";
    import * as Page from "$lib/components/page/components"

    export let ephemeralNameLength: number
    export let ephemeralFolderPath: string

    const dispatch = createEventDispatcher()
    function goUploadScreen() {
        dispatch('navigate', "upload")
    }

    function saveSettings() {
        dispatch('save', {
            urlLength: ephemeralNameLength,
            folder: ephemeralFolderPath
        })
        dispatch('navigate', 'upload')
    }
</script>
<Page.Header>
    <Page.Title>Current Session Settings</Page.Title>
    <Page.Description>
        <p>
            You can configure the ephemeral settings for this session. It will be reset when you go out of
            the uploading page, or get logged out of the session.
        </p>
    </Page.Description>
</Page.Header>
<NumericInputField
        bind:value={ephemeralNameLength}
        placeholder="Enter the length of random file name."
        icon={RulerHorizontal}
        tooltip="Length of file name"/>
<InputField
        bind:value={ephemeralFolderPath}
        placeholder="Enter the folder to place the file, leave empty for none."
        icon={Folder}
        tooltip="Folder Path"/>
<div class="justify-end my-4 flex flex-row gap-3">
    <Button variant="secondary" on:click={goUploadScreen} class="w-40">
        Cancel
    </Button>
    <Button on:click={saveSettings} class="w-40">
        Save
    </Button>
</div>