<script lang="ts">
  import Icon from "$lib/assets/icons/Icon.svelte";
  import IconTrash from "$lib/assets/icons/IconTrash.svelte";
  import type {
    Ingredient,
    Instruction,
    SubRecipe,
    TrackableIngredient,
    TrackableInstruction,
  } from "$lib/types/Recipe";
  import { createEventDispatcher, onMount } from "svelte";
  import Button from "./Button.svelte";
  import IngredientInput from "./IngredientInput.svelte";
  import InstructionInput from "./InstructionInput.svelte";
  import { generateRandomTrackingId } from "$lib/util";

  export let trackingId: string;
  export let subRecipeCount: number;

  let ingredients: TrackableIngredient[] = [];
  let instructions: TrackableInstruction[] = [];

  let titleInput: HTMLInputElement;

  let subRecipe: SubRecipe = { ingredients: [], instructions: [] };
  const dispatch = createEventDispatcher();

  const onTitleChange = (
    e: Event & {
      currentTarget: EventTarget & HTMLInputElement;
    },
  ) => {
    subRecipe.title = e.currentTarget.value;
  };

  const onIngredientsChange = (e: CustomEvent) => {
    const ingredientCount = e.detail.ingredientCount as number;
    const ingredient = e.detail.ingredient as Ingredient;

    subRecipe.ingredients[ingredientCount] = ingredient;
  };

  const onInstructionsChange = (e: CustomEvent) => {
    const instructionCount = e.detail.instructionCount as number;
    const instruction = e.detail.instruction as Instruction;

    subRecipe.instructions[instructionCount] = instruction;
  };

  $: {
    dispatch("change", { subRecipe, subRecipeCount });
  }

  onMount(() => {
    titleInput.focus();
  });

  const onDelete = () => {
    dispatch("delete", { trackingId });
  };

  const onAddIngredient = () => {
    ingredients = [
      ...ingredients,
      { trackingId: generateRandomTrackingId(), name: "" },
    ];
  };

  const onDeleteIngredient = (e: CustomEvent) => {
    ingredients = ingredients.filter((ingredient) => {
      return ingredient.trackingId != e.detail.trackingId;
    });
  };

  const onAddInstruction = () => {
    instructions = [
      ...instructions,
      { text: "", trackingId: generateRandomTrackingId() },
    ];
  };

  const onDeleteInstruction = (e: CustomEvent) => {
    instructions = instructions.filter((instruction) => {
      return instruction.trackingId != e.detail.trackingId;
    });
  };
</script>

<div class="relative bg-gray-100 p-4 rounded-md border-[1px]">
  <input type="hidden" name={trackingId} />
  <input
    name={`sub-${subRecipeCount}-title`}
    type="text"
    placeholder="Recipe Title"
    class="w-full mb-4 font-semibold text-2xl bg-transparent border-b-[1px] focus:outline-none"
    bind:this={titleInput}
    on:change={onTitleChange}
  />
  <h2 class="font-semibold mb-2 text-gray-700">Ingredients.</h2>
  <ul class="list-disc list-inside flex flex-col gap-2">
    {#each ingredients as ingredient, i (ingredient.trackingId)}
      <IngredientInput
        on:change={onIngredientsChange}
        on:blur={() => {}}
        trackingId={ingredient.trackingId}
        ingredientCount={i}
        {subRecipeCount}
        on:delete={onDeleteIngredient}
      />
    {/each}
  </ul>
  <button on:click={onAddIngredient} class="text-gray-500 mb-4" type="button"
    >+ Add Ingredient</button
  >
  <h2 class="font-semibold mb-2 text-gray-700">Instructions.</h2>
  <ol class="list-decimal list-inside flex flex-col gap-2">
    {#each instructions as instruction, i (instruction.trackingId)}
      <InstructionInput
        on:delete={onDeleteInstruction}
        trackingId={instruction.trackingId}
        {subRecipeCount}
        instructionCount={i}
        on:change={onInstructionsChange}
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
