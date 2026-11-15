package main

import (
	"errors"
	"fmt"
)

var ErrEmptyMessage = errors.New("пустое сообщение")

type Notifier interface {
	Send(msg string) error
}

type BaseNotifier struct {
	Channel string
}

func (b BaseNotifier) Log() {
	fmt.Println("Sending notification via channel:", b.Channel)
}

type EmailNotifier struct {
	BaseNotifier
	Address string
}

func (e EmailNotifier) Send(msg string) error {
	if msg == "" {
		return ErrEmptyMessage
	}
	fmt.Printf("канал %s, адрес %s: %s\n", e.Channel, e.Address, msg)
	return nil
}

var _ Notifier = EmailNotifier{}

type PushNotifier struct {
	BaseNotifier BaseNotifier
	DeviceID     uint64
}

func (p PushNotifier) Send(msg string) error {
	if msg == "" {
		return ErrEmptyMessage
	}
	if p.DeviceID == 0 {
		return &SendError{Channel: p.BaseNotifier.Channel, Reason: "не задан device id"}
	}
	fmt.Printf("канал %s, устройство %d: %s\n", p.BaseNotifier.Channel, p.DeviceID, msg)
	return nil
}

var _ Notifier = PushNotifier{}

type SendError struct {
	Channel string
	Reason  string
}

func (e *SendError) Error() string {
	return fmt.Sprintf("%s: %s", e.Channel, e.Reason)
}

var _ error = (*SendError)(nil)

type BrokenNotifier struct{}

func (BrokenNotifier) Send(string) error {
	panic("драйвер канала повреждён")
}

type Channel struct {
	Name   string
	Sender Notifier
}

func SafeSend(notifier Notifier, msg string) (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = fmt.Errorf("паника при отправке: %v", value)
		}
	}()
	return notifier.Send(msg)
}

func Dispatch(all []Channel, msg string) []error {
	var failures []error
	for _, channel := range all {
		if err := SafeSend(channel.Sender, msg); err != nil {
			failures = append(failures, fmt.Errorf("канал %s: %w", channel.Name, err))
		}
	}
	return failures
}

func main() {
	email := EmailNotifier{
		BaseNotifier: BaseNotifier{Channel: "email"},
		Address:      "example@mail.ru",
	}
	push := PushNotifier{
		BaseNotifier: BaseNotifier{Channel: "push"},
		DeviceID:     123456,
	}

	channels := []Channel{
		{Name: "email", Sender: email},
		{Name: "broken", Sender: BrokenNotifier{}},
		{Name: "push", Sender: push},
	}

	fmt.Println("Сценарий 1: рассылка с паникой одного отправителя")
	failures := Dispatch(channels, "Заказ готов")
	for _, err := range failures {
		fmt.Println("ошибка:", err)
	}
	fmt.Println("Всего ошибок:", len(failures))

	fmt.Println("\nСценарий 2: пустое сообщение")
	failures = Dispatch([]Channel{
		{Name: "email", Sender: email},
		{Name: "push", Sender: push},
	}, "")
	for _, err := range failures {
		fmt.Println("ошибка:", err)
		if errors.Is(err, ErrEmptyMessage) {
			fmt.Println("Нужно ввести текст сообщения.")
		}
		wrapped := fmt.Errorf("рассылка: %w", err)
		fmt.Println("ErrEmptyMessage после дополнительной обёртки:", errors.Is(wrapped, ErrEmptyMessage))
	}
	fmt.Println("Всего ошибок:", len(failures))

	fmt.Println("\nСценарий 3: push без device id")
	push.DeviceID = 0
	failures = Dispatch([]Channel{{Name: "push", Sender: push}}, "Заказ готов")
	for _, err := range failures {
		fmt.Println("ошибка:", err)
		var sendErr *SendError
		if errors.As(err, &sendErr) {
			fmt.Printf("Channel: %s, Reason: %s\n", sendErr.Channel, sendErr.Reason)
		}
	}
	fmt.Println("Всего ошибок:", len(failures))
}
