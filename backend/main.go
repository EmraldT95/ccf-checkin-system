package main

import (
	"ccf/check-in/common"
	"ccf/check-in/database"
	"ccf/check-in/server"
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// Parse the config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/ccf-check-in/")
	viper.AutomaticEnv()

	config, err := common.ParseConfigFile()
	if err != nil {
		panic(fmt.Errorf("Malformed config file: %w", err))
	}

	// Connect to the DB
	conn, err := common.GetDatabaseConn(context.Background(), &config.Database)
	if err != nil {
		panic(fmt.Errorf("Database connection failed: %w", err))
	}
	defer conn.Close(context.Background())

	// Run migrations
	err = database.RunMigration(conn)
	if err != nil {
		fmt.Printf("Failed to run migrations: %v", err)
		os.Exit(-1)
	}

	router := server.GetRouter(conn)
	fmt.Println("Running server on port 8000")

	err = router.Run(":8000")
	if err != nil {
		panic(fmt.Errorf("Failed to start server: %w", err))
	}
}
