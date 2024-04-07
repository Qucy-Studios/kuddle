<script lang="ts">
    export let classNames: string = ""
</script>
<div class="min-h-full w-full flex flex-col {classNames}">
    <slot/>
</div>