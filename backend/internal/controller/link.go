package controller

import (
	"errors"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/model"
	"go-svc-tpl/utils/stacktrace"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

// Interface of link controller
type ILinkController interface {
	Create(*gin.Context, *dto.LinkCreateReq) (*dto.LinkCreateResp, error)
	Delete(*gin.Context, *dto.LinkDeleteReq) (*dto.LinkDeleteResp, error)
	GetInfo(*gin.Context, *dto.GetLinkInfoReq) (*dto.GetLinkInfoResp, error)
	UpdateInfo(*gin.Context, *dto.UpdateLinkInfoReq) (*dto.UpdateLinkInfoResp, error)
	GetList(*gin.Context, *dto.GetLinkListReq) (*dto.GetLinkListResp, error)
}

// check interface implementation
var _ ILinkController = (*LinkController)(nil)

// link controller
var NewLinkController = func() *LinkController {
	return &LinkController{}
}

type LinkController struct {
	// maybe some logic config to read from viper
	// or a service dependency
}

func GetUniqueShort(ctx *gin.Context) string {
	rand.Seed(time.Now().Unix())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var start_len = 4
	var end_len = len(letters)
	for i := start_len; i <= end_len; i++ {
		s := make([]rune, i)
		for j := range s {
			r := rand.Intn(end_len)
			s[j] = letters[r]
		}
		result := dao.DB(ctx).Table(model.LinkTable).Where("short = ?", string(s)).Limit(1).Find(&model.Link{})
		if result.RowsAffected == 0 {
			logrus.Debugf("get unique short success, short: %s", string(s))
			return string(s)
		} else {
			logrus.Debugf("get unique short failed, continuing to generate...")
			continue
		}
	}
	logrus.Fatal("get unique short failed, no more letters to generate")
	return ""
}

func (c *LinkController) Create(ctx *gin.Context, req *dto.LinkCreateReq) (*dto.LinkCreateResp, error) {
	db := dao.DB(ctx).Table(model.LinkTable)
	//pr, _ := json.Marshal(req)
	//logrus.Printf("%v", string(pr))
	if req.Short != "" {
		result := db.Where("short = ?", req.Short).Limit(1).Find(&model.Link{})
		if result.RowsAffected == 0 {
			// "short" not exists
			var link = model.Link{
				ID:        0,
				Active:    true,
				Comment:   req.Comment,
				Origin:    req.Origin,
				Short:     req.Short,
				StartTime: req.StartTime,
				EndTime:   req.EndTime,
			}
			var result *gorm.DB
			userID, err := ctx.Cookie("userID")
			logrus.Println(userID)
			if err != nil {
				logrus.Info("No Login")
				result = db.Omit("owner_id").Create(&link)
			} else {
				logrus.Info("Login")
				logrus.Println(userID)
				num, _ := strconv.Atoi(userID)
				link.OwnerID = uint(num)
				result = db.Create(&link)
			}
			if result.Error != nil {
				return nil, stacktrace.PropagateWithCode(result.Error, dto.InternalError, "Failed to create short link")
			}
			return &dto.LinkCreateResp{
				Origin:    link.Origin,
				Comment:   link.Comment,
				StartTime: null.TimeFromPtr(&link.StartTime.Time),
				EndTime:   null.TimeFromPtr(&link.EndTime.Time),
				Active:    link.Active,
				Short:     link.Short}, nil
		} else {
			// already exists
			return nil, stacktrace.PropagateWithCode(errors.New("ShortLinkExist"), dto.ErrShortLinkExist, "Short link already exists")
		}

	} else {
		// need to create a shortlink
		shortlink := GetUniqueShort(ctx)
		var link = model.Link{
			ID:        0,
			Active:    true,
			Comment:   req.Comment,
			Origin:    req.Origin,
			Short:     shortlink,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		}
		var result *gorm.DB
		userID, err := ctx.Cookie("userID")
		logrus.Println(userID)
		if err != nil {
			logrus.Info("No Login")
			result = db.Omit("owner_id").Create(&link)
		} else {
			logrus.Info("Login")
			logrus.Println(userID)
			num, _ := strconv.Atoi(userID)
			link.OwnerID = uint(num)
			result = db.Create(&link)
		}
		if result.Error != nil {
			return nil, stacktrace.PropagateWithCode(result.Error, dto.InternalError, "Failed to create short link")
		}
		return &dto.LinkCreateResp{
			Origin:    link.Origin,
			Comment:   link.Comment,
			StartTime: null.TimeFromPtr(&link.StartTime.Time),
			EndTime:   null.TimeFromPtr(&link.EndTime.Time),
			Active:    link.Active,
			Short:     link.Short}, nil
	}
}

func (c *LinkController) Delete(ctx *gin.Context, req *dto.LinkDeleteReq) (*dto.LinkDeleteResp, error) {
	userID := ctx.GetInt("userID")
	db := dao.DB(ctx).Table(model.LinkTable)
	result := db.Where("short = ? AND owner_id = ?", req.Short, userID).Delete(&model.Link{})
	if result.RowsAffected == 0 {
		return nil, stacktrace.PropagateWithCode(errors.New("ErrNoShortLinkOrUserNotMatch"), dto.ErrNoShortLink, "Short link not found or This link doesn't belong to current user")
	} else {
		// delete success
		return &dto.LinkDeleteResp{}, nil
	}
}

func (c *LinkController) GetInfo(ctx *gin.Context, req *dto.GetLinkInfoReq) (*dto.GetLinkInfoResp, error) {
	db := dao.DB(ctx).Table(model.LinkTable)
	var link model.Link
	req.Short = ctx.Query("short")
	if req.Short == "" {
		return nil, stacktrace.PropagateWithCode(errors.New("ParamNoShortLink"), dto.ErrNoShortLink, "Param No ShortLink")
	}
	result := db.Where("short = ?", req.Short).Limit(1).Find(&link)
	if result.RowsAffected == 0 {
		return nil, stacktrace.PropagateWithCode(errors.New("ErrNoShortLink"), dto.ErrNoShortLink, "Short link not found")
	} else {
		// exist
		return &dto.GetLinkInfoResp{
			Short:     link.Short,
			Origin:    link.Origin,
			Comment:   link.Comment,
			StartTime: link.StartTime,
			EndTime:   link.EndTime,
			Active:    link.Active,
		}, nil
	}
}

func (c *LinkController) UpdateInfo(ctx *gin.Context, req *dto.UpdateLinkInfoReq) (*dto.UpdateLinkInfoResp, error) {
	userID := ctx.GetInt("userID")
	db := dao.DB(ctx).Table(model.LinkTable)
	result := db.Where("short = ? AND owner_id = ?", req.Short, userID).Limit(1).Updates(
		model.Link{
			Comment:   req.Comment,
			Active:    req.Active,
			Origin:    req.Origin,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		})
	if result.RowsAffected == 0 {
		return nil, stacktrace.PropagateWithCode(errors.New("ErrNoShortLinkOrUserNotMatch"), dto.ErrNoShortLink, "Short link not found or This link doesn't belong to current user")
	} else {
		return &dto.UpdateLinkInfoResp{}, nil
	}
}

func (c *LinkController) GetList(ctx *gin.Context, req *dto.GetLinkListReq) (*dto.GetLinkListResp, error) {
	db := dao.DB(ctx).Table(model.LinkTable)
	userID := ctx.GetInt("userID")
	page_number, _ := strconv.Atoi(ctx.Query("page_number"))
	page_size, _ := strconv.Atoi(ctx.Query("page_size"))
	link_min := page_number * page_size
	link_max := link_min + page_size
	var link_count int
	var links []model.Link
	result := db.Where("owner_id = ?", userID).Find(&links)
	if result.RowsAffected == 0 {
		return nil, stacktrace.PropagateWithCode(errors.New("ErrUserNotFound"), dto.ErrUserNotFound, "User not found")
	} else {
		var resps []dto.GetLinkInfoResp
		for i, link := range links {
			if !(i >= link_min && i < link_max) && page_size != -1 {
				// out of the range we need
				continue
			}
			link_count++
			resps = append(resps, dto.GetLinkInfoResp{
				Short:     link.Short,
				Comment:   link.Comment,
				Active:    link.Active,
				Origin:    link.Origin,
				StartTime: link.StartTime,
				EndTime:   link.EndTime,
			})
		}
		return &dto.GetLinkListResp{
			Total: int64(link_count),
			Links: resps,
		}, nil
	}
}
