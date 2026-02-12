# 📦 Multi-Shop Electronics Backend – Specification (Go)

## 🎯 Objectif

Développer un backend en **Go** pour un système de gestion multi-boutiques d’électronique incluant :

* Isolation complète des shops (multi-tenant strict)
* Gestion des rôles internes
* Page publique accessible sans authentification
* Redirection dynamique vers WhatsApp

---

# 🏢 Contexte métier

Chaque boutique possède :

* 👑 Super Admin
* 🧑‍💼 Admin
* 🌍 Une page publique accessible aux clients

Les clients doivent pouvoir :

* Voir les produits disponibles
* Voir le stock
* Cliquer sur un produit
* Être redirigés vers le WhatsApp du magasin

---

# 🧱 Modèle de données obligatoire

## 1️⃣ Shop (Done)

* ID
* Name
* Active
* WhatsAppNumber
* CreatedAt

---

## 2️⃣ User (Done)

* ID
* Name
* Email
* Password (bcrypt)
* Role (SuperAdmin | Admin)
* ShopID
* CreatedAt

⚠️ Aucun Guest en base de données.

---

## 3️⃣ Product (Done)

* ID
* Name
* Description
* Category
* PurchasePrice
* SellingPrice
* Stock
* ImageURL
* ShopID
* CreatedAt

---

## 4️⃣ Transaction (Done)

* ID
* Type (Sale | Expense | Withdrawal)
* ProductID
* Quantity
* Amount
* ShopID
* CreatedAt

---

# 🔐 Gestion des rôles

## 👑 SuperAdmin

Peut :

* CRUD produits (Done)
* Voir PurchasePrice (Done)
* Voir profits (Done)
* Voir dashboard (To Do)
* Gérer utilisateurs (Done)
* Modifier WhatsAppNumber du shop (To Do)

---

## 🧑‍💼 Admin

Peut :

* CRUD produits (Done)
* CRUD transactions (Done)
* Voir SellingPrice (Done)
* Voir stock (Done)

Ne peut pas :

* Voir PurchasePrice
* Voir profit
* Modifier WhatsAppNumber

---

## 👥 Guest (Client)

* Aucun compte
* Accès public uniquement

Peut :

* Voir les produits disponibles (Done)
* Voir le stock (Done)
* Cliquer pour demander information (Done)

---

# 🌐 Routes API obligatoires

## 🔑 Auth

```
POST /register (Done)
POST /login (Done)
```

* JWT obligatoire
* ShopID extrait du token pour toutes les routes privées

---

## 📦 Produits (privé)

```
GET    /products (Done)
POST   /products (Done)
PUT    /products/:id (Done)
DELETE /products/:id (Done)
```

* Filtrage obligatoire par ShopID issu du JWT

---

## 🌍 Produits publics (Guest)

```
GET /public/:shopID/products (Done)
```

Retourne uniquement :

* Name
* Description
* Category
* SellingPrice
* Stock
* ImageURL

⚠️ Ne jamais exposer PurchasePrice.

---

# 📱 Route WhatsApp (Done)

Lorsqu’un client clique sur un produit, le backend doit générer dynamiquement un lien :

Format :

```
https://wa.me/<WhatsAppNumber>?text=Bonjour%20je%20veux%20plus%20d'information%20sur%20<NomProduit>
```

Exemple :

```
https://wa.me/212600000000?text=Bonjour%20je%20veux%20plus%20d'information%20sur%20iPhone%2014
```

Ce lien doit être généré côté backend.

---

# 🧠 Logique métier obligatoire

## 1️⃣ Multi-tenant strict (Done)

* Même pour les routes publiques
* Un client ne doit voir que les produits du shop demandé

---

## 2️⃣ Gestion du stock (Done)

Si `stock = 0`, deux possibilités :

* Afficher "Out of stock"
  OU
* Ne pas afficher le produit

---

## 3️⃣ Sécurité (Done)

* PurchasePrice jamais exposé publiquement
* ShopID toujours extrait du JWT pour routes privées
* Validation des rôles via middleware

---

# 📊 Dashboard (SuperAdmin uniquement) (Done)

```
GET /reports/dashboard
```

Doit retourner :

* Total ventes
* Total dépenses
* Profit net
* Nombre de produits en stock faible (< 5)

---

# 🎨 Frontend (libre) (To Do)

Possibilités :

* React
* HTML simple
* Swagger UI

Doit démontrer :

* Page publique fonctionnelle
* Redirection WhatsApp correcte
* Respect des rôles
* Isolation multi-shop fonctionnelle

---

# 📦 Livrables

* Repository GitHub
* README
* ERD
* Diagramme d’architecture
* Collection Postman
* Vidéo de démonstration

---

# 📏 Grille d’évaluation

| Critère                  | Points |
| ------------------------ | ------ |
| Auth & JWT               | /15    |
| Multi-tenant             | /20    |
| Gestion rôles            | /15    |
| Logique stock            | /10    |
| Route publique sécurisée | /10    |
| WhatsApp dynamique       | /10    |
| Code Go propre           | /10    |
| Documentation            | /10    |

---

# 🚀 Pourquoi ce projet est intéressant

Ce projet combine :

* Backend sécurisé
* Multi-tenant réel
* Logique business concrète
* API publique + privée
* Génération dynamique d’URL
* Gestion stricte des rôles
* Séparation finance / public
