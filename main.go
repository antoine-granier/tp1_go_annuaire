package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

type Contact struct {
	Nom       string `json:"nom"`
	Telephone string `json:"telephone"`
}

type Annuaire struct {
	Contacts []Contact `json:"contacts"`
}

func chargerAnnuaire(fichier string) Annuaire {
	data, err := os.ReadFile(fichier)
	if err != nil {
		return Annuaire{}
	}
	var a Annuaire
	_ = json.Unmarshal(data, &a)
	return a
}

func (a *Annuaire) sauvegarder(fichier string) error {
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fichier, data, 0644)
}

func (a *Annuaire) rechercherIndex(nom string) (int, bool) {
	for i, c := range a.Contacts {
		if c.Nom == nom {
			return i, true
		}
	}
	return -1, false
}

func (a *Annuaire) ajouter(nom, tel string) error {
	if _, exists := a.rechercherIndex(nom); exists {
		return errors.New("un contact avec ce nom existe déjà")
	}
	a.Contacts = append(a.Contacts, Contact{Nom: nom, Telephone: tel})
	return nil
}

func (a *Annuaire) supprimer(nom string) error {
	if i, found := a.rechercherIndex(nom); found {
		a.Contacts = append(a.Contacts[:i], a.Contacts[i+1:]...)
		return nil
	}
	return errors.New("contact introuvable")
}

func (a *Annuaire) modifier(nom, nouveauTel string) error {
	if i, found := a.rechercherIndex(nom); found {
		a.Contacts[i].Telephone = nouveauTel
		return nil
	}
	return errors.New("contact introuvable")
}

func (a *Annuaire) rechercher(nom string) (*Contact, error) {
	if i, found := a.rechercherIndex(nom); found {
		return &a.Contacts[i], nil
	}
	return nil, errors.New("contact introuvable")
}

func (a *Annuaire) lister() []Contact {
	return a.Contacts
}

func main() {
	action := flag.String("action", "", "Action à effectuer: ajouter, supprimer, modifier, rechercher, lister")
	nom := flag.String("nom", "", "Nom et prénom du contact")
	tel := flag.String("tel", "", "Numéro de téléphone du contact")
	flag.Parse()

	annuaire := chargerAnnuaire("contacts.json")

	var err error
	switch *action {
	case "ajouter":
		err = annuaire.ajouter(*nom, *tel)
		if err == nil {
			fmt.Println("Contact ajouté.")
		}
	case "supprimer":
		err = annuaire.supprimer(*nom)
		if err == nil {
			fmt.Println("Contact supprimé.")
		}
	case "modifier":
		err = annuaire.modifier(*nom, *tel)
		if err == nil {
			fmt.Println("Contact modifié.")
		}
	case "rechercher":
		var contact *Contact
		contact, err = annuaire.rechercher(*nom)
		if err == nil {
			fmt.Printf("Contact trouvé: %s - %s\n", contact.Nom, contact.Telephone)
		}
	case "lister":
		for _, c := range annuaire.lister() {
			fmt.Printf("%s - %s\n", c.Nom, c.Telephone)
		}
	default:
		fmt.Println("Action inconnue. Utilisez --action avec: ajouter, supprimer, modifier, rechercher, lister.")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Erreur:", err)
		os.Exit(1)
	}

	if *action == "ajouter" || *action == "supprimer" || *action == "modifier" {
		if err := annuaire.sauvegarder("contacts.json"); err != nil {
			fmt.Println("Erreur lors de la sauvegarde:", err)
			os.Exit(1)
		}
	}
}
