<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import type { HTMLButtonAttributes } from "svelte/elements";

  type Fit = "shrink" | "expand";
  type Type = "primary" | "secondary" | "danger";

  export let fitX: Fit = "shrink";
  export let type: Type = "primary";
  export let props: HTMLButtonAttributes;

  const dispatch = createEventDispatcher();

  const onClick = () => {
    dispatch("click");
  };

  const fitMap = new Map<Fit, string>([
    ["shrink", ""],
    ["expand", "w-full"],
  ]);

  const typeMap = new Map<Type, string>([
    [
      "primary",
      "bg-green-400 hover:bg-green-500 border-green-500 hover:border-green-600 text-white fill-white",
    ],
    [
      "danger",
      "bg-red-400 hover:bg-red-500 border-red-500 hover:border-red-600 text-red-50 fill-red-50",
    ],
  ]);
</script>

<button
  class={`border-[1px] font-semibold py-2 px-4 rounded-sm flex items-center justify-center gap-2 ${fitMap.get(fitX)} ${typeMap.get(type)}`}
  on:click={onClick}
  {...props}
>
  <slot />
</button>
