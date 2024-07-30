<script lang="ts">
  import Icon from "$lib/assets/icons/Icon.svelte";
  import IconTrash from "$lib/assets/icons/IconTrash.svelte";
  import type { Ingredient } from "$lib/types/Recipe";
  import { createEventDispatcher, onMount } from "svelte";

  export let subRecipeCount: number;
  export let ingredientCount: number;
  export let trackingId: string;

  const dispatch = createEventDispatcher();

  let amount: number;
  let unit: string;
  let name: string;

  const onChange = (e: Event) => {
    const ingredient: Ingredient = { name, amount, unit };
    dispatch("change", {
      ingredient: ingredient,
      ingredientCount: ingredientCount,
    });
  };

  const deleteSelf = () => {
    dispatch("delete", { trackingId });
  };
</script>

<div class="flex w-full gap-2 items-center last:mb-4">
  <li class=""></li>
  <input
    on:change={onChange}
    name={`sub-${subRecipeCount}-ingredient-${ingredientCount}-amount`}
    class="ingredient-input w-8"
    placeholder="#"
    type="number"
    bind:value={amount}
  />
  <input
    name={`sub-${subRecipeCount}-ingredient-${ingredientCount}-unit`}
    class="ingredient-input w-16"
    placeholder="Unit"
    type="text"
    on:change={onChange}
    bind:value={unit}
  />
  <input
    name={`sub-${subRecipeCount}-ingredient-${ingredientCount}-name`}
    class="ingredient-input w-full"
    placeholder="Name"
    type="text"
    on:change={onChange}
    bind:value={name}
  />

  <button on:click={deleteSelf} type="button">
    <Icon size={14}>
      <IconTrash />
    </Icon>
  </button>
</div>
