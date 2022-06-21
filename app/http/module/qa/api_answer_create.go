package qa

import (
	"github.com/karry-11110/LM_system/app/http/middleware/auth"
	provider "github.com/karry-11110/LM_system/app/provider/qa"
	"github.com/karry-11110/ginx/framework/gin"
)

type answerCreateParam struct {
	QuestionID int64  `json:"question_id" binding:"required"`
	Context    string `json:"context" binding:"required"`
}

// AnswerCreate 代表创建回答
// @Summary 创建回答
// @Description 创建回答
// @Accept  json
// @Produce  json
// @Tags qa
// @Param answerCreateParam body answerCreateParam true "创建回答参数"
// @Success 200 string Msg "操作成功"
// @Router /answer/create [post]
func (api *QAApi) AnswerCreate(c *gin.Context) {
	qaService := c.MustMake(provider.QaKey).(provider.Service)

	param := &answerCreateParam{}
	if err := c.ShouldBind(param); err != nil {
		c.ISetStatus(400).IText(err.Error())
		return
	}

	user := auth.GetAuthUser(c)
	if user == nil {
		c.ISetStatus(500).IText("请登录后再操作")
		return
	}

	answer := &provider.Answer{
		QuestionID: param.QuestionID,
		Context:    param.Context,
		AuthorID:   user.ID,
	}
	if err := qaService.PostAnswer(c, answer); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	c.ISetOkStatus().IText("操作成功")
}
