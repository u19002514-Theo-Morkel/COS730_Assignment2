package Moderation

import (
	"Assignment2/internal/core/Models"
	"encoding/json"
	"log/slog"
)

// Process the moderation queue and add the moderation to the database
func (c Controller) readModerationQueue() {
	for {
		// Get the first item in the queue
		data, err := c.RC.BRPop(c.Ctx, 0, "comment_queue_response").Result()
		if err != nil {
			slog.Error(err.Error())
		}

		if data == nil {
			slog.Debug("No data in queue")
			continue
		}

		slog.Debug(data[1])

		// Unmarshal the data
		var moderation Models.Moderation
		err = json.Unmarshal([]byte(data[1]), &moderation)
		if err != nil {
			slog.Error(err.Error())
		}

		// Add the moderation to the database
		c.DB.Create(&moderation)
	}
}
