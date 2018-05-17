package callcount

import (
	"context"
	"fmt"

	"github.com/fnproject/fn/api/models"
)

func (c *CallCount) BeforeCall(ctx context.Context, call *models.Call) error {
	if _, ok := callCountMap[call.ID]; !ok {
		callCountMap[call.ID] = 0
	}

	return nil
}

func (c *CallCount) AfterCall(ctx context.Context, call *models.Call) error {
	callCountMap[call.ID]++
	fmt.Printf("Call nubmer: %d\n", call.ID, callCountMap[call.ID])
	return nil
}
