package http

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/redis"
	"github.com/quangdangfit/gocommon/validation"
	"gorm.io/gorm"

	"goshop/internal/product/repository"
	"goshop/internal/product/service"
	"goshop/pkg/middleware"
)

func Routes(r *gin.RouterGroup, db *gorm.DB, validator validation.Validation, cache redis.IRedis) {
	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(validator, productRepo)
	productHandler := NewProductHandler(cache, productSvc)

	authMiddleware := middleware.JWTAuth()

	productRoute := r.Group("/products")
	{
		productRoute.GET("", productHandler.ListProducts)
		productRoute.POST("", authMiddleware, productHandler.CreateProduct)
		productRoute.PUT("/:id", authMiddleware, productHandler.UpdateProduct)
		productRoute.GET("/:id", productHandler.GetProductByID)
	}
}
