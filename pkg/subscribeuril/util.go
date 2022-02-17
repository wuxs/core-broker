package subscribeuril

import (
	"fmt"
	"strconv"
	"strings"
)

const defaultSeparator = ":"

func GenerateSubscribeTopic(subscribeID uint, entityID string, opts ...Option) string {
	separator := defaultSeparator
	if len(opts) != 0 {
		separator = opts[0]()
	}
	return fmt.Sprintf("%d%s%s", subscribeID, separator, entityID)
}

func GetSubscribeID(topic string, opts ...Option) uint {
	id, _ := ParseTopic(topic, opts...)
	return id
}

func GetEntityID(topic string, opts ...Option) string {
	_, id := ParseTopic(topic, opts...)
	return id
}

func ParseTopic(topic string, opts ...Option) (subscribeID uint, entityID string) {
	separator := defaultSeparator
	if len(opts) != 0 {
		separator = opts[0]()
	}

	id, err := strconv.ParseUint(topic[:strings.Index(topic, separator)], 10, 0)
	if err == nil {
		subscribeID = uint(id)
	}
	return subscribeID, topic[strings.Index(topic, separator)+1:]
}

type Option func() string

func WithSeparator(separator string) Option {
	return func() string {
		return separator
	}
}
