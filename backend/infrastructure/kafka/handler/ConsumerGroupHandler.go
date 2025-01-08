package groupHenler

import (
	"github.com/IBM/sarama"
)

// ConsumerGroupHandler - структура для обработки сообщений из группы потребителей
type ConsumerGroupHandler struct {
	// Канал для передачи сообщений
	MessageChan chan *sarama.ConsumerMessage
}

// Setup вызывается при запуске потребителя
func (h *ConsumerGroupHandler) Setup(sess sarama.ConsumerGroupSession) error {
	// В этом методе можно делать подготовку, например, инициализацию или логирование
	return nil
}

// Cleanup вызывается при завершении работы потребителя
func (h *ConsumerGroupHandler) Cleanup(sess sarama.ConsumerGroupSession) error {
	// В этом методе можно делать очистку, например, освобождение ресурсов
	return nil
}

// ConsumeClaim основная логика обработки сообщений
func (h *ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Читаем сообщения из claim
	for msg := range claim.Messages() {
		// Отправляем сообщение в канал
		h.MessageChan <- msg

		// Помечаем сообщение как обработанное
		sess.MarkMessage(msg, "")
	}

	return nil
}
