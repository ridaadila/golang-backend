package main

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		println("You are blacklisted", name)
	} else {
		println("Welcome", name)
	}
}

func main() {
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
