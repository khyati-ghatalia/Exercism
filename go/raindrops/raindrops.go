package raindrops

import "strconv"

//Convert function Converts a number to a string, the contents of which depend on the number's factors.
func Convert(input int) string {
	RaindropNumbers := [3]int{3, 5, 7}
	RaindropsStrings := [3]string{"Pling", "Plang", "Plong"}

	var Outputstring string

	for i := 0; i < len(RaindropNumbers); i++ {
		if input%RaindropNumbers[i] == 0 {
			Outputstring = Outputstring + RaindropsStrings[i]
		}
	}
	if Outputstring == "" {
		return strconv.Itoa(input)
	}
	return Outputstring

}
