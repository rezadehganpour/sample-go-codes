package Switch

func PrintWithSwitch(p Person) (message string) {
	switch p {
	case "Reza":
		message = "Mr. " + p.Name + " " + p.Message
	default:
		message = "You dont have a best name in the world!"
	}
	return
}
