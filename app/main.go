package main

func main() {
	const (
		host         = "localhost"
		port         = 5431
		user         = "postgres"
		password     = "postgres"
		databaseName = "postgres"
		address      = ":8080"
	)

	app := App{}
	app.Initialize(host, port, user, password, databaseName)
	app.Run(address)
}
