package response

type Pagination struct {
	Current int64       `json:"current"` // 当前页码
	Data    interface{} `json:"data"`    // 分页数据
	Total   int64       `json:"total"`   // 总数据量
	Pages   int64       `json:"pages"`   // 总分页数
	Limit   int64       `json:"limit"`   // 每页条数
}

/**
 * 提供了一个构造方法
 * 需要传入当前页码和每页条数
 */
func NewPage(current, size int64) *Pagination {
	return &Pagination{Current: current, Limit: size}
}

/**
 * 设置总数据量
 */
func (p *Pagination) SetTotal(t int64) {
	p.Total = t

	// 总分页数 = 总数据量 / 每页条数
	s := p.GetTotal() / p.GetSize()

	// 如果总数据量 / 每页条数有余数，则总分页数 + 1
	if p.GetTotal()%p.GetSize() != 0 {
		s = s + 1
	}

	p.Pages = s
}

/**
 * 设置分页数据
 */
func (p *Pagination) SetRecords(data interface{}) {
	p.Data = data
}

/**
 * 获取当前页码
 * 默认当前页码 1
 */
func (p *Pagination) GetCurrent() int64 {
	if p.Current < 1 {
		p.Current = 1
	}
	return p.Current
}

/**
 * 获取数据偏移量
 */
func (p *Pagination) GetOffset() int64 {
	if p.GetCurrent() > 0 {
		return (p.GetCurrent() - 1) * p.GetSize()
	}
	return 0
}

/**
 * 获取每页条数
 * 默认每页 10 条数据
 */
func (p *Pagination) GetSize() int64 {
	if p.Limit < 1 {
		p.Limit = 10
	}
	return p.Limit
}

/**
 * 获取总数据量
 */
func (p *Pagination) GetTotal() int64 {
	return p.Total
}

/**
 * 获取总分页数
 */
func (p *Pagination) GetPages() int64 {
	return p.Pages
}

/**
 * 获取分页数据
 */
func (page *Pagination) GetRecords() interface{} {
	return page.Data
}
