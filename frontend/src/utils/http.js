export async function parseJsonSafe(response) {
  try {
    return await response.json();
  } catch (_) {
    return null;
  }
}

export function getErrorMessage(response, payload) {
  if (!payload || typeof payload !== "object") {
    return `HTTP ${response.status}`;
  }

  return (
    payload.message ||
    payload.error ||
    payload.data?.message ||
    payload.data?.error ||
    `HTTP ${response.status}`
  );
}

export function throwIfNotOk(response, payload) {
  if (response.ok) return;
  throw new Error(getErrorMessage(response, payload));
}
