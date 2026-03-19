package main

import (
	"context"
	"fmt"
	"log"
	"os"

	//internals
	"church-backend/internal/auth"
	"church-backend/internal/database"
	"church-backend/internal/repository"
	"church-backend/internal/service"

	"github.com/its-ernest/echox/store"

	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	// member routes
	"church-backend/internal/routes/members"

	// ministry routes
	"church-backend/internal/routes/ministry"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	e := echo.New()

	ctx := context.Background()
	dbURL := os.Getenv("DB_URL")
	dbPool, err := database.NewPostgresPool(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	jwtSecretStr := os.Getenv("JWT_SECRET")
	if jwtSecretStr == "" {
		log.Println("\033[33m[WARN]\033[0m JWT_SECRET not set, using insecure default")
		jwtSecretStr = "church-default-secret-2026"
	}
	jwtSecret := []byte(jwtSecretStr)

	// echo v5 logger middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogMethod:  true,
		LogURI:     true,
		LogLatency: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			statusColor := "\033[32m" // Green
			if v.Status >= 400 {
				statusColor = "\033[33m"
			}
			if v.Status >= 500 {
				statusColor = "\033[31m"
			}
			reset := "\033[0m"

			log.Printf("%s %s %s| %s%d%s | %10v | %s",
				"\033[34m[API]\033[0m", v.Method, v.URI,
				statusColor, v.Status, reset, v.Latency, c.Request().RemoteAddr)
			return nil
		},
	}))

	// initialize the MemoryStore from my echox
	memStore := store.NewMemoryStore()

	//initializze model repos
	memberRepo := &repository.MemberRepository{Pool: dbPool}
	ministryRepo := &repository.MinistryRepository{Pool: dbPool}

	authService := service.NewAuthService(memStore, memberRepo)
	authHandler := auth.NewHandler(authService, jwtSecret)

	// auth routes groups
	authGroup := e.Group("/auth")
	authHandler.Register(authGroup, memStore)

	// services declarations
	memberService := members.NewService(memberRepo)
	ministryService := ministry.NewService(ministryRepo)

	// handlers declarations
	memberHandler := members.NewHandler(memberService)
	ministryHandler := ministry.NewHandler(ministryService)

	// members route
	memberGroup := e.Group("/members")
	memberGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: jwtSecret,
		// ghost bug fixed: auth and members handlers now matches this claims
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
	}))
	members.RegisterRoutes(memberGroup, memberHandler)

	// ministries route
	ministryGroup := e.Group("/ministries")
	ministryGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: jwtSecret,
		// ghost bug fixed: auth and members handlers now matches this claims
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
	}))
	ministry.RegisterRoutes(ministryGroup, ministryHandler)

	// start server
	log.Println("Church Backend starting on :8080")
	for _, route := range e.Router().Routes() {
		fmt.Printf("\033[36m[ROUTE]\033[0m %-6s %-15s -> %s\n", route.Method, route.Path, route.Name)
	}

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
