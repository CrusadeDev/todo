package app

import (
	"todoApp/internal/app/command"
	"todoApp/internal/app/query"
)

type Application struct {
	Queries  Queries
	Commands Commands
}

type Queries struct {
	ShowAll   query.IndexHandler
	GetSingle query.GetSingleHandler
}

type Commands struct {
	Create command.CreateHandler
	Remove command.RemoveHandler
}
