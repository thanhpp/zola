package alg_test

import (
	"testing"

	"github.com/thanhpp/zola/pkg/alg"
)

func TestLongestCommonSubstring(t *testing.T) {
	var (
		str1     = "abcdfasw"
		str2     = "zbcdf"
		expected = 4
	)

	actual := alg.LongestCommonSubstring(str1, str2)

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
