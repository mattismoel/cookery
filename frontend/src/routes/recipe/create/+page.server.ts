import type { Recipe, SubRecipe } from "$lib/types/Recipe";
import type { Actions } from "./$types";

export const actions: Actions = {
  default: async ({ request }) => {
    await request.formData().then(formData => {
      let subRecipes: SubRecipe[] = []

      const subCountStr = formData.get("subCount")?.toString()
      let subCount: number;

      if (subCountStr)
        subCount = parseInt(subCountStr.toString())

      for (const pair of formData.entries()) {
        const key = pair[0]
        const val = pair[1]

        console.log(key)

        if (!key.startsWith("sub-")) {
          continue;
        }

        const subCount = parseInt(key[4])
        // subRecipes[subCount].title = "test"
      }

      console.log(subRecipes)
    });
  }
}
