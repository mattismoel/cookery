<script lang="ts">
  export let size: number = 24;
  export let extraClasses: string = "";

  $: sizePx = `${size}px`;
</script>

<svg
  class={extraClasses}
  style:width={sizePx}
  style:height={sizePx}
  viewBox={`0 0 24 24`}
>
  <slot />
</svg>
