package module

import (
	"net/http"

	"github.com/0B1t322/CP-Rosseti-Back/ent"
	"github.com/0B1t322/CP-Rosseti-Back/ent/practtest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/question"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submodule"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submoduletest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/theoreticaltest"
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
	PractTest       *CreatePractTestReq       `json:"practTest,omitempty"`
}

type CreatePractTestReq struct {
	Config map[string]interface{} `json:"config"`
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
	PractTest       *AddPractTestResp       `json:"practTest,omitempty"`
}

type AddTheoreticalTestResp struct {
	Questions []AddQuestionResp `json:"questions"`
}

type AddPractTestResp struct {
	ID     int                    `json:"id"`
	Config map[string]interface{} `json:"config"`
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

	if req.PractTest != nil {
		_, err := m.Client.PractTest.Create().
			SetSubModuleTest(test).
			SetConfig(req.PractTest.Config).
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
		WithPractTest().
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
		var test *AddTheoreticalTestResp

		if getSubModuleTest.Edges.TherTest != nil {
			var questions []AddQuestionResp
			{
				test = &AddTheoreticalTestResp{}
				for _, question := range getSubModuleTest.Edges.TherTest.Edges.Question {
					var answers []AddAnswerResp
					{
						for _, answer := range question.Edges.Answer {
							answers = append(
								answers,
								AddAnswerResp{
									ID:      answer.ID,
									Answer:  answer.Answer,
									Correct: answer.Correct,
								},
							)
						}
					}

					questions = append(
						questions,
						AddQuestionResp{
							ID:       question.ID,
							Question: question.Question,
							Answers:  answers,
						},
					)
				}
			}

			test.Questions = questions
		}

		resp.TheoretcialTest = test

		var practTest *AddPractTestResp
		{
			if getSubModuleTest.Edges.PractTest != nil {
				practTest = &AddPractTestResp{
					ID:     getSubModuleTest.Edges.PractTest.ID,
					Config: getSubModuleTest.Edges.PractTest.Config,
				}
			}
		}

		resp.PractTest = practTest
	}

	c.JSON(http.StatusCreated, resp)
}

type AddTheorTestReq struct {
	SubModuleID              int `json:"-" uri:"id"`
	CreateTheoreticalTestReq `json:",inline"`
}

type AddTheorTestResp struct {
	AddTheoreticalTestResp `json:",inline"`
}

// AddTheorTest
//
// @Tags module
//
// @Summary add theor test
//
// @Description add theor test to submodule test
//
// @Description that can do only admin
//
// @Router /v1/module/submodule/{id}/test/theor [post]
//
// @Security ApiKeyAuth
//
// @Param theotest body module.AddTheorTestReq true "Theor test info"
//
// @Param id path integer true "id of submodule"
//
// @Accept json
//
// @Produce json
//
// @Success 201 {object} module.AddTheorTestResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) AddTheorTest(c *gin.Context) {
	var req AddTheorTestReq
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

	subModuleTest, err := m.Client.SubModule.Query().
		Where(submodule.ID(req.SubModuleID), submodule.HasTest()).
		WithTest().
		Only(c)
	if ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("SubModule test or submodule not exist"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddTheorTest",
				"err":  err,
			},
		).Error("Failed add submodule theor test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed add submodule theor test"))
		c.Abort()
		return
	}

	theoTest, err := m.Client.TheoreticalTest.Create().
		SetSubModuleTestID(subModuleTest.Edges.Test.ID).
		Save(c)
	if ent.IsConstraintError(err) {
		c.JSON(http.StatusBadRequest, e.FromString("Theor test is exist"))
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

	for _, question := range req.CreateTheoreticalTestReq.Questions {
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

	theorTest, err := m.Client.TheoreticalTest.Query().
		WithQuestion(
			func(qq *ent.QuestionQuery) {
				qq.WithAnswer()
			},
		).
		Where(theoreticaltest.ID(theoTest.ID)).
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

	var resp AddTheorTestResp
	{
		theoTest := AddTheoreticalTestResp{}

		var questions []AddQuestionResp
		{
			for _, question := range theorTest.Edges.Question {
				var answers []AddAnswerResp
				{
					for _, answer := range question.Edges.Answer {
						answers = append(
							answers,
							AddAnswerResp{
								ID:      answer.ID,
								Answer:  answer.Answer,
								Correct: answer.Correct,
							},
						)
					}
				}

				questions = append(
					questions,
					AddQuestionResp{
						ID:       question.ID,
						Question: question.Question,
						Answers:  answers,
					},
				)
			}
		}
		theoTest.Questions = questions

		resp.AddTheoreticalTestResp = theoTest
	}

	c.JSON(http.StatusCreated, resp)
}

type AddPractTestReq struct {
	SubModuleID        int `json:"-" uri:"id"`
	CreatePractTestReq `json:",inline"`
}

// AddPractTest
//
// @Tags module
//
// @Summary add Pract test
//
// @Description add Pract test to submodule test
//
// @Description that can do only admin
//
// @Router /v1/module/submodule/{id}/test/pract [post]
//
// @Security ApiKeyAuth
//
// @Param theotest body module.AddPractTestReq true "Pract test info"
//
// @Param id path integer true "id of submodule"
//
// @Accept json
//
// @Produce json
//
// @Success 201 {object} module.AddPractTestResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) AddPractTest(c *gin.Context) {
	var req AddPractTestReq
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

	subModuleTest, err := m.Client.SubModule.Query().
		Where(submodule.ID(req.SubModuleID), submodule.HasTest()).
		WithTest().
		Only(c)
	if ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("SubModule test or submodule not exist"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddTheorTest",
				"err":  err,
			},
		).Error("Failed add submodule theor test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed add submodule theor test"))
		c.Abort()
		return
	}

	practTest, err := m.Client.PractTest.Create().
		SetSubModuleTestID(subModuleTest.Edges.Test.ID).
		SetConfig(req.Config).
		Save(c)
	if ent.IsConstraintError(err) {
		c.JSON(http.StatusBadRequest, e.FromString("Theor test is exist"))
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

	c.JSON(http.StatusCreated, AddPractTestResp{ID: practTest.ID, Config: practTest.Config})
}

type UpdateConfigReq struct {
	SubModuleID int                    `json:"-" uri:"id"`
	Config      map[string]interface{} `json:"config"`
}


// UpdateConfigToPractTest
//
// @Tags module
//
// @Summary update Pract test
//
// @Description update Pract test to submodule test
//
// @Description that can do only admin
//
// @Router /v1/module/submodule/{id}/test/pract [put]
//
// @Security ApiKeyAuth
//
// @Param theotest body module.UpdateConfigReq true "Pract test info"
//
// @Param id path integer true "id of submodule"
//
// @Accept json
//
// @Produce json
//
// @Success 200 {object} module.AddPractTestResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) UpdateConfigToPractTest(c *gin.Context) {
	var req UpdateConfigReq
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

	practTestID, err := m.Client.PractTest.Query().Where(
		practtest.HasSubModuleTestWith(
			submoduletest.HasSubModuleWith(
				submodule.ID(req.SubModuleID),
			),
		),
	).OnlyID(c)

	if ent.IsNotFound(err) || ent.IsConstraintError(err) {
		c.JSON(http.StatusNotFound, e.FromString("Pract test not found"))
		c.Abort()
		return
	}

	if err := m.Client.PractTest.UpdateOneID(practTestID).SetConfig(req.Config).Exec(c); err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "UpdateConfigToPractTest",
				"err":  err,
			},
		).Error("Failed add submodule test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to update submodule pract test"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, AddPractTestResp{ID: practTestID, Config: req.Config})
}

type UpdateTheorTestReq struct {
	SubModuleID		int		`json:"-" uri:"id"`
	Questions []CreateQuestionReq `json:"questions"`
}

// UpdateTheorTest
//
// @Tags module
//
// @Summary add questions to Theor test
//
// @Description add questions to Theor test
//
// @Router /v1/module/submodule/{id}/test/theor [put]
//
// @Security ApiKeyAuth
//
// @Param theortest body module.UpdateTheorTestReq true "Theor test info"
//
// @Param id path integer true "id of submodule"
//
// @Accept json
//
// @Produce json
//
// @Success 200 {object} module.AddTheorTestResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) UpdateTheorTest(c *gin.Context) {
	var req UpdateTheorTestReq
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

	theorTestID, err := m.Client.TheoreticalTest.Query().Where(
		theoreticaltest.HasSubModuleTestWith(
			submoduletest.HasSubModuleWith(
				submodule.ID(req.SubModuleID),
			),
		),
	).OnlyID(c)
	if ent.IsNotFound(err) || ent.IsConstraintError(err) {
		c.JSON(http.StatusNotFound, e.FromString("Theor test not found"))
		c.Abort()
		return
	}

	for _, question := range req.Questions {
		createdQuestion, err := m.Client.Question.Create().
			SetTheoreticalTestID(theorTestID).
			SetQuestion(question.Question).
			Save(c)
		if err != nil {
			log.WithFields(
				log.Fields{
					"pkg":  "controllers/module",
					"func": "UpdateTheorTest",
					"err":  err,
				},
			).Error("Failed update submodule theor test")
			c.JSON(http.StatusInternalServerError, e.FromString("Failed update submodule theor test"))
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
						"func": "UpdateTheorTest",
						"err":  err,
					},
				).Error("Failed update submodule theor test")
				c.JSON(http.StatusInternalServerError, e.FromString("Failed update submodule theor test"))
				c.Abort()
				return
			}
		}
	}

	theorTest, err := m.Client.TheoreticalTest.Query().
						WithQuestion(
							func(qq *ent.QuestionQuery) {
								qq.WithAnswer()
							},
						).Where(theoreticaltest.ID(theorTestID)).Only(c)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "UpdateTheorTest",
				"err":  err,
			},
		).Error("Failed update submodule theor test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed update submodule theor test"))
		c.Abort()
		return
	}

	var resp AddTheorTestResp
	{
		theoTest := AddTheoreticalTestResp{}

		var questions []AddQuestionResp
		{
			for _, question := range theorTest.Edges.Question {
				var answers []AddAnswerResp
				{
					for _, answer := range question.Edges.Answer {
						answers = append(
							answers,
							AddAnswerResp{
								ID:      answer.ID,
								Answer:  answer.Answer,
								Correct: answer.Correct,
							},
						)
					}
				}

				questions = append(
					questions,
					AddQuestionResp{
						ID:       question.ID,
						Question: question.Question,
						Answers:  answers,
					},
				)
			}
		}
		theoTest.Questions = questions

		resp.AddTheoreticalTestResp = theoTest
	}

	c.JSON(http.StatusOK, resp)
}

type AddAnswerReq struct {
	QuestionID		int					`json:"-" uri:"id"`
	Answers			[]CreateAnswerReq	`json:"answers"`
}

// AddAnsewerToQuestion
//
// @Tags module
//
// @Summary add questions to Theor test
//
// @Description add questions to Theor test
//
// @Router /v1/module/submodule/test/theor/question/:id [put]
//
// @Security ApiKeyAuth
//
// @Param answer body module.AddAnswerReq true "answer info"
//
// @Param id path integer true "id of question"
//
// @Accept json
//
// @Produce json
//
// @Success 200 {object} module.AddTheorTestResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) AddAnswerToQuestion(c *gin.Context) {
	var req AddAnswerReq
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

	var createAnswers []*ent.AnswerCreate
	{
		for _, answer := range req.Answers {
			createAnswers = append(
				createAnswers, 
				m.Client.Answer.Create().
					SetAnswer(answer.Answer).
					SetCorrect(answer.Correct).
					SetQuestionID(req.QuestionID),
			)
		}
	}

	_, err := m.Client.Answer.CreateBulk(createAnswers...).Save(c)
	if ent.IsConstraintError(err) || ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("Question not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddAnswersToQestuion",
				"err":  err,
			},
		).Error("Failed to add answers to question")
		c.JSON(http.StatusInternalServerError, e.FromString(("Failed to add answers to question")))
		c.Abort()
		return
	}

	_, err = m.Client.Question.Query().WithAnswer().Where(question.ID(req.QuestionID)).Only(c)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddAnswersToQestuion",
				"err":  err,
			},
		).Error("Failed to add answers to question")
		c.JSON(http.StatusInternalServerError, e.FromString(("Failed to add answers to question")))
		c.Abort()
		return
	}

	// var resp AddQuestionResp
	// {
		
	// }
}

type UpdateQuestionsReq struct {
	ID			int			`json:"-" uri:"id"`
	Qustion		*string		`json:"question,omitempty"`
}

type UpdateQuestionReq struct {
	ID			int			`json:"id"`
	Qustion		*string		`json:"question,omitempty"`
}

func (m ModuleController) UpdateQuestion(c *gin.Context)