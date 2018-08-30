package actions

import (
	"github.com/gianrubio/nerdalize-challenge/api/models"
	"github.com/gobuffalo/buffalo"
	"github.com/satori/go.uuid"
)

var fake models.Task = models.Task{
	DockerImage: "gianrubio/task:1",
	Environment: []string{"BLA=1", "X=3"},
	Limits: models.Limits{
		Cpu:    "200m",
		Memory: "100m",
	},
	ID:       UUID(),
	Command:  "/bin/my_binary",
	Schedule: "@every 1d",
	Volume:   "s3://my-bucket/my-dir/",
	User: models.User{
		Email: "gianrubio@gmail.com",
		Token: UUID(),
		ID:    UUID(),
	},
}

type TaskResource struct{}

func (t TaskResource) List(c buffalo.Context) error {

	return c.Render(200, r.JSON(fake))
}

func (t TaskResource) Create(c buffalo.Context) error {
	return c.Render(201, r.JSON(fake))
}

func UUID() uuid.UUID {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return uuid
}
