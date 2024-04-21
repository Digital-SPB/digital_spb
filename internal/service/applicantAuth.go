package service

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

type ApplicantAuthService struct {
	repo     repo.ApplicantAuth
	examRepo repo.Exam
}

func NewApplicantAuthService(repo repo.ApplicantAuth, examRepo repo.Exam) *ApplicantAuthService {
	return &ApplicantAuthService{
		repo:     repo,
		examRepo: examRepo,
	}
}

func (s *ApplicantAuthService) CreateApplicant(ctx context.Context, input entity.Applicant) (int, error) {
	input.Password = s.generatePasswordHash(input.Password)
	appId, err := s.repo.CreateApplicant(ctx, input)
	if err != nil {
		return 0, err
	}
	fmt.Println("0")

	for i := 0; i < len(input.Exams); i++ {
		fmt.Println("1")
		examInput := input.Exams[i]
		fmt.Println("2")
		if examInput.ExamMark != 0 {
			examInput.ApplicantId = appId
			fmt.Println("3")
			fmt.Println(examInput)
			s.examRepo.AddExam(ctx, examInput)

			fmt.Println("4")
		}
	}
	//s.examRepo.AddExam(ctx, input.Exams[0])

	fmt.Println("5")
	return appId, nil
}

func (a *ApplicantAuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
