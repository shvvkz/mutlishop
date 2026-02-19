<template>
  <div class="card border-0 shadow-sm mb-4">
    <div class="card-header bg-white">
      <strong>WhatsApp de la boutique</strong>
    </div>
    <div class="card-body">
      <form class="row g-3" @submit.prevent="submit">
        <div class="col-md-8">
          <label class="form-label">Numero WhatsApp *</label>
          <input
            v-model.trim="whatsAppNumber"
            class="form-control"
            placeholder="Ex: +2250700000000"
            required
          />
        </div>
        <div class="col-md-4 d-flex align-items-end">
          <button class="btn btn-primary w-100" :disabled="loading" type="submit">
            {{ loading ? "En cours..." : "Mettre a jour" }}
          </button>
        </div>
      </form>
      <p class="small text-muted mb-0 mt-2">
        Endpoint utilise: <code>PATCH /api/shop/whatsapp</code>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from "vue";

const props = defineProps({
  modelValue: {
    type: String,
    default: "",
  },
  loading: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["update:modelValue", "submit"]);

const whatsAppNumber = ref(props.modelValue || "");

watch(
  () => props.modelValue,
  (value) => {
    whatsAppNumber.value = value || "";
  },
);

function submit() {
  emit("update:modelValue", whatsAppNumber.value);
  emit("submit", whatsAppNumber.value);
}
</script>
