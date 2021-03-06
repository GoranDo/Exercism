package space

type Planet string

const earthSeconds = 31557600

func Age(period float64, p Planet) float64 {
	var year float64
	switch p {
	case "Mercury":
		year = 0.2408467
	case "Venus":
		year = 0.61519726
	case "Mars":
		year = 1.8808158
	case "Jupiter":
		year = 11.862615
	case "Saturn":
		year = 29.447498
	case "Uranus":
		year = 84.016846
	case "Neptune":
		year = 164.79132
	case "Earth":
		year = 1.0
	}
	return period / (earthSeconds * year)
}
