package service

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"

	log "github.com/sirupsen/logrus"
)

type EducationalDirectionService struct {
	educationalDirectionRepo repo.EducationalDirection
	applicantRepo            repo.Applicant
	vacancyRepo              repo.Vacancy
}

func NewEducationalDirectionService(educationalDirectionRepo repo.EducationalDirection, applicantRepo repo.Applicant, vacancyRepo repo.Vacancy) *EducationalDirectionService {
	return &EducationalDirectionService{
		educationalDirectionRepo: educationalDirectionRepo,
		applicantRepo:            applicantRepo,
		vacancyRepo:              vacancyRepo,
	}
}

func (s *EducationalDirectionService) CreateEducationalDirection(ctx context.Context, education entity.EducatitionalDirection) (int, error) {
	return s.educationalDirectionRepo.CreateEducationalDirection(ctx, education)
}

func (s *EducationalDirectionService) CountEducationalDirection(ctx context.Context) (int, error) {
	return s.educationalDirectionRepo.CountEducationalDirection(ctx)
}

func (s *EducationalDirectionService) GetEducationalDirectionForApplicant(ctx context.Context, applicantId int) ([]entity.EducatitionalDirection, error) {
	applicant, err := s.applicantRepo.GetApplicant(ctx, applicantId)
	if err != nil {
		log.Error("Error while getting applicant", err)
		return nil, err
	}

	exams, err := s.applicantRepo.GetExam(ctx, applicant.Id)
	if err != nil {
		log.Error("Error while getting exams", err)
		return nil, err
	}

	profession := applicant.Profession
	sumMarks := 0
	for _, exam := range exams {
		sumMarks += exam.ExamMark
	}

	educations, err := s.educationalDirectionRepo.GetEducationalDirections(ctx)
	if err != nil {
		log.Error("Error while getting educational directions", err)
		return []entity.EducatitionalDirection{}, err
	}

	recommendEducations := make([]entity.EducatitionalDirection, 0)
	for _, education := range educations {
		// баллы прошлых лет
		lastMark := education.Sum
		if lastMark > sumMarks {
			continue
		}

		// баллы
		isOkMarks := true
		for _, exam := range exams {
			if exam.ExamName == education.Subject1 && exam.ExamMark < education.Value1 {
				isOkMarks = false
				break
			}
			if exam.ExamName == education.Subject2 && exam.ExamMark < education.Value2 {
				isOkMarks = false
				break
			}
			if exam.ExamName == education.Subject3 && exam.ExamMark < education.Value3 {
				isOkMarks = false
				break
			}
		}

		// профессия
		isOkProfession := false
		vacancies, err := s.vacancyRepo.GetVacanciesByEducationId(ctx, education.Id)
		if err != nil {
			log.Error("Error while getting vacancies", err)
			return []entity.EducatitionalDirection{}, err
		}

		for _, vacancy := range vacancies {
			// профессия подходит

			if vacancy.Name == profession {
				isOkProfession = true
				break
			}
		}

		if isOkMarks && isOkProfession {
			recommendEducations = append(recommendEducations, education)
		}
	}

	return recommendEducations, nil
}
