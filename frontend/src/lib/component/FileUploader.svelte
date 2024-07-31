<script lang="ts">
  import { createEventDispatcher } from "svelte";

  export let accept: string = "";
  export let name: string;
  export let multiple: boolean = false;
  export let files: FileList | undefined = undefined;

  const dispatch = createEventDispatcher();

  $: {
    dispatch("change", { files });
  }
</script>

<input
  bind:files
  type="file"
  {multiple}
  {name}
  accept={accept !== "" ? accept : undefined}
/>
