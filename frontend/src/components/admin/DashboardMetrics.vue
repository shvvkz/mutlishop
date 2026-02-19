<template>
  <div class="card border-0 shadow-sm mb-4">
    <div class="card-header bg-white">
      <strong>Dashboard financier</strong>
    </div>
    <div class="card-body">
      <div v-if="!dashboard" class="text-muted">Aucune donnee chargee.</div>
      <div v-else class="row g-3">
        <div class="col-md-4">
          <MetricCard label="Ventes" :value="dashboard.total_sales" tone="success" />
        </div>
        <div class="col-md-4">
          <MetricCard label="Depenses" :value="dashboard.total_expenses" tone="danger" />
        </div>
        <div class="col-md-4">
          <MetricCard label="Retraits" :value="dashboard.total_withdrawals" tone="warning" />
        </div>
        <div class="col-md-6">
          <MetricCard label="Profit net" :value="dashboard.net_profit" tone="primary" />
        </div>
        <div class="col-md-6">
          <MetricCard
            label="Produits en stock faible"
            :value="dashboard.low_stock_products"
            tone="secondary"
            :is-currency="false"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import MetricCard from "./MetricCard.vue";

defineProps({
  dashboard: {
    type: Object,
    default: null,
  },
});
</script>
