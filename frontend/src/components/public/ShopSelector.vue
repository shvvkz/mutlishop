<template>
  <div class="card border-0 shadow-sm mb-4" id="catalogue">
    <div class="card-body">
      <div class="row g-3 align-items-end">
        <div class="col-sm-4">
          <label class="form-label">Shop ID</label>
          <input
            :value="modelValue"
            class="form-control"
            min="1"
            type="number"
            @input="$emit('update:modelValue', $event.target.value)"
          />
        </div>
        <div class="col-sm-8 d-flex gap-2 flex-wrap">
          <button class="btn btn-primary" :disabled="loading" @click="$emit('load')">
            {{ loading ? "Chargement..." : "Charger le catalogue" }}
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
  loading: {
    type: Boolean,
    default: false,
  },
});

defineEmits(["update:modelValue", "load"]);
</script>
