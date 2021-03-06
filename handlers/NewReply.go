package handlers

import (
	"../libs"
	"../models"
	"fmt"
	"strconv"
)

type NewReplyHandler struct {
	libs.BaseHandler
}

func (self *NewReplyHandler) Post() {
	inputs := self.Input()
	tid, _ := strconv.Atoi(inputs.Get("comment_parent"))
	sess_userid, _ := self.GetSession("userid").(int64)

	author := inputs.Get("author")
	email := inputs.Get("email")
	website := inputs.Get("website")

	rc := inputs.Get("comment")

	if author != "" && email != "" && tid != 0 && rc != "" {
		if err := models.AddReply(tid, int(sess_userid), rc, author, email, website); err != nil {
			fmt.Println(err)
		}
		self.Ctx.Redirect(302, "/view/"+inputs.Get("comment_parent"))
	} else {
		self.Ctx.Redirect(302, "/")
	}
}
