package api

import (
	"github.com/dotdancer/gogofly/service/dto"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
	}
}

// @Tag 用户管理
// @Summary 用户登录
// @Description 用户登录详情描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	m.OK(ResponseJson{
		Data: iUserLoginDTO,
	})

	//fmt.Println("Login 执行了")
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "Login Success",
	//})

	//OK(c, ResponseJson{
	//	Msg: "Login Success",
	//})

	//Fail(c, ResponseJson{
	//	Code: 9001,
	//	Msg:  "Login Failed",
	//})
}
