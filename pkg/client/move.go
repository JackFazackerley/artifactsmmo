package client

import (
	"context"
	"net/http"

	"github.com/JackFazackerley/artifactsmmo/pkg/characters"
)

const (
	pathActionMove = "my/Oster1/action/move"
)

func (c Client) Move(ctx context.Context, x, y int) (characters.Movement, error) {
	var resp Response[characters.Movement]

	payload := &characters.MovementRequest{
		X: x,
		Y: y,
	}

	if err := c.Do(ctx, pathActionMove, http.MethodPost, payload, &resp); err != nil {
		return characters.Movement{}, err
	}

	return resp.Data, nil
}
