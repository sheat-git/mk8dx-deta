package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheat-git/mkdeta/service"
)

func SokujiGet(c *gin.Context) {
	uid := c.Query("user_id")
	if uid != "" {
		s, err := service.GetSokujiByUserId(uid)
		if err != nil {
			c.String(404, "user not found")
			return
		}
		c.JSON(200, s)
		return
	}
	cid := c.Query("channel_id")
	if cid != "" {
		s, err := service.GetSokujiByChannelId(cid)
		if err != nil {
			c.String(404, "channel not found")
			return
		}
		c.JSON(200, s)
		return
	}
	c.String(400, "Must provide user_id or channel_id")
}
