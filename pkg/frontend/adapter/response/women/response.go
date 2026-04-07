package women

import "golang-clean-architecture/pkg/frontend/domain/entity"

// --- list ---

type ListResponse struct {
	Women []WomanListItem `json:"women"`
}

type WomanListItem struct {
	ID               uint                  `json:"id"`
	Name             string                `json:"name"`
	Age              *int                  `json:"age"`
	Birthplace       *string               `json:"birthplace"`
	BloodType        *string               `json:"blood_type"`
	Hobby            *string               `json:"hobby"`
	StoreAssignments []StoreAssignmentItem `json:"store_assignments"`
	Images           []ImageItem           `json:"images"`
	Blogs            []BlogListItem        `json:"blogs"`
}

type StoreAssignmentItem struct {
	ID      uint `json:"id"`
	StoreID uint `json:"store_id"`
}

type ImageItem struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type BlogListItem struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

// --- detail ---

type DetailResponse struct {
	ID               uint                  `json:"id"`
	Name             string                `json:"name"`
	Age              *int                  `json:"age"`
	Birthplace       *string               `json:"birthplace"`
	BloodType        *string               `json:"blood_type"`
	Hobby            *string               `json:"hobby"`
	StoreAssignments []StoreAssignmentItem `json:"store_assignments"`
	Images           []ImageItem           `json:"images"`
	Blogs            []BlogDetailItem      `json:"blogs"`
}

type BlogDetailItem struct {
	ID     uint        `json:"id"`
	Title  string      `json:"title"`
	Body   *string     `json:"body"`
	Photos []PhotoItem `json:"photos"`
}

type PhotoItem struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

// --- builders ---

func NewListResponse(women []entity.WomanEntity) ListResponse {
	items := make([]WomanListItem, 0, len(women))
	for _, w := range women {
		items = append(items, toWomanListItem(w))
	}
	return ListResponse{Women: items}
}

func toWomanListItem(w entity.WomanEntity) WomanListItem {
	assignments := make([]StoreAssignmentItem, 0)
	for _, a := range w.GetStoreAssignments().All() {
		assignments = append(assignments, StoreAssignmentItem{
			ID:      a.ID,
			StoreID: a.StoreID,
		})
	}

	images := make([]ImageItem, 0)
	for _, i := range w.GetImages().All() {
		images = append(images, ImageItem{
			ID:   i.ID,
			Path: i.Path,
		})
	}

	blogs := make([]BlogListItem, 0)
	for _, b := range w.GetBlogs().All() {
		blogs = append(blogs, BlogListItem{
			ID:    b.GetID(),
			Title: b.GetTitle(),
		})
	}

	return WomanListItem{
		ID:               w.GetID(),
		Name:             w.GetName(),
		Age:              w.GetAge(),
		Birthplace:       w.GetBirthplace(),
		BloodType:        w.GetBloodType(),
		Hobby:            w.GetHobby(),
		StoreAssignments: assignments,
		Images:           images,
		Blogs:            blogs,
	}
}

func NewDetailResponse(w entity.WomanEntity) DetailResponse {
	assignments := make([]StoreAssignmentItem, 0)
	for _, a := range w.GetStoreAssignments().All() {
		assignments = append(assignments, StoreAssignmentItem{
			ID:      a.ID,
			StoreID: a.StoreID,
		})
	}

	images := make([]ImageItem, 0)
	for _, i := range w.GetImages().All() {
		images = append(images, ImageItem{
			ID:   i.ID,
			Path: i.Path,
		})
	}

	blogs := make([]BlogDetailItem, 0)
	for _, b := range w.GetBlogs().All() {
		photos := make([]PhotoItem, 0)
		for _, p := range b.GetPhotos().All() {
			photos = append(photos, PhotoItem{
				ID:  p.ID,
				URL: p.URL,
			})
		}
		blogs = append(blogs, BlogDetailItem{
			ID:     b.GetID(),
			Title:  b.GetTitle(),
			Body:   b.GetBody(),
			Photos: photos,
		})
	}

	return DetailResponse{
		ID:               w.GetID(),
		Name:             w.GetName(),
		Age:              w.GetAge(),
		Birthplace:       w.GetBirthplace(),
		BloodType:        w.GetBloodType(),
		Hobby:            w.GetHobby(),
		StoreAssignments: assignments,
		Images:           images,
		Blogs:            blogs,
	}
}
