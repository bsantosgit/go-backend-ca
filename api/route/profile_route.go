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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
