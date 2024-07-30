import { error, json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";
import { SECRET_SERVER_URL } from "$env/static/private";
import type { Recipe } from "$lib/types/Recipe";
import type { APIError } from "$lib/types/APIError";

export const POST: RequestHandler = async ({ request }) => {
  const recipe = await request.json() as Recipe;
  console.log(recipe)

  await fetch(`${SECRET_SERVER_URL}/recipe`, {
    method: "PUT",
    body: JSON.stringify(recipe),
    headers: { "Content-Type": "application/json" },
  })
    .then(async res => {
      if (!res.ok) {
        const err = await res.json() as APIError
        return error(err.status, err.message)
      }
    })
    .catch(err => {
      return error(500, err);
    })

  return json({}, { status: 201 })
}
