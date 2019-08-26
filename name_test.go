package gender

import "testing"

func TestGetGender(t *testing.T) {
	tests := map[string]*Gender{
		"Edmur":    Male(),
		"João":     Male(),
		"Maria":    Female(),
		"Darci":    Unisex(),
		"Lucas":    Male(),
		"Larissa":  Female(),
		"Alice":    Female(),
		"Helen":    Female(),
		"Matheus":  Male(),
		"Nathalia": Female(),
		"Mateus":   Male(),
		"Natalia":  Female(),
	}

	for name, gender := range tests {
		result, _ := GetGender(name)
		if !result.Equals(gender) {
			t.Errorf("Name %s expected to be %s but returned %s", name, gender.String(), result.String())
		}
	}

}
