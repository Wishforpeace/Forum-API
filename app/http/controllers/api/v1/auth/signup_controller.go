package auth

import (
	v1 "Forum-API/app/http/controllers/api/v1"
	"Forum-API/app/models/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理用户身份认证相关逻辑

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) InPhoneExist(c *gin.Context) {

	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}

	request := PhoneExistRequest{}

	// 解析JSON请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回422状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}