package taskDto

import taskEntity "task-tracker/domain/entity/task"

type TaskResponseDto struct {
	Task *taskEntity.Task
}
