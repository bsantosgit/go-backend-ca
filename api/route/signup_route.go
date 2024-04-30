package route

import (
	"time"

	"github.com/bsantosgit/go-backend-ca/api/controller"
	"github.com/bsantosgit/go-backend-ca/bootstrap"
	"github.com/bsantosgit/go-backend-ca/domain"
	"github.com/bsantosgit/go-backend-ca/mongo"
	"github.com/bsantosgit/go-backend-ca/repository"
	"github.com/bsantosgit/go-backend-ca/usecase"
	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
