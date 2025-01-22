package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"infrastructure/initInfra"
	"log"
	initService "task-service/init-services"
)

const AppId = "task-service"

func main() {

	// TODO переписать с либы на os.Env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	dc := &initService.DependencyContainer{}

	err := initInfra.NewDependenciesInitializer().
		AppId(AppId).
		InfraInitChain(initService.GetInfraInitChains()).
		ServicesInitChain(initService.GetServicesInitChains(dc)).
		InfraSetter(dc).
		Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	<-dc.StopChan()
}
