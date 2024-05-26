package Comment

import (
	"Assignment2/internal/core/Models"
	"encoding/json"
	"log/slog"
)

// addCommentToQueue adds a comment to the queue
func (c Controller) AddCommentToQueue(comment *Models.Comment) {

	// Marshal the comment
	data, err := json.Marshal(comment)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// Push the comment to the queue
	err = c.RC.LPush(c.Ctx, "comment_queue", data).Err()
	if err != nil {
		slog.Error(err.Error())
	}
}
