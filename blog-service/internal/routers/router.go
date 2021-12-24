package routers

import (
	_ "blog-service/docs"
	"blog-service/global"
	"blog-service/internal/middleware"
	"blog-service/internal/routers/api"
	v1 "blog-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	//url :=  ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	{
		article := v1.NewArticle()
		tag := v1.NewTag()

		// 创建标签
		apiv1.POST("/tags", tag.Create)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiv1.PUT("/tags/:id", tag.Update)
		// 获取标签列表
		apiv1.GET("/tags", tag.List)

		// 创建文章
		apiv1.POST("/articles", article.Create)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", article.Delete)
		// 更新指定文章
		apiv1.PUT("/articles/:id", article.Update)
		// 获取指定文章
		apiv1.GET("/articles/:id", article.Get)
		// 获取文章列表
		apiv1.GET("/articles", article.List)
	}

	return r
}
