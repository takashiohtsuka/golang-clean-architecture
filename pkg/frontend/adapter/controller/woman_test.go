package controller_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang-clean-architecture/pkg/apperror"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/adapter/controller"
	responseWomen "golang-clean-architecture/pkg/frontend/adapter/response/women"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/input"
	"golang-clean-architecture/pkg/usecase/query"
)

// --- mock Context ---

type mockContext struct {
	params      map[string]string
	queryParams map[string]string
	statusCode  int
	body        any
}

func (m *mockContext) JSON(code int, i any) error {
	m.statusCode = code
	m.body = i
	return nil
}
func (m *mockContext) Bind(i any) error          { return nil }
func (m *mockContext) Param(name string) string          { return m.params[name] }
func (m *mockContext) QueryParam(name string) string     { return m.queryParams[name] }
func (m *mockContext) Request() *http.Request            { return &http.Request{} }

// --- mock WomanUsecase ---

type mockWomanUsecase struct {
	getListInput  input.GetWomanListInput
	getListReturn collection.Collection[entity.WomanEntity]
	getListError  error

	getDetailInput  input.GetWomanDetailInput
	getDetailReturn entity.WomanEntity
	getDetailError  error
}

func (m *mockWomanUsecase) GetList(_ context.Context, i input.GetWomanListInput) (collection.Collection[entity.WomanEntity], error) {
	m.getListInput = i
	return m.getListReturn, m.getListError
}

func (m *mockWomanUsecase) GetDetail(i input.GetWomanDetailInput) (entity.WomanEntity, error) {
	m.getDetailInput = i
	return m.getDetailReturn, m.getDetailError
}

// --- GetWomanList tests ---

func TestWomanController_GetWomanList_WithNoStoreID_Returns200(t *testing.T) {
	uc := &mockWomanUsecase{
		getListReturn: collection.NewCollection[entity.WomanEntity](nil),
	}
	ctx := &mockContext{queryParams: map[string]string{}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	assert.Nil(t, uc.getListInput.StoreID)
}

func TestWomanController_GetWomanList_WithValidStoreID_PassesStoreIDToUsecase(t *testing.T) {
	uc := &mockWomanUsecase{
		getListReturn: collection.NewCollection[entity.WomanEntity](nil),
	}
	ctx := &mockContext{queryParams: map[string]string{"store_id": "1"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	require.NotNil(t, uc.getListInput.StoreID)
	assert.Equal(t, uint(1), *uc.getListInput.StoreID)
}

func TestWomanController_GetWomanList_WithInvalidStoreID_Returns400(t *testing.T) {
	uc := &mockWomanUsecase{}
	ctx := &mockContext{queryParams: map[string]string{"store_id": "invalid"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, ctx.statusCode)
}

func TestWomanController_GetWomanList_WhenUsecaseFails_Returns500(t *testing.T) {
	uc := &mockWomanUsecase{
		getListError: errors.New("db error"),
	}
	ctx := &mockContext{queryParams: map[string]string{}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.statusCode)
}

// --- GetWomanDetail tests ---

func TestWomanController_GetWomanDetail_WithValidID_Returns200(t *testing.T) {
	uc := &mockWomanUsecase{
		getDetailReturn: &entity.Woman{ID: 1, Name: "女性1"},
	}
	ctx := &mockContext{params: map[string]string{"id": "1"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	assert.Equal(t, uint(1), uc.getDetailInput.WomanID)
}

func TestWomanController_GetWomanDetail_WithInvalidID_Returns400(t *testing.T) {
	uc := &mockWomanUsecase{}
	ctx := &mockContext{params: map[string]string{"id": "invalid"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, ctx.statusCode)
}

func TestWomanController_GetWomanDetail_WhenNotFound_Returns404(t *testing.T) {
	uc := &mockWomanUsecase{
		getDetailError: apperror.NewNotFoundException("woman not found"),
	}
	ctx := &mockContext{params: map[string]string{"id": "99999"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, ctx.statusCode)
}

func TestWomanController_GetWomanDetail_WhenUsecaseFails_Returns500(t *testing.T) {
	uc := &mockWomanUsecase{
		getDetailError: errors.New("db error"),
	}
	ctx := &mockContext{params: map[string]string{"id": "1"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.statusCode)
}

// --- GetWomanList response body tests ---

func TestWomanController_GetWomanList_ResponseBodyIsListResponse(t *testing.T) {
	uc := &mockWomanUsecase{
		getListReturn: collection.NewCollection[entity.WomanEntity]([]entity.WomanEntity{
			&entity.Woman{
				ID:               1,
				Name:             "女性1",
				StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
				Images:           collection.NewCollection[entity.WomanImage](nil),
				Blogs:            collection.NewCollection[entity.BlogEntity](nil),
			},
		}),
	}
	ctx := &mockContext{queryParams: map[string]string{}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanList(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	body, ok := ctx.body.(responseWomen.ListResponse)
	require.True(t, ok, "body は ListResponse 型であること")
	require.Len(t, body.Women, 1)
	assert.Equal(t, uint(1), body.Women[0].ID)
	assert.Equal(t, "女性1", body.Women[0].Name)
}

// --- GetWomanDetail response body tests ---

func TestWomanController_GetWomanDetail_ResponseBodyIsDetailResponse(t *testing.T) {
	uc := &mockWomanUsecase{
		getDetailReturn: &entity.Woman{
			ID:               1,
			Name:             "女性1",
			StoreAssignments: collection.NewCollection[entity.WomanStoreAssignment](nil),
			Images:           collection.NewCollection[entity.WomanImage](nil),
			Blogs:            collection.NewCollection[entity.BlogEntity](nil),
		},
	}
	ctx := &mockContext{params: map[string]string{"id": "1"}}

	c := controller.NewWomanController(uc)
	err := c.GetWomanDetail(ctx)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, ctx.statusCode)
	body, ok := ctx.body.(responseWomen.DetailResponse)
	require.True(t, ok, "body は DetailResponse 型であること")
	assert.Equal(t, uint(1), body.ID)
	assert.Equal(t, "女性1", body.Name)
}

// --- query.Condition の検証ヘルパー（参考） ---

var _ = query.Where // query パッケージが使われていることの確認
