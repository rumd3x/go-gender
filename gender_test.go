package gender

import "testing"

func TestGendersString(t *testing.T) {
	genders := map[string]*Gender{
		"Unisex":  Unisex(),
		"Male":    Male(),
		"Female":  Female(),
		"Unknown": &Gender{},
	}

	for name, g := range genders {
		if g.String() != name {
			t.Errorf("Expected %s but gender string representation returned %s", name, g.String())
		}
	}
}

func TestGendersEquals(t *testing.T) {
	pairs := map[*Gender]*Gender{
		Male():                      Female(),
		Female():                    Unisex(),
		Unisex():                    Male(),
		&Gender{}:                   Unisex(),
		&Gender{false, false, true}: Unisex(),
		&Gender{true, true, false}:  Unisex(),
		&Gender{false, true, false}: Unisex(),
		&Gender{true, false, false}: Unisex(),
	}

	for g1, g2 := range pairs {
		if g1.Equals(g2) {
			t.Errorf("Expected genders %s and %s to be different", g1.String(), g2.String())
		}
	}
}
