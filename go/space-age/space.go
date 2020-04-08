package space

//Age takes as a input seconds and planet to give the number of years on that planet
//Earth: orbital period 365.25 Earth days, or 31557600 seconds
//Mercury: orbital period 0.2408467 Earth years
//Venus: orbital period 0.61519726 Earth years
//Mars: orbital period 1.8808158 Earth years
//Jupiter: orbital period 11.862615 Earth years
//Saturn: orbital period 29.447498 Earth years
//Uranus: orbital period 84.016846 Earth years
//Neptune: orbital period 164.79132 Earth years

const Earthrevolution = 31557600

type Planet string

func Age(seconds float64, planets Planet) float64 {
	if seconds < 0 {
		return 0
	}
	switch planets {
	case "Earth":
		return seconds / Earthrevolution
	case "Mercury":
		return seconds / (Earthrevolution * 0.2408467)
	case "Venus":
		return seconds / (Earthrevolution * 0.61519726)
	case "Mars":
		return seconds / (Earthrevolution * 1.8808158)
	case "Jupiter":
		return seconds / (Earthrevolution * 11.862615)
	case "Saturn":
		return seconds / (Earthrevolution * 29.447498)
	case "Uranus":
		return seconds / (Earthrevolution * 84.016846)
	case "Neptune":
		return seconds / (Earthrevolution * 164.79132)

	}
	return 0
}
