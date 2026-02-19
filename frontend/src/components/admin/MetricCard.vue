<template>
  <div class="card h-100 border-0" :class="bgClass">
    <div class="card-body">
      <p class="mb-1 text-muted">{{ label }}</p>
      <p class="mb-0 fw-bold fs-5">{{ formattedValue }}</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  label: {
    type: String,
    required: true,
  },
  value: {
    type: [String, Number],
    default: 0,
  },
  tone: {
    type: String,
    default: "secondary",
  },
  isCurrency: {
    type: Boolean,
    default: true,
  },
});

const bgClass = computed(() => `bg-${props.tone}-subtle`);

const formattedValue = computed(() => {
  const numeric = Number(props.value || 0);
  if (!props.isCurrency) return `${numeric}`;
  return `${numeric.toFixed(2)} EUR`;
});
</script>
