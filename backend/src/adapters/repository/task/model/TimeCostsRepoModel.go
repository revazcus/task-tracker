package taskRepoModel

import taskTimeCosts "task-tracker/domain/entity/task/cost"

type TimeCostsRepoModel struct {
	TotalMinutes int                  `bson:"total_minutes"`
	Entries      []TimeEntryRepoModel `bson:"entries"`
}

type TimeEntryRepoModel struct {
	Minutes int    `bson:"minutes"`
	Date    int64  `bson:"date"`
	UserId  string `bson:"user_id"`
}

func TimeCostsToRepoModel(timeCosts *taskTimeCosts.TimeCosts) *TimeCostsRepoModel {
	timeCostsRepoModel := TimeCostsRepoModel{TotalMinutes: timeCosts.TotalMinutes()}
	for _, timeEntry := range timeCosts.TimeEntries() {
		timeCostsRepoModel.Entries = append(timeCostsRepoModel.Entries, TimeEntryRepoModel{
			Minutes: timeEntry.Minutes(),
			Date:    timeEntry.Date().UnixNano(),
			UserId:  timeEntry.UserId().String(),
		})
	}
	return &timeCostsRepoModel
}

func (m *TimeCostsRepoModel) GetObject() (*taskTimeCosts.TimeCosts, error) {
	timeCosts := taskTimeCosts.NewTimeCosts()
	for _, timeEntryRepoModel := range m.Entries {
		if err := timeCosts.AddEntry(timeEntryRepoModel.Minutes, timeEntryRepoModel.UserId); err != nil {
			return nil, err
		}
	}
	return timeCosts, nil
}
