<script lang="ts">
  import Button from "$lib/component/Button.svelte";
  import InputText from "$lib/component/InputText.svelte";
  import SubRecipeCard from "$lib/component/SubRecipeCard.svelte";
  import type {
    Recipe,
    SubRecipe,
    TrackableSubRecipe,
  } from "$lib/types/Recipe";
  import { generateRandomTrackingId } from "$lib/util";
  import { onMount } from "svelte";

  let subRecipes: TrackableSubRecipe[] = [];

  let recipe: Recipe = {} as Recipe;

  onMount(() => {
    recipe.subRecipes = [];
  });

  const onSubmit = async (e: CustomEvent) => {
    await fetch("/recipe/create", {
      method: "POST",
      body: JSON.stringify(recipe),
      headers: {
        "Content-Type": "application/json",
      },
    });
  };

  const addPart = () => {
    subRecipes = [
      ...subRecipes,
      {
        trackingId: generateRandomTrackingId(),
        title: "",
        ingredients: [],
        instructions: [],
      },
    ];
  };

  const deleteSubRecipe = (e: CustomEvent) => {
    const trackingId = e.detail.trackingId;
    subRecipes = subRecipes.filter(
      (subRecipe) => subRecipe.trackingId !== trackingId,
    );
  };

  const onSubRecipeChange = (e: CustomEvent) => {
    if (!e.detail.subRecipe || e.detail.subRecipeCount == null) {
      return;
    }

    const subRecipe = e.detail.subRecipe as SubRecipe;
    const subRecipeCount = e.detail.subRecipeCount as number;
    recipe.subRecipes[subRecipeCount] = subRecipe;
  };

  const onTitleChange = (e: CustomEvent) => {
    recipe.title = e.detail.value;
  };

  const onCookMinutesChange = (e: CustomEvent) => {
    recipe.cookMinutes = parseInt(e.detail.value);
  };

  const onTotalMinutesChange = (e: CustomEvent) => {
    recipe.totalMinutes = parseInt(e.detail.value);
  };
</script>

<main class="h-[calc(100vh-var(--nav-height))] mx-responsive my-responsive">
  <form method="POST">
    <input type="hidden" name="subCount" value={subRecipes.length} />
    <h1 class="font-semibold text-xl mb-4">Create Recipe.</h1>
    <div class="space-y-2">
      <InputText
        on:change={onTitleChange}
        align="center"
        props={{ type: "text", name: "title", placeholder: "Title" }}
      />
      <div class="w-full grid grid-cols-2 gap-2">
        <InputText
          fitX="expand"
          props={{
            type: "number",
            name: "cookMinutes",
            placeholder: "Cook Minutes",
          }}
          on:change={onCookMinutesChange}
        />
        <InputText
          fitX="expand"
          props={{
            type: "number",
            name: "cookMinutes",
            placeholder: "Total Minutes",
          }}
          on:change={onTotalMinutesChange}
        />
      </div>
      <h2 class="font-semibold">Parts.</h2>
      <div class="flex flex-col space-y-2">
        {#each subRecipes as subRecipe, i (subRecipe.trackingId)}
          <SubRecipeCard
            on:change={onSubRecipeChange}
            on:delete={deleteSubRecipe}
            subRecipeCount={i}
            trackingId={subRecipe.trackingId}
          />
        {/each}
      </div>
      <button
        on:click={addPart}
        class="w-full bg-gray-100 border-[1px] rounded-sm py-2"
        type="button">+</button
      >
      {#if subRecipes.length > 0}
        <Button on:click={onSubmit} fitX="expand" props={{ type: "button" }}
          >Upload Recipe</Button
        >
      {/if}
    </div>
  </form>
</main>
