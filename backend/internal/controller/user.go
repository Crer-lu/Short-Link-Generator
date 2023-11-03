package controller

import (
	"encoding/json"
	"errors"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/model"
	"go-svc-tpl/utils/stacktrace"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/dchest/captcha"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var token_mapping = map[int]string{}

// Interface of link controller
type IUserController interface {
	Register(*gin.Context, *dto.UserRegisterReq) (*dto.UserRegisterResp, error)
	Captcha(*gin.Context, *dto.UserCaptchaReq) (*dto.UserCaptchaResp, error)
	Login(*gin.Context, *dto.UserLoginReq) (*dto.UserLoginResp, error)
	Logout(*gin.Context, *dto.UserLogoutReq) (*dto.UserLogoutResp, error)
	GetInfo(*gin.Context, *dto.UserGetInfoReq) (*dto.UserGetInfoResp, error)
	UpdateInfo(*gin.Context, *dto.UserUpdateInfoReq) (*dto.UserUpdateInfoResp, error)
	UpdatePasswd(*gin.Context, *dto.UserUpdatePasswdReq) (*dto.UserUpdatePasswdResp, error)
}

// check interface implementation
var _ IUserController = (*UserController)(nil)

// link controller
var NewUserController = func() *UserController {
	return &UserController{}
}

type UserController struct {
	// maybe some logic config to read from viper
	// or a service dependency
}

// 邮箱正则表达式
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// 判断字符串是否符合邮箱格式
func isEmailValid(email string) bool {
	return emailRegex.MatchString(email)
}

// Register user
func (c *UserController) Register(ctx *gin.Context, req *dto.UserRegisterReq) (*dto.UserRegisterResp, error) {
	db := dao.DB(ctx).Table(model.UserTable)
	// pr, _ := json.Marshal(req)
	// logrus.Printf("%v", string(pr))
	if isEmailValid(req.Email) {
		user := &model.User{
			Email:    req.Email,
			Name:     req.Name,
			Password: req.Password,
		}
		result := db.Create(&user)
		if result.Error != nil {
			return nil, stacktrace.PropagateWithCode(result.Error, dto.ErrEmailExist, "Error Email Exist")
		} else {
			return &dto.UserRegisterResp{}, nil
		}
	} else {
		return nil, stacktrace.PropagateWithCode(errors.New("ErrorEmailFormat"), dto.BadReqeust, "Error Email Format")
	}
}

// Captcha user
func (c *UserController) Captcha(ctx *gin.Context, req *dto.UserCaptchaReq) (*dto.UserCaptchaResp, error) {
	capt := ctx.Query("captcha_id")
	if capt == "" {
		// create new captcha
		capt := captcha.NewLen(4)
		return &dto.UserCaptchaResp{
			CAPTCHAID:  capt,
			CAPTCHAURL: ctx.Request.RequestURI + "?captcha_id=" + capt,
		}, nil
	} else {
		err := captcha.WriteImage(ctx.Writer, capt, 90, 50)
		if err != nil {
			return nil, stacktrace.PropagateWithCode(err, dto.BadReqeust, "Captcha ID Invalid")
		} else {
			return nil, nil
		}
	}

}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(str_len int) string {
	rand.Seed(time.Now().UnixNano())
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

// Login user
func (c *UserController) Login(ctx *gin.Context, req *dto.UserLoginReq) (*dto.UserLoginResp, error) {
	db := dao.DB(ctx).Table(model.UserTable)
	pr, _ := json.Marshal(req)
	logrus.Printf("%v", string(pr))
	var user model.User
	result := db.Where("email = ?", req.Email).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		return nil, stacktrace.PropagateWithCode(errors.New("EmailNotExist"), dto.ErrUserNotFound, "Error Email Not Exist")
	} else {
		if user.Password != req.Password {
			return nil, stacktrace.PropagateWithCode(errors.New("PasswordNotMatch"), dto.ErrPassword, "Error Password Not Match")
		} else {
			// verify captcha
			ver_ans := captcha.VerifyString(req.CAPTCHAID, req.CAPTCHAValue)
			if !ver_ans {
				return nil, stacktrace.PropagateWithCode(errors.New("CaptchaNotMatch"), dto.ErrCaptcha, "Error Captcha Not Match")
			}
			// password correct -> set cookie & return
			token_mapping[int(user.ID)] = randStr(10) // todo
			ctx.SetCookie("token", token_mapping[int(user.ID)], 3600, "/", "localhost", false, true)
			ctx.SetCookie("userID", strconv.Itoa(int(user.ID)), 3600, "/", "localhost", false, true)
			return &dto.UserLoginResp{}, nil
		}

	}
}

// Logout user
func (c *UserController) Logout(ctx *gin.Context, req *dto.UserLogoutReq) (*dto.UserLogoutResp, error) {
	// delete cookie
	ctx.SetCookie("userID", strconv.Itoa(ctx.GetInt("userID")), -1, "/", "localhost", false, true)
	ctx.SetCookie("token", ctx.GetString("token"), -1, "/", "localhost", false, true)
	delete(token_mapping, ctx.GetInt("userID"))
	return &dto.UserLogoutResp{}, nil
}

// GetInfo user
func (c *UserController) GetInfo(ctx *gin.Context, req *dto.UserGetInfoReq) (*dto.UserGetInfoResp, error) {
	db := dao.DB(ctx).Table(model.UserTable)
	var userID = ctx.GetInt("userID")
	var user model.User
	result := db.Where("id = ?", userID).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		return nil, stacktrace.PropagateWithCode(errors.New("ErrorUserNotFound"), dto.ErrUserNotFound, "Error User Not Found")
	} else {
		return &dto.UserGetInfoResp{
			ID:    int64(user.ID),
			Name:  user.Name,
			Email: user.Email,
		}, nil
	}
}

// UpdateInfo user
func (c *UserController) UpdateInfo(ctx *gin.Context, req *dto.UserUpdateInfoReq) (*dto.UserUpdateInfoResp, error) {
	if !isEmailValid(req.Email) {
		return nil, stacktrace.PropagateWithCode(errors.New("ErrorEmailInvalid"), dto.BadReqeust, "Error Email Invalid")
	}
	db := dao.DB(ctx).Table(model.UserTable)
	result := db.Where("id = ?", req.ID).Limit(1).Updates(model.User{
		Name:  req.Name,
		Email: req.Email,
	})
	if result.RowsAffected == 0 {
		return nil, stacktrace.PropagateWithCode(errors.New("ErrorUserNotFound"), dto.ErrUserNotFound, "Error User Not Found")
	} else {
		return &dto.UserUpdateInfoResp{}, nil
	}
}

// UpdatePwd user
func (c *UserController) UpdatePasswd(ctx *gin.Context, req *dto.UserUpdatePasswdReq) (*dto.UserUpdatePasswdResp, error) {
	db := dao.DB(ctx).Table(model.UserTable)
	var user model.User
	var userID = ctx.GetInt("userID")
	result := db.Where("id = ?", userID).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		return nil, stacktrace.PropagateWithCode(errors.New("ErrorNotFound"), dto.ErrUserNotFound, "Error User Not Found")
	} else {
		if user.Password != req.OldPwd {
			return nil, stacktrace.PropagateWithCode(errors.New("PasswordIncorrect"), dto.ErrPassword, "Error Password Incorrect")
		} else {
			result := db.Where("id = ?", user.ID).Limit(1).Updates(&model.User{
				Password: req.NewPwd,
			})
			if result.Error != nil {
				return nil, stacktrace.PropagateWithCode(result.Error, dto.InternalError, "Update Password Failed")
			} else {
				return &dto.UserUpdatePasswdResp{}, nil
			}
		}
	}

}
