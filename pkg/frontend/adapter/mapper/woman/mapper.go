package woman

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	fvo "golang-clean-architecture/pkg/frontend/domain/valueobject"
	"golang-clean-architecture/pkg/helper"
)

// MapToAggregate は flat な rows をエリアコンテキスト付き Woman 集約のコレクションに変換する。
func MapToAggregate(rows []map[string]any) collection.Collection[entity.WomanEntity] {
	womanOrder := make([]uint, 0)
	womanMap := make(map[uint]*entity.Woman)
	seenAssignments := make(map[uint]map[uint]bool)
	seenImages := make(map[uint]map[uint]bool)
	seenBlogs := make(map[uint]map[uint]bool)

	for _, row := range rows {
		womanID := helper.ToUint(row["woman_id"])

		if _, exists := womanMap[womanID]; !exists {
			womanOrder = append(womanOrder, womanID)
			womanMap[womanID] = &entity.Woman{
				ID:         womanID,
				District:   fvo.NewDistrict(helper.ToUint(row["district_id"]), helper.ToString(row["district_name"])),
				Prefecture: fvo.NewPrefecture(helper.ToUint(row["prefecture_id"]), helper.ToString(row["prefecture_name"])),
				Region:     fvo.NewRegion(helper.ToUint(row["region_id"]), helper.ToString(row["region_name"])),
				Name:       helper.ToString(row["woman_name"]),
				Age:        helper.ToIntPtr(row["age"]),
				Birthplace: helper.ToStringPtr(row["birthplace"]),
				BloodType:  helper.ToStringPtr(row["blood_type"]),
				Hobby:      helper.ToStringPtr(row["hobby"]),
			}
			seenAssignments[womanID] = make(map[uint]bool)
			seenImages[womanID] = make(map[uint]bool)
			seenBlogs[womanID] = make(map[uint]bool)
		}

		assignmentID := helper.ToUint(row["assignment_id"])
		if assignmentID != 0 && !seenAssignments[womanID][assignmentID] {
			seenAssignments[womanID][assignmentID] = true
			current := womanMap[womanID].StoreAssignments.All()
			current = append(current, entity.WomanStoreAssignment{
				ID:      assignmentID,
				StoreID: helper.ToUint(row["assignment_store_id"]),
			})
			womanMap[womanID].StoreAssignments = collection.NewCollection(current)
		}

		imageID := helper.ToUint(row["image_id"])
		if imageID != 0 && !seenImages[womanID][imageID] {
			seenImages[womanID][imageID] = true
			current := womanMap[womanID].Images.All()
			current = append(current, entity.WomanImage{
				ID:   imageID,
				Path: helper.ToString(row["image_path"]),
			})
			womanMap[womanID].Images = collection.NewCollection(current)
		}

		blogID := helper.ToUint(row["blog_id"])
		if blogID != 0 && !seenBlogs[womanID][blogID] {
			seenBlogs[womanID][blogID] = true
			current := womanMap[womanID].Blogs.All()
			current = append(current, &entity.Blog{
				ID:          blogID,
				WomanID:     womanID,
				Title:       helper.ToString(row["blog_title"]),
				IsPublished: true,
				Photos:      collection.NewCollection[entity.Photo](nil),
			})
			womanMap[womanID].Blogs = collection.NewCollection(current)
		}
	}

	items := make([]entity.WomanEntity, 0, len(womanOrder))
	for _, wid := range womanOrder {
		items = append(items, womanMap[wid])
	}
	return collection.NewCollection(items)
}
