package allergies

var allergiesList = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(code int) []string {
	var allergies []string
	for i, allergy := range allergiesList {
		if 0 != (code & (1 << uint(i))) {
			allergies = append(allergies, allergy)
		}
	}
	return allergies
}

func AllergicTo(code int, allergen string) bool {
	for i, allergy := range allergiesList {
		if allergy == allergen && 0 != (code&(1<<uint(i))) {
			return true
		}
	}
	return false
}
