package jackbox

import (
	"fmt"
)

func tracerr(ctx string, err error) error {
	return fmt.Errorf("%w\n -> %s", err, ctx)
}
