package women

import "golang-clean-architecture/pkg/frontend/domain/entity"

type DistrictListResponse struct {
	Women []WomanListItem `json:"women"`
}

func NewDistrictListResponse(women []entity.WomanEntity) DistrictListResponse {
	items := make([]WomanListItem, 0, len(women))
	for _, w := range women {
		items = append(items, toWomanListItem(w))
	}
	return DistrictListResponse{Women: items}
}
