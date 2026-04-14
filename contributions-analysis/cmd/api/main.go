package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wnn-dev/contributions-analysis/config"
	"github.com/wnn-dev/contributions-analysis/database/postgres"
	analysisHandler "github.com/wnn-dev/contributions-analysis/handlers/analysis"
	contributionHandler "github.com/wnn-dev/contributions-analysis/handlers/contribution"
	contributorHandler "github.com/wnn-dev/contributions-analysis/handlers/contributor"
	htmlcssHandler "github.com/wnn-dev/contributions-analysis/handlers/htmlcss"
	jsonServer "github.com/wnn-dev/contributions-analysis/server/json"
)

func main() {
	log.Println("Starting the Contributions Analysis Service!")

	var envFile string
	flag.StringVar(&envFile, "env", "", "Environment Variable File")
	flag.Parse()

	if envFile == "" {
		log.Fatal("no ENV file provided")
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("error loading env file: ", err)
	}

	initContext := &gin.Context{}

	configuration, err := config.LoadEnvFile()
	if err != nil {
		log.Println("error loading configuration: ", err)
	}

	log.Println("Configuration Variables:", configuration)

	log.Println("Connecting to postgres")
	database := postgres.New(initContext, *configuration)
	defer database.Close()

	// Services
	contributorService := postgres.NewContributorService(initContext, database)
	contributionService := postgres.NewContributionService(initContext, database)
	analysisResultService := postgres.NewAnalysisResultService(initContext, database)
	htmlCssSubmissionService := postgres.NewHtmlCssSubmissionService(initContext, database)

	// Handlers
	contribHandler := contributorHandler.NewHandler(contributorService, contributionService)
	contribtnHandler := contributionHandler.NewHandler(contributionService)
	analysisHandler := analysisHandler.NewHandler(analysisResultService)
	htmlCssHandler := htmlcssHandler.NewHandler(htmlCssSubmissionService, analysisResultService)

	// JSON Servers
	contributorJsonServer := jsonServer.NewContributorServer(contribHandler)
	contributionJsonServer := jsonServer.NewContributionServer(contribtnHandler)
	analysisJsonServer := jsonServer.NewAnalysisServer(analysisHandler)
	htmlCssJsonServer := jsonServer.NewHtmlCssServer(htmlCssHandler)

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Authorization", "Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"}
	corsConfig.AllowCredentials = true

	router := gin.New()
	router.Use(cors.New(corsConfig))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setUpRoutes(router, contributorJsonServer, contributionJsonServer, analysisJsonServer, htmlCssJsonServer)

	srv := &http.Server{
		Addr:    configuration.Address,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen: %s\n", err)
		}
		log.Println("Contributions Analysis service is listening on address", srv.Addr)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panicf("server shutdown: err=%v", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 3 seconds.")
	}
	log.Println("Server exiting")
}

func setUpRoutes(
	router *gin.Engine,
	contributorJsonServer *jsonServer.ContributorJsonServer,
	contributionJsonServer *jsonServer.ContributionJsonServer,
	analysisJsonServer *jsonServer.AnalysisJsonServer,
	htmlCssJsonServer *jsonServer.HtmlCssJsonServer,
) {
	api := router.Group("/contributions-analysis/api/v1")

	api.Use()
	{
		// Contributor routes
		api.POST("/contributor/register", contributorJsonServer.SignUp())
		api.PUT("/contributor/login", contributorJsonServer.Login())
		api.GET("/contributor", contributorJsonServer.GetContributor())
		api.GET("/contributors", contributorJsonServer.GetContributors())

		// Contribution routes
		api.POST("/contribution/create", contributionJsonServer.CreateContribution())
		api.GET("/contribution", contributionJsonServer.GetContribution())
		api.GET("/contributions", contributionJsonServer.GetContributions())

		// Analysis routes
		api.POST("/analysis/submit", analysisJsonServer.Submit())
		api.GET("/analysis", analysisJsonServer.GetAnalysisResult())
		api.GET("/analyses", analysisJsonServer.GetAnalysisResults())
		api.GET("/analysis/by-contributor", analysisJsonServer.GetByContributor())
		api.GET("/analysis/by-contribution", analysisJsonServer.GetByContribution())
		api.PATCH("/analysis/status", analysisJsonServer.UpdateStatus())

		// HTML/CSS test routes
		api.POST("/test/html-css/submit", htmlCssJsonServer.Submit())
		api.GET("/test/html-css/submissions", htmlCssJsonServer.GetSubmissionsByContributor())
	}
}
