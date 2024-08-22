package client

import (
	"context"
	"net/http"

	"github.com/JackFazackerley/artifactsmmo/pkg/characters"
)

const (
	pathActionFight = "my/Oster1/action/fight"
)

func (c Client) Fight(ctx context.Context) (characters.CharacterFightDataSchema, error) {
	var resp Response[characters.CharacterFightDataSchema]

	if err := c.Do(ctx, pathActionFight, http.MethodPost, nil, &resp); err != nil {
		return characters.CharacterFightDataSchema{}, err
	}

	return resp.Data, nil
}
