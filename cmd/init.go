package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"schedule/internal/config"
	"schedule/internal/domain/kafkaListeners"
	"schedule/internal/domain/mappers"
	"schedule/internal/domain/services"
	"schedule/internal/infrastructure/repositories"
	"schedule/internal/presentation/controllers"
	"schedule/internal/presentation/utils"
	"schedule/kafka"
)

func Init() {
	cfg := config.GetConfig()

	// Init db
	db, err := sqlx.Connect(cfg.Driver, fmt.Sprintf("dbname=%s sslmode=disable", cfg.DbName))
	if err != nil {
		log.Fatalln(err)
	}

	// Init repositories
	classRepository := repositories.NewClassRepository(db)
	groupRepository := repositories.NewGroupRepository(db)
	userRepository := repositories.NewUserRepository(db)

	// Init mappers
	classMapper := mappers.NewClassMapper()
	userMapper := mappers.NewUserMapper()

	// Init services
	classService := services.NewClassService(classRepository, classMapper)
	groupService := services.NewGroupService(groupRepository)
	userService := services.NewUserService(userRepository, userMapper)

	// Init Kafka
	brokers := []string{fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)}
	consumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
	producer, err := kafka.NewProducer([]string{fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	// Init Kafka listeners
	groupListener := kafkaListeners.NewGroupListener(consumer, groupService)
	groupListener.Listen(cfg.GroupsTopic)

	// Init controllers
	router := mux.NewRouter()
	router.Use(utils.Recovery)
	classController := controllers.NewClassController(classService, userService)
	groupController := controllers.NewGroupController(producer)
	classController.SetupRoutes(router)
	groupController.SetupRoutes(router)

	// Starting app
	log.Printf("Running on http://%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), router))
}
