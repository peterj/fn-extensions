package callcount

import (
	"context"
	"fmt"

	"github.com/fnproject/fn/api/models"
)

func (c *CallCount) BeforeCall(ctx context.Context, call *models.Call) error {
	if _, ok := callCountMap[call.AppID]; !ok {
		callCountMap[call.AppID] = 0
	}

	return nil
}

func (c *CallCount) AfterCall(ctx context.Context, call *models.Call) error {
	callCountMap[call.AppID]++
	fmt.Printf("Call number: %d\n", callCountMap[call.AppID])
	return nil
}
