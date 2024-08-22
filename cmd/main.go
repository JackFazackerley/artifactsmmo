package main

import (
	"context"
	"fmt"

	"github.com/JackFazackerley/artifactsmmo/pkg/client"
	"go.uber.org/zap"
)

const token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6Ik9zdGVyciIsInBhc3N3b3JkX2NoYW5nZWQiOiIifQ.dlZlV8OpIAwmGo33Rt9jyO6YDTVOpKoAL2AkElk4wrw"

func main() {
	logger, err := zap.NewProduction(
		zap.IncreaseLevel(zap.InfoLevel),
		zap.AddStacktrace(zap.FatalLevel),
	)
	if err != nil {
		panic(err)
	}

	client, err := client.NewClient(token)
	if err != nil {
		logger.Error("creating client", zap.Error(err))
	}

	// if _, err := client.Move(context.Background(), 0, 1); err != nil {
	// 	logger.Error("moving character", zap.Error(err))
	// }

	fight, err := client.Fight(context.Background())
	if err != nil {
		logger.Error("fighting", zap.Error(err))
	}

	fmt.Println(fight)

	characters, err := client.CharacterRequest(context.Background())
	if err != nil {
		logger.Error("getting character", zap.Error(err))
	}
	fmt.Println(characters)
}
