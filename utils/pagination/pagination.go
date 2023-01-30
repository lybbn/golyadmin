package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

const (
	page_size                = 10
	max_page_size            = 9999
	page_size_query_param    = "limit"
	current_page_query_param = "page"
)

// 标准分页结构体，接收最原始的DO
// 建议在外部再建一个字段一样的结构体，用以将DO转换成DTO或VO
type Page[T any] struct {
	CurrentPage int64 `json:"page"`  // 当前页码
	PageSize    int64 `json:"limit"` // 每页条数
	Total       int64 `json:"total"` // 总数据量
	Pages       int64 `json:"pages"` // 总分页数
	Data        []T   `json:"data"`  // 分页数据
}

// 各种查询条件先在query设置好后再放进来
func (page *Page[T]) PaginateQuery(query *gorm.DB, c *gin.Context) (e error) {
	e = nil
	current_page, _ := strconv.ParseInt(c.Query(current_page_query_param), 10, 64)
	page_size, _ := strconv.ParseInt(c.Query(page_size_query_param), 10, 64)
	page.CurrentPage = current_page
	page.PageSize = page_size
	var model T
	query.Model(&model).Count(&page.Total)
	if page.Total == 0 {
		page.Data = []T{}
		return
	}
	e = query.Model(&model).Scopes(Paginate(page)).Find(&page.Data).Error
	return
}

func Paginate[T any](page *Page[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.CurrentPage <= 0 {
			page.CurrentPage = 1
		}
		switch {
		case page.PageSize > max_page_size:
			page.PageSize = max_page_size // 限制一下分页大小
		case page.PageSize <= 0:
			page.PageSize = page_size
		}
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		p := page.CurrentPage
		if page.CurrentPage > page.Pages {
			// p = page.Pages
			return db.Limit(0)
		}
		size := page.PageSize
		offset := int((p - 1) * size)
		return db.Offset(offset).Limit(int(size))
	}
}
