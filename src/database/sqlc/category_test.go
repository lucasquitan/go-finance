package db

import (
	"context"
	"testing"

	"github.com/lucasquitan/go-finance/src/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)
	arg := CreateCategoryParams{
		UserID:      user.ID,
		Title:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Description, category.Description)
	require.NotEmpty(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	cat1 := createRandomCategory(t)
	cat2, err := testQueries.getCategory(context.Background(), cat1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, cat2)

	require.Equal(t, cat1.ID, cat2.ID)
	require.Equal(t, cat1.UserID, cat2.UserID)
	require.Equal(t, cat1.Title, cat2.Title)
	require.Equal(t, cat1.Type, cat2.Type)
	require.Equal(t, cat1.Description, cat2.Description)

	require.NotEmpty(t, cat2.CreatedAt)
}

func TestDeleteCategory(t *testing.T) {
	category := createRandomCategory(t)
	err := testQueries.deleteCategory(context.Background(), category.ID)

	require.NoError(t, err)
}

func TestUpdateGategory(t *testing.T) {
	category1 := createRandomCategory(t)

	arg := updateCategoryParams{
		ID:          category1.ID,
		Title:       util.RandomString(6),
		Description: util.RandomString(18),
	}

	category2, err := testQueries.updateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Title, category2.Title)
	require.Equal(t, arg.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestListCategories(t *testing.T) {
	lastCategory := createRandomCategory(t)

	args := getCategoriesParams{
		UserID:      lastCategory.UserID,
		Type:        lastCategory.Type,
		Title:       lastCategory.Title,
		Description: lastCategory.Description,
	}

	categories, err := testQueries.getCategories(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.Equal(t, lastCategory.ID, category.ID)
		require.Equal(t, lastCategory.UserID, category.UserID)
		require.Equal(t, lastCategory.Title, category.Title)
		require.Equal(t, lastCategory.Description, category.Description)
		require.NotEmpty(t, lastCategory.CreatedAt)
	}
}
