package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"schedule/internal/config"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/mappers"
	"schedule/internal/domain/services"
	"schedule/internal/infrastructure/repositories"
	"schedule/internal/kafka"
	"schedule/internal/presentation/api/routes"
	"schedule/internal/presentation/utils"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println(cfg)

	db, err := sqlx.Connect("postgres", "dbname=schedule sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	var repository = repositories.NewGroupRepository(db)
	var service = services.NewGroupService(repository)
	var er = service.RemoveClassFromGroup(&dto.RemoveClassFromGroup{ClassId: 5, GroupId: 1})
	if er != nil {
		log.Fatal(er)
	}

	brokers := []string{"localhost:9092"}
	consumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	err = consumer.Consume("classes", func(message *sarama.ConsumerMessage) {
		log.Printf("Received message: %s", string(message.Value))
	})
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	router := mux.NewRouter()
	classRepository := repositories.NewClassRepository(db)
	classMapper := mappers.NewClassMapper()
	producer, err := kafka.NewProducer(brokers)
	classService := services.NewClassService(classRepository, classMapper, producer)
	routes.SetupRoutes(router, classService)
	router.Use(utils.Recovery)
}
