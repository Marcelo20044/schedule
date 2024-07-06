package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"schedule/internal/config"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/services"
	"schedule/internal/infrastructure/repositories"
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
}
