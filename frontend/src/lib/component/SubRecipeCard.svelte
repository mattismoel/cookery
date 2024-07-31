<script lang="ts">
  import Icon from "$lib/assets/icons/Icon.svelte";
  import IconTrash from "$lib/assets/icons/IconTrash.svelte";
  import type { Ingredient, Instruction, SubRecipe } from "$lib/types/Recipe";
  import { createEventDispatcher, onMount } from "svelte";
  import Button from "./Button.svelte";
  import IngredientInput from "./IngredientInput.svelte";
  import InstructionInput from "./InstructionInput.svelte";

  export let subIdx: number;

  let titleInput: HTMLInputElement;
  export let subRecipe: SubRecipe = { ingredients: [], instructions: [] };

  const dispatch = createEventDispatcher();

  onMount(() => {
    titleInput.focus();
  });

  $: {
    dispatch("change", { subRecipe, subIdx });
  }

  const onDelete = () => {
    dispatch("delete", { subIdx });
  };

  const onAddIngredient = () => {
    subRecipe.ingredients = [...subRecipe.ingredients, { name: "" }];
  };

  const onDeleteIngredient = (e: CustomEvent) => {
    const ingredientIdx = e.detail.ingredientIdx as number;
    subRecipe.ingredients = subRecipe.ingredients.filter(
      (_ingredient, idx) => idx !== ingredientIdx,
    );
  };

  const onIngredientChange = (e: CustomEvent) => {
    const ingredientIdx = e.detail.ingredientIdx as number;
    const ingredient = e.detail.ingredient as Ingredient;
    subRecipe.ingredients[ingredientIdx] = ingredient;
  };

  const onAddInstruction = () => {
    subRecipe.instructions = [...subRecipe.instructions, { text: "" }];
  };

  const onDeleteInstruction = (e: CustomEvent) => {
    const instructionIdx = e.detail.instructionIdx as number;
    subRecipe.instructions = subRecipe.instructions.filter(
      (_instruction, idx) => {
        idx !== instructionIdx;
      },
    );
  };

  const onInstructionChange = (e: CustomEvent) => {
    const instructionIdx = e.detail.instructionIdx as number;
    const instruction = e.detail.instruction as Instruction;
    subRecipe.instructions[instructionIdx] = instruction;
  };
</script>

<div class="relative bg-gray-100 p-4 rounded-md border-[1px]">
  <input type="hidden" />
  <input
    name={`sub-${subIdx}-title`}
    type="text"
    placeholder="Recipe Title"
    class="w-full mb-4 font-semibold text-2xl bg-transparent border-b-[1px] focus:outline-none"
    bind:this={titleInput}
    on:change={(e) => (subRecipe.title = e.currentTarget.value)}
  />
  <h2 class="font-semibold mb-2 text-gray-700">Ingredients.</h2>
  <ul class="list-disc list-inside flex flex-col gap-2">
    {#each subRecipe.ingredients as ingredient, ingredientIdx}
      <IngredientInput
        on:change={onIngredientChange}
        on:delete={onDeleteIngredient}
        {subIdx}
        {ingredientIdx}
      />
    {/each}
  </ul>
  <button on:click={onAddIngredient} class="text-gray-500 mb-4" type="button"
    >+ Add Ingredient</button
  >
  <h2 class="font-semibold mb-2 text-gray-700">Instructions.</h2>
  <ol class="list-decimal list-inside flex flex-col gap-2">
    {#each subRecipe.instructions as instruction, instructionIdx}
      <InstructionInput
        on:delete={onDeleteInstruction}
        on:change={onInstructionChange}
        {subIdx}
        {instructionIdx}
      />
    {/each}
  </ol>
  <button on:click={onAddInstruction} class="text-gray-500 mb-4" type="button"
    >+ Add Instruction</button
  >
  <Button
    on:click={onDelete}
    fitX="expand"
    type="danger"
    props={{ type: "button" }}
  >
    <Icon size={18}>
      <IconTrash />
    </Icon>
  </Button>
</div>
