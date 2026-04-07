package women_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/adapter/response/women"
	"golang-clean-architecture/pkg/frontend/domain/entity"
)

// --- helpers ---

func strPtr(s string) *string { return &s }
func intPtr(i int) *int       { return &i }

// --- NewListResponse tests ---

func TestNewListResponse_WithEmptySlice_ReturnsEmptyList(t *testing.T) {
	result := women.NewListResponse([]entity.WomanEntity{})

	require.NotNil(t, result.Women)
	assert.Empty(t, result.Women)
}

func TestNewListResponse_MapsBasicFields(t *testing.T) {
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		Age:              intPtr(25),
		Birthplace:       strPtr("東京"),
		BloodType:        strPtr("A"),
		Hobby:            strPtr("読書"),
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection[entity.BlogEntity](nil),
	}

	result := women.NewListResponse([]entity.WomanEntity{w})

	require.Len(t, result.Women, 1)
	item := result.Women[0]
	assert.Equal(t, uint(1), item.ID)
	assert.Equal(t, "女性1", item.Name)
	assert.Equal(t, intPtr(25), item.Age)
	assert.Equal(t, strPtr("東京"), item.Birthplace)
	assert.Equal(t, strPtr("A"), item.BloodType)
	assert.Equal(t, strPtr("読書"), item.Hobby)
}

func TestNewListResponse_MapsStoreAssignments(t *testing.T) {
	assignments := []entity.WomanStoreAssignment{
		{ID: 10, StoreID: 100},
		{ID: 11, StoreID: 101},
	}
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		StoreAssignments: collection.NewCollection(assignments),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection[entity.BlogEntity](nil),
	}

	result := women.NewListResponse([]entity.WomanEntity{w})

	require.Len(t, result.Women[0].StoreAssignments, 2)
	assert.Equal(t, uint(10), result.Women[0].StoreAssignments[0].ID)
	assert.Equal(t, uint(100), result.Women[0].StoreAssignments[0].StoreID)
	assert.Equal(t, uint(11), result.Women[0].StoreAssignments[1].ID)
	assert.Equal(t, uint(101), result.Women[0].StoreAssignments[1].StoreID)
}

func TestNewListResponse_MapsImages(t *testing.T) {
	images := []entity.WomanImage{
		{ID: 20, Path: "images/photo1.jpg"},
	}
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection(images),
		Blogs:            collection.NewCollection[entity.BlogEntity](nil),
	}

	result := women.NewListResponse([]entity.WomanEntity{w})

	require.Len(t, result.Women[0].Images, 1)
	assert.Equal(t, uint(20), result.Women[0].Images[0].ID)
	assert.Equal(t, "images/photo1.jpg", result.Women[0].Images[0].Path)
}

func TestNewListResponse_MapsBlogs(t *testing.T) {
	blogs := []entity.BlogEntity{
		&entity.Blog{ID: 30, Title: "ブログ1", Photos: collection.NewCollection[entity.Photo](nil)},
		&entity.Blog{ID: 31, Title: "ブログ2", Photos: collection.NewCollection[entity.Photo](nil)},
	}
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection(blogs),
	}

	result := women.NewListResponse([]entity.WomanEntity{w})

	require.Len(t, result.Women[0].Blogs, 2)
	assert.Equal(t, uint(30), result.Women[0].Blogs[0].ID)
	assert.Equal(t, "ブログ1", result.Women[0].Blogs[0].Title)
	assert.Equal(t, uint(31), result.Women[0].Blogs[1].ID)
	assert.Equal(t, "ブログ2", result.Women[0].Blogs[1].Title)
}

func TestNewListResponse_WithMultipleWomen_MapsAll(t *testing.T) {
	w1 := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection[entity.BlogEntity](nil),
	}
	w2 := &entity.Woman{
		ID:               2,
		Name:             "女性2",
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection[entity.BlogEntity](nil),
	}

	result := women.NewListResponse([]entity.WomanEntity{w1, w2})

	require.Len(t, result.Women, 2)
	assert.Equal(t, uint(1), result.Women[0].ID)
	assert.Equal(t, uint(2), result.Women[1].ID)
}

// --- NewDetailResponse tests ---

func TestNewDetailResponse_MapsBasicFields(t *testing.T) {
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		Age:              intPtr(30),
		Birthplace:       strPtr("大阪"),
		BloodType:        strPtr("B"),
		Hobby:            strPtr("映画"),
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection[entity.BlogEntity](nil),
	}

	result := women.NewDetailResponse(w)

	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "女性1", result.Name)
	assert.Equal(t, intPtr(30), result.Age)
	assert.Equal(t, strPtr("大阪"), result.Birthplace)
	assert.Equal(t, strPtr("B"), result.BloodType)
	assert.Equal(t, strPtr("映画"), result.Hobby)
}

func TestNewDetailResponse_WithNilOptionalFields_MapsAsNil(t *testing.T) {
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection[entity.BlogEntity](nil),
	}

	result := women.NewDetailResponse(w)

	assert.Nil(t, result.Age)
	assert.Nil(t, result.Birthplace)
	assert.Nil(t, result.BloodType)
	assert.Nil(t, result.Hobby)
}

func TestNewDetailResponse_MapsBlogsWithPhotos(t *testing.T) {
	body := "本文"
	blogs := []entity.BlogEntity{
		&entity.Blog{
			ID:     30,
			Title:  "ブログ1",
			Body:   &body,
			Photos: collection.NewCollection([]entity.Photo{{ID: 50, URL: "photos/photo1.jpg"}}),
		},
	}
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection(blogs),
	}

	result := women.NewDetailResponse(w)

	require.Len(t, result.Blogs, 1)
	blog := result.Blogs[0]
	assert.Equal(t, uint(30), blog.ID)
	assert.Equal(t, "ブログ1", blog.Title)
	assert.Equal(t, &body, blog.Body)
	require.Len(t, blog.Photos, 1)
	assert.Equal(t, uint(50), blog.Photos[0].ID)
	assert.Equal(t, "photos/photo1.jpg", blog.Photos[0].URL)
}

func TestNewDetailResponse_MapsMultipleBlogsWithPhotos(t *testing.T) {
	blogs := []entity.BlogEntity{
		&entity.Blog{
			ID:     30,
			Title:  "ブログ1",
			Photos: collection.NewCollection([]entity.Photo{{ID: 50, URL: "photos/1.jpg"}, {ID: 51, URL: "photos/2.jpg"}}),
		},
		&entity.Blog{
			ID:     31,
			Title:  "ブログ2",
			Photos: collection.NewCollection[entity.Photo](nil),
		},
	}
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection(blogs),
	}

	result := women.NewDetailResponse(w)

	require.Len(t, result.Blogs, 2)
	assert.Len(t, result.Blogs[0].Photos, 2)
	assert.Empty(t, result.Blogs[1].Photos)
}

func TestNewDetailResponse_MapsImages(t *testing.T) {
	images := []entity.WomanImage{
		{ID: 20, Path: "images/photo1.jpg"},
		{ID: 21, Path: "images/photo2.jpg"},
	}
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection(images),
		Blogs:            collection.NewCollection[entity.BlogEntity](nil),
	}

	result := women.NewDetailResponse(w)

	require.Len(t, result.Images, 2)
	assert.Equal(t, uint(20), result.Images[0].ID)
	assert.Equal(t, "images/photo1.jpg", result.Images[0].Path)
	assert.Equal(t, uint(21), result.Images[1].ID)
	assert.Equal(t, "images/photo2.jpg", result.Images[1].Path)
}

func TestNewDetailResponse_WithEmptyCollections_ReturnsEmptySlices(t *testing.T) {
	w := &entity.Woman{
		ID:               1,
		Name:             "女性1",
		StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
		Images:           collection.NewCollection[entity.WomanImage](nil),
		Blogs:            collection.NewCollection[entity.BlogEntity](nil),
	}

	result := women.NewDetailResponse(w)

	assert.Empty(t, result.StoreAssignments)
	assert.Empty(t, result.Images)
	assert.Empty(t, result.Blogs)
}
