package hydrators

import (
	"github.com/dadrus/heimdall/internal/heimdall"
	"github.com/dadrus/heimdall/internal/pipeline/subject"
)

type Hydrator interface {
	Execute(heimdall.Context, *subject.Subject) error
	WithConfig(config map[any]any) (Hydrator, error)
}
