<script lang="ts">
    import type {UploadResult} from "$lib/types";
    import {onMount} from "svelte";
    import * as Page from "$lib/components/page/components";
    import {getRecentUploads} from "$lib/recents";
    import {encryptionKey} from "$lib/states";
    import Copy from "$lib/components/Copy.svelte";
    import ProgrammaticableRedirect from "$lib/components/ProgrammaticableRedirect.svelte";
    import RecentUpload from "$lib/components/recents/RecentUpload.svelte";

    let recentUploads: UploadResult[] = []


    onMount(async () => {
        if ($encryptionKey == null) {
            goHome()
            return
        }

        recentUploads = await getRecentUploads()
    })

    function goHome() {
        document.getElementById("_home")?.click()
    }
</script>

<Page.Container classNames="mb-8">
    <ProgrammaticableRedirect id="_home" href="/"/>
    <Page.Header>
        <Page.Title>Recent Uploads</Page.Title>
        <Page.Description>
            <p>
                You can find and copy your recently uploaded files here. By pressing the <b>Copy Icon</b> beside the
                entry, you can copy the link that was generated for the image.
            </p>
        </Page.Description>
    </Page.Header>
    <div class="flex flex-col gap-2 mt-4">
        {#each recentUploads as upload}
            <RecentUpload upload={upload}/>
        {/each}
    </div>
</Page.Container>