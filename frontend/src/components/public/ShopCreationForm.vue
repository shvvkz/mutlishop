<template>
  <div class="card border-0 shadow-sm mb-4">
    <div class="card-header bg-white d-flex justify-content-between align-items-center">
      <strong>Creer une boutique</strong>
      <button class="btn btn-sm btn-outline-secondary" type="button" @click="toggleOpen">
        {{ open ? "Fermer" : "Nouveau shop" }}
      </button>
    </div>

    <div v-if="open" class="card-body">
      <form class="row g-3" @submit.prevent="submit">
        <div class="col-md-6">
          <label class="form-label">Nom boutique *</label>
          <input v-model.trim="form.name" class="form-control" required />
        </div>
        <div class="col-md-6">
          <label class="form-label">WhatsApp *</label>
          <input v-model.trim="form.whatsapp_number" class="form-control" required />
        </div>
        <div class="col-md-6">
          <label class="form-label">Email SuperAdmin *</label>
          <input v-model.trim="form.email" class="form-control" type="email" required />
        </div>
        <div class="col-md-6">
          <label class="form-label">Mot de passe *</label>
          <input v-model="form.password" class="form-control" type="password" minlength="6" required />
        </div>
        <div class="col-12">
          <button class="btn btn-success" :disabled="loading" type="submit">
            {{ loading ? "Creation..." : "Creer shop + SuperAdmin" }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from "vue";

defineProps({
  loading: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["submit"]);

const open = ref(false);
const form = reactive({
  name: "",
  whatsapp_number: "",
  email: "",
  password: "",
});

function toggleOpen() {
  open.value = !open.value;
}

function submit() {
  emit("submit", {
    name: form.name,
    whatsapp_number: form.whatsapp_number,
    email: form.email,
    password: form.password,
  });
}
</script>
