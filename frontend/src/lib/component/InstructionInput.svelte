<script lang="ts">
  import Icon from "$lib/assets/icons/Icon.svelte";
  import IconTrash from "$lib/assets/icons/IconTrash.svelte";
  import type { Instruction } from "$lib/types/Recipe";
  import { createEventDispatcher, onMount } from "svelte";

  export let subIdx: number;
  export let instructionIdx: number;

  let textArea: HTMLTextAreaElement;
  let instruction: Instruction = {} as Instruction;

  const dispatch = createEventDispatcher();

  onMount(() => textArea.focus());

  $: {
    dispatch("change", { instructionIdx, instruction });
  }

  const resizeTextArea = () => {
    const { style } = textArea;
    style.height = style.minHeight = "auto";
    style.minHeight = `${Math.min(textArea.scrollHeight + 4, parseInt(style.maxHeight))}px`;
    style.height = `${textArea.scrollHeight}px`;
  };

  const deleteSelf = () => {
    dispatch("delete", { instructionIdx });
  };
</script>

<div class="flex gap-2 last:mb-4">
  <li class="w-4" />
  <textarea
    name={`sub-${subIdx}-instruction-${instructionIdx}-text`}
    rows="1"
    on:change={(e) => {
      instruction.text = e.currentTarget.value;
    }}
    bind:this={textArea}
    on:input={resizeTextArea}
    class="resize-none overflow-y-hidden w-full bg-transparent border-b-[1px] focus:outline-none rounded-sm"
    value=""
  />
  <button type="button" on:click={deleteSelf}>
    <Icon size={14}>
      <IconTrash />
    </Icon>
  </button>
</div>
