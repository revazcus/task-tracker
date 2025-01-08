package kafkaTopic

type TopicName string

func (t TopicName) String() string {
	return string(t)
}
