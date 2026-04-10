package interactor

import (
	"context"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/input"
	"golang-clean-architecture/pkg/backend/usecase/inputport"
	backendoutputport "golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/outputport"
)

type CompanyUsecase struct {
	companyRepository backendoutputport.CompanyRepository
	uow               outputport.UnitOfWork
}

func NewCompanyUsecase(
	companyRepository backendoutputport.CompanyRepository,
	uow outputport.UnitOfWork,
) inputport.CompanyUsecase {
	return &CompanyUsecase{companyRepository, uow}
}

func (u *CompanyUsecase) Create(ctx context.Context, i input.CreateCompanyInput) error {
	company := &entity.Company{
		Name:     i.Name,
		Rank:     i.Rank,
		IsActive: i.IsActive,
	}
	return u.uow.Do(ctx, func() error {
		return u.companyRepository.Create(ctx, company)
	})
}
