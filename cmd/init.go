package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"schedule/internal/config"
	"schedule/internal/domain/mappers"
	"schedule/internal/domain/services"
	"schedule/internal/infrastructure/repositories"
	"schedule/internal/presentation/api/controllers"
	"schedule/internal/presentation/utils"
)

func Init() {
	cfg := config.GetConfig()

	// Init db
	db, err := sqlx.Connect(cfg.Driver, fmt.Sprintf("dbname=%s sslmode=disable", cfg.DbName))
	if err != nil {
		log.Fatalln(err)
	}

	// Init Kafka
	brokers := []string{fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)}

	// Init repositories
	classRepository := repositories.NewClassRepository(db)
	groupRepository := repositories.NewGroupRepository(db)
	userRepository := repositories.NewUserRepository(db)

	// Init mappers
	classMapper := mappers.NewClassMapper()
	userMapper := mappers.NewUserMapper()

	// Init services
	classService := services.NewClassService(classRepository, classMapper)
	groupService := services.NewGroupService(groupRepository, brokers)
	groupService.StartConsuming()
	userService := services.NewUserService(userRepository, userMapper)

	// Init controllers
	router := mux.NewRouter()
	classController := controllers.NewClassController(classService, userService)
	groupController := controllers.NewGroupController()
	classController.SetupRoutes(router)
	groupController.SetupGroupRoutes(router)
	router.Use(utils.Recovery)
	log.Printf("Running on http://%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), router))
}
