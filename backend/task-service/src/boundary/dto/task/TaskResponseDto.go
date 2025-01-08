package taskDto

import taskEntity "task-service/src/domain/entity/task"

type TaskResponseDto struct {
	Task *taskEntity.Task
}
