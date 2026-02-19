<template>
  <article class="card border-0 shadow-sm h-100">
    <div class="card-body d-flex flex-column">
      <div class="d-flex justify-content-between align-items-start mb-2">
        <h5 class="card-title mb-0">{{ product.name }}</h5>
        <span class="badge text-bg-light">{{ product.category || "General" }}</span>
      </div>

      <p class="text-muted small mb-2" v-if="showStock">Stock: {{ product.stock }}</p>
      <p class="fw-semibold fs-5 mb-1">{{ Number(product.selling_price || 0).toFixed(2) }} EUR</p>
      <p v-if="showPurchasePrice && product.purchase_price !== undefined" class="small text-muted mb-3">
        Achat: {{ Number(product.purchase_price || 0).toFixed(2) }} EUR
      </p>

      <div v-if="$slots.actions" class="mt-auto d-flex flex-wrap gap-2">
        <slot name="actions" />
      </div>
    </div>
  </article>
</template>

<script setup>
defineProps({
  product: {
    type: Object,
    required: true,
  },
  showStock: {
    type: Boolean,
    default: true,
  },
  showPurchasePrice: {
    type: Boolean,
    default: false,
  },
});
</script>
