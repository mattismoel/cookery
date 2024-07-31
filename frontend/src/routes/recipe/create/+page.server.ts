import { SECRET_SERVER_URL } from "$env/static/private";
import type { APIError } from "$lib/types/APIError";
import type { Ingredient, Instruction, Recipe, SubRecipe } from "$lib/types/Recipe";
import { error } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions: Actions = {
  default: async ({ request }) => {
    const formData = await request.formData()

    await fetch(`${SECRET_SERVER_URL}/recipe`, {
      method: "PUT",
      body: formData,
    })
      .then(async res => {
        if (!res.ok) {
          const err = await res.json() as APIError;
          return error(err.status, err.message)
        }
      })
      .catch(err => console.log(err))

  }
}
