package interactor_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang-clean-architecture/pkg/apperror"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/input"
	"golang-clean-architecture/pkg/frontend/usecase/interactor"
	"golang-clean-architecture/pkg/usecase/query"
)

// --- mock ---

type mockWomanRepository struct {
	findAllConditions []query.Condition
	findAllReturn     collection.Collection[entity.WomanEntity]
	findAllError      error

	findOneConditions []query.Condition
	findOneReturn     entity.WomanEntity
	findOneError      error
}

func (m *mockWomanRepository) FindAll(_ context.Context, conditions []query.Condition) (collection.Collection[entity.WomanEntity], error) {
	m.findAllConditions = conditions
	return m.findAllReturn, m.findAllError
}

func (m *mockWomanRepository) FindOne(conditions []query.Condition) (entity.WomanEntity, error) {
	m.findOneConditions = conditions
	return m.findOneReturn, m.findOneError
}

// --- GetList tests ---

func TestWomanUsecase_GetList_WithNoStoreID_PassesEmptyConditions(t *testing.T) {
	mock := &mockWomanRepository{
		findAllReturn: collection.NewCollection[entity.WomanEntity](nil),
	}
	u := interactor.NewWomanUsecase(mock)

	_, err := u.GetList(context.Background(), input.GetWomanListInput{})

	require.NoError(t, err)
	assert.Empty(t, mock.findAllConditions)
}

func TestWomanUsecase_GetList_WithStoreID_PassesStoreIDCondition(t *testing.T) {
	mock := &mockWomanRepository{
		findAllReturn: collection.NewCollection[entity.WomanEntity](nil),
	}
	u := interactor.NewWomanUsecase(mock)

	storeID := uint(1)
	_, err := u.GetList(context.Background(), input.GetWomanListInput{StoreID: &storeID})

	require.NoError(t, err)
	require.Len(t, mock.findAllConditions, 1)
	assert.Equal(t, "wsa.store_id", mock.findAllConditions[0].Column)
	assert.Equal(t, storeID, mock.findAllConditions[0].Value)
}

func TestWomanUsecase_GetList_WhenRepositoryFails_ReturnsError(t *testing.T) {
	mock := &mockWomanRepository{
		findAllError: errors.New("db error"),
	}
	u := interactor.NewWomanUsecase(mock)

	_, err := u.GetList(context.Background(), input.GetWomanListInput{})

	assert.Error(t, err)
}

// --- GetDetail tests ---

func TestWomanUsecase_GetDetail_PassesWomanIDCondition(t *testing.T) {
	mock := &mockWomanRepository{
		findOneReturn: &entity.Woman{ID: 42, Name: "女性1"},
	}
	u := interactor.NewWomanUsecase(mock)

	_, err := u.GetDetail(input.GetWomanDetailInput{WomanID: 42})

	require.NoError(t, err)
	require.Len(t, mock.findOneConditions, 1)
	assert.Equal(t, "w.id", mock.findOneConditions[0].Column)
	assert.Equal(t, uint(42), mock.findOneConditions[0].Value)
}

func TestWomanUsecase_GetDetail_WhenNotFound_ReturnsNotFoundException(t *testing.T) {
	mock := &mockWomanRepository{
		findOneReturn: &entity.NilWoman{},
	}
	u := interactor.NewWomanUsecase(mock)

	_, err := u.GetDetail(input.GetWomanDetailInput{WomanID: 99999})

	require.Error(t, err)
	var nfe *apperror.NotFoundException
	assert.True(t, errors.As(err, &nfe), "NotFoundException が返ること")
}

func TestWomanUsecase_GetDetail_WhenRepositoryFails_ReturnsError(t *testing.T) {
	mock := &mockWomanRepository{
		findOneError: errors.New("db error"),
	}
	u := interactor.NewWomanUsecase(mock)

	_, err := u.GetDetail(input.GetWomanDetailInput{WomanID: 1})

	assert.Error(t, err)
}
