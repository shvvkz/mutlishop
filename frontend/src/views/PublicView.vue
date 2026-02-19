<template>
  <section>
    <ToastMessage
      v-model:show="toastVisible"
      :message="toastMessage"
      :type="toastType"
      :duration="3200"
    />

    <HeroSection />

    <ShopSelector v-model="shopId" :loading="loading" @load="loadProducts" />

    <ProductStats :products="products" />

    <div class="row g-3">
      <div class="col-md-6 col-xl-4" v-for="product in products" :key="product.id">
        <ProductCard :product="product">
          <template #actions>
            <button class="btn btn-success btn-sm" @click="openWhatsApp(product.id)">
              Commander sur WhatsApp
            </button>
          </template>
        </ProductCard>
      </div>
    </div>
  </section>
</template>

<script setup>
import { onMounted, ref } from "vue";
import ProductCard from "../components/ProductCard.vue";
import ToastMessage from "../components/ToastMessage.vue";
import HeroSection from "../components/public/HeroSection.vue";
import ProductStats from "../components/public/ProductStats.vue";
import ShopSelector from "../components/public/ShopSelector.vue";
import { getPublicProducts, getPublicWhatsApp } from "../services/api";

const shopId = ref("1");
const products = ref([]);
const loading = ref(false);
const toastVisible = ref(false);
const toastMessage = ref("");
const toastType = ref("info");

function showToast(message, type = "info") {
  toastMessage.value = message;
  toastType.value = type;
  toastVisible.value = false;
  window.setTimeout(() => {
    toastVisible.value = true;
  }, 0);
}

function extractArray(payload) {
  if (Array.isArray(payload)) return payload;
  return payload?.data || [];
}

async function loadProducts() {
  loading.value = true;

  try {
    const payload = await getPublicProducts(shopId.value);
    products.value = extractArray(payload);
  } catch (error) {
    showToast(`Erreur: ${error.message}`, "error");
  } finally {
    loading.value = false;
  }
}

async function openWhatsApp(productId) {
  try {
    const payload = await getPublicWhatsApp(shopId.value, productId);
    const url = payload?.data?.whatsapp_url || payload?.whatsapp_url;

    if (!url) {
      showToast("Lien WhatsApp introuvable", "error");
      return;
    }

    window.open(url, "_blank");
  } catch (error) {
    showToast(`Erreur: ${error.message}`, "error");
  }
}

onMounted(loadProducts);
</script>
