<script lang="ts">
  import type { Ingredient } from "$lib/types/Recipe";

  export let ingredients: Ingredient[];
</script>

<ul class="list-disc list-inside text-gray-600">
  {#each ingredients as ingredient}
    <li>
      {`${ingredient.amount} ${ingredient.unit ? ingredient.unit : ""} ${ingredient.name}${ingredient.note ? `, ${ingredient.note}` : ""}`}
    </li>
  {/each}
</ul>
