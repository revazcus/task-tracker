package userRepoModel

import (
	agreementPrimitive "task-tracker/domain/entity/user/agreement"
	commonTime "task-tracker/infrastructure/tools/time"
)

type AgreementRepoModel struct {
	Accepted     bool  `bson:"accepted"`
	AcceptedDate int64 `bson:"accepted_date"`
}

func AgreementToRepoModel(agreement *agreementPrimitive.Agreement) *AgreementRepoModel {
	return &AgreementRepoModel{
		Accepted:     agreement.IsAccepted(),
		AcceptedDate: agreement.AcceptedDate().UnixNano(),
	}
}

func (m *AgreementRepoModel) GetPrimitive() (*agreementPrimitive.Agreement, error) {
	agreement, err := agreementPrimitive.NewBuilder().
		Accepted(m.Accepted).
		AcceptedDate(commonTime.FromUnixNano(m.AcceptedDate)).
		Build()
	if err != nil {
		return nil, err
	}

	return agreement, nil
}
