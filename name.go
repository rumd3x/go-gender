package gender

import (
	"errors"
	"sync"
)

//GetGender takes person first name and returns its gender prediction
func GetGender(name string) (*Gender, error) {
	var wg sync.WaitGroup
	var malesErr error
	var femalesErr error
	var errMsg string

	males := 0
	females := 0

	wg.Add(2)

	go func() {
		males, malesErr = FindMales(name)
		if malesErr != nil {
			errMsg = errMsg + malesErr.Error()
		}
		wg.Done()
	}()

	go func() {
		females, femalesErr = FindFemales(name)
		if femalesErr != nil {
			errMsg = errMsg + femalesErr.Error()
		}
		wg.Done()
	}()

	if errMsg != "" {
		return nil, errors.New(errMsg)
	}

	wg.Wait()

	genderRate := float64(males) / float64(females)

	if genderRate > 2.5 {
		return Male(), nil
	}

	if genderRate <= 0.4 {
		return Female(), nil
	}

	return Unisex(), nil
}
