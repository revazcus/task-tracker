package taskRepoModel

import (
	shortUserRepoModel "task-tracker/common/repoModel/shortUser"
	taskTimeCosts "task-tracker/domain/entity/task/cost"
	commonTime "task-tracker/infrastructure/tools/time"
)

type TimeCostsRepoModel struct {
	TimeInvestmentRepoModels []TimeInvestmentRepoModel `bson:"time_investments"`
}

type TimeInvestmentRepoModel struct {
	Date    int64                                  `bson:"date"`
	Minutes int                                    `bson:"minutes"`
	Worker  *shortUserRepoModel.ShortUserRepoModel `bson:"worker"`
}

func TimeInvestmentsToRepoModels(timeCosts *taskTimeCosts.TimeCosts) *TimeCostsRepoModel {
	timeCostsRepoModel := TimeCostsRepoModel{TimeInvestmentRepoModels: make([]TimeInvestmentRepoModel, 0)}
	for _, timeInvestment := range timeCosts.TimeInvestments() {
		timeCostsRepoModel.TimeInvestmentRepoModels = append(timeCostsRepoModel.TimeInvestmentRepoModels, *TimeInvestmentToRepoModel(timeInvestment))
	}
	return &timeCostsRepoModel
}

func TimeInvestmentToRepoModel(timeCost *taskTimeCosts.TimeInvestment) *TimeInvestmentRepoModel {
	return &TimeInvestmentRepoModel{
		Worker:  shortUserRepoModel.ShortUserToRepoModel(timeCost.Worker()),
		Date:    timeCost.Date().UnixNano(),
		Minutes: timeCost.Minutes(),
	}
}

func (m *TimeCostsRepoModel) GetObject() (*taskTimeCosts.TimeCosts, error) {
	timeCosts := taskTimeCosts.NewTimeCosts()
	for _, timeCostRepoModel := range m.TimeInvestmentRepoModels {
		worker, err := timeCostRepoModel.Worker.GetObject()
		if err != nil {
			return nil, err
		}
		if err := timeCosts.AddTimeCost(worker, commonTime.FromUnixNano(timeCostRepoModel.Date), timeCostRepoModel.Minutes); err != nil {
			return nil, err
		}
	}
	return timeCosts, nil
}
