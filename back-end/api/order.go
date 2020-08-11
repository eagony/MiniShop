package api

import (
	"MiniShop/extensions"
	"MiniShop/middlewares"
	"MiniShop/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// OrderAPI 订单相关的API
type OrderAPI struct{}

// Register 注册路由
func (o *OrderAPI) Register(g *gin.RouterGroup) {
	g.GET("/orders/:id", o.GetOrder)
	g.POST("/orders", o.CreateOrder)

	// 需要验证
	g.GET("/orders", middlewares.AuthMiddleware(), o.GetAllOrders)
	g.PUT("/orders/:id", middlewares.AuthMiddleware(), o.UpdateOrder)
	g.DELETE("/orders/:id", middlewares.AuthMiddleware(), o.DeleteOrder)
}

// CreateOrder 新建一个订单
func (o *OrderAPI) CreateOrder(c *gin.Context) {
	order := models.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bind error": err.Error(),
		})
		return
	}
	if order.Detail == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Data error": "Invalid order, no detail.",
		})
		return
	}
	// fmt.Println(order)
	// 处理事务
	tx := extensions.GetDB().Begin()
	defer tx.Close()

	if err := tx.Create(&order).Error; err != nil {
		// 处理错误
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Create error": err.Error(),
		})
		return
	}
	tx.Commit()

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "订单创建成功！",
		"data":    order,
	})
}

// DeleteOrder 删除一个订单
func (o *OrderAPI) DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Param error": err.Error(),
		})
		return
	}

	orderToBeDeleted := models.Order{}

	// 处理事务
	tx := extensions.GetDB().Begin()
	defer tx.Close()

	// 查询
	if err := tx.First(&orderToBeDeleted, id).Error; err != nil {
		// 处理错误
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}

	// 执行删除
	if err = tx.Delete(&orderToBeDeleted).Error; err != nil {
		// 处理删除错误
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
		"data":    orderToBeDeleted,
	})
}

// UpdateOrder 更新一个订单
func (o *OrderAPI) UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Param error": err.Error(),
		})
		return
	}

	tx := extensions.GetDB().Begin()
	defer tx.Close()

	orderToBeUpdate := models.Order{}
	if err = tx.First(&orderToBeUpdate, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&orderToBeUpdate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Bing error": err.Error(),
		})
		return
	}

	tx.Save(&orderToBeUpdate)
	tx.Commit()

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "订单信息修改成功！",
		"data":    orderToBeUpdate,
	})
}

// GetAllOrders 返回所有订单
func (o *OrderAPI) GetAllOrders(c *gin.Context) {
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
			"Query error": perPageErr.Error(),
		})
		return
	}

	// 查询事务
	tx := extensions.GetDB().Begin()
	defer tx.Close()

	var data []models.Order
	tx.Limit(perPage).Offset((page - 1) * perPage).Order("created_at desc").Find(&data)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
		"total":  len(data),
	})
}

// GetOrder 返回单个订单
func (o *OrderAPI) GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Param error": err.Error(),
		})
		return
	}
	order := models.Order{}

	tx := extensions.GetDB().Begin()
	defer tx.Close()

	if err = tx.First(&order, id).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Query error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    order,
	})
}
