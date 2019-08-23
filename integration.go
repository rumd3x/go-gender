package gender

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// IBGEResponse maps endpoint JSON response body into struct
type IBGEResponse []struct {
	Name   string `json:"nome"`
	Gender string `json:"sexo"`
	Locale string `json:"localidade"`
	Res    []struct {
		Period    string `json:"periodo"`
		Frequency int    `json:"frequencia"`
	} `json:"res"`
}

// FindMales takes person first name and returns name frequency for males
func FindMales(name string) (int, error) {
	return FindBySex(name, "m")
}

// FindFemales takes person first name and returns name frequency for females
func FindFemales(name string) (int, error) {
	return FindBySex(name, "f")
}

// FindBySex takes person first name and sex and returns name frequency for the specified gender
func FindBySex(name, gender string) (int, error) {
	response, err := http.Get(fmt.Sprintf("https://servicodados.ibge.gov.br/api/v2/censos/nomes/%s?sexo=%s", name, gender))

	if err != nil {
		return 0, err
	}

	responseIBGE := IBGEResponse{}
	err = json.NewDecoder(response.Body).Decode(&responseIBGE)
	if err != nil {
		return 0, err
	}

	if len(responseIBGE) == 0 {
		return 0, nil
	}

	frequency := 0
	for _, f := range responseIBGE[0].Res {
		frequency += f.Frequency
	}

	return frequency, nil
}
