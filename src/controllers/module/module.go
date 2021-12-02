package module

import (
	"net/http"

	"github.com/0B1t322/CP-Rosseti-Back/ent"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submoduletest"
	e "github.com/0B1t322/CP-Rosseti-Back/models/err"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ModuleController struct {
	Client *ent.Client
}

func New(client *ent.Client) *ModuleController {
	return &ModuleController{
		Client: client,
	}
}

type CreateModuleReq struct {
	Name string `json:"name"`
}

type CreateModuleResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CreateModule
//
// @Tags module
//
// @Summary create module
//
// @Description create module
//
// @Description that can do only admin
//
// @Router /v1/module [post]
//
// @Security ApiKeyAuth
//
// @Param module body module.CreateModuleReq true "Module info"
//
// @Accept json
//
// @Produce json
//
// @Success 201 {object} module.CreateModuleResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) CreateModule(c *gin.Context) {
	var req CreateModuleReq
	{
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("Unexpected body"))
			c.Abort()
			return
		}
	}

	created, err := m.Client.Module.Create().SetName(req.Name).Save(c)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "CreateModule",
				"err":  err,
			},
		).Error("Failed create module")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to create module"))
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, CreateModuleResp{ID: created.ID, Name: created.Name})
}

type AddSubModuleReq struct {
	ModuleID int    `json:"-" uri:"id"`
	Name     string `json:"name"`
	Text     string `json:"text"`
}

type AddSubModuleResp struct {
	ID       int    `json:"id"`
	ModuleID int    `json:"moduleID"`
	Name     string `json:"name"`
	Text     string `json:"text"`
}

// AddSubModule
//
// @Tags module
//
// @Summary add submodule
//
// @Description add submodule
//
// @Description that can do only admin
//
// @Router /v1/module/{id}/submodule [post]
//
// @Security ApiKeyAuth
//
// @Param submodule body module.AddSubModuleReq true "SubModule info"
//
// @Param id path integer true "id of module"
//
// @Accept json
//
// @Produce json
//
// @Success 201 {object} module.AddSubModuleResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) AddSubModule(c *gin.Context) {
	var req AddSubModuleReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
			c.Abort()
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("Unexpected body"))
			c.Abort()
			return
		}
	}

	creater, err := m.Client.SubModule.Create().
		SetModuleID(req.ModuleID).
		SetName(req.Name).
		SetText(req.Text).
		Save(c)

	if ent.IsNotFound(err) || ent.IsConstraintError(err) {
		c.JSON(http.StatusNotFound, e.FromString("Module not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddSubModule",
				"err":  err,
			},
		).Error("Failed add submodule")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to add submodule"))
		c.Abort()
		return
	}

	c.JSON(
		http.StatusCreated,
		AddSubModuleResp{
			ID:       creater.ID,
			ModuleID: req.ModuleID,
			Name:     creater.Name,
			Text:     creater.Text,
		},
	)

}

type AddSubModuleTestReq struct {
	SubModuleID     int                       `json:"-" uri:"id"`
	TheoreticalTest *CreateTheoreticalTestReq `json:"theoreticalTest,omitempty"`
}

type CreateTheoreticalTestReq struct {
	Questions []CreateQuestionReq `json:"questions"`
}

type CreateQuestionReq struct {
	Question string            `json:"question"`
	Answers  []CreateAnswerReq `json:"answers"`
}

type CreateAnswerReq struct {
	Answer  string `json:"answer"`
	Correct bool   `json:"correct"`
}

type AddSubModuleTestResp struct {
	ID              int                     `json:"id"`
	TheoretcialTest *AddTheoreticalTestResp `json:"theoreticalTest,omitempty"`
}

type AddTheoreticalTestResp struct {
	Questions []AddQuestionResp `json:"questions"`
}

type AddQuestionResp struct {
	ID       int             `json:"id"`
	Question string          `json:"question"`
	Answers  []AddAnswerResp `json:"answers"`
}

type AddAnswerResp struct {
	ID      int    `json:"id"`
	Answer  string `json:"answer"`
	Correct bool   `json:"correct"`
}

// AddSubModuleTest
//
// @Tags module
//
// @Summary add submodule test
//
// @Description add submodule test
//
// @Description that can do only admin
//
// @Router /v1/module/submodule/{id}/test [post]
//
// @Security ApiKeyAuth
//
// @Param submoduletest body module.AddSubModuleTestReq true "SubModuleTest info"
//
// @Param id path integer true "id of submodule"
//
// @Accept json
//
// @Produce json
//
// @Success 201 {object} module.AddSubModuleTestResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) AddSubModuleTest(c *gin.Context) {
	var req AddSubModuleTestReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
			c.Abort()
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("Unexpected body"))
			c.Abort()
			return
		}
	}

	test, err := m.Client.SubModuleTest.Create().SetSubModuleID(req.SubModuleID).Save(c)
	if ent.IsConstraintError(err) {
		log.Info("err", err)
		c.JSON(http.StatusNotFound, e.FromString("SubModule not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddSubModuleTest",
				"err":  err,
			},
		).Error("Failed add submodule test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to add submodule test"))
		c.Abort()
		return
	}

	if req.TheoreticalTest != nil {
		theoTest, err := m.Client.TheoreticalTest.Create().SetSubModuleTest(test).Save(c)
		if err != nil {
			log.WithFields(
				log.Fields{
					"pkg":  "controllers/module",
					"func": "AddSubModuleTest",
					"err":  err,
				},
			).Error("Failed add submodule test")
			c.JSON(http.StatusInternalServerError, e.FromString("Failed to add submodule test"))
			c.Abort()
			return
		}

		for _, question := range req.TheoreticalTest.Questions {
			createdQuestion, err := m.Client.Question.Create().
				SetTheoreticalTest(theoTest).
				SetQuestion(question.Question).
				Save(c)
			if err != nil {
				log.WithFields(
					log.Fields{
						"pkg":  "controllers/module",
						"func": "AddSubModuleTest",
						"err":  err,
					},
				).Error("Failed add submodule test")
				c.JSON(http.StatusInternalServerError, e.FromString("Failed to add submodule test"))
				c.Abort()
				return
			}

			for _, answer := range question.Answers {
				_, err := m.Client.Answer.Create().
					SetAnswer(answer.Answer).
					SetCorrect(answer.Correct).
					SetQuestuion(createdQuestion).
					Save(c)
				if err != nil {
					log.WithFields(
						log.Fields{
							"pkg":  "controllers/module",
							"func": "AddSubModuleTest",
							"err":  err,
						},
					).Error("Failed add submodule test")
					c.JSON(http.StatusInternalServerError, e.FromString("Failed to add submodule test"))
					c.Abort()
					return
				}
			}
		}
	}

	getSubModuleTest, err := m.Client.SubModuleTest.
		Query().
		WithTherTest(
			func(ttq *ent.TheoreticalTestQuery) {
				ttq.WithQuestion(
					func(qq *ent.QuestionQuery) {
						qq.WithAnswer()
					},
				)
			},
		).
		Where(submoduletest.ID(test.ID)).
		Only(c)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddSubModuleTest",
				"err":  err,
			},
		).Error("Failed add submodule test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to add submodule test"))
		c.Abort()
		return
	}

	var resp AddSubModuleTestResp
	{
		resp.ID = getSubModuleTest.ID
		test := &AddTheoreticalTestResp{}
		
		var questions []AddQuestionResp
		{
			for _, question := range getSubModuleTest.Edges.TherTest.Edges.Question {
				var answers []AddAnswerResp
				{
					for _, answer := range question.Edges.Answer {
						answers = append(
							answers,
							AddAnswerResp{
								ID: answer.ID,
								Answer: answer.Answer,
								Correct: answer.Correct,
							},
						)
					}
				}

				questions = append(
					questions, 
					AddQuestionResp{
						ID: question.ID,
						Question: question.Question,
						Answers: answers,
					},
				)
			}
		}

		test.Questions = questions

		resp.TheoretcialTest = test
	}

	c.JSON(http.StatusCreated, resp)
}
