<template>
  <div class="card border-0 shadow-sm mb-4">
    <div class="card-header bg-white d-flex justify-content-between align-items-center">
      <strong>{{ mode === "create" ? "Nouveau produit" : "Modifier produit" }}</strong>
      <button v-if="mode === 'edit'" class="btn btn-sm btn-outline-secondary" @click="$emit('cancel')">
        Annuler
      </button>
    </div>

    <div class="card-body">
      <form class="row g-3" @submit.prevent="submit">
        <div class="col-md-6">
          <label class="form-label">Nom *</label>
          <input v-model.trim="form.name" class="form-control" required />
        </div>

        <div class="col-md-6">
          <label class="form-label">Categorie</label>
          <input v-model.trim="form.category" class="form-control" />
        </div>

        <div class="col-md-4">
          <label class="form-label">Prix achat *</label>
          <input v-model.number="form.purchase_price" class="form-control" type="number" min="0" step="0.01" required />
        </div>

        <div class="col-md-4">
          <label class="form-label">Prix vente *</label>
          <input v-model.number="form.selling_price" class="form-control" type="number" min="0" step="0.01" required />
        </div>

        <div class="col-md-4">
          <label class="form-label">Stock *</label>
          <input v-model.number="form.stock" class="form-control" type="number" min="0" step="1" required />
        </div>

        <div class="col-12">
          <label class="form-label">Description</label>
          <textarea v-model.trim="form.description" class="form-control" rows="2"></textarea>
        </div>

        <div class="col-12">
          <label class="form-label">Image URL</label>
          <input v-model.trim="form.image_url" class="form-control" />
        </div>

        <div class="col-12 d-flex gap-2">
          <button class="btn btn-primary" :disabled="loading" type="submit">
            {{ loading ? "En cours..." : mode === "create" ? "Creer le produit" : "Enregistrer" }}
          </button>
          <button type="button" class="btn btn-outline-secondary" @click="resetForm">Reinitialiser</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive, watch } from "vue";

const props = defineProps({
  mode: {
    type: String,
    default: "create",
  },
  initialProduct: {
    type: Object,
    default: null,
  },
  loading: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["submit", "cancel"]);

const defaultForm = () => ({
  name: "",
  category: "",
  description: "",
  image_url: "",
  purchase_price: 0,
  selling_price: 0,
  stock: 0,
});

const form = reactive(defaultForm());

function hydrateFromProduct(product) {
  const source = product || {};
  form.name = source.name || "";
  form.category = source.category || "";
  form.description = source.description || "";
  form.image_url = source.image_url || "";
  form.purchase_price = Number(source.purchase_price || 0);
  form.selling_price = Number(source.selling_price || 0);
  form.stock = Number(source.stock || 0);
}

function resetForm() {
  if (props.mode === "edit") {
    hydrateFromProduct(props.initialProduct);
    return;
  }
  hydrateFromProduct(null);
}

function submit() {
  emit("submit", {
    name: form.name,
    category: form.category,
    description: form.description,
    image_url: form.image_url,
    purchase_price: Number(form.purchase_price || 0),
    selling_price: Number(form.selling_price || 0),
    stock: Number(form.stock || 0),
  });
}

watch(
  () => props.initialProduct,
  (product) => {
    if (props.mode === "edit") {
      hydrateFromProduct(product);
      return;
    }
    resetForm();
  },
  { immediate: true },
);
</script>
