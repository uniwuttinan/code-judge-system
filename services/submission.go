package services

import (
	"github.com/wuttinanhi/code-judge-system/entities"
	"github.com/wuttinanhi/code-judge-system/repositories"
)

type SubmissionService interface {
	CreateSubmission(submission *entities.Submission) (*entities.Submission, error)
	DeleteSubmission(submission *entities.Submission) error
	GetSubmissionByID(submissionID uint) (*entities.Submission, error)
	GetSubmissionByUser(user *entities.User) ([]entities.Submission, error)
	GetSubmissionByChallenge(challenge *entities.Challenge) ([]entities.Submission, error)
	CreateSubmissionTestcase(submissionTestcase *entities.SubmissionTestcase) (*entities.SubmissionTestcase, error)
	GetSubmissionTestcaseBySubmission(submission *entities.Submission) ([]entities.SubmissionTestcase, error)
}

type submissionService struct {
	submissionRepository repositories.SubmissionRepository
}

// CreateSubmission implements SubmissionService.
func (s *submissionService) CreateSubmission(submission *entities.Submission) (*entities.Submission, error) {
	submission, err := s.submissionRepository.CreateSubmission(submission)
	return submission, err
}

// CreateSubmissionTestcase implements SubmissionService.
func (s *submissionService) CreateSubmissionTestcase(submissionTestcase *entities.SubmissionTestcase) (*entities.SubmissionTestcase, error) {
	submissionTestcase, err := s.submissionRepository.CreateSubmissionTestcase(submissionTestcase)
	return submissionTestcase, err
}

// DeleteSubmission implements SubmissionService.
func (s *submissionService) DeleteSubmission(submission *entities.Submission) error {
	err := s.submissionRepository.DeleteSubmission(submission)
	return err
}

// GetSubmissionByChallenge implements SubmissionService.
// GetSubmissionByChallenge implements SubmissionService.
func (s *submissionService) GetSubmissionByChallenge(challenge *entities.Challenge) ([]entities.Submission, error) {
	submissions, err := s.submissionRepository.GetSubmissionByChallenge(challenge)
	return submissions, err
}

// GetSubmissionByID implements SubmissionService.
func (s *submissionService) GetSubmissionByID(submissionID uint) (*entities.Submission, error) {
	submission, err := s.submissionRepository.GetSubmissionByID(submissionID)
	return submission, err
}

// GetSubmissionByUser implements SubmissionService.
func (s *submissionService) GetSubmissionByUser(user *entities.User) ([]entities.Submission, error) {
	submissions, err := s.submissionRepository.GetSubmissionByUser(user)
	return submissions, err
}

// GetSubmissionTestcaseBySubmission implements SubmissionService.
func (s *submissionService) GetSubmissionTestcaseBySubmission(submission *entities.Submission) ([]entities.SubmissionTestcase, error) {
	submissionTestcases, err := s.submissionRepository.GetSubmissionTestcaseBySubmission(submission)
	return submissionTestcases, err
}

func NewSubmissionService(submissionRepository repositories.SubmissionRepository) SubmissionService {
	return &submissionService{
		submissionRepository: submissionRepository,
	}
}
