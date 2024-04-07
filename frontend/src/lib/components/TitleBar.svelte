<script lang="ts">
    import {Quit, WindowMinimise} from "$lib/wailsjs/runtime/runtime.js";
    import {Cross2, Dash, Home, LockClosed, Person, Plus, Trash} from "@steeze-ui/radix-icons"
    import {currentSessionName, encryptionKey, sessionNames} from "$lib/states";
    import {Icon} from "@steeze-ui/svelte-icon";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
    import {Separator} from "$lib/components/ui/separator";
    import {onMount} from "svelte";
    import {Edit, History} from "@steeze-ui/lucide-icons";
    import ProgrammaticableRedirect from "$lib/components/ProgrammaticableRedirect.svelte";
    import {goto, pushState, replaceState} from "$app/navigation";

    function lockSession() {
        $encryptionKey = null
        $sessionNames = null
    }

    function changeSession(name: string) {
        $currentSessionName = name
    }

    function redirect(path: "_home" | "_newsession" | "_recents" | "_edit" | "_delete") {
        document.getElementById(path)?.click()
    }
</script>
<div class="pb-4" style="--wails-draggable:drag">
    <ProgrammaticableRedirect id="_newsession" href="/sessions/new"/>
    <ProgrammaticableRedirect id="_recents" href="/recents"/>
    <ProgrammaticableRedirect id="_home" href="/"/>
    <ProgrammaticableRedirect id="_edit" href={"/sessions/edit"}/>
    <ProgrammaticableRedirect id="_delete" href={"/sessions/delete"}/>
    <div class="w-full flex flex-row items-center py-8 pb-4 hover:cursor-move" style="--wails-draggable:drag">
        {#if $sessionNames != null}
            <DropdownMenu.Root>
                <DropdownMenu.Trigger>
                    <div class="flex  flex-row items-center gap-4">
                        <div class="rounded-full p-2 bg-zinc-300">
                            <Icon src={Person} size="16"/>
                        </div>
                        <h2 class="font-bold text-xs">{$currentSessionName}</h2>
                    </div>
                </DropdownMenu.Trigger>
                <DropdownMenu.Content class="w-56">
                    <DropdownMenu.Group>
                        <DropdownMenu.Label>{$currentSessionName ?? "Unknown"}</DropdownMenu.Label>
                        <DropdownMenu.Separator/>
                        <DropdownMenu.Item class="flex flex-row items-center w-full" on:click={() => redirect('_home')}>
                            <Icon src={Home} size="16" class="mr-2"/>
                            <span>Home</span>
                        </DropdownMenu.Item>
                        <DropdownMenu.Item  class="flex flex-row items-center w-full" on:click={() => redirect('_newsession')}>
                            <Icon src={Plus} size="16" class="mr-2"/>
                            <span>New Session</span>
                        </DropdownMenu.Item>
                        <DropdownMenu.Item  class="flex flex-row items-center w-full" on:click={() => redirect('_recents')}>
                            <Icon src={History} size="16" class="mr-2"/>
                            <span>Recent Uploads</span>
                        </DropdownMenu.Item>
                        {#if ($sessionNames ?? []).length > 1}
                            <DropdownMenu.Sub>
                                <DropdownMenu.SubTrigger>
                                    <Icon src={Person} size="16" class="mr-2"/>
                                    <span>Other sessions</span>
                                </DropdownMenu.SubTrigger>
                                <DropdownMenu.SubContent>
                                    {#each $sessionNames ?? [] as session}
                                        {#if session !== $currentSessionName}
                                            <DropdownMenu.Item on:click={() => changeSession(session)}>
                                                <Icon src={Person} size="16" class="mr-2"/>
                                                <span>{session}</span>
                                            </DropdownMenu.Item>
                                        {/if}
                                    {/each}
                                </DropdownMenu.SubContent>
                            </DropdownMenu.Sub>
                        {/if}
                        <DropdownMenu.Item on:click={() => redirect("_edit")}>
                            <Icon src={Edit} size="16" class="mr-2"/>
                            <span>Edit current session</span>
                        </DropdownMenu.Item>
                        <DropdownMenu.Item on:click={lockSession}>
                            <Icon src={LockClosed} size="16" class="mr-2"/>
                            <span>Lock Session</span>
                        </DropdownMenu.Item>
                        {#if ($sessionNames ?? []).length > 1}
                            <DropdownMenu.Item class="text-red-500" on:click={() => redirect("_delete")}>
                                <Icon src={Trash} size="16" class="mr-2"/>
                                <span>Delete current session</span>
                            </DropdownMenu.Item>
                        {/if}
                        </DropdownMenu.Group>
                </DropdownMenu.Content>
            </DropdownMenu.Root>
        {:else}
            <a href="/" class="rounded-full p-2 bg-zinc-300" style="--wails-draggable:drag">
                <Icon src={LockClosed} size="16"/>
            </a>
        {/if}
        <div class="absolute right-0 mx-4 mr-8" style="--wails-draggable:drag">
            <div class="gap-2 flex flex-row items-center">
                <button on:click={WindowMinimise}
                        class="font-bold uppercase w-fit p-1 px-[0.84rem] hover:bg-zinc-200 transition hover:opacity-80 duration-300 ease-in-out text-xs">
                    <Icon src={Dash} size="16"/>
                </button>
                <button on:click={Quit}
                        class="font-bold uppercase text-red-500 w-fit p-1 px-[0.84rem] hover:bg-zinc-200 transition hover:opacity-80 duration-300 ease-in-out text-xs">
                    <Icon src={Cross2} size="16"/>
                </button>
            </div>
        </div>
    </div>
    <Separator/>
</div>
