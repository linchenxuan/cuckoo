package valobj

type ArticleType int8

const (
	Original ArticleType = iota
	Transshipment
	Translation
)

type ArticleStatus int8

const (
	PUBLIC ArticleStatus = iota // 公开
	SECRET                      // 私密
	DRAFT                       // 草稿
)

type GetArticleListReq struct {
	PageSize   int `json:"page_size"`
	PageNum    int `json:"page_num"`
	CategoryId int `json:"category_id"`
	TagId      int `json:"tag_id"`
}

type CreateArticleReq struct {
	Title       string        `json:"title" validate:"required"`
	Content     string        `json:"content" validate:"required"`
	Type        ArticleType   `json:"type" validate:"required"`
	Status      ArticleStatus `json:"status" validate:"required"`
	IsTop       bool          `json:"is_top" validate:"required"`
	Desc        string        `json:"desc"`
	CoverImage  string        `json:"cover_image"`
	OriginalUrl string        `json:"original_url"`
	Tags        []string      `json:"tags"`
	Category    string        `json:"category"`
}
