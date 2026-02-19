<template>
  <div class="card border-0 shadow-sm mb-4">
    <div class="card-header bg-white">
      <strong>Nouvel utilisateur Admin</strong>
    </div>
    <div class="card-body">
      <form class="row g-3" @submit.prevent="submit">
        <div class="col-md-4">
          <label class="form-label">Nom *</label>
          <input v-model.trim="form.name" class="form-control" required />
        </div>

        <div class="col-md-4">
          <label class="form-label">Email *</label>
          <input v-model.trim="form.email" class="form-control" type="email" required />
        </div>

        <div class="col-md-4">
          <label class="form-label">Mot de passe *</label>
          <input v-model="form.password" class="form-control" type="password" minlength="6" required />
        </div>

        <div class="col-12 d-flex gap-2">
          <button class="btn btn-primary" :disabled="loading" type="submit">
            {{ loading ? "En cours..." : "Creer admin" }}
          </button>
          <button class="btn btn-outline-secondary" type="button" @click="resetForm">Reinitialiser</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive } from "vue";

defineProps({
  loading: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["submit"]);

const form = reactive({
  name: "",
  email: "",
  password: "",
});

function resetForm() {
  form.name = "";
  form.email = "";
  form.password = "";
}

function submit() {
  emit("submit", {
    name: form.name,
    email: form.email,
    password: form.password,
  });
}
</script>
