package taskRepoModel

import (
	taskTimeCosts "task-tracker/domain/entity/task/cost"
	commonTime "task-tracker/infrastructure/tools/time"
)

type TimeCostsRepoModel struct {
	TimeCosts []TimeCostRepoModel `bson:"time_costs"`
}

type TimeCostRepoModel struct {
	UserId  string `bson:"user_id"`
	Date    int64  `bson:"date"`
	Minutes int    `bson:"minutes"`
}

func TimeCostsToRepoModel(timeCosts *taskTimeCosts.TimeCosts) *TimeCostsRepoModel {
	timeCostsRepoModel := TimeCostsRepoModel{TimeCosts: make([]TimeCostRepoModel, 0)}
	for _, timeCost := range timeCosts.TimeCosts() {
		timeCostsRepoModel.TimeCosts = append(timeCostsRepoModel.TimeCosts, *TimeCostToRepoModel(timeCost))
	}
	return &timeCostsRepoModel
}

func TimeCostToRepoModel(timeCost *taskTimeCosts.TimeCost) *TimeCostRepoModel {
	return &TimeCostRepoModel{
		UserId:  timeCost.UserId().String(),
		Date:    timeCost.Date().UnixNano(),
		Minutes: timeCost.Minutes(),
	}
}

func (m *TimeCostsRepoModel) GetObject() (*taskTimeCosts.TimeCosts, error) {
	timeCosts := taskTimeCosts.NewTimeCosts()
	for _, timeCostRepoModel := range m.TimeCosts {
		if err := timeCosts.AddTimeCost(timeCostRepoModel.UserId, commonTime.FromUnixNano(timeCostRepoModel.Date), timeCostRepoModel.Minutes); err != nil {
			return nil, err
		}
	}
	return timeCosts, nil
}
