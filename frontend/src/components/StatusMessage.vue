<template>
  <div v-if="message" :class="['alert', alertClass]" role="alert">
    {{ message }}
  </div>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  message: {
    type: String,
    default: "",
  },
  type: {
    type: String,
    default: "info",
  },
});

const alertClass = computed(() => {
  if (props.type === "error") return "alert-danger";
  if (props.type === "success") return "alert-success";
  return "alert-secondary";
});
</script>
