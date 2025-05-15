Groupe:
- DERENSY Dany
- GRANIER Antoine

# Annuaire – Gestion de contacts (CLI)

Ce programme permet de gérer un annuaire de contacts via la ligne de commande. Les contacts sont stockés dans un fichier `contacts.json`.

---

## Commandes disponibles

### Ajouter un contact

```bash
go run main.go --action=ajouter --nom="Jean Dupont" --tel="0612345678"
```
### Supprimer un contact

```bash
go run main.go --action=supprimer --nom="Jean Dupont"
```
### Modifier un contact

```bash
go run main.go --action=modifier --nom="Jean Dupont" --tel="0699999999"
```
### Rechercher un contact

```bash
go run main.go --action=rechercher --nom="Jean Dupont"
```
### Lister tous les contacts

```bash
go run main.go --action=lister
```

## Lancer les tests
```bash
go test -v
```
