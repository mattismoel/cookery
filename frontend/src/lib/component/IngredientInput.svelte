<script lang="ts">
  import Icon from "$lib/assets/icons/Icon.svelte";
  import IconTrash from "$lib/assets/icons/IconTrash.svelte";
  import type { Ingredient } from "$lib/types/Recipe";
  import { createEventDispatcher, onMount } from "svelte";

  export let subIdx: number;
  export let ingredientIdx: number;

  let amountInput: HTMLInputElement;

  const dispatch = createEventDispatcher();

  let ingredient: Ingredient = {} as Ingredient;

  onMount(() => {
    amountInput.focus();
  });

  $: {
    dispatch("change", { ingredientIdx, ingredient });
  }

  const deleteSelf = () => {
    dispatch("delete", { ingredientIdx });
  };
</script>

<div class="flex w-full gap-2 items-center last:mb-4">
  <li class=""></li>
  <input
    name={`sub-${subIdx}-ingredient-${ingredientIdx}-amount`}
    class="ingredient-input w-8"
    placeholder="#"
    type="number"
    on:change={(e) => (ingredient.amount = e.currentTarget.valueAsNumber)}
    bind:this={amountInput}
  />
  <input
    name={`sub-${subIdx}-ingredient-${ingredientIdx}-unit`}
    class="ingredient-input w-16"
    placeholder="Unit"
    type="text"
    on:change={(e) => (ingredient.unit = e.currentTarget.value)}
  />
  <input
    name={`sub-${subIdx}-ingredient-${ingredientIdx}-name`}
    class="ingredient-input w-full"
    placeholder="Name"
    type="text"
    on:change={(e) => (ingredient.name = e.currentTarget.value)}
  />

  <button on:click={deleteSelf} type="button">
    <Icon size={14}>
      <IconTrash />
    </Icon>
  </button>
</div>
