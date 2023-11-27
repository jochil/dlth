package generator

import (
	"github.com/CodeIntelligenceTesting/dlth/pkg/candidate"
	"github.com/CodeIntelligenceTesting/dlth/pkg/types"
)

func Render(c *candidate.Candidate) string {
	switch c.Language {
	case types.Go:
		return renderGoUnitTest(c)
	case types.Java:
		return renderJavaFuzzTest(c)
	}

	return ""
}
