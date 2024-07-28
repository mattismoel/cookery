<script>
  import Button from "$lib/component/Button.svelte";
  import InputText from "$lib/component/InputText.svelte";
</script>

<main class="mx-responsive flex items-center h-[calc(100vh-var(--nav-height))]">
  <form method="post">
    <h1 class="font-semibold text-2xl mb-4">Welcome back.</h1>
    <div class="space-y-2">
      <InputText
        align="center"
        props={{ type: "email", placeholder: "Email", name: "email" }}
      />
      <InputText
        align="center"
        props={{ type: "password", placeholder: "Password", name: "password" }}
      />
      <Button props={{ type: "submit" }} fitX="expand">Login</Button>
      <div class="text-base text-gray-500">
        <span>Don't have an account?</span>
        <a class="font-semibold" href="/register">Create one here.</a>
      </div>
    </div>
  </form>
</main>
