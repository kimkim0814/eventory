package controller

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"github.com/tikasan/eventory/models"
	"github.com/tikasan/eventory/utility"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	db *gorm.DB
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, db *gorm.DB) *UsersController {
	return &UsersController{
		Controller: service.NewController("UsersController"),
		db:         db,
	}
}

// Login runs the login action.
func (c *UsersController) Login(ctx *app.LoginUsersContext) error {
	// UsersController_Login: start_implement

	// Put your logic here

	// UsersController_Login: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// RegularCreate runs the regular create action.
func (c *UsersController) RegularCreate(ctx *app.RegularCreateUsersContext) error {
	// UsersController_RegularCreate: start_implement

	// Put your logic here
	// 使用している端末の仮ユーザーが存在するか
	userTerminalDB := models.NewUserTerminalDB(c.db)
	userTerminal, err := userTerminalDB.GetByIdentifier(ctx.Context, ctx.Identifier)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// メールアドレスが既に登録されているか、登録されていれば既に登録されているユーザーに紐付ける
	userDB := models.NewUserDB(c.db)
	currentUser, err := userDB.GetByEmail(ctx.Context, ctx.Email)
	if err == nil {
		// 既に存在しているユーザーと端末情報を紐付ける
		userTerminal.UserID = currentUser.ID
		userTerminalDB.Update(ctx.Context, userTerminal)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		message := "alreadyExists"
		return ctx.OK(&app.Message{&message})
	}
	// メールアドレスとユーザーを紐付ける
	newUser := &models.User{}
	newUser.ID = userTerminal.UserID
	newUser.Email = ctx.Email
	err = userDB.Update(ctx.Context, newUser)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// UsersController_RegularCreate: end_implement
	message := "ok"
	return ctx.OK(&app.Message{&message})
}

// TmpCreate runs the tmp create action.
func (c *UsersController) TmpCreate(ctx *app.TmpCreateUsersContext) error {
	// UsersController_TmpCreate: start_implement

	// Put your logic here
	// 既にユーザーが存在すればtoken情報を返す
	userTerminalDB := models.NewUserTerminalDB(c.db)
	tmpUserTerminal, err := userTerminalDB.GetByIdentifier(ctx.Context, ctx.Identifier)
	if err == nil {
		return ctx.OK(&app.Token{&tmpUserTerminal.Token})
	}
	// ユーザー作成
	t := c.db.Begin()
	userDB := models.NewUserDB(t)
	user := &models.User{}
	err = userDB.Add(ctx.Context, user)
	if err != nil {
		t.Rollback()
		return fmt.Errorf("%v", err)
	}
	userTerminalDB = models.NewUserTerminalDB(t)
	userTerminal := &models.UserTerminal{}
	userTerminal.UserID = user.ID
	userTerminal.Token = utility.CreateToken(userTerminal.Identifier)
	userTerminal.Identifier = ctx.Identifier
	userTerminal.Platform = ctx.Platform
	userTerminal.ClientVersion = ctx.ClientVersion
	err = userTerminalDB.Add(ctx.Context, userTerminal)
	if err != nil {
		t.Rollback()
		return fmt.Errorf("%v", err)
	}
	t.Commit()
	// UsersController_TmpCreate: end_implement
	return ctx.OK(&app.Token{&userTerminal.Token})
}

// Status runs the status action.
func (c *UsersController) Status(ctx *app.StatusUsersContext) error {
	// UsersController_Status: start_implement

	// Put your logic here
	userID, err := utility.GetToken(ctx.Context)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	userTerminal := &models.UserTerminal{}
	userTerminal.UserID = userID
	userTerminal.ClientVersion = ctx.ClientVersion
	userTerminal.Platform = ctx.Platform
	userTerminalDB := models.NewUserTerminalDB(c.db)
	err = userTerminalDB.Update(ctx.Context, userTerminal)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// UsersController_Status: end_implement
	return nil
}
