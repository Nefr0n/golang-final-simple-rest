package app

import (
	"fmt"
	"net"
	"os"

	"github.com/Skapar/simple-rest/config"
	"github.com/Skapar/simple-rest/internal/repository"
	"github.com/Skapar/simple-rest/internal/routes"
	"github.com/Skapar/simple-rest/internal/server"
	"github.com/Skapar/simple-rest/internal/service"
	"github.com/Skapar/simple-rest/pkg"
	profile "github.com/Skapar/simple-rest/proto"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type App struct {
	Log    *zap.SugaredLogger
	DB     *gorm.DB
	Cfg    *config.Config
	Router *gin.Engine
}

func NewApp() *App {
	godotenv.Load()
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "trace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}
	zlogger := zap.New(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel),
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	log := zlogger.Sugar()
	cfg := config.New()
	cfg.Init()

	dbService := &pkg.GormDatabase{}
	db, err := dbService.Connect(cfg, log)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/health"}}))
	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())

	authRepo := repository.NewAuthRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	authService := service.NewAuthService(authRepo, cfg)
	userService := service.NewUserService(authRepo)
	profileService := service.NewProfileService(profileRepo)

	routes.SetupRoutes(router, authService, userService)

	// GRPC
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.ListenGRPCPort)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to listen: %s", err))
	}

	s := grpc.NewServer()

	grpcServer := server.New(profileService)
	profile.RegisterProfileServiceServer(s, grpcServer)

	go func() {
		if err := s.Serve(l); err != nil {
			panic(fmt.Sprintf("failed to serve: %s", err))
		}
	}()

	log.Infof("Listening on port grpc: %v", addr)

	return &App{
		Log:    log,
		DB:     db,
		Cfg:    cfg,
		Router: router,
	}
}
