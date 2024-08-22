package client

import (
	"context"
	"net/http"

	"github.com/JackFazackerley/artifactsmmo/pkg/characters"
)

const (
	pathCharacters = "/my/characters"
)

func (c Client) CharacterRequest(ctx context.Context) ([]characters.Character, error) {
	var resp Response[[]characters.Character]

	if err := c.Do(ctx, pathCharacters, http.MethodGet, nil, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
