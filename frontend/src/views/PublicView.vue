<template>
  <section>
    <ToastMessage
      v-model:show="toastVisible"
      :message="toastMessage"
      :type="toastType"
      :duration="3200"
    />

    <HeroSection />

    <ShopCreationForm :loading="creatingShop" @submit="handleCreateShop" />

    <ShopSelector
      v-model="shopId"
      :loading="loading"
      :shops="shops"
      @load="loadProducts"
      @reload-shops="loadShops"
    />

    <ProductStats :products="products" />

    <div class="row g-3">
      <div class="col-md-6 col-xl-4" v-for="product in products" :key="product.id">
        <ProductCard :product="product" :show-stock="false">
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
import ShopCreationForm from "../components/public/ShopCreationForm.vue";
import ShopSelector from "../components/public/ShopSelector.vue";
import {
  createPublicShop,
  getPublicProducts,
  getPublicShops,
  getPublicWhatsApp,
} from "../services/api";

const shopId = ref("");
const shops = ref([]);
const products = ref([]);
const loading = ref(false);
const creatingShop = ref(false);
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

function pickDefaultShop(list) {
  if (!Array.isArray(list) || !list.length) return "";
  const activeShop = list.find((shop) => shop.active);
  return String((activeShop || list[0]).id);
}

async function loadShops() {
  try {
    const payload = await getPublicShops();
    const list = extractArray(payload);
    shops.value = list;

    if (!shopId.value || !list.some((shop) => String(shop.id) === shopId.value)) {
      shopId.value = pickDefaultShop(list);
    }
  } catch (error) {
    showToast(`Erreur: ${error.message}`, "error");
  }
}

async function loadProducts() {
  if (!shopId.value) {
    products.value = [];
    showToast("Selectionne une boutique", "info");
    return;
  }

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

async function handleCreateShop(payload) {
  creatingShop.value = true;
  try {
    const response = await createPublicShop(payload);
    showToast("Shop et SuperAdmin crees", "success");
    await loadShops();
    const createdShopId = String(response?.data?.shop_id || response?.shop_id || "");
    if (createdShopId) {
      shopId.value = createdShopId;
      await loadProducts();
    }
  } catch (error) {
    showToast(`Erreur: ${error.message}`, "error");
  } finally {
    creatingShop.value = false;
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

onMounted(async () => {
  await loadShops();
  if (shopId.value) {
    await loadProducts();
  }
});
</script>
