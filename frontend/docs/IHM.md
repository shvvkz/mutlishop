# Documentation IHM Frontend (Vue.js)

## Contexte et objectif
Le frontend de MultiShop a ete concu comme une IHM simple et exploitable rapidement pour consommer les endpoints de l'API Go documentes via Swagger.

L'objectif n'etait pas de produire un design complexe, mais une interface claire, testable et orientee flux metier:
- consultation catalogue public,
- authentification admin,
- operations de gestion (produits, transactions, utilisateurs, dashboard).

## Realisation avec l'aide de l'IA
L'implementation a ete acceleree avec un assistant IA pour:
- proposer une architecture frontend modulaire,
- generer des composants Vue reutilisables,
- adapter rapidement l'IHM quand les endpoints backend ont evolue,
- securiser l'integration via validations et gestion d'erreurs.

L'IA a servi de copilote technique, mais les choix metier (roles, parcours, priorites) ont ete valides au fil de l'eau pendant le developpement.

## Pourquoi Vue.js
Vue.js a ete retenu pour les raisons suivantes:
- prise en main rapide et DX simple,
- composants lisibles et faciles a maintenir,
- integration naturelle avec Vite pour un cycle dev/build rapide,
- bonne adequation avec un projet API-first.

Dans ce projet, Vue a permis de decouper clairement l'IHM en blocs fonctionnels (Header, formulaires, tableaux, toasts, etc.) sans surcharger les vues.

## Liaison Frontend / Backend
L'integration front/back repose sur 4 principes:

1. **Client API centralise**
- `frontend/src/services/api.js`
- toutes les requetes HTTP y sont regroupees (public + prive JWT), y compris l'evolution recente des transactions:
  - `/api/transactions/sale`
  - `/api/transactions/expense`
  - `/api/transactions/withdrawal`

2. **Gestion d'erreur uniforme**
- `frontend/src/utils/http.js`
- parsing safe des reponses + messages d'erreur homogenes.

3. **Auth JWT cote client**
- stockage token local,
- extraction des claims (`role`, `shop_id`) pour adapter l'IHM,
- guards de navigation sur `/admin`.

4. **Proxy de dev Vite**
- `frontend/vite.config.js`
- en dev, `/api/*` est proxy vers `http://localhost:8080`.

## Experience utilisateur recherchee
L'UX ciblee est pragmatique: "simple, rapide, sans friction".

- **Vitrine publique**
  - affichage du catalogue,
  - CTA WhatsApp direct pour convertir la consultation en prise de contact.

- **Espace admin**
  - navigation par onglets (Dashboard, Produits, Transactions, Utilisateurs),
  - formulaires explicites plutot que boites de dialogue systeme,
  - feedback immediat via toasts (succes/erreur),
  - restrictions d'interface selon role (`Admin` vs `SuperAdmin`) pour eviter les actions interdites.

- **Lisibilite technique**
  - composants petits et reutilisables,
  - separation claire entre presentation et appels API,
  - adaptation rapide aux changements backend sans rework massif.

## Etat actuel
L'IHM couvre les principaux endpoints du Swagger et est fonctionnelle pour les parcours metier standards.

Des ameliorations futures restent possibles (pagination, recherche, edition inline, tests E2E), mais la base actuelle est stable et exploitable.
