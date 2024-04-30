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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
