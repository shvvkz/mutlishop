<template>
  <div class="card border-0 shadow-sm mb-4">
    <div class="card-header bg-white">
      <strong>Nouvelle transaction</strong>
    </div>
    <div class="card-body">
      <form class="row g-3" @submit.prevent="submit">
        <div class="col-md-4">
          <label class="form-label">Type *</label>
          <select v-model="form.type" class="form-select" required>
            <option value="Sale">Sale</option>
            <option value="Expense">Expense</option>
            <option value="Withdrawal">Withdrawal</option>
          </select>
        </div>

        <div class="col-md-4">
          <label class="form-label">Montant *</label>
          <input v-model.number="form.amount" class="form-control" type="number" min="0" step="0.01" required />
        </div>

        <div class="col-md-4">
          <label class="form-label">Produit</label>
          <select v-model.number="form.product_id" class="form-select" :disabled="form.type !== 'Sale'">
            <option :value="0">Aucun</option>
            <option v-for="product in products" :key="product.id" :value="product.id">
              {{ product.name }} (ID {{ product.id }})
            </option>
          </select>
        </div>

        <div class="col-md-4">
          <label class="form-label">Quantite</label>
          <input
            v-model.number="form.quantity"
            class="form-control"
            type="number"
            min="1"
            step="1"
            :disabled="form.type !== 'Sale'"
          />
        </div>

        <div class="col-12 d-flex gap-2">
          <button class="btn btn-primary" :disabled="loading" type="submit">
            {{ loading ? "En cours..." : "Creer transaction" }}
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
  products: {
    type: Array,
    default: () => [],
  },
  loading: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["submit"]);

const form = reactive({
  type: "Sale",
  amount: 0,
  product_id: 0,
  quantity: 1,
});

function resetForm() {
  form.type = "Sale";
  form.amount = 0;
  form.product_id = 0;
  form.quantity = 1;
}

function submit() {
  const payload = {
    type: form.type,
    amount: Number(form.amount || 0),
  };

  if (form.type === "Sale") {
    payload.product_id = Number(form.product_id || 0);
    payload.quantity = Number(form.quantity || 1);
  }

  emit("submit", payload);
}
</script>
