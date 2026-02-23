<template>
  <div class="card border-0 shadow-sm mb-4">
    <div class="card-header bg-white">
      <strong>Dashboard operationnel</strong>
    </div>
    <div class="card-body">
      <div class="row g-3">
        <div class="col-md-4">
          <MetricCard
            label="Produits"
            :value="productsCount"
            tone="secondary"
            :is-currency="false"
          />
        </div>
        <div class="col-md-4">
          <MetricCard
            label="Transactions"
            :value="transactionsCount"
            tone="primary"
            :is-currency="false"
          />
        </div>
        <div class="col-md-4">
          <MetricCard
            label="Stock faible (<5)"
            :value="lowStockProducts"
            tone="warning"
            :is-currency="false"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import MetricCard from "./MetricCard.vue";

const props = defineProps({
  products: {
    type: Array,
    default: () => [],
  },
  transactions: {
    type: Array,
    default: () => [],
  },
});

const productsCount = computed(() => props.products.length);
const transactionsCount = computed(() => props.transactions.length);
const lowStockProducts = computed(() => {
  return props.products.filter((product) => Number(product.stock || 0) < 5).length;
});
</script>
