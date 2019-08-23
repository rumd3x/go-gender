package gender

// Gender describes gender implementation behavior
type Gender struct {
	male   bool
	female bool
	unisex bool
}

// Male Returns a new Male Gender
func Male() *Gender {
	return &Gender{male: true}
}

// Female Returns a new Female Gender
func Female() *Gender {
	return &Gender{female: true}
}

// Unisex Returns a new Unisex Gender
func Unisex() *Gender {
	return &Gender{true, true, true}
}

// Male Returns false if its certain its a famale, true otherwise
func (g *Gender) Male() bool {
	return g.male
}

// Female Returns false if its certain its a male, true otherwise
func (g *Gender) Female() bool {
	return g.female
}

// Unisex Returns false if its either male or female, true otherwise
func (g *Gender) Unisex() bool {
	return g.unisex
}

// String returns gender representation as a string
func (g *Gender) String() string {
	if g.Unisex() {
		return "Unisex"
	}

	if g.Male() {
		return "Male"
	}

	if g.Female() {
		return "Female"
	}

	return "Unknown"
}

// Equals compares to genders to see if they are equal
func (g *Gender) Equals(gender *Gender) bool {
	if g.Male() != gender.Male() {
		return false
	}

	if g.Female() != gender.Female() {
		return false
	}

	if g.Unisex() != gender.Unisex() {
		return false
	}

	return true
}
