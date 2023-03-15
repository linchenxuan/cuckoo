package service

import (
	"cuckoo/internal/domain/entity"
	"cuckoo/internal/domain/po"
	"cuckoo/internal/domain/repository"
	"cuckoo/pkg/cuckoo_error"
	"cuckoo/pkg/valobj"
)

type ArticleService struct {
	articleRepo  repository.IArticleRepository
	tagRepo      repository.ITagRepository
	categoryRepo repository.ICategoryRepository
}

func NewArticleService(articleRepo repository.IArticleRepository, tagRepo repository.ITagRepository, categoryRepo repository.ICategoryRepository) *ArticleService {
	return &ArticleService{
		articleRepo:  articleRepo,
		tagRepo:      tagRepo,
		categoryRepo: categoryRepo,
	}
}

func (srv *ArticleService) GetArticle(id int64) (*entity.Article, cuckoo_error.IErrorStatus) {
	pArticle, err := srv.articleRepo.Get(id)
	if err != nil {
		return nil, cuckoo_error.InternalError
	}
	return &entity.Article{Article: *pArticle}, nil
}

func (srv *ArticleService) GetArticleList(req valobj.GetArticleListReq) ([]*entity.Article, int64, cuckoo_error.IErrorStatus) {
	pArticles, total, err := srv.articleRepo.FindsByPage(req.PageNum, req.PageSize, req.CategoryId, req.TagId)
	if err != nil {
		return nil, 0, cuckoo_error.InternalError
	}
	articles := make([]*entity.Article, len(pArticles))
	for i, article := range pArticles {
		articles[i] = &entity.Article{Article: *article}
	}

	return articles, total, nil
}

func (srv *ArticleService) CreateArticle(uid int64, req valobj.CreateArticleReq) (*entity.Article, cuckoo_error.IErrorStatus) {
	tags := make([]int64, 0)
	for _, tagName := range req.Tags {
		tag, _ := srv.tagRepo.New(tagName)
		if tag != nil {
			tags = append(tags, tag.Id)
		}
	}

	category, _ := srv.categoryRepo.New(req.Category)
	if category == nil {
		return nil, cuckoo_error.InternalError
	}

	pArticle := po.Article{
		CategoryId:  category.Id,
		UserId:      uid,
		Title:       req.Title,
		Description: req.Desc,
		Content:     req.Content,
		CoverImage:  req.CoverImage,
		Type:        req.Type,
		Status:      req.Status,
		OriginalUrl: req.OriginalUrl,
		IsTop:       req.IsTop,
		IsDelete:    false,
		Tags:        tags,
	}

	if srv.articleRepo.New(&pArticle) != nil {
		return nil, cuckoo_error.InternalError
	}

	return &entity.Article{Article: pArticle}, nil
}
