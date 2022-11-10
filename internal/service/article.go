package service

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`

	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"required,min=3,max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=3,max=255"`
	Content       string `form:"content" binding:"required"`
	CreatedBy     string `form:"created_by" binding:"required, min=3, max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,get=1"`
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"required,min=3,max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=3,max=255"`
	Content       string `form:"content" binding:"required"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,get=1"`
}
