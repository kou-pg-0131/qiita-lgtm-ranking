package presenters

import (
	"strings"
	"testing"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
	"github.com/stretchr/testify/assert"
)

/*
 * NewReportsPresenter()
 */

func Test_NewReportsPresenter(t *testing.T) {
	p := NewReportsPresenter()

	assert.Equal(t, new(ReportsPresenter), p)
}

/*
 * ReportsPresenter.Weekly()
 */

func TestReportsPresenter_Weekly(t *testing.T) {
	p := &ReportsPresenter{}

	body, err := p.Weekly(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{
		{Title: "TITLE_HOGE", URL: "https://example.com/HOGE", LikesCount: 10, Tags: domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}}, User: domain.User{ID: "USER_HOGE"}, CreatedAt: time.Date(2020, 1, 1, 12, 1, 0, 0, time.UTC)},
		{Title: "TITLE_FUGA", URL: "https://example.com/FUGA", LikesCount: 13, Tags: domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}}, User: domain.User{ID: "USER_FUGA"}, CreatedAt: time.Date(2020, 1, 7, 4, 36, 0, 0, time.UTC)},
		{Title: "TITLE_FOO", URL: "https://example.com/FOO", Tags: domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}}, LikesCount: 8, User: domain.User{ID: "USER_FOO"}, CreatedAt: time.Date(2020, 1, 3, 8, 28, 0, 0, time.UTC)},
		{Title: "TITLE_BAR", URL: "https://example.com/BAR", Tags: domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}}, LikesCount: 0, User: domain.User{ID: "USER_HOGE"}, CreatedAt: time.Date(2019, 12, 1, 12, 36, 0, 0, time.UTC)},
	})

	assert.Equal(t, strings.Join([]string{
		"# タグ別 LGTM 数ランキング",
		"",
		"[`AWS`](https://qiita.com/items/e24b6279326a462d456c) [`Android`](https://qiita.com/items/8b3af051428d746f26c5) [`Docker`](https://qiita.com/items/ae11fca7d2eba445b037) [`Git`](https://qiita.com/items/74eacdbf363e260981c3) [`Go`](https://qiita.com/items/49d4537d95f878b3e91a) [`iOS`](https://qiita.com/items/e61a29a383d0403e92fc) [`Java`](https://qiita.com/items/4c3f84836bfdbb137226) [`JavaScript`](https://qiita.com/items/eaa7ac5b62a0a723edbb) [`Linux`](https://qiita.com/items/362e81e53c3f9dee22f1) [`Node.js`](https://qiita.com/items/66ed7ad8f7c9673e9d50) [`PHP`](https://qiita.com/items/3318cbdbc45c6ebd4014) [`Python`](https://qiita.com/items/9d7f2ffeafb36cf59a77) [`Rails`](https://qiita.com/items/93b9e7f7d143e9ce650e) [`React`](https://qiita.com/items/f9712f8acace22815b99) [`Ruby`](https://qiita.com/items/72c3d2e896bdc3e1a6b3) [`Swift`](https://qiita.com/items/e2b6f0645e29f0e2b761) [`TypeScript`](https://qiita.com/items/25b7c0870afa6d41d19b) [`Vim`](https://qiita.com/items/f5361177baef95e447d1) [`Vue.js`](https://qiita.com/items/2774e02c6eea5c830d99) [`初心者`](https://qiita.com/items/402899ec543aff109505)",
		"",
		"# 集計について",
		"",
		"- 期間: 2020-01-01 ~ 2020-01-08",
		"- 条件: ストック数が **10** 以上の記事",
		"",
		"ソースコード:",
		`<a href="https://github.com/kou-pg-0131/qiita-lgtm-ranking"><img src="https://github-link-card.s3.ap-northeast-1.amazonaws.com/kou-pg-0131/qiita-lgtm-ranking.png" width="460px"></a>`,
		"",
		"# LGTM 数ランキング",
		"",
		"## 1 位: [TITLE_FUGA](https://example.com/FUGA)",
		"",
		"[`TAG_1`](https://qiita.com/tags/TAG_1) [`TAG_2`](https://qiita.com/tags/TAG_2)",
		"",
		"**13** LGTM",
		"[@USER_FUGA](https://qiita.com/USER_FUGA) さん ( 2020-01-07 04:36 に投稿 )",
		"",
		"## 2 位: [TITLE_HOGE](https://example.com/HOGE)",
		"",
		"[`TAG_1`](https://qiita.com/tags/TAG_1) [`TAG_2`](https://qiita.com/tags/TAG_2)",
		"",
		"**10** LGTM",
		"[@USER_HOGE](https://qiita.com/USER_HOGE) さん ( 2020-01-01 12:01 に投稿 )",
		"",
		"## 3 位: [TITLE_FOO](https://example.com/FOO)",
		"",
		"[`TAG_1`](https://qiita.com/tags/TAG_1) [`TAG_2`](https://qiita.com/tags/TAG_2)",
		"",
		"**8** LGTM",
		"[@USER_FOO](https://qiita.com/USER_FOO) さん ( 2020-01-03 08:28 に投稿 )",
		"",
	}, "\n"), body)
	assert.Nil(t, err)
}

/*
 * ReportsPresenter.WeeklyByTag()
 */

func TestReportsPresenter_WeeklyByTag(t *testing.T) {
	p := &ReportsPresenter{}

	body, err := p.WeeklyByTag(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{
		{Title: "TITLE_HOGE", URL: "https://example.com/HOGE", LikesCount: 10, Tags: domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}}, User: domain.User{ID: "USER_HOGE"}, CreatedAt: time.Date(2020, 1, 1, 12, 1, 0, 0, time.UTC)},
		{Title: "TITLE_FUGA", URL: "https://example.com/FUGA", LikesCount: 13, Tags: domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}}, User: domain.User{ID: "USER_FUGA"}, CreatedAt: time.Date(2020, 1, 7, 4, 36, 0, 0, time.UTC)},
		{Title: "TITLE_FOO", URL: "https://example.com/FOO", Tags: domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}}, LikesCount: 8, User: domain.User{ID: "USER_FOO"}, CreatedAt: time.Date(2020, 1, 3, 8, 28, 0, 0, time.UTC)},
		{Title: "TITLE_BAR", URL: "https://example.com/BAR", Tags: domain.Tags{{Name: "TAG_1"}, {Name: "TAG_2"}}, LikesCount: 0, User: domain.User{ID: "USER_HOGE"}, CreatedAt: time.Date(2019, 12, 1, 12, 36, 0, 0, time.UTC)},
	}, "TAG")

	assert.Equal(t, strings.Join([]string{
		"# 他のタグ",
		"",
		"[`AWS`](https://qiita.com/items/e24b6279326a462d456c) [`Android`](https://qiita.com/items/8b3af051428d746f26c5) [`Docker`](https://qiita.com/items/ae11fca7d2eba445b037) [`Git`](https://qiita.com/items/74eacdbf363e260981c3) [`Go`](https://qiita.com/items/49d4537d95f878b3e91a) [`iOS`](https://qiita.com/items/e61a29a383d0403e92fc) [`Java`](https://qiita.com/items/4c3f84836bfdbb137226) [`JavaScript`](https://qiita.com/items/eaa7ac5b62a0a723edbb) [`Linux`](https://qiita.com/items/362e81e53c3f9dee22f1) [`Node.js`](https://qiita.com/items/66ed7ad8f7c9673e9d50) [`PHP`](https://qiita.com/items/3318cbdbc45c6ebd4014) [`Python`](https://qiita.com/items/9d7f2ffeafb36cf59a77) [`Rails`](https://qiita.com/items/93b9e7f7d143e9ce650e) [`React`](https://qiita.com/items/f9712f8acace22815b99) [`Ruby`](https://qiita.com/items/72c3d2e896bdc3e1a6b3) [`Swift`](https://qiita.com/items/e2b6f0645e29f0e2b761) [`TypeScript`](https://qiita.com/items/25b7c0870afa6d41d19b) [`Vim`](https://qiita.com/items/f5361177baef95e447d1) [`Vue.js`](https://qiita.com/items/2774e02c6eea5c830d99) [`初心者`](https://qiita.com/items/402899ec543aff109505)",
		"",
		"# 集計について",
		"",
		"- 期間: 2020-01-01 ~ 2020-01-08",
		"- 条件: ストック数が **2** 以上の記事",
		"",
		"ソースコード:",
		`<a href="https://github.com/kou-pg-0131/qiita-lgtm-ranking"><img src="https://github-link-card.s3.ap-northeast-1.amazonaws.com/kou-pg-0131/qiita-lgtm-ranking.png" width="460px"></a>`,
		"",
		"# LGTM 数ランキング",
		"",
		"## 1 位: [TITLE_FUGA](https://example.com/FUGA)",
		"",
		"[`TAG_1`](https://qiita.com/tags/TAG_1) [`TAG_2`](https://qiita.com/tags/TAG_2)",
		"",
		"**13** LGTM",
		"[@USER_FUGA](https://qiita.com/USER_FUGA) さん ( 2020-01-07 04:36 に投稿 )",
		"",
		"## 2 位: [TITLE_HOGE](https://example.com/HOGE)",
		"",
		"[`TAG_1`](https://qiita.com/tags/TAG_1) [`TAG_2`](https://qiita.com/tags/TAG_2)",
		"",
		"**10** LGTM",
		"[@USER_HOGE](https://qiita.com/USER_HOGE) さん ( 2020-01-01 12:01 に投稿 )",
		"",
		"## 3 位: [TITLE_FOO](https://example.com/FOO)",
		"",
		"[`TAG_1`](https://qiita.com/tags/TAG_1) [`TAG_2`](https://qiita.com/tags/TAG_2)",
		"",
		"**8** LGTM",
		"[@USER_FOO](https://qiita.com/USER_FOO) さん ( 2020-01-03 08:28 に投稿 )",
		"",
	}, "\n"), body)
	assert.Nil(t, err)
}
