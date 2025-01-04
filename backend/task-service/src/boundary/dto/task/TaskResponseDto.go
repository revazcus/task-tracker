package taskDto

import taskEntity "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task"

type TaskResponseDto struct {
	Task *taskEntity.Task
}
