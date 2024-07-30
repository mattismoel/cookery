<script lang="ts">
  import Icon from "$lib/assets/icons/Icon.svelte";
  import IconTrash from "$lib/assets/icons/IconTrash.svelte";
  import type { Instruction } from "$lib/types/Recipe";
  import { createEventDispatcher, onMount } from "svelte";

  export let subRecipeCount: number;
  export let instructionCount: number;
  export let trackingId: string;

  let textArea: HTMLTextAreaElement;

  const dispatch = createEventDispatcher();

  onMount(() => textArea.focus());

  const resizeTextArea = () => {
    const { style } = textArea;
    style.height = style.minHeight = "auto";
    style.minHeight = `${Math.min(textArea.scrollHeight + 4, parseInt(style.maxHeight))}px`;
    style.height = `${textArea.scrollHeight}px`;
  };

  const deleteSelf = () => {
    dispatch("delete", { trackingId });
  };

  const onChange = (e: Event) => {
    const instruction: Instruction = {
      text: textArea.value,
      imageUrl: "no implementation",
    };

    dispatch("change", {
      instructionCount: instructionCount,
      instruction,
    });
  };
</script>

<div class="flex gap-2 last:mb-4">
  <li class="w-4" />
  <textarea
    name={`sub-${subRecipeCount}-instruction-${instructionCount}`}
    bind:this={textArea}
    on:input={resizeTextArea}
    on:change={onChange}
    rows="1"
    class="resize-none overflow-y-hidden w-full bg-transparent border-b-[1px] focus:outline-none rounded-sm"
    value=""
  />
  <button type="button" on:click={deleteSelf}>
    <Icon size={14}>
      <IconTrash />
    </Icon>
  </button>
</div>
