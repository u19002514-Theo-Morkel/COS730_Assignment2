package Page

import (
	"Assignment2/internal/core/Models"
	"encoding/json"
	"log/slog"
)

func (c Controller) AddPageToQueue(page *Models.Page) {

	data, err := json.Marshal(page)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = c.RC.LPush(c.Ctx, "page_queue", data).Err()
	if err != nil {
		slog.Error(err.Error())
	}
}
