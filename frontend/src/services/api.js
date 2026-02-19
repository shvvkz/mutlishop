import { parseJsonSafe, throwIfNotOk } from "../utils/http";

const TOKEN_KEY = "ms_token";
const AUTH_EVENT = "ms-auth-changed";

function toJsonHeaders(headers = {}) {
  return {
    "Content-Type": "application/json",
    ...headers,
  };
}

function withAuth(headers = {}) {
  const token = localStorage.getItem(TOKEN_KEY);
  if (!token) return headers;
  return { ...headers, Authorization: `Bearer ${token}` };
}

async function request(path, { method = "GET", body, headers = {}, auth = false } = {}) {
  const finalHeaders = body
    ? toJsonHeaders(auth ? withAuth(headers) : headers)
    : auth
      ? withAuth(headers)
      : headers;

  const response = await fetch(path, {
    method,
    headers: finalHeaders,
    body: body ? JSON.stringify(body) : undefined,
  });

  const payload = await parseJsonSafe(response);
  throwIfNotOk(response, payload);

  return payload;
}

export function getToken() {
  return localStorage.getItem(TOKEN_KEY) || "";
}

export function setToken(token) {
  localStorage.setItem(TOKEN_KEY, token);
  window.dispatchEvent(new Event(AUTH_EVENT));
}

export function clearToken() {
  localStorage.removeItem(TOKEN_KEY);
  window.dispatchEvent(new Event(AUTH_EVENT));
}

export function onAuthChanged(callback) {
  window.addEventListener(AUTH_EVENT, callback);
  return () => window.removeEventListener(AUTH_EVENT, callback);
}

function decodeBase64Url(input) {
  const padded = input.replace(/-/g, "+").replace(/_/g, "/");
  const pad = padded.length % 4;
  const normalized = pad ? padded + "=".repeat(4 - pad) : padded;
  return atob(normalized);
}

export function getAuthInfo() {
  const token = getToken();
  if (!token) return null;

  const parts = token.split(".");
  if (parts.length < 2) return null;

  try {
    const payload = JSON.parse(decodeBase64Url(parts[1]));
    return {
      userId: payload.user_id,
      shopId: payload.shop_id,
      role: payload.role || "",
      exp: payload.exp,
    };
  } catch (_) {
    return null;
  }
}

export function login(email, password) {
  return request("/api/login", { method: "POST", body: { email, password } });
}

export function getPublicProducts(shopId) {
  return request(`/api/public/${shopId}/products`);
}

export function getPublicWhatsApp(shopId, productId) {
  return request(`/api/public/${shopId}/products/${productId}/whatsapp`);
}

export function getProducts() {
  return request("/api/products", { auth: true });
}

export function createProduct(payload) {
  return request("/api/products", {
    method: "POST",
    auth: true,
    body: payload,
  });
}

export function updateProduct(id, payload) {
  return request(`/api/products/${id}`, {
    method: "PUT",
    auth: true,
    body: payload,
  });
}

export function deleteProduct(id) {
  return request(`/api/products/${id}`, {
    method: "DELETE",
    auth: true,
  });
}

export function getTransactions() {
  return request("/api/transactions", { auth: true });
}

export function createTransaction(payload) {
  return request("/api/transactions", {
    method: "POST",
    auth: true,
    body: payload,
  });
}

export function deleteTransaction(id) {
  return request(`/api/transactions/${id}`, {
    method: "DELETE",
    auth: true,
  });
}

export function getUsers() {
  return request("/api/users", { auth: true });
}

export function createUser(payload) {
  return request("/api/users", {
    method: "POST",
    auth: true,
    body: payload,
  });
}

export function updateUserRole(id, role) {
  return request(`/api/users/${id}`, {
    method: "PATCH",
    auth: true,
    body: { role },
  });
}

export function deleteUser(id) {
  return request(`/api/users/${id}`, {
    method: "DELETE",
    auth: true,
  });
}

export function getDashboard() {
  return request("/api/reports/dashboard", { auth: true });
}

export function updateShopWhatsApp(whatsappNumber) {
  return request("/api/shop/whatsapp", {
    method: "PATCH",
    auth: true,
    body: { whatsapp_number: whatsappNumber },
  });
}
