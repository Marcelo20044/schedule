package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os/exec"
	"schedule/internal/config"
	"schedule/internal/domain/kafkaListeners"
	"schedule/internal/domain/mappers"
	"schedule/internal/domain/services"
	"schedule/internal/infrastructure/repositories"
	"schedule/internal/presentation/controllers"
	"schedule/internal/presentation/utils"
	"schedule/kafka"
	"time"
)

func Init() {
	cfg := config.GetConfig()

	// Start Docker containers
	cmd := exec.Command("docker-compose", "-f", "./docker-compose.yml", "up", "-d")
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}

	// Wait for Kafka and init it
	brokers := []string{fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)}
	consumer, producer, err := initKafka(brokers, 2*time.Minute)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka: %v", err)
	}

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

func initKafka(brokers []string, timeout time.Duration) (*kafka.Consumer, *kafka.Producer, error) {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		consumer, err := kafka.NewConsumer(brokers)
		if err == nil {
			producer, err := kafka.NewProducer(brokers)
			if err == nil {
				return consumer, producer, nil
			}
		}

		log.Printf("Waiting for Kafka to be ready...")
		time.Sleep(5 * time.Second)
	}

	return nil, nil, fmt.Errorf("timed out waiting for Kafka to be ready")
}
