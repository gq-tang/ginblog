package pagination

import (
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type Paginator struct {
	Request     *http.Request
	PerPageNums int
	MaxPages    int

	nums      int64
	pageRange []int
	pageNums  int
	page      int
}

// PageNums returns the total number of pages.
func (p *Paginator) PageNums() int {
	if p.pageNums != 0 {
		return p.pageNums
	}
	pageNums := math.Ceil(float64(p.nums) / float64(p.PerPageNums))
	if p.MaxPages > 0 {
		pageNums = math.Min(pageNums, float64(p.MaxPages))
	}
	p.pageNums = int(pageNums)
	return p.pageNums
}

// Nums returns the total number of items
func (p *Paginator) Nums() int64 {
	return p.nums
}

// SetNums sets the total number of items
func (p *Paginator) SetNums(nums interface{}) {
	p.nums, _ = toInt64(nums)
}

// Page returns the current page.
func (p *Paginator) Page() int {
	if p.page != 0 {
		return p.page
	}
	if p.Request.Form == nil {
		p.Request.ParseForm()
	}
	p.page, _ = strconv.Atoi(p.Request.Form.Get("p"))
	if p.page > p.PageNums() {
		p.page = p.PageNums()
	}
	if p.page <= 0 {
		p.page = 1
	}
	return p.page
}

//
func (p *Paginator) Pages() []int {
	if p.pageRange == nil && p.nums > 0 {
		var pages []int
		pageNums := p.PageNums()
		page := p.Page()
		switch {
		case page >= pageNums-4 && pageNums > 9:
			start := pageNums - 9 + 1
			pages = make([]int, 9)
			for i := range pages {
				pages[i] = start + i
			}
		case page >= 5 && pageNums > 9:
			start := page - 5 + 1
			pages = make([]int, int(math.Min(9, float64(page+4+1))))
			for i := range pages {
				pages[i] = start + i
			}
		default:
			pages = make([]int, int(math.Min(9, float64(pageNums))))
			for i := range pages {
				pages[i] = i + 1
			}
		}
		p.pageRange = pages
	}
	return p.pageRange
}

// PageLink returns url for a given page index.
func (p *Paginator) PageLink(page int) string {
	link, _ := url.ParseRequestURI(p.Request.URL.String())
	values := link.Query()
	if page == 1 {
		values.Del("p")
	} else {
		values.Set("p", strconv.Itoa(page))
	}
	link.RawQuery = values.Encode()
	return link.String()
}

// PageLinkPrev return url to the previous page.
func (p *Paginator) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page() - 1)
	}
	return
}

// PageLinkNext return url to the next page.
func (p *Paginator) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page() + 1)
	}
	return link
}

// PageLinkFirst return url to the first page.
func (p *Paginator) PageLinkFirst() string {
	return p.PageLink(1)
}

// PageLinkLast return url to the last page.
func (p *Paginator) PageLinkLast() string {
	return p.PageLink(p.PageNums())
}

// HasPrev returns true if the current page has a predecessor.
func (p *Paginator) HasPrev() bool {
	return p.Page() > 1
}

// HasNext returns true if the current page has a successor.
func (p *Paginator) HasNext() bool {
	return p.Page() < p.PageNums()
}

// IsActive returns true if the given page index points to the current page.
func (p *Paginator) IsActive(page int) bool {
	return p.Page() == page
}

// Offset returns the current offset.
func (p *Paginator) Offset() int {
	return (p.Page() - 1) * p.PerPageNums
}

// HasPages returns true if there is more than one page.
func (p *Paginator) HasPages() bool {
	return p.PageNums() > 1
}

// NewPaginator Instantiates a paginator struct for the current http request.
func NewPaginator(req *http.Request, per int, nums interface{}) *Paginator {
	p := Paginator{}
	p.Request = req
	if per <= 0 {
		per = 10
	}
	p.PerPageNums = per
	p.SetNums(nums)
	return &p
}
