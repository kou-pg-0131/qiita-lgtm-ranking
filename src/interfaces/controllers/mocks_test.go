package controllers

import (
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
	"github.com/stretchr/testify/mock"
)

type mockItemsRepository struct {
	mock.Mock
}

func (m *mockItemsRepository) GetAll(query string) (*domain.Items, error) {
	args := m.Called(query)
	return args.Get(0).(*domain.Items), args.Error(1)
}

type mockReportsRepository struct {
	mock.Mock
}

func (m *mockReportsRepository) Update(id, title, body string, tags domain.Tags) error {
	args := m.Called(id, title, body, tags)
	return args.Error(0)
}

type mockReportsPresenter struct {
	mock.Mock
}

func (m *mockReportsPresenter) Weekly(from time.Time, items *domain.Items) (string, error) {
	args := m.Called(from, items)
	return args.String(0), args.Error(1)
}

func (m *mockReportsPresenter) WeeklyByTag(from time.Time, items *domain.Items, tag string) (string, error) {
	args := m.Called(from, items, tag)
	return args.String(0), args.Error(1)
}

type mockOS struct {
	mock.Mock
}

func (m *mockOS) Getenv(s string) string {
	args := m.Called(s)
	return args.String(0)
}
