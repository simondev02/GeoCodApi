package main

import (
	"SimonBK_Vehiculo/domain/models"
	"SimonBK_Vehiculo/infra/db"
	"fmt"
	"log"
)

func RunMigratye() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	err = db.DBConn.AutoMigrate(&models.VehicleGroup{})
	if err != nil {
		log.Fatalf("Failed to migrate table Vehicle: %v", err)
	}

	fmt.Println("Migration successful")
}
