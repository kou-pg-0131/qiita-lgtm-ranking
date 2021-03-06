package gateways

import (
	"errors"
	"testing"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
	"github.com/stretchr/testify/assert"
)

/*
 * NewReportsRepository()
 */

func Test_NewReportsRepository(t *testing.T) {
	mapi := new(mockQiitaAPI)
	r := NewReportsRepository(mapi)

	assert.NotNil(t, r)
	assert.Equal(t, mapi, r.qiitaAPI)
}

/*
 * ReportsRepository.Update()
 */

func TestReportsRepository_UpdateReturnNilWhenSucceeded(t *testing.T) {
	mapi := new(mockQiitaAPI)
	mapi.On("UpdateItem", "ID", "TITLE", "BODY", domain.Tags{{Name: "Qiita"}, {Name: "lgtm"}, {Name: "ランキング"}, {Name: "TAG_1"}, {Name: "TAG_2"}}).Return(nil)

	r := &ReportsRepository{qiitaAPI: mapi}

	err := r.Update("ID", "TITLE", "BODY", domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}})

	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "UpdateItem", 1)
}

func TestReportsRepository_UpdateReturnErrorWhenUpdateItemFailed(t *testing.T) {
	mapi := new(mockQiitaAPI)
	mapi.On("UpdateItem", "ID", "TITLE", "BODY", domain.Tags{{Name: "Qiita"}, {Name: "lgtm"}, {Name: "ランキング"}, {Name: "TAG_1"}, {Name: "TAG_2"}}).Return(errors.New("SOMETHING_WRONG"))

	r := &ReportsRepository{qiitaAPI: mapi}

	err := r.Update("ID", "TITLE", "BODY", domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}})

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mapi.AssertNumberOfCalls(t, "UpdateItem", 1)
}
