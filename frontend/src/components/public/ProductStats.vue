<template>
  <div class="row g-3 mb-4">
    <div class="col-sm-4">
      <div class="card border-0 bg-primary-subtle h-100">
        <div class="card-body">
          <p class="text-muted mb-1">Produits visibles</p>
          <p class="fs-4 fw-bold mb-0">{{ products.length }}</p>
        </div>
      </div>
    </div>
    <div class="col-sm-4">
      <div class="card border-0 bg-success-subtle h-100">
        <div class="card-body">
          <p class="text-muted mb-1">Valeur stock (vente)</p>
          <p class="fs-4 fw-bold mb-0">{{ stockValue }}</p>
        </div>
      </div>
    </div>
    <div class="col-sm-4">
      <div class="card border-0 bg-warning-subtle h-100">
        <div class="card-body">
          <p class="text-muted mb-1">Rupture / stock nul</p>
          <p class="fs-4 fw-bold mb-0">{{ outOfStock }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  products: {
    type: Array,
    default: () => [],
  },
});

const stockValue = computed(() => {
  const total = props.products.reduce((sum, product) => {
    return sum + Number(product.selling_price || 0) * Number(product.stock || 0);
  }, 0);
  return `${total.toFixed(2)} EUR`;
});

const outOfStock = computed(() => {
  return props.products.filter((product) => Number(product.stock || 0) <= 0).length;
});
</script>
