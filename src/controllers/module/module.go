package module

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/0B1t322/CP-Rosseti-Back/ent"
	"github.com/0B1t322/CP-Rosseti-Back/ent/module"
	"github.com/0B1t322/CP-Rosseti-Back/ent/moduledependcies"
	"github.com/0B1t322/CP-Rosseti-Back/ent/moduletest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/practtest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submodule"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submoduletest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/test"
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

func (m ModuleController) GetOrCreateTestForSubModule(
	ctx context.Context,
	id int,
) (*ent.Test, error) {
	var get *ent.Test

	if test, err := m.Client.Test.Query().Where(
		test.HasSubmoduleTestWith(
			submoduletest.HasSubModuleWith(
				submodule.ID(id),
			),
		),
	).Only(ctx); ent.IsNotFound(err) {
		test, err := m.Client.Test.Create().SetTestType("submodule").Save(ctx)
		if err != nil {
			return nil, err
		}
		get = test
	} else if err != nil {
		return nil, err
	} else {
		get = test
	}

	return get, nil
}

func (m ModuleController) GetOrCreateTestForModule(
	ctx context.Context,
	id int,
) (*ent.Test, error) {
	var get *ent.Test

	if test, err := m.Client.Test.Query().Where(
		test.HasModuleTestWith(
			moduletest.HasModuleWith(
				module.ID(id),
			),
		),
	).Only(ctx); ent.IsNotFound(err) {
		test, err := m.Client.Test.Create().SetTestType("module").Save(ctx)
		if err != nil {
			return nil, err
		}
		get = test
	} else if err != nil {
		return nil, err
	} else {
		get = test
	}

	return get, nil
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
	Correct *bool  `json:"correct,omitempty"`
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

	Test, err := m.GetOrCreateTestForSubModule(c, req.SubModuleID)
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

	subTest, err := m.Client.SubModuleTest.Create().
		SetSubModuleID(req.SubModuleID).
		SetTest(Test).
		Save(c)
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
		theoTest, err := m.Client.TheoreticalTest.Create().SetTest(Test).Save(c)
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
			SetTest(Test).
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

	var getTheorTest *ent.TheoreticalTest
	{
		get, err := m.Client.TheoreticalTest.Query().
			WithQuestion(
				func(qq *ent.QuestionQuery) {
					qq.WithAnswer()
				},
			).Where(
			theoreticaltest.HasTestWith(
				test.HasSubmoduleTestWith(
					submoduletest.ID(subTest.ID),
				),
			),
		).Only(c)
		if ent.IsNotFound(err) {
			// Ignore
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
		} else {
			getTheorTest = get
		}
	}

	var getPractTest *ent.PractTest
	{
		get, err := m.Client.PractTest.Query().
			Where(
				practtest.HasTestWith(
					test.HasSubmoduleTestWith(
						submoduletest.ID(subTest.ID),
					),
				),
			).Only(c)
		if ent.IsNotFound(err) {
			// Ignore
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
		} else {
			getPractTest = get
		}
	}

	var resp AddSubModuleTestResp
	{
		resp.ID = subTest.ID
		var test *AddTheoreticalTestResp

		if getTheorTest != nil {
			var questions []AddQuestionResp
			{
				test = &AddTheoreticalTestResp{}
				for _, question := range getTheorTest.Edges.Question {
					var answers []AddAnswerResp
					{
						for _, answer := range question.Edges.Answer {
							answers = append(
								answers,
								AddAnswerResp{
									ID:      answer.ID,
									Answer:  answer.Answer,
									Correct: &answer.Correct,
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
			if getPractTest != nil {
				practTest = &AddPractTestResp{
					ID:     getPractTest.ID,
					Config: getPractTest.Config,
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
		WithTest(
			func(smtq *ent.SubModuleTestQuery) {
				smtq.WithTest()
			},
		).
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
		SetTestID(subModuleTest.Edges.Test.Edges.Test.ID).
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
								Correct: &answer.Correct,
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
		WithTest(
			func(smtq *ent.SubModuleTestQuery) {
				smtq.WithTest()
			},
		).
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
		SetTestID(subModuleTest.Edges.Test.Edges.Test.ID).
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
		practtest.HasTestWith(
			test.HasSubmoduleTestWith(
				submoduletest.HasSubModuleWith(
					submodule.ID(req.SubModuleID),
				),
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

type UpdateModuleConfigReq struct {
	ModuleID int                    `json:"-" uri:"id"`
	Config   map[string]interface{} `json:"config"`
	Duration *int64                 `json:"duration"`
}

// UpdateConfigToPractTest
//
// @Tags module
//
// @Summary update Pract test
//
// @Description update Pract test to module test
//
// @Description that can do only admin
//
// @Router /v1/module/{id}/test/pract [put]
//
// @Security ApiKeyAuth
//
// @Param theotest body module.UpdateModuleConfigReq true "Pract test info"
//
// @Param id path integer true "id of module"
//
// @Accept json
//
// @Produce json
//
// @Success 200 {object} module.AddModulePractTestResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) UpdateModuleConfigToPractTest(c *gin.Context) {
	var req UpdateModuleConfigReq
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
		practtest.HasTestWith(
			test.HasModuleTestWith(
				moduletest.HasModuleWith(
					module.ID(req.ModuleID),
				),
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
	SubModuleID  int                 `json:"-" uri:"id"`
	Questions    []UpdateQuestionReq `json:"questions"`
	NewQuestions []CreateQuestionReq `json:"newQuestions,omitempty"`
}

type UpdateQuestionReq struct {
	ID         int                `json:"id"`
	Question   *string            `json:"question,omitempty"`
	Answers    []UpdateAnswersReq `json:"answers,omitempty"`
	NewAnswers []CreateAnswerReq  `json:"newAnswers,omitempty"`
	// DeleteAnswers []int              `json:"deleteAnswers,omitempty"`
}

type UpdateAnswersReq struct {
	ID      int     `json:"id"`
	Answer  *string `json:"answer"`
	Correct *bool   `json:"correct"`
}

// UpdateTheorTest
//
// @Tags module
//
// @Summary update Theor test
//
// @Description update Theor test
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
		theoreticaltest.HasTestWith(
			test.HasSubmoduleTestWith(
				submoduletest.HasSubModuleWith(
					submodule.ID(req.SubModuleID),
				),
			),
		),
	).OnlyID(c)
	if ent.IsNotFound(err) || ent.IsConstraintError(err) {
		c.JSON(http.StatusNotFound, e.FromString("Theor test not found"))
		c.Abort()
		return
	}

	for _, question := range req.NewQuestions {
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

	for _, question := range req.Questions {
		qustuinBuilder := m.Client.Question.UpdateOneID(question.ID)
		if question.Question != nil {
			qustuinBuilder.SetQuestion(*question.Question)
		}
		qustuinBuilder.Exec(c)

		for _, answer := range question.Answers {
			answerBuilder := m.Client.Answer.UpdateOneID(answer.ID)
			if answer.Answer != nil {
				answerBuilder.SetAnswer(*answer.Answer)
			}
			if answer.Correct != nil {
				answerBuilder.SetCorrect(*answer.Correct)
			}
			answerBuilder.Exec(c)
		}

		for _, answer := range question.NewAnswers {
			m.Client.Answer.Create().
				SetAnswer(answer.Answer).
				SetCorrect(answer.Correct).
				SetQuestuionID(question.ID).
				Save(c)
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
								Correct: &answer.Correct,
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

type UpdateModuleTheorTestReq struct {
	ModuleID     int                       `json:"-" uri:"id"`
	Duration     *int64                    `json:"duration"`
	Questions    []UpdateModuleQuestionReq `json:"questions"`
	NewQuestions []CreateModuleQuestionReq `json:"newQuestions,omitempty"`
}

type UpdateModuleQuestionReq struct {
	ID         int                      `json:"id"`
	Question   *string                  `json:"question,omitempty"`
	Answers    []UpdateModuleAnswersReq `json:"answers,omitempty"`
	NewAnswers []CreateModuleAnswerReq  `json:"newAnswers,omitempty"`
	// DeleteAnswers []int              `json:"deleteAnswers,omitempty"`
}

type UpdateModuleAnswersReq struct {
	ID      int     `json:"id"`
	Answer  *string `json:"answer"`
	Correct *bool   `json:"correct"`
}

// UpdateModuleTheorTest
//
// @Tags module
//
// @Summary update Theor test
//
// @Description update Theor test
//
// @Router /v1/module/{id}/test/theor [put]
//
// @Security ApiKeyAuth
//
// @Param theortest body module.UpdateModuleTheorTestReq true "Theor test info"
//
// @Param id path integer true "id of module"
//
// @Accept json
//
// @Produce json
//
// @Success 200 {object} module.AddModuleTheorTestResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) UpdateModuleTheorTest(c *gin.Context) {
	var req UpdateModuleTheorTestReq
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
		theoreticaltest.HasTestWith(
			test.HasModuleTestWith(
				moduletest.HasModuleWith(
					module.ID(req.ModuleID),
				),
			),
		),
	).OnlyID(c)
	if ent.IsNotFound(err) || ent.IsConstraintError(err) {
		c.JSON(http.StatusNotFound, e.FromString("Theor test not found"))
		c.Abort()
		return
	}

	if req.Duration != nil {
		m.Client.TheoreticalTest.UpdateOneID(theorTestID).SetDuration(int(*req.Duration)).Exec(c)
	}

	for _, question := range req.NewQuestions {
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

	for _, question := range req.Questions {
		qustuinBuilder := m.Client.Question.UpdateOneID(question.ID)
		if question.Question != nil {
			qustuinBuilder.SetQuestion(*question.Question)
		}
		qustuinBuilder.Exec(c)

		for _, answer := range question.Answers {
			answerBuilder := m.Client.Answer.UpdateOneID(answer.ID)
			if answer.Answer != nil {
				answerBuilder.SetAnswer(*answer.Answer)
			}
			if answer.Correct != nil {
				answerBuilder.SetCorrect(*answer.Correct)
			}
			answerBuilder.Exec(c)
		}

		for _, answer := range question.NewAnswers {
			m.Client.Answer.Create().
				SetAnswer(answer.Answer).
				SetCorrect(answer.Correct).
				SetQuestuionID(question.ID).
				Save(c)
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

	var resp AddModuleTheorTestResp
	{
		theoTest := AddModuleTheorTestResp{}

		var questions []AddModuleQuestionResp
		{
			for _, question := range theorTest.Edges.Question {
				var answers []AddModuleAnswerResp
				{
					for _, answer := range question.Edges.Answer {
						answers = append(
							answers,
							AddModuleAnswerResp{
								ID:      answer.ID,
								Answer:  answer.Answer,
								Correct: &answer.Correct,
							},
						)
					}
				}

				questions = append(
					questions,
					AddModuleQuestionResp{
						ID:       question.ID,
						Question: question.Question,
						Answers:  answers,
					},
				)
			}
		}
		theoTest.Questions = questions
		theoTest.Duratione = int64(theorTest.Duration)
		resp = theoTest
	}

	c.JSON(http.StatusOK, resp)
}

type DeleteTheorTestReq struct {
	SubModuleID int `json:"-" uri:"id"`
}

// DeleteTheorTest
//
// @Tags module
//
// @Summary delete Theor test
//
// @Description delete Theor test
//
// @Router /v1/module/submodule/{id}/test/theor [delete]
//
// @Security ApiKeyAuth
//
// @Param id path integer true "id of submodule"
//
// @Produce json
//
// @Success 200
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) DeleteTheorTest(c *gin.Context) {
	var req DeleteTheorTestReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
			c.Abort()
			return
		}
	}

	if _, err := m.Client.TheoreticalTest.Delete().Where(
		theoreticaltest.HasTestWith(
			test.HasSubmoduleTestWith(
				submoduletest.HasSubModuleWith(
					submodule.ID(req.SubModuleID),
				),
			),
		),
	).Exec(c); ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("Theor test not found"))
		c.Abort()
		return
	} else if ent.IsConstraintError(err) {
		log.Info(err)
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "DeleteTheorTest",
				"err":  err,
			},
		).Error("Failed delete submodule theor test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed delete submodule theor test"))
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

type DeletePractTestReq struct {
	SubModuleID int `json:"-" uri:"id"`
}

// DeletePractTest
//
// @Tags module
//
// @Summary delete Pract test
//
// @Description delete Pract test
//
// @Router /v1/module/submodule/{id}/test/pract [delete]
//
// @Security ApiKeyAuth
//
// @Param id path integer true "id of submodule"
//
// @Produce json
//
// @Success 200
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) DeletePractTest(c *gin.Context) {
	var req DeletePractTestReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
			c.Abort()
			return
		}
	}

	if _, err := m.Client.PractTest.Delete().Where(
		practtest.HasTestWith(
			test.HasSubmoduleTestWith(
				submoduletest.HasSubModuleWith(
					submodule.ID(req.SubModuleID),
				),
			),
		),
	).Exec(c); ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("Pract test not found"))
		c.Abort()
		return
	} else if ent.IsConstraintError(err) {
		log.Info(err)
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "DeletePractTest",
				"err":  err,
			},
		).Error("Failed delete submodule pract test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed delete submodule pract test"))
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

type DeleteModuleTheorTestReq struct {
	ModuleID int `json:"-" uri:"id"`
}

// DeleteTheorTest
//
// @Tags module
//
// @Summary delete Theor test
//
// @Description delete Theor test
//
// @Router /v1/module/{id}/test/theor [delete]
//
// @Security ApiKeyAuth
//
// @Param id path integer true "id of module"
//
// @Produce json
//
// @Success 200
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) DeleteModuleTheorTest(c *gin.Context) {
	var req DeleteModuleTheorTestReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
			c.Abort()
			return
		}
	}

	if _, err := m.Client.TheoreticalTest.Delete().Where(
		theoreticaltest.HasTestWith(
			test.HasModuleTestWith(
				moduletest.HasModuleWith(
					module.ID(req.ModuleID),
				),
			),
		),
	).Exec(c); ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("Theor test not found"))
		c.Abort()
		return
	} else if ent.IsConstraintError(err) {
		log.Info(err)
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "DeleteTheorTest",
				"err":  err,
			},
		).Error("Failed delete submodule theor test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed delete submodule theor test"))
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

type DeleteModulePractTestReq struct {
	ModuleID int `json:"-" uri:"id"`
}

// DeletePractTest
//
// @Tags module
//
// @Summary delete Pract test
//
// @Description delete Pract test
//
// @Router /v1/module/{id}/test/pract [delete]
//
// @Security ApiKeyAuth
//
// @Param id path integer true "id of module"
//
// @Produce json
//
// @Success 200
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) DeleteModulePractTest(c *gin.Context) {
	var req DeleteModulePractTestReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
			c.Abort()
			return
		}
	}

	if _, err := m.Client.PractTest.Delete().Where(
		practtest.HasTestWith(
			test.HasModuleTestWith(
				moduletest.HasModuleWith(
					module.ID(req.ModuleID),
				),
			),
		),
	).Exec(c); ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("Pract test not found"))
		c.Abort()
		return
	} else if ent.IsConstraintError(err) {
		log.Info(err)
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "DeletePractTest",
				"err":  err,
			},
		).Error("Failed delete submodule pract test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed delete submodule pract test"))
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

type GetSubModuleReq struct {
	SubModuleID int `json:"-" uri:"id"`
}

type GetSubModuleResp struct {
	AddSubModuleResp `json",inline"`
	Test             *AddSubModuleTestResp `json:"test"`
}

// GetSubModule
//
// @Tags module
//
// @Summary get sub module
//
// @Description get sub module
//
// @Router /v1/module/submodule/{id} [get]
//
// @Security ApiKeyAuth
//
// @Param id path integer true "id of submodule"
//
// @Produce json
//
// @Success 200 {object} GetSubModuleResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) GetSubModule(c *gin.Context) {
	var req GetSubModuleReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
			c.Abort()
			return
		}
	}

	var getSubModule *ent.SubModule
	{
		get, err := m.Client.SubModule.Query().
			WithTest().
			Where(submodule.ID(req.SubModuleID)).
			Only(c)
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, e.FromString("submodule not found"))
			c.Abort()
			return
		} else if err != nil {
			log.WithFields(
				log.Fields{
					"pkg":  "controllers/module",
					"func": "GetSubModule",
					"err":  err,
				},
			).Error("Failed to get sub module")
			c.JSON(http.StatusInternalServerError, e.FromString("Failed to get sub module"))
			c.Abort()
			return
		} else {
			getSubModule = get
		}
	}

	var getTheorTest *ent.TheoreticalTest
	{
		get, err := m.Client.TheoreticalTest.Query().
			WithQuestion(
				func(qq *ent.QuestionQuery) {
					qq.WithAnswer()
				},
			).Where(
			theoreticaltest.HasTestWith(
				test.HasSubmoduleTestWith(
					submoduletest.HasSubModuleWith(
						submodule.ID(req.SubModuleID),
					),
				),
			),
		).Only(c)
		if ent.IsNotFound(err) {
			// Ignore
		} else if err != nil {
			log.WithFields(
				log.Fields{
					"pkg":  "controllers/module",
					"func": "GetSubModule",
					"err":  err,
				},
			).Error("Failed to get sub module")
			c.JSON(http.StatusInternalServerError, e.FromString("Failed to get sub module"))
			c.Abort()
			return
		} else {
			getTheorTest = get
		}
	}

	var getPractTest *ent.PractTest
	{
		get, err := m.Client.PractTest.Query().
			Where(
				practtest.HasTestWith(
					test.HasSubmoduleTestWith(
						submoduletest.HasSubModuleWith(
							submodule.ID(req.SubModuleID),
						),
					),
				),
			).Only(c)
		if ent.IsNotFound(err) {
			// Ignore
		} else if err != nil {
			log.WithFields(
				log.Fields{
					"pkg":  "controllers/module",
					"func": "GetSubModule",
					"err":  err,
				},
			).Error("Failed to get sub module")
			c.JSON(http.StatusInternalServerError, e.FromString("Failed to get sub module"))
			c.Abort()
			return
		} else {
			getPractTest = get
		}
	}

	var resp GetSubModuleResp
	{
		resp.ID = req.SubModuleID
		resp.Name = getSubModule.Name
		resp.Text = getSubModule.Text
		var test *AddTheoreticalTestResp

		if getTheorTest != nil {
			var questions []AddQuestionResp
			{
				test = &AddTheoreticalTestResp{}
				for _, question := range getTheorTest.Edges.Question {
					var answers []AddAnswerResp
					{
						for _, answer := range question.Edges.Answer {
							answers = append(
								answers,
								AddAnswerResp{
									ID:      answer.ID,
									Answer:  answer.Answer,
									Correct: &answer.Correct,
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

		var practTest *AddPractTestResp
		{
			if getPractTest != nil {
				practTest = &AddPractTestResp{
					ID:     getPractTest.ID,
					Config: getPractTest.Config,
				}
			}
		}
		if getSubModule.Edges.Test != nil {
			resp.Test = &AddSubModuleTestResp{
				ID:              getSubModule.Edges.Test.ID,
				TheoretcialTest: test,
				PractTest:       practTest,
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

type GetModuleReq struct {
	ID int `json:"-" uri:"id"`
}

type GetModuleResp struct {
	ID         int                `json:"id"`
	Name       string             `json:"name"`
	SubModules []AddSubModuleResp `json:"subModules,omitempty"`
	DependOn   []int              `json:"dependOn,omitempty"`
	Tests       *AddModuleTestResp  `json:"tests"`
}

func (m ModuleController) GetModule(ctx context.Context, id int) (*GetModuleResp, error) {
	getModule, err := m.Client.Module.Query().
		WithModuleDependcies(
			func(mdq *ent.ModuleDependciesQuery) {
				mdq.Where(
					moduledependcies.DependentID(id),
				)
			},
		).
		WithSubModules().
		WithTest().
		Where(
			module.ID(id),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	var getTheorTest *ent.TheoreticalTest
	{
		get, err := m.Client.TheoreticalTest.Query().
			WithQuestion(
				func(qq *ent.QuestionQuery) {
					qq.WithAnswer()
				},
			).Where(
			theoreticaltest.HasTestWith(
				test.HasModuleTestWith(
					moduletest.HasModuleWith(
						module.ID(id),
					),
				),
			),
		).Only(ctx)
		if ent.IsNotFound(err) {
			// Ignore
		} else if err != nil {
			return nil, err
		} else {
			getTheorTest = get
		}
	}

	var getPractTest *ent.PractTest
	{
		get, err := m.Client.PractTest.Query().
			Where(
				practtest.HasTestWith(
					test.HasModuleTestWith(
						moduletest.HasModuleWith(
							module.ID(id),
						),
					),
				),
			).Only(ctx)
		if ent.IsNotFound(err) {
			// Ignore
		} else if err != nil {
			return nil, err
		} else {
			getPractTest = get
		}
	}

	var resp GetModuleResp
	{
		resp.ID = getModule.ID
		resp.Name = getModule.Name
		if getModule.Edges.SubModules != nil {
			for _, subModule := range getModule.Edges.SubModules {
				resp.SubModules = append(
					resp.SubModules,
					AddSubModuleResp{
						ID:       subModule.ID,
						ModuleID: getModule.ID,
						Name:     subModule.Name,
						Text:     subModule.Text,
					},
				)
			}
		}

		if getModule.Edges.ModuleDependcies != nil {
			for _, dependOn := range getModule.Edges.ModuleDependcies {
				resp.DependOn = append(resp.DependOn, dependOn.DependentOnID)
			}
		}

		var task *AddModuleTestResp

		var theorTest *AddModuleTheorTestResp
		var practTest *AddModulePractTestResp

		if getTheorTest != nil {
			theorTest = &AddModuleTheorTestResp{
				Duratione: int64(getTheorTest.Duration),
			}

			var questions []AddModuleQuestionResp
			{
				for _, question := range getTheorTest.Edges.Question {
					var answers []AddModuleAnswerResp
					{
						for _, answer := range question.Edges.Answer {
							answers = append(
								answers,
								AddModuleAnswerResp{
									ID:      answer.ID,
									Answer:  answer.Answer,
									Correct: &answer.Correct,
								},
							)
						}
					}

					questions = append(
						questions,
						AddModuleQuestionResp{
							ID:       question.ID,
							Question: question.Question,
							Answers:  answers,
						},
					)
				}
			}

			theorTest.Questions = questions
		}

		if getPractTest != nil {
			practTest = &AddModulePractTestResp{
				Duratione: int64(getPractTest.Duration),
				ID: getPractTest.ID,
				Config: getPractTest.Config,
			}
		}

		if getPractTest != nil || getTheorTest != nil {
			task = &AddModuleTestResp{
				TheoretcialTest: theorTest,
				PractTest: practTest,
			}

			if getPractTest != nil {
				task.ID = getPractTest.TestID
			} else {
				task.ID = getTheorTest.ID
			}
		}

		resp.Tests = task
	}

	return &resp, nil
}

// GetModule
//
// @Tags module
//
// @Summary get  module
//
// @Description get  module
//
// @Router /v1/module/{id} [get]
//
// @Security ApiKeyAuth
//
// @Param id path integer true "id of module"
//
// @Produce json
//
// @Success 200 {object} module.GetModuleResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) HTTPGetModule(c *gin.Context) {
	var req GetModuleReq
	{
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
			c.Abort()
			return
		}
	}

	resp, err := m.GetModule(c, req.ID)
	if ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("Module not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "GetModule",
				"err":  err,
			},
		).Error("Failed to get module")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to get module"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (m ModuleController) AddModuleDependecy(
	ctx context.Context,
	id int,
	depend_id int,
) error {
	_, err := m.Client.ModuleDependcies.Create().
		SetDependentID(id).
		SetDependentOnID(depend_id).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m ModuleController) DeleteModuleDependecy(
	ctx context.Context,
	id int,
	depend_id int,
) error {
	_, err := m.Client.ModuleDependcies.Delete().
		Where(
			moduledependcies.DependentID(id),
			moduledependcies.DependentOnID(depend_id),
		).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

type AddModuleDependecyReq struct {
	ID      int `json:"-" uri:"id"`
	DepenOn int `json:"depenOn"`
}

// AddModuleDependecy
//
// @Tags module
//
// @Summary add module dependecy
//
// @Description Add module dependecy
//
// @Router /v1/module/{id}/dependecy [put]
//
// @Security ApiKeyAuth
//
// @Param id path integer true "id of module"
//
// @Param body body module.AddModuleDependecyReq true "body"
//
// @Accept json
//
// @Produce json
//
// @Success 200 {object} module.GetModuleResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) HTTPAddModuleDependecy(
	c *gin.Context,
) {
	var req AddModuleDependecyReq
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

	if err := m.AddModuleDependecy(c, req.ID, req.DepenOn); ent.IsConstraintError(err) {
		c.JSON(http.StatusNotFound, e.FromString("one of module not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddModuleDependecy",
				"err":  err,
			},
		).Error("Failed to add module dependecy")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to add module dependecy"))
		c.Abort()
		return
	}

	resp, err := m.GetModule(c, req.ID)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddModuleDependecy",
				"err":  err,
			},
		).Error("Failed to add module dependecy")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to add module dependecy"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteModuleDependecy
//
// @Tags module
//
// @Summary Delete module dependecy
//
// @Description Delete module dependecy
//
// @Router /v1/module/{id}/dependecy [delete]
//
// @Security ApiKeyAuth
//
// @Param id path integer true "id of module"
//
// @Param body body module.AddModuleDependecyReq true "body"
//
// @Accept json
//
// @Produce json
//
// @Success 200 {object} module.GetModuleResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) HTTPDeleteModuleDependecy(
	c *gin.Context,
) {
	var req AddModuleDependecyReq
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

	if err := m.DeleteModuleDependecy(c, req.ID, req.DepenOn); ent.IsConstraintError(err) {
		c.JSON(http.StatusNotFound, e.FromString("one of module not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "DeleteModuleDependecy",
				"err":  err,
			},
		).Error("Failed to Delete module dependecy")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to Delete module dependecy"))
		c.Abort()
		return
	}

	resp, err := m.GetModule(c, req.ID)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "DeleteModuleDependecy",
				"err":  err,
			},
		).Error("Failed to Delete module dependecy")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to Delete module dependecy"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resp)
}

type GetModulesReq struct {
}

type GetModulesResp struct {
	Modules []GetModuleResp `json:"modules"`
}

func (m ModuleController) GetModules(
	ctx context.Context,
	nameMatch string,
	offset, limit int,
) (*GetModulesResp, error) {
	builder := m.Client.Module.Query().
		WithModuleDependcies().
		WithSubModules()
	if nameMatch != "" {
		builder.Where(module.NameContains(nameMatch))
	}
	if offset != 0 {
		builder.Offset(offset)
	}
	if limit != 0 {
		builder.Limit(limit)
	}

	modules, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	var resp GetModulesResp
	{
		for _, module := range modules {
			var moduleResp GetModuleResp
			{
				moduleResp.ID = module.ID
				moduleResp.Name = module.Name
				if module.Edges.SubModules != nil {
					for _, subModule := range module.Edges.SubModules {
						moduleResp.SubModules = append(
							moduleResp.SubModules,
							AddSubModuleResp{
								ID:       subModule.ID,
								ModuleID: module.ID,
								Name:     subModule.Name,
								Text:     subModule.Text,
							},
						)
					}
				}

				if module.Edges.ModuleDependcies != nil {
					for _, dependOn := range module.Edges.ModuleDependcies {
						moduleResp.DependOn = append(moduleResp.DependOn, dependOn.DependentOnID)
					}
				}
			}
			resp.Modules = append(resp.Modules, moduleResp)
		}
	}
	return &resp, nil
}

// HTTPGetModules
//
// @Tags module
//
// @Summary get modules
//
// @Description get modules
//
// @Router /v1/module [get]
//
// @Security ApiKeyAuth
//
// @Param namemMatch query string false "name filter"
//
// @Param offset query integer false "name filter"
//
// @Param limit query integer false "name filter"
//
// @Produce json
//
// @Success 200 {object} module.GetModulesResp
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) HTTPGetModules(c *gin.Context) {
	var (
		nameMatch = c.Query("nameMatch")
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")

		offset int
		limit  int
	)

	if offsetStr != "" {
		offset, _ = strconv.Atoi(offsetStr)
	}

	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	resp, err := m.GetModules(c, nameMatch, offset, limit)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "GetModules",
				"err":  err,
			},
		).Error("Failed to Get Modules module dependecy")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to Get Modules module dependecy"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, resp)
}

type AddModuleTestReq struct {
	ModuleID        int                             `json:"-" uri:"id"`
	TheoreticalTest *CreateModuleTheoreticalTestReq `json:"theoreticalTest,omitempty"`
	PractTest       *CreateModulePractTestReq       `json:"practTest,omitempty"`
}

type CreateModulePractTestReq struct {
	Config    map[string]interface{} `json:"config"`
	Duratione int64                  `json:"duration"`
}

type CreateModuleTheoreticalTestReq struct {
	Questions []CreateModuleQuestionReq `json:"questions"`
	Duratione int64                     `json:"duration"`
}

type CreateModuleQuestionReq struct {
	Question string                  `json:"question"`
	Answers  []CreateModuleAnswerReq `json:"answers"`
}

type CreateModuleAnswerReq struct {
	Answer  string `json:"answer"`
	Correct bool   `json:"correct"`
}

type AddModuleTestResp struct {
	ID              int                     `json:"id"`
	TheoretcialTest *AddModuleTheorTestResp `json:"theoreticalTest,omitempty"`
	PractTest       *AddModulePractTestResp `json:"practTest,omitempty"`
}

type AddModuleTheorTestResp struct {
	Questions []AddModuleQuestionResp `json:"questions"`
	Duratione int64                   `json:"duration"`
}

type AddModulePractTestResp struct {
	ID        int                    `json:"id"`
	Config    map[string]interface{} `json:"config"`
	Duratione int64                  `json:"duration"`
}

type AddModuleQuestionResp struct {
	ID       int                   `json:"id"`
	Question string                `json:"question"`
	Answers  []AddModuleAnswerResp `json:"answers"`
}

type AddModuleAnswerResp struct {
	ID      int    `json:"id"`
	Answer  string `json:"answer"`
	Correct *bool  `json:"correct,omitempty"`
}

// AddModuleTest
//
// @Tags module
//
// @Summary add module test
//
// @Description add module test
//
// @Description that can do only admin
//
// @Router /v1/module/{id}/test [post]
//
// @Security ApiKeyAuth
//
// @Param moduletest body module.AddModuleTestReq true "ModuleTest info"
//
// @Param id path integer true "id of module"
//
// @Accept json
//
// @Produce json
//
// @Success 201 {object} module.AddModuleTestResp
//
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 404 {object} e.Error "module not found"
//
// @Failure 500 {object} e.Error "internal"
//
// @Failure 401 {object} e.Error "not auth"
func (m ModuleController) AddModuleTest(c *gin.Context) {
	var req AddModuleTestReq
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

	Test, err := m.GetOrCreateTestForModule(c, req.ModuleID)
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

	subTest, err := m.Client.ModuleTest.Create().
		SetModuleID(req.ModuleID).
		SetTest(Test).
		Save(c)
	if ent.IsConstraintError(err) {
		log.Info("err", err)
		c.JSON(http.StatusNotFound, e.FromString("SubModule not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg":  "controllers/module",
				"func": "AddModuleTest",
				"err":  err,
			},
		).Error("Failed add module test")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to add module test"))
		c.Abort()
		return
	}

	if req.TheoreticalTest != nil {
		theoTest, err := m.Client.TheoreticalTest.Create().
			SetTest(Test).
			SetDuration(int(req.TheoreticalTest.Duratione)).
			Save(c)
		if err != nil {
			log.WithFields(
				log.Fields{
					"pkg":  "controllers/module",
					"func": "AddModuleTest",
					"err":  err,
				},
			).Error("Failed add module test")
			c.JSON(http.StatusInternalServerError, e.FromString("Failed to add module test"))
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
						"func": "AddModuleTest",
						"err":  err,
					},
				).Error("Failed add module test")
				c.JSON(http.StatusInternalServerError, e.FromString("Failed to add module test"))
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
							"func": "AddModuleTest",
							"err":  err,
						},
					).Error("Failed add module test")
					c.JSON(http.StatusInternalServerError, e.FromString("Failed to add module test"))
					c.Abort()
					return
				}
			}
		}
	}

	if req.PractTest != nil {
		_, err := m.Client.PractTest.Create().
			SetTest(Test).
			SetConfig(req.PractTest.Config).
			Save(c)
		if err != nil {
			log.WithFields(
				log.Fields{
					"pkg":  "controllers/module",
					"func": "AddModuleTest",
					"err":  err,
				},
			).Error("Failed add module test")
			c.JSON(http.StatusInternalServerError, e.FromString("Failed to add module test"))
			c.Abort()
			return
		}
	}

	var getTheorTest *ent.TheoreticalTest
	{
		get, err := m.Client.TheoreticalTest.Query().
			WithQuestion(
				func(qq *ent.QuestionQuery) {
					qq.WithAnswer()
				},
			).Where(
			theoreticaltest.HasTestWith(
				test.HasModuleTestWith(
					moduletest.ID(subTest.ID),
				),
			),
		).Only(c)
		if ent.IsNotFound(err) {
			// Ignore
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
		} else {
			getTheorTest = get
		}
	}

	var getPractTest *ent.PractTest
	{
		get, err := m.Client.PractTest.Query().
			Where(
				practtest.HasTestWith(
					test.HasModuleTestWith(
						moduletest.ID(subTest.ID),
					),
				),
			).Only(c)
		if ent.IsNotFound(err) {
			// Ignore
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
		} else {
			getPractTest = get
		}
	}

	var resp AddModuleTestResp
	{
		resp.ID = subTest.ID
		var test *AddModuleTheorTestResp

		if getTheorTest != nil {
			var questions []AddModuleQuestionResp
			{
				test = &AddModuleTheorTestResp{}
				for _, question := range getTheorTest.Edges.Question {
					var answers []AddModuleAnswerResp
					{
						for _, answer := range question.Edges.Answer {
							answers = append(
								answers,
								AddModuleAnswerResp{
									ID:      answer.ID,
									Answer:  answer.Answer,
									Correct: &answer.Correct,
								},
							)
						}
					}

					questions = append(
						questions,
						AddModuleQuestionResp{
							ID:       question.ID,
							Question: question.Question,
							Answers:  answers,
						},
					)
				}
			}

			test.Questions = questions
			test.Duratione = time.Duration(getTheorTest.Duration).Microseconds()
		}

		resp.TheoretcialTest = test

		var practTest *AddModulePractTestResp
		{
			if getPractTest != nil {
				practTest = &AddModulePractTestResp{
					ID:        getPractTest.ID,
					Config:    getPractTest.Config,
					Duratione: time.Duration(getPractTest.Duration).Microseconds(),
				}
			}
		}

		resp.PractTest = practTest
	}

	c.JSON(http.StatusCreated, resp)
}
