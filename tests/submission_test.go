package tests_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/wuttinanhi/code-judge-system/controllers"
	"github.com/wuttinanhi/code-judge-system/databases"
	"github.com/wuttinanhi/code-judge-system/entities"
	"github.com/wuttinanhi/code-judge-system/services"
	"github.com/wuttinanhi/code-judge-system/tests"
)

func TestSubmissionRoute(t *testing.T) {
	db := databases.NewTempSQLiteDatabase()
	testServiceKit := services.CreateTestServiceKit(db)
	app := controllers.SetupAPI(testServiceKit)

	// create user
	user, err := testServiceKit.UserService.Register("test-submission-route@example.com", "testpassword", "test-submission-route")
	if err != nil {
		t.Error(err)
	}

	// get user access token1
	userAccessToken, err := testServiceKit.JWTService.GenerateToken(*user)
	if err != nil {
		t.Error(err)
	}

	SUBMISSION_LANGUAGE := "go"
	SUBMISSION_SOURCE_CODE := "test source code"

	// create challenge
	challenge, err := testServiceKit.ChallengeService.CreateChallenge(&entities.Challenge{
		Name:        "Test Challenge",
		Description: "Test Description",
		Testcases: []*entities.ChallengeTestcase{
			{Input: "1", ExpectedOutput: "1", LimitMemory: 1, LimitTimeMs: 1},
			{Input: "2", ExpectedOutput: "2", LimitMemory: 2, LimitTimeMs: 2},
			{Input: "3", ExpectedOutput: "3", LimitMemory: 3, LimitTimeMs: 3},
		},
	})
	if err != nil {
		t.Error(err)
	}

	t.Run("/submission/submit", func(t *testing.T) {
		dto := entities.SubmissionCreateDTO{
			ChallengeID: challenge.ID,
			Language:    SUBMISSION_LANGUAGE,
			SourceCode:  SUBMISSION_SOURCE_CODE,
		}
		requestBody, _ := json.Marshal(dto)

		request, _ := http.NewRequest(http.MethodPost, "/submission/submit", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+userAccessToken)

		response, err := app.Test(request, -1)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %v", response.StatusCode)
		}

		// get submission in server-side
		submission, err := testServiceKit.SubmissionService.GetSubmissionByID(1)
		if err != nil {
			t.Error(err)
		}
		if submission.ChallengeID != challenge.ID {
			t.Errorf("Expected challenge id %v, got %v", challenge.ID, submission.ChallengeID)
		}
		if submission.UserID != user.ID {
			t.Errorf("Expected user id %v, got %v", user.ID, submission.UserID)
		}
		if submission.Language != dto.Language {
			t.Errorf("Expected language %v, got %v", dto.Language, submission.Language)
		}
		if submission.SourceCode != dto.SourceCode {
			t.Errorf("Expected source code %v, got %v", dto.SourceCode, submission.SourceCode)
		}

		// validate submission testcases
		submissionTestcases, err := testServiceKit.SubmissionService.GetSubmissionTestcaseBySubmission(submission)
		if err != nil {
			t.Error(err)
		}
		if len(submissionTestcases) != len(challenge.Testcases) {
			t.Errorf("Expected %v submission testcases, got %v", len(challenge.Testcases), len(submissionTestcases))
		}
		for i := range challenge.Testcases {
			submissionTestcase := submissionTestcases[i]
			if submissionTestcase.ID != uint(i+1) {
				t.Errorf("Expected challenge testcase id %v, got %v", uint(i+1), submissionTestcase.ID)
			}
			if submissionTestcase.Status != entities.SubmissionStatusPending {
				t.Errorf("Expected status %v, got %v", entities.SubmissionStatusPending, submissionTestcase.Status)
			}
			if submissionTestcase.Output != "" {
				t.Errorf("Expected output %v, got %v", "", submissionTestcase.Output)
			}
		}
	})

	t.Run("/submission/get/submission/:id", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/submission/get/submission/1", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+userAccessToken)

		response, err := app.Test(request, -1)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %v", response.StatusCode)
		}

		bodyBytes := tests.ResponseBodyToBytes(response)

		var submission entities.Submission
		err = json.Unmarshal(bodyBytes, &submission)
		if err != nil {
			t.Error(err)
		}

		if submission.ChallengeID != challenge.ID {
			t.Errorf("Expected challenge id %v, got %v", challenge.ID, submission.ChallengeID)
		}
		if submission.UserID != user.ID {
			t.Errorf("Expected user id %v, got %v", user.ID, submission.UserID)
		}
		if submission.Language != "go" {
			t.Errorf("Expected language %v, got %v", "go", submission.Language)
		}
		if submission.SourceCode != SUBMISSION_SOURCE_CODE {
			t.Errorf("Expected source code %v, got %v", SUBMISSION_SOURCE_CODE, submission.SourceCode)
		}
	})

	t.Run("/submission/get/user", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/submission/get/user", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+userAccessToken)

		response, err := app.Test(request, -1)
		if err != nil {
			t.Error(err)
		}

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %v", response.StatusCode)
		}

		var submissions []entities.Submission
		err = json.Unmarshal(bodyBytes, &submissions)
		if err != nil {
			t.Error(err)
		}

		if len(submissions) != 1 {
			t.Errorf("Expected %v submissions, got %v", 1, len(submissions))
		}

		submission := submissions[0]
		if submission.ChallengeID != challenge.ID {
			t.Errorf("Expected challenge id %v, got %v", challenge.ID, submission.ChallengeID)
		}
		if submission.UserID != user.ID {
			t.Errorf("Expected user id %v, got %v", user.ID, submission.UserID)
		}
		if submission.Language != "go" {
			t.Errorf("Expected language %v, got %v", "go", submission.Language)
		}
		if submission.SourceCode != SUBMISSION_SOURCE_CODE {
			t.Errorf("Expected source code %v, got %v", SUBMISSION_SOURCE_CODE, submission.SourceCode)
		}
	})

	t.Run("/submission/get/challenge/:id", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/submission/get/challenge/"+strconv.Itoa(int(challenge.ID)), nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+userAccessToken)

		response, err := app.Test(request, -1)
		if err != nil {
			t.Error(err)
		}

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %v", response.StatusCode)
		}

		var submissions []entities.Submission
		err = json.Unmarshal(bodyBytes, &submissions)
		if err != nil {
			t.Error(err)
		}

		if len(submissions) != 1 {
			t.Errorf("Expected %v submissions, got %v", 1, len(submissions))
		}

		submission := submissions[0]
		if submission.ChallengeID != challenge.ID {
			t.Errorf("Expected challenge id %v, got %v", challenge.ID, submission.ChallengeID)
		}
		if submission.UserID != user.ID {
			t.Errorf("Expected user id %v, got %v", user.ID, submission.UserID)
		}
		if submission.Language != "go" {
			t.Errorf("Expected language %v, got %v", "go", submission.Language)
		}
		if submission.SourceCode != SUBMISSION_SOURCE_CODE {
			t.Errorf("Expected source code %v, got %v", SUBMISSION_SOURCE_CODE, submission.SourceCode)
		}
	})
}
