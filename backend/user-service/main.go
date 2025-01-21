package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"infrastructure/initInfra"
	"log"
	initService "user-service/init-service"
)

const AppId = "user-service"

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

	//// Открывает стартовую страницу в браузере (работает только с Windows)
	//if err := exec.Command("explorer", "http://localhost:8081").Run(); err != nil {
	//	dc.Logger.Error(context.Background(), err)
	//}

}
