package main

func main() {
	config := NewConfig()

	config.ReadConfig()
	config.SetUser("igor")
	config.ReadConfig()
}
