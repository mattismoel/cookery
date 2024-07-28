import { error, redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";
import { SECRET_SERVER_URL } from "$env/static/private";
import type { APIError } from "$lib/types/APIError";
import type { User } from "$lib/types/User";

export const actions: Actions = {
  default: async ({ request }) => {
    const formData = await request.formData();

    const resp = await fetch(`${SECRET_SERVER_URL}/login`, { method: "POST", body: formData })
    if (!resp.ok) {
      const errorData: APIError = await resp.json();
      error(errorData.status, errorData.message);
    }

    const user: User = await resp.json()

    console.log(user)

    redirect(302, "/home")
  },
}
