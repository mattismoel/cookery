<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import type { HTMLButtonAttributes } from "svelte/elements";

  type Fit = "shrink" | "expand";

  export let fitX: Fit = "shrink";
  export let props: HTMLButtonAttributes;

  const dispatch = createEventDispatcher();

  const onClick = () => {
    dispatch("click");
  };

  const fitMap = new Map<Fit, string>([
    ["shrink", ""],
    ["expand", "w-full"],
  ]);
</script>

<button
  class={`bg-green-400 hover:bg-green-500 border-[1px] border-green-500 hover:border-green-600 font-semibold text-white py-2 px-4 rounded-sm ${fitMap.get(fitX)}`}
  on:click={onClick}
  {...props}
>
  <slot />
</button>
