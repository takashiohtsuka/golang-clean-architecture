package women

import "golang-clean-architecture/pkg/frontend/domain/entity"

type AreaItem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type WomanDistrictListItem struct {
	ID         uint         `json:"id"`
	District   AreaItem     `json:"district"`
	Prefecture AreaItem     `json:"prefecture"`
	Region     AreaItem     `json:"region"`
	Name       string       `json:"name"`
	Age        *int         `json:"age"`
	Birthplace *string      `json:"birthplace"`
	BloodType  *string      `json:"blood_type"`
	Hobby      *string      `json:"hobby"`
	Images     []ImageItem  `json:"images"`
	Blogs      []BlogListItem `json:"blogs"`
}

type DistrictListResponse struct {
	Women []WomanDistrictListItem `json:"women"`
}

func NewDistrictListResponse(women []entity.WomanEntity) DistrictListResponse {
	items := make([]WomanDistrictListItem, 0, len(women))
	for _, w := range women {
		items = append(items, toWomanDistrictListItem(w))
	}
	return DistrictListResponse{Women: items}
}

func toWomanDistrictListItem(w entity.WomanEntity) WomanDistrictListItem {
	images := make([]ImageItem, 0)
	for _, i := range w.GetImages().All() {
		images = append(images, ImageItem{ID: i.ID, Path: i.Path})
	}

	blogs := make([]BlogListItem, 0)
	for _, b := range w.GetBlogs().All() {
		blogs = append(blogs, BlogListItem{ID: b.GetID(), Title: b.GetTitle()})
	}

	return WomanDistrictListItem{
		ID:         w.GetID(),
		District:   AreaItem{ID: w.GetDistrict().GetID(), Name: w.GetDistrict().GetName()},
		Prefecture: AreaItem{ID: w.GetPrefecture().GetID(), Name: w.GetPrefecture().GetName()},
		Region:     AreaItem{ID: w.GetRegion().GetID(), Name: w.GetRegion().GetName()},
		Name:       w.GetName(),
		Age:        w.GetAge(),
		Birthplace: w.GetBirthplace(),
		BloodType:  w.GetBloodType(),
		Hobby:      w.GetHobby(),
		Images:     images,
		Blogs:      blogs,
	}
}
