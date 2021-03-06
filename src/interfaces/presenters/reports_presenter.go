package presenters

import (
	"fmt"
	"strings"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
)

// IReportsPresenter .
type IReportsPresenter interface {
	Weekly(from time.Time, items *domain.Items) (string, error)
	WeeklyByTag(from time.Time, items *domain.Items, tag string) (string, error)
}

// ReportsPresenter .
type ReportsPresenter struct{}

// NewReportsPresenter .
func NewReportsPresenter() *ReportsPresenter {
	return new(ReportsPresenter)
}

// Weekly .
func (p *ReportsPresenter) Weekly(from time.Time, items *domain.Items) (string, error) {
	rows := []string{}

	rows = append(rows, "# タグ別 LGTM 数ランキング")
	rows = append(rows, "")
	rows = append(rows, p.tagsMarkdown())
	rows = append(rows, "")

	rows = append(rows, "# 集計について")
	rows = append(rows, "")
	rows = append(rows, p.aboutAggregateMarkdown(from, from.AddDate(0, 0, 7), 10))
	rows = append(rows, "")

	rows = append(rows, "# LGTM 数ランキング")
	rows = append(rows, "")
	rows = append(rows, p.itemsToMarkdown(items))

	return strings.Join(rows, "\n"), nil
}

// WeeklyByTag .
func (p *ReportsPresenter) WeeklyByTag(from time.Time, items *domain.Items, tag string) (string, error) {
	rows := []string{}

	rows = append(rows, "# 他のタグ")
	rows = append(rows, "")
	rows = append(rows, p.tagsMarkdown())
	rows = append(rows, "")

	rows = append(rows, "# 集計について")
	rows = append(rows, "")
	rows = append(rows, p.aboutAggregateMarkdown(from, from.AddDate(0, 0, 7), 2))
	rows = append(rows, "")

	rows = append(rows, "# LGTM 数ランキング")
	rows = append(rows, "")
	rows = append(rows, p.itemsToMarkdown(items))

	return strings.Join(rows, "\n"), nil
}

func (p *ReportsPresenter) aboutAggregateMarkdown(from, to time.Time, minStocks int) string {
	return fmt.Sprintf(`- 期間: %s ~ %s
- 条件: ストック数が **%d** 以上の記事

ソースコード:
<a href="https://github.com/kou-pg-0131/qiita-lgtm-ranking"><img src="https://github-link-card.s3.ap-northeast-1.amazonaws.com/kou-pg-0131/qiita-lgtm-ranking.png" width="460px"></a>`,
		from.Format("2006-01-02"), to.Format("2006-01-02"), minStocks,
	)
}

func (p *ReportsPresenter) tagsMarkdown() string {
	return strings.Join([]string{
		"[`AWS`](https://qiita.com/items/e24b6279326a462d456c)",
		"[`Android`](https://qiita.com/items/8b3af051428d746f26c5)",
		"[`Docker`](https://qiita.com/items/ae11fca7d2eba445b037)",
		"[`Git`](https://qiita.com/items/74eacdbf363e260981c3)",
		"[`Go`](https://qiita.com/items/49d4537d95f878b3e91a)",
		"[`iOS`](https://qiita.com/items/e61a29a383d0403e92fc)",
		"[`Java`](https://qiita.com/items/4c3f84836bfdbb137226)",
		"[`JavaScript`](https://qiita.com/items/eaa7ac5b62a0a723edbb)",
		"[`Linux`](https://qiita.com/items/362e81e53c3f9dee22f1)",
		"[`Node.js`](https://qiita.com/items/66ed7ad8f7c9673e9d50)",
		"[`PHP`](https://qiita.com/items/3318cbdbc45c6ebd4014)",
		"[`Python`](https://qiita.com/items/9d7f2ffeafb36cf59a77)",
		"[`Rails`](https://qiita.com/items/93b9e7f7d143e9ce650e)",
		"[`React`](https://qiita.com/items/f9712f8acace22815b99)",
		"[`Ruby`](https://qiita.com/items/72c3d2e896bdc3e1a6b3)",
		"[`Swift`](https://qiita.com/items/e2b6f0645e29f0e2b761)",
		"[`TypeScript`](https://qiita.com/items/25b7c0870afa6d41d19b)",
		"[`Vim`](https://qiita.com/items/f5361177baef95e447d1)",
		"[`Vue.js`](https://qiita.com/items/2774e02c6eea5c830d99)",
		"[`初心者`](https://qiita.com/items/402899ec543aff109505)",
	}, " ")
}

func (p *ReportsPresenter) itemsToMarkdown(items *domain.Items) string {
	rows := []string{}

	items.SortByLikesCount()
	for i, item := range *items.HasLGTM() {
		rows = append(rows, fmt.Sprintf("## %d 位: [%s](%s)", i+1, item.Title, item.URL))
		rows = append(rows, "")

		tags := []string{}
		for _, tag := range item.Tags {
			tags = append(tags, fmt.Sprintf("[`%s`](https://qiita.com/tags/%s)", tag.Name, tag.Name))
		}
		rows = append(rows, strings.Join(tags, " "))
		rows = append(rows, "")

		rows = append(rows, fmt.Sprintf("**%d** LGTM", item.LikesCount))
		rows = append(rows, fmt.Sprintf("[@%s](https://qiita.com/%s) さん ( %s に投稿 )", item.User.ID, item.User.ID, item.CreatedAt.Format("2006-01-02 15:04")))
		rows = append(rows, "")
	}

	return strings.Join(rows, "\n")
}
