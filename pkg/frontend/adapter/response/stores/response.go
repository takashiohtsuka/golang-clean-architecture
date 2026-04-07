package stores

import "golang-clean-architecture/pkg/frontend/domain/entity"

// --- list ---

type ListResponse struct {
	Stores []StoreListItem `json:"stores"`
}

type StoreListItem struct {
	ID           uint           `json:"id"`
	Name         string         `json:"name"`
	BusinessType string         `json:"business_type"`
	Women        []WomanItem    `json:"women"`
}

type WomanItem struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Blogs []BlogItem `json:"blogs"`
}

type BlogItem struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

// --- detail ---

type DetailResponse struct {
	ID           uint        `json:"id"`
	Name         string      `json:"name"`
	BusinessType string      `json:"business_type"`
	Women        []WomanItem `json:"women"`
}

// --- builders ---

func NewListResponse(stores []entity.StoreEntity) ListResponse {
	items := make([]StoreListItem, 0, len(stores))
	for _, s := range stores {
		items = append(items, toStoreListItem(s))
	}
	return ListResponse{Stores: items}
}

func toStoreListItem(s entity.StoreEntity) StoreListItem {
	return StoreListItem{
		ID:           s.GetID(),
		Name:         s.GetName(),
		BusinessType: s.GetBusinessType().GetCode(),
		Women:        toWomenItems(s),
	}
}

func NewDetailResponse(s entity.StoreEntity) DetailResponse {
	return DetailResponse{
		ID:           s.GetID(),
		Name:         s.GetName(),
		BusinessType: s.GetBusinessType().GetCode(),
		Women:        toWomenItems(s),
	}
}

func toWomenItems(s entity.StoreEntity) []WomanItem {
	women := make([]WomanItem, 0)
	for _, w := range s.GetWomen().All() {
		blogs := make([]BlogItem, 0)
		for _, b := range w.GetBlogs().All() {
			blogs = append(blogs, BlogItem{
				ID:    b.GetID(),
				Title: b.GetTitle(),
			})
		}
		women = append(women, WomanItem{
			ID:    w.GetID(),
			Name:  w.GetName(),
			Blogs: blogs,
		})
	}
	return women
}
