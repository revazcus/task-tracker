package taskDto

import taskEntity "task-service/domain/entity/task"

type TaskResponseDto struct {
	Task *taskEntity.Task
}
