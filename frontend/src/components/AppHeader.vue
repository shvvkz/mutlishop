<template>
  <header class="border-bottom bg-white sticky-top shadow-sm">
    <nav class="navbar navbar-expand-lg container py-3">
      <router-link class="navbar-brand fw-bold d-flex align-items-center gap-2" to="/">
        <span class="badge text-bg-primary">MS</span>
        <span>MultiShop</span>
      </router-link>

      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#app-nav"
        aria-controls="app-nav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <div id="app-nav" class="collapse navbar-collapse">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link class="nav-link" to="/">Vitrine</router-link>
          </li>
          <li class="nav-item" v-if="isAuthenticated">
            <router-link class="nav-link" :to="adminLink">Admin</router-link>
          </li>
        </ul>

        <div class="d-flex align-items-center gap-2 flex-wrap justify-content-end">
          <span class="badge" :class="isAuthenticated ? 'text-bg-success' : 'text-bg-secondary'">
            {{ isAuthenticated ? "Connecte" : "Invite" }}
          </span>

          <span v-if="isAuthenticated" class="badge text-bg-light">
            Role: {{ authRole }}
          </span>

          <span v-if="isAuthenticated && authInfo?.shopId" class="badge text-bg-light">
            Shop: {{ authInfo.shopId }}
          </span>

          <router-link v-if="!isAuthenticated" class="btn btn-outline-primary btn-sm" to="/login">
            Login
          </router-link>
          <button v-else class="btn btn-outline-danger btn-sm" @click="logout">Logout</button>
        </div>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { computed, onBeforeUnmount, ref } from "vue";
import { useRouter } from "vue-router";
import { clearToken, getAuthInfo, getToken, onAuthChanged } from "../services/api";

const router = useRouter();

const token = ref(getToken());
const authInfo = ref(getAuthInfo());

const isAuthenticated = computed(() => Boolean(token.value));
const authRole = computed(() => authInfo.value?.role || "Inconnu");

const adminLink = computed(() => {
  if (!isAuthenticated.value) return "/login";
  const tab = authRole.value === "SuperAdmin" ? "dashboard" : "products";
  return `/admin?tab=${tab}`;
});

const unsubscribe = onAuthChanged(() => {
  token.value = getToken();
  authInfo.value = getAuthInfo();
});

onBeforeUnmount(() => {
  unsubscribe();
});

function logout() {
  clearToken();
  router.push("/login");
}
</script>
