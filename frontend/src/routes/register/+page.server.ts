import { SECRET_SERVER_URL } from "$env/static/private";
import type { Actions } from "./$types";

export const actions: Actions = {
  default: async ({ request }) => {
    await fetch(`${SECRET_SERVER_URL}/register`, { method: "PUT", body: await request.formData() })
      .catch(errResp => { console.log(errResp) })
  }
}
