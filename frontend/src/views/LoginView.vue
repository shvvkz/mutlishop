<template>
  <section class="row justify-content-center">
    <ToastMessage
      v-model:show="toastVisible"
      :message="toastMessage"
      :type="toastType"
      :duration="3200"
    />

    <div class="col-md-7 col-lg-5">
      <div class="card border-0 shadow-sm">
        <div class="card-body p-4">
          <h2 class="h4 mb-1">Connexion admin</h2>
          <p class="text-muted mb-4">Endpoint utilise: <code>POST /api/login</code></p>

          <form class="d-grid gap-3" @submit.prevent="submit">
            <div>
              <label class="form-label">Email</label>
              <input
                v-model.trim="email"
                class="form-control"
                type="email"
                autocomplete="username"
                required
              />
            </div>

            <div>
              <label class="form-label">Mot de passe</label>
              <input
                v-model="password"
                class="form-control"
                type="password"
                autocomplete="current-password"
                required
              />
            </div>

            <button :disabled="loading" class="btn btn-primary" type="submit">
              {{ loading ? "Connexion..." : "Se connecter" }}
            </button>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import ToastMessage from "../components/ToastMessage.vue";
import { login, setToken } from "../services/api";

const router = useRouter();
const route = useRoute();

const email = ref("super1@admin.com");
const password = ref("superadmin");
const loading = ref(false);
const toastVisible = ref(false);
const toastMessage = ref("");
const toastType = ref("info");

function showToast(message, type = "info") {
  toastMessage.value = message;
  toastType.value = type;
  toastVisible.value = false;
  window.setTimeout(() => {
    toastVisible.value = true;
  }, 0);
}

async function submit() {
  loading.value = true;

  try {
    const payload = await login(email.value, password.value);
    const token = payload?.data?.token || payload?.token;

    if (!token) {
      throw new Error("Token JWT absent de la reponse");
    }

    setToken(token);
    showToast("Connexion reussie", "success");

    const redirect = typeof route.query.redirect === "string" ? route.query.redirect : "/admin";
    router.push(redirect);
  } catch (error) {
    showToast(`Erreur: ${error.message}`, "error");
  } finally {
    loading.value = false;
  }
}
</script>
