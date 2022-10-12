package reviuw

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {

	kontext := context.Background()
	fmt.Println(kontext)

	kontext2 := context.TODO()
	fmt.Println(kontext2)
}
