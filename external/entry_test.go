package external

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ExecSuite struct {
	suite.Suite
}

func (e *ExecSuite) SetupSuite() {
	CreateFile()
}

func (e *ExecSuite) TestExec() {
	err := ExecFile()
	require.NoError(e.T(), err, "exec")
}

func Test_expsuite(t *testing.T) {
	suite.Run(t, new(ExecSuite))
}
