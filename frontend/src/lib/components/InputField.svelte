<script lang="ts">
    import {type IconSource} from "@steeze-ui/radix-icons";
    import {Icon} from "@steeze-ui/svelte-icon";
    import * as Tooltip from "$lib/components/ui/tooltip";

    export let icon: IconSource
    export let value: string
    export let placeholder: string
    export let tooltip: string | undefined = undefined
</script>

<div class="group rounded border-zinc-300 border hover:border-black bg-white transition duration-500 py-3 px-5 flex flex-row gap-4 items-center">
    <Tooltip.Root>
        <Tooltip.Trigger>
            <Icon src={icon} size="18" class="text-zinc-400 group-hover:text-black transition duration-500"/>
        </Tooltip.Trigger>
        <Tooltip.Content>
            <p>{tooltip ?? placeholder}</p>
        </Tooltip.Content>
    </Tooltip.Root>
    <input bind:value
           class="w-full bg-transparent outline-none peer border-none placeholder:roboto"
           placeholder={placeholder}
           type="text"/>
</div>