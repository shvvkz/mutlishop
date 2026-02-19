<template>
  <section>
    <div class="mb-3">
      <div>
        <h2 class="h4 mb-1">Espace Admin</h2>
        <p class="text-muted mb-0">
          Role courant: <strong>{{ currentRole || "Inconnu" }}</strong>
        </p>
      </div>
    </div>

    <ToastMessage
      v-model:show="toastVisible"
      :message="toastMessage"
      :type="toastType"
      :duration="3200"
    />

    <ul class="nav nav-pills mb-3">
      <li class="nav-item" v-for="tab in allowedTabs" :key="tab">
        <button
          class="nav-link"
          :class="{ active: activeTab === tab }"
          type="button"
          @click="switchTab(tab)"
        >
          {{ tabLabel(tab) }}
        </button>
      </li>
    </ul>

    <div v-if="activeTab === 'products'">
      <ProductForm
        :mode="formMode"
        :initial-product="editingProduct"
        :loading="submittingProduct"
        @submit="handleProductSubmit"
        @cancel="cancelEdit"
      />

      <div class="card border-0 shadow-sm mb-4">
        <div class="card-header bg-white d-flex justify-content-between align-items-center">
          <strong>Produits</strong>
          <span class="badge text-bg-light">{{ products.length }}</span>
        </div>
        <div class="card-body">
          <div class="row g-3">
            <div class="col-md-6 col-xl-4" v-for="product in products" :key="product.id">
              <ProductCard :product="product" :show-purchase-price="showPurchasePrice">
                <template #actions>
                  <button class="btn btn-outline-secondary btn-sm" @click="editProduct(product)">
                    Modifier
                  </button>
                  <button
                    v-if="isSuperAdmin"
                    class="btn btn-outline-danger btn-sm"
                    @click="removeProduct(product.id)"
                  >
                    Supprimer
                  </button>
                </template>
              </ProductCard>
            </div>
            <div v-if="!products.length" class="col-12 text-muted">Aucun produit charge.</div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="activeTab === 'transactions'">
      <TransactionForm
        :products="products"
        :loading="submittingTransaction"
        @submit="handleTransactionSubmit"
      />
      <TransactionsTable :transactions="transactions" @delete="removeTransaction" />
    </div>

    <div v-if="activeTab === 'users' && isSuperAdmin">
      <UserForm :loading="submittingUser" @submit="handleUserSubmit" />
      <UsersTable :users="users" @change-role="changeUserRole" @delete="removeUser" />
    </div>

    <div v-if="activeTab === 'dashboard' && isSuperAdmin">
      <ShopWhatsAppForm
        v-model="shopWhatsAppNumber"
        :loading="submittingWhatsApp"
        @submit="submitShopWhatsApp"
      />
      <DashboardMetrics :dashboard="dashboard" />
    </div>
  </section>
</template>

<script setup>
import { computed, onBeforeUnmount, onMounted, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import DashboardMetrics from "../components/admin/DashboardMetrics.vue";
import ProductForm from "../components/admin/ProductForm.vue";
import ShopWhatsAppForm from "../components/admin/ShopWhatsAppForm.vue";
import TransactionForm from "../components/admin/TransactionForm.vue";
import TransactionsTable from "../components/admin/TransactionsTable.vue";
import UserForm from "../components/admin/UserForm.vue";
import UsersTable from "../components/admin/UsersTable.vue";
import ProductCard from "../components/ProductCard.vue";
import ToastMessage from "../components/ToastMessage.vue";
import {
  createProduct,
  createTransaction,
  createUser,
  deleteProduct,
  deleteTransaction,
  deleteUser,
  getAuthInfo,
  getDashboard,
  getProducts,
  getTransactions,
  getUsers,
  onAuthChanged,
  updateShopWhatsApp,
  updateUserRole,
  updateProduct,
} from "../services/api";

const route = useRoute();
const router = useRouter();
const ALL_TABS = ["dashboard", "products", "transactions", "users"];

const authInfo = ref(getAuthInfo());
const products = ref([]);
const transactions = ref([]);
const users = ref([]);
const dashboard = ref(null);
const toastVisible = ref(false);
const toastMessage = ref("");
const toastType = ref("info");
const activeTab = ref("dashboard");
const formMode = ref("create");
const editingProduct = ref(null);
const submittingProduct = ref(false);
const submittingTransaction = ref(false);
const submittingUser = ref(false);
const submittingWhatsApp = ref(false);
const shopWhatsAppNumber = ref("");

const currentRole = computed(() => authInfo.value?.role || "");
const isSuperAdmin = computed(() => currentRole.value === "SuperAdmin");
const allowedTabs = computed(() => (isSuperAdmin.value ? ALL_TABS : ["products", "transactions"]));

const showPurchasePrice = computed(() => {
  return products.value.some((product) => product.purchase_price !== undefined);
});

const unsubscribeAuth = onAuthChanged(() => {
  authInfo.value = getAuthInfo();
});

onBeforeUnmount(() => {
  unsubscribeAuth();
});

function tabLabel(tab) {
  if (tab === "dashboard") return "Dashboard";
  if (tab === "products") return "Produits";
  if (tab === "transactions") return "Transactions";
  return "Utilisateurs";
}

function getDefaultTab() {
  return isSuperAdmin.value ? "dashboard" : "products";
}

function extractArray(payload) {
  if (Array.isArray(payload)) return payload;
  return payload?.data || [];
}

function setLoading(message) {
  // no-op: loading is represented by per-section button states
}

function setError(error) {
  showToast(`Erreur: ${error?.message || "operation impossible"}`, "error");
}

function setSuccess(message) {
  showToast(message, "success");
}

function clearStatus() {
  // no-op
}

function showToast(message, type = "info") {
  toastMessage.value = message;
  toastType.value = type;
  toastVisible.value = false;
  window.setTimeout(() => {
    toastVisible.value = true;
  }, 0);
}

async function switchTab(tab, syncUrl = true) {
  const nextTab = allowedTabs.value.includes(tab) ? tab : getDefaultTab();
  activeTab.value = nextTab;

  if (syncUrl) {
    await router.replace({ query: { ...route.query, tab: nextTab } });
  }

  if (nextTab === "products") {
    await loadProducts();
    return;
  }

  if (nextTab === "transactions") {
    await Promise.all([loadTransactions(), loadProducts()]);
    return;
  }

  if (nextTab === "users" && isSuperAdmin.value) {
    await loadUsers();
    return;
  }

  if (nextTab === "dashboard" && isSuperAdmin.value) {
    await loadDashboard();
  }
}

async function initTabFromUrl() {
  const queryTab = typeof route.query.tab === "string" ? route.query.tab : "";
  const initialTab = allowedTabs.value.includes(queryTab) ? queryTab : getDefaultTab();
  await switchTab(initialTab, true);
}

async function loadProducts() {
  setLoading("Chargement des produits...");
  try {
    const payload = await getProducts();
    products.value = extractArray(payload);
    clearStatus();
  } catch (error) {
    setError(error);
  }
}

async function loadTransactions() {
  setLoading("Chargement des transactions...");
  try {
    const payload = await getTransactions();
    transactions.value = extractArray(payload);
    clearStatus();
  } catch (error) {
    setError(error);
  }
}

async function loadUsers() {
  if (!isSuperAdmin.value) return;
  setLoading("Chargement des utilisateurs...");
  try {
    const payload = await getUsers();
    users.value = extractArray(payload);
    clearStatus();
  } catch (error) {
    setError(error);
  }
}

async function loadDashboard() {
  if (!isSuperAdmin.value) return;
  setLoading("Chargement du dashboard...");
  try {
    const payload = await getDashboard();
    dashboard.value = payload?.data || payload;
    clearStatus();
  } catch (error) {
    setError(error);
  }
}

async function removeProduct(id) {
  if (!isSuperAdmin.value) return;
  if (!window.confirm("Supprimer ce produit ?")) return;

  try {
    await deleteProduct(id);
    setSuccess("Produit supprime");
    await loadProducts();
  } catch (error) {
    setError(error);
  }
}

function editProduct(product) {
  formMode.value = "edit";
  editingProduct.value = { ...product };
  window.scrollTo({ top: 0, behavior: "smooth" });
}

function cancelEdit() {
  formMode.value = "create";
  editingProduct.value = null;
}

async function handleProductSubmit(payload) {
  submittingProduct.value = true;
  try {
    if (formMode.value === "create") {
      await createProduct(payload);
      setSuccess("Produit cree");
    } else {
      await updateProduct(editingProduct.value.id, payload);
      setSuccess("Produit modifie");
      cancelEdit();
    }
    await loadProducts();
  } catch (error) {
    setError(error);
  } finally {
    submittingProduct.value = false;
  }
}

async function handleTransactionSubmit(payload) {
  submittingTransaction.value = true;
  try {
    if (payload.type === "Sale" && !payload.product_id) {
      throw new Error("Selectionne un produit pour une transaction Sale");
    }
    await createTransaction(payload);
    setSuccess("Transaction creee");
    await Promise.all([loadTransactions(), loadProducts()]);
    if (isSuperAdmin.value) await loadDashboard();
  } catch (error) {
    setError(error);
  } finally {
    submittingTransaction.value = false;
  }
}

async function removeTransaction(id) {
  if (!window.confirm("Supprimer cette transaction ?")) return;

  try {
    await deleteTransaction(id);
    setSuccess("Transaction supprimee");
    await Promise.all([loadTransactions(), loadProducts()]);
    if (isSuperAdmin.value) await loadDashboard();
  } catch (error) {
    setError(error);
  }
}

async function handleUserSubmit(payload) {
  if (!isSuperAdmin.value) return;
  submittingUser.value = true;
  try {
    await createUser(payload);
    setSuccess("Utilisateur admin cree");
    await loadUsers();
  } catch (error) {
    setError(error);
  } finally {
    submittingUser.value = false;
  }
}

async function changeUserRole(input) {
  if (!isSuperAdmin.value) return;
  try {
    await updateUserRole(input.id, input.role);
    setSuccess("Role utilisateur mis a jour");
    await loadUsers();
  } catch (error) {
    setError(error);
  }
}

async function removeUser(id) {
  if (!isSuperAdmin.value) return;
  if (!window.confirm("Supprimer cet utilisateur ?")) return;

  try {
    await deleteUser(id);
    setSuccess("Utilisateur supprime");
    await loadUsers();
  } catch (error) {
    setError(error);
  }
}

async function submitShopWhatsApp(number) {
  if (!isSuperAdmin.value) return;
  submittingWhatsApp.value = true;
  try {
    await updateShopWhatsApp(number);
    setSuccess("Numero WhatsApp de la boutique mis a jour");
    shopWhatsAppNumber.value = number;
  } catch (error) {
    setError(error);
  } finally {
    submittingWhatsApp.value = false;
  }
}

onMounted(async () => {
  await loadProducts();
  await loadTransactions();
  if (isSuperAdmin.value) {
    await loadUsers();
    await loadDashboard();
  }
  await initTabFromUrl();
});

watch(
  () => route.query.tab,
  async (value) => {
    if (typeof value !== "string") return;
    if (value === activeTab.value) return;
    await switchTab(value, false);
  },
);

watch(
  () => currentRole.value,
  async () => {
    const safeTab = allowedTabs.value.includes(activeTab.value) ? activeTab.value : getDefaultTab();
    await switchTab(safeTab, true);
  },
);
</script>
