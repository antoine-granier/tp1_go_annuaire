package main

import (
	"testing"
)

func TestAnnuaire(t *testing.T) {
	type step struct {
		name       string
		action     func(*Annuaire) (any, error)
		expectErr  bool
		expectData any
	}

	tests := []step{
		{
			name: "Ajouter Alice",
			action: func(a *Annuaire) (any, error) {
				return nil, a.ajouter("Alice Dupont", "0123456789")
			},
		},
		{
			name: "Ajouter doublon Alice",
			action: func(a *Annuaire) (any, error) {
				return nil, a.ajouter("Alice Dupont", "0987654321")
			},
			expectErr: true,
		},
		{
			name: "Modifier Alice",
			action: func(a *Annuaire) (any, error) {
				return nil, a.modifier("Alice Dupont", "0000000000")
			},
		},
	}

	a := Annuaire{
		Contacts: []Contact{},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.action(&a)

			if tc.expectErr && err == nil {
				t.Errorf("attendait une erreur")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("erreur inattendue: %q", err)
			}
			if tc.expectData != nil && err == nil {
				if got != tc.expectData {
					t.Errorf("valeur attendue %q, obtenue %q", tc.expectData, got)
				}
			}
		})
	}
}
