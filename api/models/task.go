package models

import (
	"github.com/satori/go.uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	DockerImage string    `json:"docker-image`
	Environment []string  `json:"env`
	Command     string    `json:command`
	Schedule    string    `json:schedule`
	Limits      Limits    `json:limits`
	Volume      string    `json:volume`
	User        User      `json:user`
}

type Limits struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}
