<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import {Key} from "@steeze-ui/lucide-icons";
    import {EyeClosed, EyeOpen, type IconSource} from "@steeze-ui/radix-icons";
    import {Icon} from "@steeze-ui/svelte-icon";

    export let value: string
    export let placeholder: string
    export let icon: IconSource

    export let showSecrets: boolean
    export let disabled: boolean = false
    const dispatcher = createEventDispatcher()

    function toggleShowSecrets() {
        dispatcher('toggle')
    }

    let borderClasses = disabled ? "hover:border-red-500" : "hover:border-black";
    let iconClasses = disabled ? "group-hover:text-red-500" : "hover:text-black"
</script>

<div class="group rounded border-zinc-300 border {borderClasses} bg-white transition duration-500 py-3 px-5 flex flex-row gap-4 items-center">
    <Icon src={icon} size="18" class="text-zinc-400 {iconClasses} transition duration-500"/>
    {#if showSecrets}
        <input bind:value
               class="w-full bg-transparent outline-none peer border-none placeholder:roboto"
               placeholder={placeholder}
               disabled={disabled}
               type="text"/>
    {:else}
        <input bind:value
               class="w-full bg-transparent outline-none peer border-none placeholder:roboto"
               placeholder={placeholder}
               disabled={disabled}
               type="password"/>
    {/if}
    <button on:mousedown={toggleShowSecrets} on:mouseup={toggleShowSecrets}>
        <Icon src={showSecrets ? EyeClosed : EyeOpen} size="18" class="text-zinc-400 {iconClasses} transition duration-500"/>
    </button>
</div>