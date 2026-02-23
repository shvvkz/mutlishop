<template>
  <div class="card border-0 shadow-sm mb-4" id="catalogue">
    <div class="card-body">
      <div class="row g-3 align-items-end">
        <div class="col-sm-4">
          <label class="form-label">Boutique</label>
          <select
            class="form-select"
            :value="modelValue"
            :disabled="loading || !shops.length"
            @change="$emit('update:modelValue', $event.target.value)"
          >
            <option value="" disabled>{{ shops.length ? "Selectionner une boutique" : "Aucune boutique" }}</option>
            <option v-for="shop in shops" :key="shop.id" :value="String(shop.id)" :disabled="!shop.active">
              #{{ shop.id }} - {{ shop.name }}{{ shop.active ? "" : " (inactive)" }}
            </option>
          </select>
        </div>
        <div class="col-sm-8 d-flex gap-2 flex-wrap">
          <button class="btn btn-primary" :disabled="loading" @click="$emit('load')">
            {{ loading ? "Chargement..." : "Charger le catalogue" }}
          </button>
          <button class="btn btn-outline-secondary" :disabled="loading" @click="$emit('reload-shops')">
            Rafraichir les boutiques
          </button>
          <span class="align-self-center text-muted small">
            Endpoint: <code>/api/public/{{ modelValue }}/products</code>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  modelValue: {
    type: String,
    required: true,
  },
  shops: {
    type: Array,
    default: () => [],
  },
  loading: {
    type: Boolean,
    default: false,
  },
});

defineEmits(["update:modelValue", "load", "reload-shops"]);
</script>
