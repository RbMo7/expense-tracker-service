package database

import (
	"fmt"
	"net"
	"time"
)

func CheckConnection() error {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", "localhost:5432", timeout)
	if err != nil {
		fmt.Println("❌ Postgres is NOT reachable on port 5432:", err)
		return err
	}
	conn.Close()
	fmt.Println("✅ Postgres is reachable on port 5432")
	return nil
}
