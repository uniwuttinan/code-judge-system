package cmds

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wuttinanhi/code-judge-system/controllers"
	"github.com/wuttinanhi/code-judge-system/services"
)

func SetupWeb(serviceKit *services.ServiceKit) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: controllers.ErrorHandler,
	})

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	authHandler := controllers.NewAuthHandler(serviceKit)
	challengeHandler := controllers.NewChallengeHandler(serviceKit)
	challengeTestcaseHandler := controllers.NewChallengeTestcaseHandler(serviceKit)
	submissionHandler := controllers.NewSubmissionHandler(serviceKit)

	userGroup := app.Group("/user")
	userGroup.Post("/register", authHandler.Register)
	userGroup.Post("/login", authHandler.Login)

	challengeGroup := app.Group("/challenge")
	challengeGroup.Use(controllers.UserMiddleware(serviceKit))
	challengeGroup.Post("/create", challengeHandler.CreateChallengeWithTestcase)
	challengeGroup.Get("/all", challengeHandler.GetAllChallenges)
	challengeGroup.Get("/get/:id", challengeHandler.GetChallengeByID)
	challengeGroup.Put("/update", challengeHandler.UpdateChallenge)
	challengeGroup.Delete("/delete/:id", challengeHandler.DeleteChallenge)

	testcaseGroup := app.Group("/testcase")
	testcaseGroup.Use(controllers.UserMiddleware(serviceKit))
	testcaseGroup.Post("/create", challengeTestcaseHandler.CreateTestcase)
	testcaseGroup.Get("/get/:id", challengeTestcaseHandler.GetTestcaseByID)
	testcaseGroup.Put("/update", challengeTestcaseHandler.UpdateTestcase)
	testcaseGroup.Delete("/delete/:id", challengeTestcaseHandler.DeleteTestcase)

	submissionGroup := app.Group("/submission")
	submissionGroup.Use(controllers.UserMiddleware(serviceKit))
	submissionGroup.Post("/submit", submissionHandler.SubmitSubmission)
	submissionGroup.Get("/get/user", submissionHandler.GetSubmissionByUser)
	submissionGroup.Get("/get/challenge/:id", submissionHandler.GetSubmissionByChallenge)
	submissionGroup.Get("/get/submission/:id", submissionHandler.GetSubmissionByID)

	return app
}
