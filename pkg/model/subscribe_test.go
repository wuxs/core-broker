package model

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUpdateEntitySubscribeEndpoint(t *testing.T) {
	subscribeAddr := ""
	data := "Default Title@4@amqp://tkeel.io:5672/AAa1yvw7dYJGkuQU,123132@5@amqp://tkeel.io:5672/soV8UVBhdyLakMpR,1@6@amqp://tkeel.io:5672/Zwm1ihXdD7Q7eGcg,1test@7@amqp://tkeel.io:5672/ORc25nkwSMOUHSDe"
	endpoint := "Default Title@4@amqp://tkeel.io:5672/AAa1yvw7dYJGkuQU"

	addrs := strings.Split(data, ",")
	validAddresses := make([]string, 0, len(addrs))
	for i := range addrs {
		if addrs[i] != endpoint {
			validAddresses = append(validAddresses, addrs[i])
		}
	}
	if len(validAddresses) != 0 {
		subscribeAddr = strings.Join(validAddresses, ",")
	} else {
		subscribeAddr = ""
	}

	assert.Equal(t, subscribeAddr, "123132@5@amqp://tkeel.io:5672/soV8UVBhdyLakMpR,1@6@amqp://tkeel.io:5672/Zwm1ihXdD7Q7eGcg,1test@7@amqp://tkeel.io:5672/ORc25nkwSMOUHSDe")
}
