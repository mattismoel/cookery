import { redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions: Actions = {
  default: async ({ request }) => {
    const formData = await request.formData();

    redirect(302, "/home")

    const email = formData.get("email")
    const password = formData.get("password")

    console.log(email, password)
  },
}
