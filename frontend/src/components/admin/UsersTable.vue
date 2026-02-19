<template>
  <div class="card border-0 shadow-sm mb-4">
    <div class="card-header bg-white d-flex justify-content-between align-items-center">
      <strong>Utilisateurs</strong>
      <span class="badge text-bg-light">{{ users.length }}</span>
    </div>
    <div class="table-responsive">
      <table class="table table-hover mb-0">
        <thead>
          <tr>
            <th>ID</th>
            <th>Email</th>
            <th>Role</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="!users.length">
            <td colspan="4" class="text-muted">Aucun utilisateur charge.</td>
          </tr>
          <tr v-for="user in users" :key="user.id">
            <td>{{ user.id }}</td>
            <td>{{ user.email }}</td>
            <td>
              <select
                class="form-select form-select-sm"
                :value="user.role"
                @change="$emit('change-role', { id: user.id, role: $event.target.value })"
              >
                <option value="Admin">Admin</option>
                <option value="SuperAdmin">SuperAdmin</option>
              </select>
            </td>
            <td>
              <button class="btn btn-sm btn-outline-danger" @click="$emit('delete', user.id)">
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
  users: {
    type: Array,
    default: () => [],
  },
});

defineEmits(["change-role", "delete"]);
</script>
