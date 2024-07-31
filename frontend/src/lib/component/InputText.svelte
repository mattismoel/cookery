<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import type { HTMLInputAttributes } from "svelte/elements";

  type Fit = "expand" | "shrink";

  export let props: HTMLInputAttributes;
  export let align: "left" | "center" | "right" = "left";
  export let fitX: Fit = "expand";

  export let value: string = "";

  const dispatch = createEventDispatcher();

  const onChange = (e: Event) => {
    dispatch("change", { value });
  };

  $: fitMap = new Map<Fit, string>([
    ["expand", "w-full"],
    ["shrink", "w-min"],
  ]);

  $: alignMap = new Map<string, string>([
    ["left", "text-left"],
    ["center", "text-center"],
    ["right", "text-right"],
  ]);
</script>

<input
  class={`bg-gray-50 border-[1px] px-4 py-2 rounded-sm ${alignMap.get(align)} ${fitMap.get(fitX)}`}
  bind:value
  on:change={onChange}
  {...props}
/>
