<template>
  <div class="d-flex flex-column min-vh-100 bg-body-tertiary">
    <AppHeader />

    <main class="container py-4 flex-grow-1">
      <router-view />
    </main>

    <AppFooter />
  </div>
</template>

<script setup>
import AppFooter from "./components/layout/AppFooter.vue";
import AppHeader from "./components/AppHeader.vue";
</script>
