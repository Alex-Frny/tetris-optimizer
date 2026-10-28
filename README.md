# Tetris Optimizer

Projet Go : assembler les tétriminos dans le plus petit carré possible grâce au backtracking.

## Prérequis

Go 1.25.0 ou supérieur. Aucune dépendance externe.

## Lancement

Depuis la racine du projet :

```sh
go run . sample.txt
```

Le programme affiche une grille carrée : les pièces sont identifiées par les lettres A à Z et les cases vides par des points. Leur orientation reste celle du fichier fourni.

Chaque pièce est décrite dans une grille de 4 lignes de 4 caractères (`#` ou `.`). Les pièces sont séparées par une ligne vide. Une pièce doit contenir quatre cases reliées par leurs côtés. Au maximum 26 pièces sont acceptées.

Une entrée invalide, un fichier introuvable ou un mauvais nombre d’arguments affiche `ERROR`.

## Tests

```sh
go test ./...
```

Les tests couvrent le parsing, la validation et le solveur, notamment le placement et le retour arrière.

## Organisation

- `main.go` : lecture du fichier et affichage de la solution.
- `parser.go` et `validator.go` : lecture et validation des pièces.
- `tetromino.go` : structures des pièces.
- `solver.go` : recherche de la plus petite grille.

## Préparation pour le portfolio

Dans cette copie, la commande principale a été raccordée au solveur déjà présent : elle affiche désormais la solution au lieu du seul nombre de pièces valides. Les sources originales sous Ubuntu ne sont pas modifiées.
