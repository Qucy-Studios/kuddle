<script>
import {Symbol} from "@steeze-ui/radix-icons";
import {Icon} from "@steeze-ui/svelte-icon";
</script>

<div class="h-screen w-full flex flex-col justify-center align-middle items-center">
    <div class="mb-2 flex flex-col items-center animate-pulse">
        <Icon src={Symbol} size="18" class="animate-spin mb-2"></Icon>
        <p class="text-xs text-zinc-400 sofia">Loading</p>
    </div>
</div>