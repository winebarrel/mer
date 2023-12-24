package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/mer/internal/util"
)

func TestComma(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		num      string
		expected string
	}{
		{num: "1", expected: "1"},
		{num: "12", expected: "12"},
		{num: "123", expected: "123"},
		{num: "1234", expected: "1,234"},
		{num: "12345", expected: "12,345"},
		{num: "123456", expected: "123,456"},
		{num: "1234567", expected: "1,234,567"},
		{num: "1.123", expected: "1.123"},
		{num: "12.123", expected: "12.123"},
		{num: "123.123", expected: "123.123"},
		{num: "1234.123", expected: "1,234.123"},
		{num: "12345.123", expected: "12,345.123"},
		{num: "123456.123", expected: "123,456.123"},
		{num: "1234567.123", expected: "1,234,567.123"},
	}

	for _, t := range tt {
		c := util.Comma(t.num)
		assert.Equal(t.expected, c)
	}
}
