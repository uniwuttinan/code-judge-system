package services

import (
	"github.com/spf13/viper"
	"github.com/wuttinanhi/code-judge-system/entities"
	"github.com/wuttinanhi/code-judge-system/repositories"
	"gorm.io/gorm"
)

type ServiceKit struct {
	JWTService        JWTService
	UserService       UserService
	ChallengeService  ChallengeService
	SubmissionService SubmissionService
	SandboxService    SandboxService
	KafkaService      KafkaService
}

func CreateServiceKit(db *gorm.DB) *ServiceKit {
	userRepo := repositories.NewUserRepository(db)
	challengeRepo := repositories.NewChallengeRepository(db)
	submissionRepo := repositories.NewSubmissionRepository(db)

	// read env var "JWT_SECRET" and pass it to JWTService
	// if JWT_SECRET is empty, use default value
	jwtSecret := viper.GetString("AUTH_JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "secret"
	}

	kafkaHost := viper.GetString("KAFKA_HOST")

	maxMemoryLimit := viper.GetUint("SANDBOX_MAX_MEMORY_MB")
	maxRuntimeMs := viper.GetUint("SANDBOX_MAX_TIME_MS")

	jwtService := NewJWTService(jwtSecret)
	userService := NewUserService(userRepo)
	sandboxService := NewSandboxService(maxMemoryLimit, maxRuntimeMs)
	challengeService := NewChallengeService(challengeRepo, sandboxService)
	submissionService := NewSubmissionService(submissionRepo, challengeService, sandboxService)
	kafkaService := NewKafkaService(kafkaHost)

	return &ServiceKit{
		JWTService:        jwtService,
		UserService:       userService,
		ChallengeService:  challengeService,
		SubmissionService: submissionService,
		SandboxService:    sandboxService,
		KafkaService:      kafkaService,
	}
}

func CreateTestServiceKit(db *gorm.DB) *ServiceKit {
	userRepo := repositories.NewUserRepository(db)
	challengeRepo := repositories.NewChallengeRepository(db)
	submissionRepo := repositories.NewSubmissionRepository(db)

	maxMemoryLimit := entities.SandboxMemoryMB * 256
	maxRuntimeMs := uint(10000)

	jwtService := NewJWTService("test")
	userService := NewUserService(userRepo)
	sandboxService := NewSandboxService(maxMemoryLimit, maxRuntimeMs)
	challengeService := NewChallengeService(challengeRepo, sandboxService)
	submissionService := NewSubmissionService(submissionRepo, challengeService, sandboxService)
	kafkaService := NewKafkaMockService()

	return &ServiceKit{
		JWTService:        jwtService,
		UserService:       userService,
		ChallengeService:  challengeService,
		SubmissionService: submissionService,
		SandboxService:    sandboxService,
		KafkaService:      kafkaService,
	}
}
