<script lang="ts">
  import IconStar from "$lib/assets/icons/IconStar.svelte";
  import IconStarFilled from "$lib/assets/icons/IconStarFilled.svelte";
  import { createEventDispatcher } from "svelte";

  export let favourite: boolean;

  const dispatch = createEventDispatcher();

  const onClick = () => {
    favourite = !favourite;
    dispatch("toggle", { favourite });
  };
</script>

<button on:click={onClick}>
  {#if favourite}
    <IconStarFilled />
  {:else}
    <IconStar />
  {/if}
</button>
