package api

import (
	"MiniShop/extensions"
	"MiniShop/middlewares"
	"MiniShop/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GoodAPI 商品相关的API
type GoodAPI struct{}

// Register 注册路由
func (g *GoodAPI) Register(e *gin.RouterGroup) {
	e.GET("/goods/:id", g.GetGood)
	e.GET("/goods", g.GetAllGoods)
	// 需要验证
	e.POST("/goods", middlewares.AuthMiddleware(), g.PublishGood)
	e.PUT("/goods/:id", middlewares.AuthMiddleware(), g.UpdateGood)
	e.DELETE("/goods/:id", middlewares.AuthMiddleware(), g.DeleteGood)
}

// PublishGood 发布一个新商品
func (g *GoodAPI) PublishGood(c *gin.Context) {
	good := models.Good{}
	if err := c.ShouldBindJSON(&good); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bind error": err.Error(),
		})
		return
	}
	// 处理事务
	tx := extensions.GetDB().Begin()
	defer tx.Close()

	if err := tx.Create(&good).Error; err != nil {
		// Handle error
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Create error": err.Error(),
		})
		return
	}
	tx.Commit()

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "商品发布成功！",
		"data":    good,
	})
}

// DeleteGood 删除一个商品
func (g *GoodAPI) DeleteGood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Param error": err.Error(),
		})
		return
	}

	goodToDeleted := models.Good{}

	// 处理事务
	tx := extensions.GetDB().Begin()
	defer tx.Close()

	// 找出要删除的商品
	if err = tx.First(&goodToDeleted, id).Error; err != nil {
		// 处理错误
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}

	// 执行删除
	if err = tx.Delete(&goodToDeleted).Error; err != nil {
		// 处理错误
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Delete error": err.Error(),
		})
		return
	}
	tx.Commit()

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "删除商品成功！",
		"data":    goodToDeleted,
	})
}

// UpdateGood 更新一个商品
func (g *GoodAPI) UpdateGood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Param error": err.Error(),
		})
		return
	}

	tx := extensions.GetDB().Begin()
	defer tx.Close()

	goodToBeUpdated := models.Good{}
	if err = tx.First(&goodToBeUpdated, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&goodToBeUpdated); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bind error": err.Error(),
		})
		return
	}

	tx.Save(&goodToBeUpdated)
	tx.Commit()

	// 处理成功
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "商品信息修改成功！",
		"data":    goodToBeUpdated,
	})
}

// GetAllGoods 返回所有商品列表
func (g *GoodAPI) GetAllGoods(c *gin.Context) {
	// 获取参数
	page, pageErr := strconv.Atoi(c.DefaultQuery("page", "1"))
	if pageErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Param error": pageErr.Error(),
		})
		return
	}

	perPage, perPageErr := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if perPageErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Param error": perPageErr.Error(),
		})
		return
	}

	// 查询事务
	tx := extensions.GetDB().Begin()
	defer tx.Close()

	var data []models.Good
	tx.Limit(perPage).Offset((page - 1) * perPage).Order("created_at desc").Find(&data)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
		"total":  len(data),
	})
}

// GetGood 返回单个商品信息
func (g *GoodAPI) GetGood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Param error": err.Error(),
		})
		return
	}
	good := models.Good{}

	// 处理事务
	tx := extensions.GetDB().Begin()
	defer tx.Close()

	if err = tx.First(&good, id).Error; err != nil {
		// 处理错误
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    good,
	})
}
