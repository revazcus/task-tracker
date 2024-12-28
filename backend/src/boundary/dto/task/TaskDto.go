package taskDto

type TaskDto struct {
	Id          string
	Title       string
	Description string
	Status      string
	Priority    string
	Tags        []string
	CreatorId   string
	PerformerId string
	DeadLine    string
	Comments    []string
	Assessment  int
	TimeCosts   int
}
