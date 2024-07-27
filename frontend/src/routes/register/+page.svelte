<script>
  import Button from "$lib/component/Button.svelte";
  import InputText from "$lib/component/InputText.svelte";
</script>

<main class="mx-responsive flex items-center h-[calc(100vh-var(--nav-height))]">
  <form class="w-full flex flex-col gap-2" method="post">
    <h1 class="text-2xl font-semibold mb-4">Join the cookery.</h1>
    <InputText
      props={{ type: "email", placeholder: "Email", name: "email" }}
      align="center"
    />
    <InputText
      props={{ type: "text", placeholder: "Firstname", name: "firstname" }}
      align="center"
    />
    <InputText
      props={{ type: "text", placeholder: "Lastname", name: "lastname" }}
      align="center"
    />
    <InputText
      props={{ type: "password", placeholder: "Password", name: "password" }}
      align="center"
    />
    <InputText
      props={{
        type: "password",
        placeholder: "Repeat Password",
        name: "passwordRepeat",
      }}
      align="center"
    />
    <Button props={{ type: "submit" }} fitX="expand">Register</Button>
    <div class="text-gray-500">
      Already have an account? <a class="font-semibold" href="/login"
        >Login here.</a
      >
    </div>
  </form>
</main>
