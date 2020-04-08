package twofer

//ShareWith is a method which takes null or a name and returns a string
// If the input parameter is null it returns "One for you, one for me."
// if the input parameter is a value e.g Alice it returns "One for Alice, one for me."
func ShareWith(name string) string {

	if len(name) < 1 {
		return ("One for you, one for me.")
	}
	return ("One for " + name + ", one for me.")
}
