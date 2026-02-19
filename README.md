# 🛒 Electronic Multi-Shop API – Bootcamp Go

## 📌 Description du Projet

Electronic Multi-Shop API est un backend développé en **Go** permettant de gérer plusieurs boutiques d’électronique avec une isolation stricte des données entre chaque boutique.

Chaque boutique possède :

* Des Super Admins
* Des Admins
* Une page publique des produits (sans authentification)

Le système garantit :

* Une isolation complète multi-tenant
* Un contrôle d’accès basé sur les rôles
* Une authentification sécurisée via JWT
* La gestion des stocks
* Une redirection WhatsApp pour les clients

---

## 🧱 Stack Technique

* **Go 1.25+**
* **Gin** (framework HTTP)
* **GORM** (ORM)
* **PostgreSQL 16**
* **Authentification JWT**
* **Docker & Docker Compose**

---

## 🖥 Prérequis

Vous avez uniquement besoin de :

* Docker
* Docker Compose

Rien d’autre.

Vous n’avez PAS besoin d’installer manuellement :

* Go
* PostgreSQL
* Une quelconque dépendance

---

## ⚙️ Variables d’Environnement

Un fichier `.env` est requis à la racine du projet.

Exemple :

```
DB_HOST=postgres
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin123
DB_NAME=multishop

JWT_SECRET=your_secret_key_here
JWT_EXPIRE_HOURS=24
```

Nous avons mit à disposition un fichier example.env.

Pour créer simplement votre fichier .env veuillez executer cette commande:
```
mv example.env .env
``` 

---

## 🚀 Installation & Exécution (Étapes)

### 1️⃣ Cloner le repository

```
git clone https://github.com/shvvkz/mutlishop.git
cd mutlishop
```

---

### 2️⃣ Créer le fichier .env

Créer un fichier `.env` à la racine du projet et y copier les variables d’environnement indiquées ci-dessus.

---

### 3️⃣ Lancer l’application

```
docker compose up --build
```

Cette commande va :

* Construire l’API Go
* Démarrer PostgreSQL dans un conteneur
* Démarrer le conteneur de l’API
* Connecter automatiquement l’API à la base de données

---

### 4️⃣ Accéder à l’API

Une fois l’application démarrée avec succès :

```
http://localhost:8080
```

L’API est alors prête à être utilisée.

---

## 🔐 Authentification

L’authentification est gérée via JWT.

Après connexion :

* Un token JWT est généré
* Ce token doit être inclus dans les requêtes :

```
Authorization: Bearer <your_token>
```

Les restrictions sont appliquées selon les rôles :

* SuperAdmin
* Admin

Les routes publiques ne nécessitent aucune authentification.

---

## 🐳 Architecture Docker

Le projet fonctionne avec deux services :

* `postgres` → Base de données PostgreSQL
* `api` → Backend Go

PostgreSQL fonctionne en interne sur le port `5432`.
L’API est exposée sur `localhost:8080`.

Aucune installation manuelle de base de données n’est nécessaire.

---

## 🏪 Données Initiales (Auto Seed)

Lors du premier démarrage, la base de données est automatiquement initialisée avec :

* **2 Shops**
* **1 SuperAdmin par shop**

Cela permet de tester immédiatement :

* L’isolation multi-tenant
* Les permissions basées sur les rôles
* Les flux d’authentification

Chaque SuperAdmin est strictement lié à sa propre boutique, garantissant que :

* Les données d’une boutique ne sont jamais accessibles par une autre
* Toutes les opérations restent isolées par tenant

---

## 📚 Documentation API (Swagger)

Le projet expose une interface de documentation interactive Swagger.

Une fois l’application démarrée, elle est accessible à l’adresse :

```
http://localhost:8080/swagger/index.html#/
```

Swagger permet de :

* Explorer tous les endpoints disponibles
* Tester les routes authentifiées
* Comprendre les formats requête/réponse
* Vérifier les restrictions liées aux rôles

Cela rend l’API entièrement testable sans outil externe.

---

## 🏗 Architecture du Projet

### Comment nous avons structuré le projet

Le projet est organisé selon une architecture en couches claire afin de séparer les responsabilités :

* `main.go` → Point d’entrée de l’application
* `config/` → Configuration (connexion base de données, variables d’environnement)
* `models/` → Définition des entités (Shop, User, Product, Transaction, etc.)
* `controllers/` → Gestion des requêtes HTTP
* `services/` → Logique métier
* `middlewares/` → Gestion JWT et contrôle des rôles
* `utils/` → Fonctions utilitaires

Cette séparation permet :

* Une meilleure lisibilité du code
* Une maintenance plus simple
* Une évolution facilitée
* Une isolation claire entre logique métier et couche HTTP

---

### Comment nous avons planifié l’architecture

L’architecture a été pensée dès le départ autour de trois contraintes principales :

1. L’isolation multi-tenant
2. La gestion stricte des rôles (Admin / SuperAdmin)
3. La séparation claire des responsabilités

Nous avons d’abord défini le modèle de données (Shop, User, Product, Transaction) afin de structurer la base.
Ensuite, nous avons conçu les règles d’accès selon les rôles pour garantir que chaque utilisateur ne puisse agir que dans le périmètre de sa boutique.

Enfin, nous avons construit l’API en respectant cette structure, en ajoutant progressivement :

* La gestion JWT
* Les middlewares d’autorisation
* La validation des accès par shop

Cette planification a permis de conserver une architecture cohérente et évolutive tout au long du projet.

---

## Difficultés rencontrées

Durant le projet, nous n'avons presque pas rencontré de difficultés techniques majeures.

La principale difficulté a été de maintenir correctement la gestion des droits entre les rôles **Admin** et **SuperAdmin** tout au long du développement.

Il a fallu s’assurer que chaque action respecte strictement les permissions associées au rôle et que l’isolation multi-tenant reste cohérente à chaque évolution du code.

Le reste du développement s’est déroulé de manière fluide.

---

## Utilisation de l'IA

Pour voir l'utilisation de l'IA dans notre projet, nous avons créer un document que vous pouvez consulter ici:
[Voir le document IA](./ai.md)