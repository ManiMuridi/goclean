package brokerservice

import (
	"fmt"

	"github.com/google/uuid"
)

type Storage interface {
	Store(message *Message) (*Message, error)
	GetById(id uuid.UUID) (*Message, error)
	GetAll() []*Message
}

type memoryStorage struct {
	messages []*Message
}

func (m *memoryStorage) Store(message *Message) (*Message, error) {
	m.messages = append(m.messages, message)
	return message, nil
}

func (m *memoryStorage) GetById(id uuid.UUID) (*Message, error) {
	return nil, nil
}

func (m *memoryStorage) GetAll() []*Message {
	return m.messages
}

func NewMemoryStorage() Storage {
	fmt.Println("MEMORY STORAGE: DO NOT USE FOR PRODUCTION")
	return &memoryStorage{make([]*Message, 0)}
}
