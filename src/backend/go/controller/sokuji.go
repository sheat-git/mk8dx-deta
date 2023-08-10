package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheat-git/mkdeta/service"
)

func SokujiGet(c *gin.Context) {
	var s map[string]interface{}
	var err error
	uid := c.Query("user_id")
	if uid != "" {
		s, err = service.GetSokujiByUserId(uid)
		if err != nil {
			c.String(404, "user not found")
			return
		}
	}
	cid := c.Query("channel_id")
	if cid != "" {
		s, err = service.GetSokujiByChannelId(cid)
		if err != nil {
			c.String(404, "channel not found")
			return
		}
	}
	if s == nil {
		c.String(404, "sokuji not found")
		return
	}
	delete(s, "key")
	c.JSON(200, s)
}
