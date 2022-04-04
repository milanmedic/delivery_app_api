package main

import (
	"fmt"
	"os"

	customer_controller "delivery_app_api.mmedic.com/m/v2/src/controllers"
	deliveryAppDb "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
	addrsqldb "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/addr_sql_db"
	customerqldb "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/customer_sql_db"
	addr_repository "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/addr_repository"
	customer_repo "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/customer_repository"
	customer_route "delivery_app_api.mmedic.com/m/v2/src/routes"
	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	customer_service "delivery_app_api.mmedic.com/m/v2/src/services/customer_service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//**************************************************************************
	// DATABASE CREATION
	db, err := deliveryAppDb.CreateDeliveryAppDb()
	HandleError(err)
	err = db.CheckConnection()
	HandleError(err)
	defer db.CloseConnection()

	//**************************************************************************
	// LOAD ENV VARIABLES
	err = godotenv.Load()

	//**************************************************************************
	// SERVER SETUP & GLOBAL MIDDLEWARE SETUP
	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	//**************************************************************************
	// USER SERVICES, CONTROLLERS & ROUTES SETUP
	adb := addrsqldb.CreateAddrDb(db)
	ar := addr_repository.CreateAddrRepository(adb)
	as := addr_service.CreateAddrService(ar)

	cdb := customerqldb.CreateCustomerDb(db)
	cr := customer_repo.CreateCustomerRepository(cdb)
	cs := customer_service.CreateCustomerService(cr)
	cc := customer_controller.CreateCustomerController(cs, as)
	customer_route.SetupCustomerRoutes(router, cc)

	//**************************************************************************
	// RUN SERVER
	PORT := GetEnvVar("PORT")
	err = router.Run(fmt.Sprintf(":%s", PORT))
	HandleError(err)

}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
	}
}
