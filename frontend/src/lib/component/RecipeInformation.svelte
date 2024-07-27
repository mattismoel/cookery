<script lang="ts">
  import IconTimer from "$lib/assets/icons/IconTimer.svelte";
  import type { Recipe } from "$lib/types/Recipe";
  import FavoriteButton from "./FavoriteButton.svelte";

  export let recipe: Recipe;
  export let favourite: boolean;
  const { author, cookMinutes, totalMinutes } = recipe;
</script>

<div class="flex justify-between">
  <div>
    <span class="text-gray-400">{`By ${author}`}</span>
    <div class="flex gap-4 mt-2 text-gray-600 fill-gray-600">
      <div class="flex gap-2">
        <IconTimer />
        <span>{cookMinutes} min</span>
      </div>
      <span>|</span>
      <span>{totalMinutes} min</span>
    </div>
  </div>
  <FavoriteButton {favourite} />
</div>
