import { SECRET_SERVER_URL } from "$env/static/private";
import type { Recipe } from "$lib/types/Recipe";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params }) => {
  const url = `${SECRET_SERVER_URL}/recipe/${params.id}`
  const recipe: Recipe = await fetch(url).then(res => {
    return res.json()
  }).catch(err => {
    console.log(err)
  })

  console.log(recipe)


  return {
    recipe,
  }
}
