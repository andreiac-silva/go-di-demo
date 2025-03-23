package setup

import "os"

var PostgresDSN string

func init() {
	PostgresDSN = os.Getenv("POSTGRES_DSN")
	if PostgresDSN == "" {
		panic("POSTGRES_DSN is required")
	}
}
