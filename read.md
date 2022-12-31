# Simulation de la salle des urgences
## Fonctionnement du projet
### Types des agents
#### Patients
Chaque patient se caractérise principalement par

- Le type de maladie qu'il/elle a. Il est de type string.
- La gravité de la maladie, ce qui est liée à l'identification des ressources requises pour le traitement(Niveau et nombre de médecins nécessaire). Il est de type int.
- Le temps de traitement estimé de type int.
- Le temps d'attente toléré du patient de type int.

#### Infirmier
Les infirmiers se caractérisent principalement par

- La disponibilité de type bool