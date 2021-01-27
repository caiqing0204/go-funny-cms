package validates

import (
	"github.com/gin-gonic/gin"
	"gocms/app/validates/validate"
	"gocms/pkg/logger"
	"gocms/pkg/response"
	"net/http"
)

// 登陆需要的参数
type LoginParams struct {
	Account  string `validate:"required" json:"account"`
	Password string `validate:"required" json:"password"`
}

type LoginAction struct {
}

func (*LoginAction) Validate(c *gin.Context, params *LoginParams) bool {
	err := c.ShouldBind(&params)
	if err != nil {
		response.ErrorResponse(http.StatusUnauthorized, err.Error()).WriteTo(c)
		return false
	}

	// 自定义错误码和消息
	return validate.WithResponseMsg(params, c, "账号或者密码错误")
}

// --- 注册相关
type RegisterParams struct {
	Account  string `validate:"username" json:"account"`
	Password string `validate:"required" json:"password"`
	Email    string `validate:"email" json:"email"`
	Captcha  string `validate:"numeric,len=5" json:"captcha"`
}

func (that *RegisterParams) Validate(c *gin.Context, params interface{}) bool {
	err := c.BindJSON(params)
	if err != nil {
		logger.PanicError(err, "注册参数验证", false)
		return false
	}
	return validate.WithResponse(params, 403, "请检查参数", c)
}

type RegisterAction struct {
}

func (that *RegisterAction) Validate(c *gin.Context, params interface{}) bool {
	err := c.BindJSON(params)
	if err != nil {
		logger.PanicError(err, "注册参数验证", false)
		return false
	}
	return validate.WithResponse(params, 403, "请检查参数", c)
}

// --- 用户创建、更新相关
type EditParams struct {
	Account  string `validate:"username" json:"account"`
	Password string `validate:"required" json:"password"`
	Phone    string `validate:"required" json:"Phone"`
	Email    string `validate:"email" json:"email"`
}
