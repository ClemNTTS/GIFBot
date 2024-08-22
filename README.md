```markdown
# GIFBot

GIFBot est un bot Discord développé en Go qui répond aux messages des utilisateurs avec des GIFs pertinents basés sur le thème des messages. Il utilise la bibliothèque `discordgo` pour interagir avec l'API Discord et `prose` pour le traitement du langage naturel.

## Fonctionnalités

- **Détection de Thèmes** : Identifie le thème principal des messages reçus en utilisant des n-grammes.
- **Réponse Automatique** : Envoie une réponse automatique basée sur le thème détecté.

## Prérequis

Avant de commencer, assurez-vous d'avoir les éléments suivants :

- [Go](https://golang.org/doc/install) installé sur votre machine.
- Un compte Discord et un bot Discord avec un jeton (token) valide.
- [Git](https://git-scm.com/) pour cloner le dépôt.

## Installation

1. **Cloner le Dépôt**

   ```bash
   git clone https://github.com/ClemNTTS/GIFBot.git
   cd GIFBot
   ```

2. **Installer les Dépendances**

   Assurez-vous que vous avez les modules Go nécessaires installés. Utilisez `go mod` pour gérer les dépendances :

   ```bash
   go mod tidy
   ```

3. **Configurer le Bot**

   - Remplacez le jeton du bot dans le code source par le vôtre.
   - Mettez à jour le fichier de configuration si nécessaire (pour les permissions ou les paramètres de l'API).

4. **Exécuter le Bot**

   Lancez le bot en utilisant la commande suivante :

   ```bash
   go run main.go
   ```

   Assurez-vous que votre bot est invité dans un serveur Discord avec les bonnes permissions.

## Utilisation

Lorsque le bot est en ligne, il surveille les messages dans les serveurs Discord où il est présent. Lorsqu'un message est reçu, le bot analyse le texte pour identifier le thème principal et répond avec un message prédéfini (par exemple, "HA"). 

### Exemple de Message

Si vous envoyez un message tel que "Clément est tellement beau", le bot peut analyser et répondre selon le thème détecté.

## Contribuer

Les contributions sont les bienvenues ! Si vous souhaitez ajouter des fonctionnalités ou corriger des bugs, veuillez suivre les étapes ci-dessous :

1. **Forker le Dépôt** : Cliquez sur "Fork" en haut à droite de cette page.
2. **Créer une Branche** : Créez une branche pour votre fonctionnalité ou correction de bug.
   ```bash
   git checkout -b ma-nouvelle-fonctionnalité
   ```
3. **Apporter des Modifications** : Faites vos modifications et ajoutez des tests si nécessaire.
4. **Committer vos Changements** :
   ```bash
   git add .
   git commit -m "Description de ma modification"
   ```
5. **Pousser la Branche** :
   ```bash
   git push origin ma-nouvelle-fonctionnalité
   ```
6. **Créer une Pull Request** : Ouvrez une pull request sur GitHub.

## License

Distribué sous la licence MIT.

## Contact

- **Nom** : Clément [@ClemNTTS](https://github.com/ClemNTTS)
- **Email** : [clément@example.com](mailto:clément@example.com)

---

Merci d'avoir utilisé GIFBot ! N'hésitez pas à poser des questions ou à demander de l'aide si nécessaire.
```
