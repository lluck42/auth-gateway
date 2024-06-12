package admin

import (
	"crypto/md5"
	"errors"
	"fmt"
	"gq_volunteer_platform/model"
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Admin struct {
	ID             uint `gorm:"primarykey;->"`
	OrganizationId uint
	Name           string
	Mobile         string
	Password       string `json:"-"`
	Role           string
	RoleID         uint
	Token          string
	UserId         uint

	model.BaseTime
}

// 密码转为Md5处理
func (a *Admin) PasswordMd5() error {

	if s, err := Md5(a.Password); err != nil {
		return err
	} else {
		a.Password = s
	}

	return nil
}

// 生成 token 处理
func (a *Admin) NewToken() error {

	var str = strconv.Itoa(int(a.ID)) + time.Now().Format(time.DateTime)

	if s, err := Md5(str); err != nil {
		return err
	} else {
		a.Token = s
	}

	return nil
}

// 密码转为Md5处理
func Md5(s string) (string, error) {
	h := md5.New()
	if _, err := io.WriteString(h, s); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// gin.Context get user
func (one *Admin) Get(c *gin.Context) error {

	if a, exists := c.Get("admin"); !exists {
		return errors.New("ctx不存在该用户")
	} else {
		if u, b := a.(Admin); !b {
			return errors.New("断言error")
		} else {
			*one = u
		}
	}

	return nil
}
