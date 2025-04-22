package app

import (
	"os"

	"github.com/Skapar/simple-rest/config"
	"github.com/Skapar/simple-rest/internal/repository"
	"github.com/Skapar/simple-rest/internal/routes"
	"github.com/Skapar/simple-rest/internal/service"
	"github.com/Skapar/simple-rest/pkg"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	authService := service.NewAuthService(authRepo, cfg)
	userService := service.NewUserService(authRepo)

	routes.SetupRoutes(router, authService, userService)

	return &App{
		Log:    log,
		DB:     db,
		Cfg:    cfg,
		Router: router,
	}
}
