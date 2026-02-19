<template>
  <div class="card border-0 shadow-sm mb-4">
    <div class="card-header bg-white d-flex justify-content-between align-items-center">
      <strong>Transactions</strong>
      <span class="badge text-bg-light">{{ transactions.length }}</span>
    </div>
    <div class="table-responsive">
      <table class="table table-striped mb-0">
        <thead>
          <tr>
            <th>ID</th>
            <th>Type</th>
            <th>Montant</th>
            <th>Quantite</th>
            <th>Produit</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="!transactions.length">
            <td colspan="6" class="text-muted">Aucune transaction chargee.</td>
          </tr>
          <tr v-for="tx in transactions" :key="tx.id">
            <td>{{ tx.id }}</td>
            <td>{{ tx.type }}</td>
            <td>{{ Number(tx.amount || 0).toFixed(2) }} EUR</td>
            <td>{{ tx.quantity || 0 }}</td>
            <td>{{ tx.product_id || "-" }}</td>
            <td>
              <button class="btn btn-sm btn-outline-danger" @click="$emit('delete', tx.id)">
                Supprimer
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
defineProps({
  transactions: {
    type: Array,
    default: () => [],
  },
});

defineEmits(["delete"]);
</script>
