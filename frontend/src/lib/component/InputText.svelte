<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";

  export let props: HTMLInputAttributes;
  export let align: "left" | "center" | "right" = "left";

  $: alignMap = new Map<string, string>([
    ["left", "text-left"],
    ["center", "text-center"],
    ["right", "text-right"],
  ]);
</script>

<input
  class={`w-full bg-gray-50 border-[1px] px-4 py-2 rounded-sm ${alignMap.get(align)}`}
  {...props}
/>
