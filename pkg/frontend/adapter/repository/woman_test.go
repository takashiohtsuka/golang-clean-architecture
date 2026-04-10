package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"golang-clean-architecture/pkg/frontend/adapter/repository"
	"golang-clean-architecture/pkg/infrastructure/datastore"
	"golang-clean-architecture/pkg/usecase/query"
)

// --- db helper ---

func sqlDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := datastore.NewTestDB().DB()
	require.NoError(t, err)
	return db
}

func exec(t *testing.T, db *sql.DB, query string, args ...any) int64 {
	t.Helper()
	res, err := db.Exec(query, args...)
	require.NoError(t, err)
	id, err := res.LastInsertId()
	require.NoError(t, err)
	return id
}

// --- cleanup helper ---

func deletePhotos(db *sql.DB)                { db.Exec("DELETE FROM photos") }
func deleteBlogs(db *sql.DB)                 { db.Exec("DELETE FROM blogs") }
func deleteWomanImages(db *sql.DB)           { db.Exec("DELETE FROM woman_images") }
func deleteWomanStoreAssignments(db *sql.DB) { db.Exec("DELETE FROM woman_store_assignments") }
func deleteWomen(db *sql.DB)                 { db.Exec("DELETE FROM women") }
func deleteStores(db *sql.DB)                { db.Exec("DELETE FROM stores") }
func deleteCompanies(db *sql.DB)             { db.Exec("DELETE FROM companies") }
func deleteContractPlans(db *sql.DB)         { db.Exec("DELETE FROM contract_plans") }
func deleteBusinessTypes(db *sql.DB)         { db.Exec("DELETE FROM business_types") }

func cleanupAll(db *sql.DB) {
	deletePhotos(db)
	deleteBlogs(db)
	deleteWomanImages(db)
	deleteWomanStoreAssignments(db)
	deleteWomen(db)
	deleteStores(db)
	deleteCompanies(db)
	deleteContractPlans(db)
	deleteBusinessTypes(db)
}

// --- tests ---

func TestWomanRepository_FindAll_ReturnsAllActiveWomen(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
	exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性2', true)", companyID)
	exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '非アクティブ女性', false)", companyID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindAll(context.Background(), []query.Condition{})

	require.NoError(t, err)
	require.Equal(t, 2, len(result.All()))
}

func TestWomanRepository_FindAll_ExcludesInactiveWoman(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	inactiveID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '非アクティブ女性', false)", companyID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindAll(context.Background(), []query.Condition{})

	require.NoError(t, err)

	ids := make([]uint, 0)
	for _, w := range result.All() {
		ids = append(ids, w.GetID())
	}
	assert.NotContains(t, ids, uint(inactiveID))
}

func TestWomanRepository_FindAll_FilterByStoreID(t *testing.T) {
	tests := []struct {
		name          string
		targetStore   string
		expectedCount int
	}{
		{
			name:          "Store1：woman1とwoman2が返る",
			targetStore:   "store1",
			expectedCount: 2,
		},
		{
			name:          "Store2：woman2のみ返る",
			targetStore:   "store2",
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := sqlDB(t)
			t.Cleanup(func() { cleanupAll(db) })

			btID := exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
			cpID := exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
			companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
			store1ID := exec(t, db, "INSERT INTO stores (company_id, business_type_id, contract_plan_id, name, is_active, open_status) VALUES (?, ?, ?, '店舗1', true, 'open')", companyID, btID, cpID)
			store2ID := exec(t, db, "INSERT INTO stores (company_id, business_type_id, contract_plan_id, name, is_active, open_status) VALUES (?, ?, ?, '店舗2', true, 'open')", companyID, btID, cpID)
			woman1ID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
			woman2ID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性2', true)", companyID)
			exec(t, db, "INSERT INTO woman_store_assignments (woman_id, store_id) VALUES (?, ?)", woman1ID, store1ID)
			exec(t, db, "INSERT INTO woman_store_assignments (woman_id, store_id) VALUES (?, ?)", woman2ID, store1ID)
			exec(t, db, "INSERT INTO woman_store_assignments (woman_id, store_id) VALUES (?, ?)", woman2ID, store2ID)

			targetStoreID := store1ID
			if tt.targetStore == "store2" {
				targetStoreID = store2ID
			}

			repo := repository.NewWomanRepository(datastore.NewTestDB())
			result, err := repo.FindAll(context.Background(), []query.Condition{
				query.Where("wsa.store_id", uint(targetStoreID)),
			})

			require.NoError(t, err)
			require.Equal(t, tt.expectedCount, len(result.All()))
		})
	}
}

func TestWomanRepository_FindAll_MapsImages(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	womanID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
	exec(t, db, "INSERT INTO woman_images (woman_id, path) VALUES (?, 'images/photo.jpg')", womanID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindAll(context.Background(), []query.Condition{})

	require.NoError(t, err)
	require.Equal(t, 1, len(result.All()))
	assert.Equal(t, 1, len(result.All()[0].GetImages().All()))
}

func TestWomanRepository_FindAll_MapsPublishedBlogsOnly(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	womanID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ1', true)", womanID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ2', true)", womanID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, '未公開ブログ', false)", womanID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindAll(context.Background(), []query.Condition{})

	require.NoError(t, err)
	require.Equal(t, 1, len(result.All()))
	assert.Equal(t, 2, len(result.All()[0].GetBlogs().All()), "公開ブログ2件のみ返る")
}

func TestWomanRepository_FindOne_ReturnsWomanWithBlogsAndPhotos(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	womanID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
	exec(t, db, "INSERT INTO woman_images (woman_id, path) VALUES (?, 'images/photo.jpg')", womanID)
	blog1ID := exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ1', true)", womanID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ2', true)", womanID)
	exec(t, db, "INSERT INTO photos (blog_id, url) VALUES (?, 'photos/photo1.jpg')", blog1ID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindOne(context.Background(), []query.Condition{query.Where("w.id", uint(womanID))})

	require.NoError(t, err)
	require.False(t, result.IsNil())
	require.Equal(t, 2, len(result.GetBlogs().All()))
	require.Equal(t, 1, len(result.GetImages().All()))

	totalPhotos := 0
	for _, b := range result.GetBlogs().All() {
		totalPhotos += len(b.GetPhotos().All())
	}
	assert.Equal(t, 1, totalPhotos)
}

func TestWomanRepository_FindOne_ReturnsNilWhenNotFound(t *testing.T) {
	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindOne(context.Background(), []query.Condition{query.Where("w.id", uint(99999))})

	require.NoError(t, err)
	assert.True(t, result.IsNil())
}

func TestWomanRepository_FindAll_MapsStoreAssignments(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	btID := exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	cpID := exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	store1ID := exec(t, db, "INSERT INTO stores (company_id, business_type_id, contract_plan_id, name, is_active, open_status) VALUES (?, ?, ?, '店舗1', true, 'open')", companyID, btID, cpID)
	store2ID := exec(t, db, "INSERT INTO stores (company_id, business_type_id, contract_plan_id, name, is_active, open_status) VALUES (?, ?, ?, '店舗2', true, 'open')", companyID, btID, cpID)
	womanID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
	exec(t, db, "INSERT INTO woman_store_assignments (woman_id, store_id) VALUES (?, ?)", womanID, store1ID)
	exec(t, db, "INSERT INTO woman_store_assignments (woman_id, store_id) VALUES (?, ?)", womanID, store2ID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindAll(context.Background(), []query.Condition{})

	require.NoError(t, err)
	require.Equal(t, 1, len(result.All()))
	assert.Equal(t, 2, len(result.All()[0].GetStoreAssignments().All()))
}

func TestWomanRepository_FindAll_LimitsBlogsToThree(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	womanID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ1', true)", womanID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ2', true)", womanID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ3', true)", womanID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ4', true)", womanID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindAll(context.Background(), []query.Condition{})

	require.NoError(t, err)
	require.Equal(t, 1, len(result.All()))
	assert.Equal(t, 3, len(result.All()[0].GetBlogs().All()), "womanBlogsLimit=3 で3件に絞られること")
}

func TestWomanRepository_FindAll_DeduplicatesWhenImagesAndBlogsAndAssignmentsExist(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	btID := exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	cpID := exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	storeID := exec(t, db, "INSERT INTO stores (company_id, business_type_id, contract_plan_id, name, is_active, open_status) VALUES (?, ?, ?, '店舗1', true, 'open')", companyID, btID, cpID)
	womanID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
	exec(t, db, "INSERT INTO woman_store_assignments (woman_id, store_id) VALUES (?, ?)", womanID, storeID)
	exec(t, db, "INSERT INTO woman_images (woman_id, path) VALUES (?, 'images/photo.jpg')", womanID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ1', true)", womanID)
	exec(t, db, "INSERT INTO blogs (woman_id, title, is_published) VALUES (?, 'ブログ2', true)", womanID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindAll(context.Background(), []query.Condition{})

	require.NoError(t, err)
	require.Equal(t, 1, len(result.All()), "womanは1件のみ返ること（JOIN重複排除）")
	w := result.All()[0]
	assert.Equal(t, 1, len(w.GetStoreAssignments().All()), "StoreAssignmentsが重複しないこと")
	assert.Equal(t, 1, len(w.GetImages().All()), "Imagesが重複しないこと")
	assert.Equal(t, 2, len(w.GetBlogs().All()), "Blogsが重複しないこと")
}

func TestWomanRepository_FindAll_ReturnsErrorWhenDBFailed(t *testing.T) {
	db := datastore.NewTestDB()
	sqlDB, err := db.DB()
	require.NoError(t, err)
	sqlDB.Close()

	repo := repository.NewWomanRepository(db)
	_, err = repo.FindAll(context.Background(), []query.Condition{})

	assert.Error(t, err)
}

func TestWomanRepository_FindOne_ReturnsErrorWhenDBFailed(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery("SELECT").WillReturnError(errors.New("db error"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	repo := repository.NewWomanRepository(gormDB)
	_, err = repo.FindOne(context.Background(), []query.Condition{query.Where("w.id", uint(1))})

	assert.Error(t, err)
}

func TestWomanRepository_FindOne_MapsStoreAssignments(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	btID := exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	cpID := exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	store1ID := exec(t, db, "INSERT INTO stores (company_id, business_type_id, contract_plan_id, name, is_active, open_status) VALUES (?, ?, ?, '店舗1', true, 'open')", companyID, btID, cpID)
	store2ID := exec(t, db, "INSERT INTO stores (company_id, business_type_id, contract_plan_id, name, is_active, open_status) VALUES (?, ?, ?, '店舗2', true, 'open')", companyID, btID, cpID)
	womanID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
	exec(t, db, "INSERT INTO woman_store_assignments (woman_id, store_id) VALUES (?, ?)", womanID, store1ID)
	exec(t, db, "INSERT INTO woman_store_assignments (woman_id, store_id) VALUES (?, ?)", womanID, store2ID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindOne(context.Background(), []query.Condition{query.Where("w.id", uint(womanID))})

	require.NoError(t, err)
	require.False(t, result.IsNil())
	assert.Equal(t, 2, len(result.GetStoreAssignments().All()))
}

func TestWomanRepository_FindOne_MapsMultipleImages(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	womanID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '女性1', true)", companyID)
	exec(t, db, "INSERT INTO woman_images (woman_id, path) VALUES (?, 'images/photo1.jpg')", womanID)
	exec(t, db, "INSERT INTO woman_images (woman_id, path) VALUES (?, 'images/photo2.jpg')", womanID)
	exec(t, db, "INSERT INTO woman_images (woman_id, path) VALUES (?, 'images/photo3.jpg')", womanID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindOne(context.Background(), []query.Condition{query.Where("w.id", uint(womanID))})

	require.NoError(t, err)
	require.False(t, result.IsNil())
	assert.Equal(t, 3, len(result.GetImages().All()))
}

func TestWomanRepository_FindOne_ReturnsNilForInactiveWoman(t *testing.T) {
	db := sqlDB(t)
	t.Cleanup(func() { cleanupAll(db) })

	exec(t, db, "INSERT INTO business_types (code) VALUES ('A')")
	exec(t, db, "INSERT INTO contract_plans (code) VALUES ('standard')")
	companyID := exec(t, db, "INSERT INTO companies (name, is_active) VALUES ('会社1', true)")
	womanID := exec(t, db, "INSERT INTO women (company_id, name, is_active) VALUES (?, '非アクティブ女性', false)", companyID)

	repo := repository.NewWomanRepository(datastore.NewTestDB())
	result, err := repo.FindOne(context.Background(), []query.Condition{query.Where("w.id", uint(womanID))})

	require.NoError(t, err)
	assert.True(t, result.IsNil())
}
