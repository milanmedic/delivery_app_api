package main

import (
	"fmt"

	"delivery_app_api.mmedic.com/m/v2/src/controllers"
	sql_driver "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
	addr_sql_db "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/addr_sql_db"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/admin_sql_db"
	article_sql_db "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/article_db_impls"
	customer_sql_db "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/customer_sql_db"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/deliverer_sql_db"
	addr_repository "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/addr_repository"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/admin_repository"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/article_repository"
	customer_repo "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/customer_repository"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/deliverer_repository"
	routes "delivery_app_api.mmedic.com/m/v2/src/routes"
	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	admin_service "delivery_app_api.mmedic.com/m/v2/src/services/admin_service"
	"delivery_app_api.mmedic.com/m/v2/src/services/article_service"
	customer_service "delivery_app_api.mmedic.com/m/v2/src/services/customer_service"
	"delivery_app_api.mmedic.com/m/v2/src/services/deliverer_service"
	"delivery_app_api.mmedic.com/m/v2/src/utils/env_utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//TODO: Refactor ROUTE paths
func main() {
	//**************************************************************************
	// DATABASE CREATION
	db, err := sql_driver.CreateDeliveryAppDb()
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
	// GENERAL ROUTES
	routes.SetupGeneralRoutes(router)

	//**************************************************************************
	// ARTICLE SERVICE
	ardb := article_sql_db.CreateArticleDb(db)
	arr := article_repository.CreateArticleRepository(ardb)
	ars := article_service.CreateArticleService(arr)
	arc := controllers.CreateArticleController(ars)
	routes.SetupArticleRoutes(router, arc)

	//**************************************************************************
	// USER SERVICES, CONTROLLERS & ROUTES SETUP
	adb := addr_sql_db.CreateAddrDb(db)
	ar := addr_repository.CreateAddrRepository(adb)
	as := addr_service.CreateAddrService(ar)

	cdb := customer_sql_db.CreateCustomerDb(db)
	cr := customer_repo.CreateCustomerRepository(cdb)
	cs := customer_service.CreateCustomerService(cr)
	cc := controllers.CreateCustomerController(cs, as)
	routes.SetupCustomerRoutes(router, cc)

	//**************************************************************************
	// DELIVERER ROUTES
	ddb := deliverer_sql_db.CreateDelivererDb(db)
	dr := deliverer_repository.CreateDelivererRepository(ddb)
	ds := deliverer_service.CreateDelivererService(dr)
	dc := controllers.CreateDelivererController(ds, as)
	routes.SetupDelivererRoutes(router, dc)

	//**************************************************************************
	// ADMIN ROUTES
	admdb := admin_sql_db.CreateAdminDb(db)
	admr := admin_repository.CreateAdminRepository(admdb)
	ads := admin_service.CreateAdminService(admr)
	adc := controllers.CreateAdminController(ads, cs, as, ds)
	routes.SetupAdminRoutes(router, adc)
	//**************************************************************************

	// RUN SERVER
	PORT := env_utils.GetEnvVar("PORT")
	err = router.Run(fmt.Sprintf(":%s", PORT))
	HandleError(err)

}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
	}
}
