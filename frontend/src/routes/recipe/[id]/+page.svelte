<script lang="ts">
  import IngredientList from "$lib/component/IngredientList.svelte";
  import Instructions from "$lib/component/Instructions.svelte";
  import RecipeInformation from "$lib/component/RecipeInformation.svelte";
  import type { PageServerData } from "./$types";

  export let data: PageServerData;
  const { recipe } = data;
</script>

<main class="min-h-[calc(100vh-var(--nav-height))]">
  <!-- Header -->
  <div class="flex flex-col h-[calc(100vh-var(--nav-height))]">
    <img
      src={recipe.bannerUrl}
      alt="Recipe."
      class="object-cover h-full w-full"
    />
    <div class="mx-responsive my-responsive">
      <h1 class="font-semibold uppercase text-4xl mb-2">{recipe.title}</h1>
      <RecipeInformation {recipe} favourite={false} />
    </div>
  </div>
  <article class="mx-responsive my-responsive">
    <h2 class="font-semibold text-xl mb-4">Ingredients.</h2>
    <div class="flex flex-col gap-8 mb-8">
      {#each recipe.subRecipes as subRecipe}
        <div>
          {#if subRecipe.title && recipe.subRecipes.length > 1}
            <h3 class="text-gray-800 mb-2">{subRecipe.title}</h3>
          {/if}
          <IngredientList ingredients={subRecipe.ingredients} />
        </div>
      {/each}
    </div>
    <h2 class="font-semibold text-xl mb-4">Instructions.</h2>
    <div class="flex flex-col gap-8 mb-8">
      {#each recipe.subRecipes as subRecipe}
        <div>
          {#if subRecipe.title && recipe.subRecipes.length > 1}
            <h3 class="text-gray-800 mb-2">{subRecipe.title}</h3>
          {/if}
          <Instructions instructions={subRecipe.instructions} />
        </div>
      {/each}
    </div>
  </article>
</main>
