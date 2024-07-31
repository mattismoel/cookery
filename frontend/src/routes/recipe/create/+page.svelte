<script lang="ts">
  import Button from "$lib/component/Button.svelte";
  import FileUploader from "$lib/component/FileUploader.svelte";
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
  let bannerFile: File;

  onMount(() => {
    recipe.subRecipes = [];
  });

  const onSubmit = async (e: CustomEvent) => {
    const formData = new FormData();
    formData.append("recipe", JSON.stringify(recipe));
    formData.append("bannerFile", bannerFile);

    await fetch("/recipe/create", {
      method: "POST",
      body: formData,
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

  const onDeleteSubRecipe = (e: CustomEvent) => {
    const subIdx = e.detail.subIdx as number;
    subRecipes = subRecipes.filter((_subRecipe, i) => i !== subIdx);
  };

  const onBannerChange = (e: CustomEvent) => {
    if (!e.detail.files) {
      return;
    }

    const files = e.detail.files as FileList;
  };

  const onSubRecipeChange = (e: CustomEvent) => {
    if (!e.detail.subRecipe || e.detail.subIdx == null) {
      return;
    }

    const subRecipe = e.detail.subRecipe as SubRecipe;
    const subIdx = e.detail.subIdx as number;
    recipe.subRecipes[subIdx] = subRecipe;
  };
</script>

<main class="h-[calc(100vh-var(--nav-height))] mx-responsive my-responsive">
  <form method="POST" enctype="multipart/form-data">
    <input type="hidden" name="subCount" value={subRecipes.length} />
    <h1 class="font-semibold text-xl mb-4">Create Recipe.</h1>
    <div class="space-y-2">
      <InputText
        on:change={(e) => (recipe.title = e.detail.value.toString())}
        align="center"
        props={{ type: "text", name: "title", placeholder: "Title" }}
      />
      <div class="w-full grid grid-cols-2 gap-2">
        <InputText
          on:change={(e) => (recipe.cookMinutes = parseFloat(e.detail.value))}
          fitX="expand"
          props={{
            type: "number",
            name: "cookMinutes",
            placeholder: "Cook Minutes",
          }}
        />
        <InputText
          on:change={(e) => (recipe.totalMinutes = parseFloat(e.detail.value))}
          fitX="expand"
          props={{
            type: "number",
            name: "totalMinutes",
            placeholder: "Total Minutes",
          }}
        />
      </div>
      <FileUploader
        on:change={(e) => {
          e.detail.files && e.detail.files.length > 0
            ? (bannerFile = e.detail.files[0])
            : undefined;
        }}
        name="bannerUrl"
        accept="image/png, image/webp, image/jpeg"
      />
      <h2 class="font-semibold">Parts.</h2>
      <div class="flex flex-col space-y-2">
        {#each subRecipes as subRecipe, subIdx (subRecipe.trackingId)}
          <SubRecipeCard
            on:change={onSubRecipeChange}
            on:delete={onDeleteSubRecipe}
            {subIdx}
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
