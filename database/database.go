package database

import "fmt"

var connection string

// function dengan nama `init` auto ke run
func init() {
	connection = "MySQL"
	fmt.Println("database init terjalankan")
}

func GetDatabase() string {
	return connection //"MySQL"
}
