# 🤖 Utilisation de l’Intelligence Artificielle dans le Projet

## 2️⃣ Outils IA utilisés

Dans le cadre de ce projet, nous avons utilisé principalement :

* **ChatGPT**

Nous n’avons pas utilisé Copilot ou Cursor de manière active pour générer du code automatique dans l’éditeur.

L’IA a été utilisée comme un outil d’assistance et d’accélération, et non comme un générateur autonome de projet.

---

## 📌 Rôle de l’IA dans le projet

L’IA a été utilisée principalement pour :

* Générer la logique métier (couche `services/`)
* Aider à structurer certaines règles métier
* Générer la documentation Swagger
* Clarifier certaines implémentations techniques

Elle n’a PAS été utilisée pour :

* Définir les modèles (qui étaient fournis)
* Écrire l’intégralité des controllers
* Structurer l’architecture globale
* Concevoir la logique d’isolation multi-tenant

Les controllers ont été développés manuellement.
Les models étaient déjà définis dans le cahier des charges.
L’architecture et la séparation en couches ont été pensées et mises en place par nous-mêmes.

---

## 📝 Exemple de Prompt utilisé (Documentation Swagger)

Exemple de fonction :

```go
func GetTransactions(c *gin.Context) {

	shopID := c.GetUint("shop_id")

	transactions, err := services.GetTransactions(shopID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, transactions)
}
```

Prompt utilisé :

> "Génère les annotations Swagger (Swaggo) pour cette fonction Gin en respectant les bonnes pratiques : résumé, description, tags, sécurité JWT, codes de retour, et route associée."

L’IA a généré la structure des commentaires Swagger (`@Summary`, `@Description`, `@Tags`, `@Security`, `@Success`, `@Failure`, `@Router`).

Nous avons ensuite :

* Vérifié chaque annotation
* Adapté les types retournés
* Corrigé les éventuelles incohérences

---

## ⚙️ Méthode de Travail avec l’IA

L’IA a été utilisée comme un outil pour produire du code plus rapidement.

Cependant, chaque ligne générée a été :

* Lue
* Compris
* Vérifiée
* Testée

Nous avons systématiquement relu le code avant de l’intégrer au projet afin d’éviter toute hallucination ou génération incorrecte.

Aucun code n’a été intégré sans validation humaine.

---

## 7️⃣ Analyse Critique de l’IA

### ✅ Là où l’IA nous a fait gagner du temps

* Génération rapide de logique métier répétitive
* Création des annotations Swagger
* Rappels de bonnes pratiques
* Structuration initiale de certaines fonctions

Elle a permis d’accélérer considérablement la phase d’implémentation des services.

---

### ⚠️ Là où l’IA nous a fait perdre du temps

* Certaines réponses généraient du code incorrect ou incomplet
* Des incohérences dans les types ou les signatures
* Des oublis liés à la gestion multi-tenant
* Des approximations sur les permissions Admin / SuperAdmin

Cela nécessitait une vérification approfondie.

---

### 🔎 Ce que nous avons dû corriger manuellement

* Adaptation précise des règles de permissions
* Correction de certaines requêtes GORM
* Ajustement des structures de retour
* Intégration correcte avec notre architecture existante

L’IA a été utilisée comme un accélérateur, mais la responsabilité finale du code, de l’architecture et de la sécurité reste entièrement humaine.

---

## 🎯 Conclusion

L’IA a été un outil d’assistance et de productivité.

Elle n’a pas remplacé la réflexion architecturale ni la validation technique.

Le projet a été conçu, structuré et vérifié manuellement, l’IA ayant uniquement servi à accélérer certaines parties spécifiques comme la logique métier et la documentation.
