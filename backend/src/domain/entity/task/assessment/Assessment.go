package assessmentPrimitive

type Assessment int

func AssessmentFrom(value int) (*Assessment, error) {
	if value <= 0 {
		return nil, ErrInvalidAssessment
	}
	assessment := Assessment(value)
	return &assessment, nil
}

func (a Assessment) Int() int {
	return int(a)
}
