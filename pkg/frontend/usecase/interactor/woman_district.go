package interactor

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/input"
	"golang-clean-architecture/pkg/frontend/usecase/inputport"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"
)

type womanDistrictUsecase struct {
	womanDistrictRepository outputport.WomanDistrictRepository
}

func NewWomanDistrictUsecase(womanDistrictRepository outputport.WomanDistrictRepository) inputport.WomanDistrictUsecase {
	return &womanDistrictUsecase{womanDistrictRepository}
}

func (u *womanDistrictUsecase) GetList(ctx context.Context, i input.GetWomanDistrictListInput) (collection.Collection[entity.WomanEntity], error) {
	return u.womanDistrictRepository.FindAllByDistrict(ctx, i.DistrictID)
}
