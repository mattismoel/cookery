<script lang="ts">
  import type { Instruction } from "$lib/types/Recipe";

  export let instructions: Instruction[];
</script>

<ol class="list-decimal list-inside text-gray-600">
  {#each instructions as instruction}
    <li class="py-1">
      {`${instruction.text}`}
    </li>
  {/each}
</ol>
