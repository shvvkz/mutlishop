<template>
  <section class="bg-white border rounded-4 p-4 p-md-5 mb-4 shadow-sm">
    <div class="row g-4 align-items-center">
      <div class="col-lg-8">
        <p class="text-uppercase text-primary fw-semibold mb-2">Catalogue public</p>
        <h1 class="display-6 fw-bold mb-3">Vos produits en ligne, commandes directes via WhatsApp</h1>
        <p class="lead text-muted mb-4">
          Cette vitrine consomme les endpoints publics documentes dans Swagger:
          <code>GET /api/public/:shopID/products</code> et
          <code>GET /api/public/:shopID/products/:productID/whatsapp</code>.
        </p>
        <div class="d-flex gap-2 flex-wrap">
          <router-link class="btn btn-primary" to="/login">Espace admin</router-link>
          <a class="btn btn-outline-secondary" href="#catalogue">Voir le catalogue</a>
        </div>
      </div>
      <div class="col-lg-4">
        <div class="p-3 rounded-3 bg-light border">
          <div class="small text-muted mb-2">Parcours client</div>
          <ol class="mb-0">
            <li>Selection du shop</li>
            <li>Consultation des produits</li>
            <li>Contact WhatsApp en un clic</li>
          </ol>
        </div>
      </div>
    </div>
  </section>
</template>
