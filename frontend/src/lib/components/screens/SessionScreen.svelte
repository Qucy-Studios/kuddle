<script lang="ts">
    import type {AlertError, SessionSettings, UploadResult} from "$lib/types";
    import {Button} from "$lib/components/ui/button";
    import {CleanUp, Download, DownloadBase64, GetSession, RequestScreenshot, Upload} from "$lib/wailsjs/go/kuddle/App";
    import {currentSessionName} from "$lib/states";
    import ErrorAlert from "$lib/components/ErrorAlert.svelte";
    import {ClipboardSetText} from "$lib/wailsjs/runtime";
    import {convertToBase64} from "$lib/utils";
    import LZString from "lz-string"
    import {getRecentUploads, pushRecentUpload} from "$lib/recents";
    import SuccessfulUploadAlert from "$lib/components/alerts/SuccessfulUploadAlert.svelte";
    import {Icon} from "@steeze-ui/svelte-icon";
    import {Gear, RulerHorizontal} from "@steeze-ui/radix-icons";
    import NumericInputField from "$lib/components/NumericInputField.svelte";
    import InputField from "$lib/components/InputField.svelte";
    import {Folder} from "@steeze-ui/lucide-icons";
    import {extractImageSrc, isValidBase64Image} from "$lib/image";
    import EphemeralSettingsScreen from "$lib/components/screens/EphemeralSettingsScreen.svelte";
    import UploadScreen from "$lib/components/screens/UploadScreen.svelte";

    // page-scoped settings
    let ephemeralSettings: SessionSettings = {
        urlLength: 10,
        folder: "kd"
    }

    function saveSettings(evt: CustomEvent<SessionSettings>) {
        ephemeralSettings = evt.detail
    }

    function navigate(evt: CustomEvent<string>) {
        currentScreen = evt.detail
    }

    // upload-scoped
    let error: AlertError | null = null
    let uploadResult: UploadResult | null = null

    let imageSrc: string | null = null
    let imageKind: "file" | "link" | "none" = "none"

    let currentScreen: string = "upload"

    function discard() {
        imageSrc = null
        imageKind = "none"
    }

    async function onPaste(evt: ClipboardEvent & {currentTarget: (EventTarget & Window)}) {
        evt.preventDefault()

        const data = evt.clipboardData!
        if (data.files.length === 0 && data.items.length === 0 && data.types.length > 0) {
            try {
                const base64 = await RequestScreenshot()
                const err = isValidBase64Image(base64)
                if (err != null) {
                    error = err
                    return
                }
                imageSrc = base64
                imageKind = "file"
            } catch (e) {
                error = {
                    title: "Unable to capture screenshot.",
                    description: "We encountered an error while trying to capture screenshot data: " + e
                }
            }
            return
        }

        if (data.files.length > 0) {
            const file = data.files[0]
            if (!file.type.startsWith("image")) {
                console.info("unknown type", file.type)
                return
            }
            const fileData = (await convertToBase64(file)).result?.toString() ?? null
            const err = isValidBase64Image(fileData)
            if (err != null) {
                error = err
                return
            }
            imageSrc = fileData
            return
        }

        if (data.items.length > 0) {
            const item = data.items[0]
            if (item.kind === "string" && item.type == "text/html") {
                item.getAsString((val) => {
                    const matches = extractImageSrc(val)
                    if (matches.length < 1) {
                        error = {
                            title: "Invalid file format.",
                            description: "Kuddles only accept image files, if this is an image file, then please report this as a bug."
                        }
                        return;
                    }

                    imageKind = "link"
                    imageSrc = matches[0].replaceAll("amp;", "")
                })
            }
            return
        }
    }

    let isLoading = false
    async function upload() {
        if  (isLoading || imageSrc == null) {
            return
        }

        isLoading = true
        try {
            let contentType = "";
            let srcPath = imageSrc
            let extension = imageSrc.split(".").pop() ?? "";
            try {
                if (imageKind === "link") {
                    let downloadResult = await Download(imageSrc)

                    srcPath = downloadResult.ImagePath
                    extension = downloadResult.Extension.split(".").pop()
                    contentType = downloadResult.ContentType
                } else {
                    srcPath = await DownloadBase64(imageSrc)
                }
            } catch (e) {
                error = {
                    title: "Failed to temporarily download image.",
                    description: "It seems like we couldn't download the image because of the error: " + e
                }
                return
            }

            let session = await GetSession($currentSessionName!)

            if (extension == "") {
                extension = "png"
            }
            if (imageKind === "file") {
                contentType = imageSrc.split(';')[0].replace("data:", "")
                extension = contentType.split("/")[1]
            } else if (contentType == "") {
                contentType = "image/" + extension
            }

            let link = await Upload(session, srcPath, extension, contentType, ephemeralSettings)
            await pushRecentUpload(session, link)
            await ClipboardSetText(link)

            discard()
        } catch (e) {
            console.error(e)
            error = {
                title: "Failed to upload image.",
                description: "We encountered an issue while trying to upload image: " + e
            }
        } finally {
            isLoading = false
            CleanUp()
                .catch(reason => console.error("failed to clean temporary dir: ", reason))
        }
    }

    function onKeyPress(evt: KeyboardEvent) {
        if (evt.key === "Enter") {
            if (imageSrc != null && imageKind != "none") {
                upload()
            }
        }
    }
</script>
<svelte:window on:paste={onPaste} on:keypress={onKeyPress}/>
<div class="min-h-full w-full flex flex-col gap-2">
    <ErrorAlert bind:error/>
    {#if currentScreen === "upload"}
        <UploadScreen
            bind:imageSrc
            bind:isLoading
            bind:uploadResult
            on:navigate={navigate}
            on:upload={upload}
            on:discard={discard}
        />
    {:else if currentScreen === "settings"}
        <!-- Settings is integrated here because it is only applicable to the current page -->
        <EphemeralSettingsScreen
            ephemeralFolderPath={ephemeralSettings.folder}
            ephemeralNameLength={ephemeralSettings.urlLength}
            on:navigate={navigate}
            on:save={saveSettings}/>
    {/if}
</div>