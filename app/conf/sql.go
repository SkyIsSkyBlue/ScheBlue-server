package conf

import "os"

func SqlDatabase() string {
	return os.Getenv("MARIADB_DATABASE")
}

func SqlUsername() string {
	return os.Getenv("MARIADB_USERNAME")
}

func SqlPassword() string {
	return os.Getenv("MARIADB_PASSWORD")
}

func SqlHostname() string {
	return os.Getenv("MARIADB_HOSTNAME")
}

func SqlPort() uint {
	return 3306
}
