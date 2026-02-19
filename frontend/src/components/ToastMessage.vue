<template>
  <div class="toast-container position-fixed top-0 end-0 p-3">
    <div v-if="show" class="toast show border-0 shadow-sm" :class="bgClass" role="alert">
      <div class="d-flex">
        <div class="toast-body">{{ message }}</div>
        <button
          type="button"
          class="btn-close me-2 m-auto"
          :class="closeClass"
          aria-label="Close"
          @click="emit('update:show', false)"
        ></button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, watch } from "vue";

const props = defineProps({
  show: {
    type: Boolean,
    default: false,
  },
  message: {
    type: String,
    default: "",
  },
  type: {
    type: String,
    default: "info",
  },
  duration: {
    type: Number,
    default: 3000,
  },
});

const emit = defineEmits(["update:show"]);

const bgClass = computed(() => {
  if (props.type === "error") return "text-bg-danger";
  if (props.type === "success") return "text-bg-success";
  return "text-bg-primary";
});

const closeClass = computed(() => {
  return props.type === "error" || props.type === "success" || props.type === "info"
    ? "btn-close-white"
    : "";
});

watch(
  () => props.show,
  (visible) => {
    if (!visible) return;
    window.setTimeout(() => {
      emit("update:show", false);
    }, props.duration);
  },
);
</script>
