package common_test

import (
	"testing"

	"github.com/userosettadev/rosetta-cli/internal/common"
	"github.com/stretchr/testify/require"
)

func TestCountTokens(t *testing.T) {

	count, err := common.CountTokens(getTestPrompt())
	require.NoError(t, err)
	require.Equal(t, 112, count)
}

func getTestPrompt() string {

	return `In a futuristic society where artificial intelligence governs every aspect of daily life, a brilliant scientist discovers a hidden flaw in the system. As chaos ensues, explore the consequences of this revelation on individuals and society. Consider the ethical dilemmas, interpersonal relationships, and the struggle for freedom in a world where technology dictates fate. Write a compelling narrative that delves into the intricacies of humanity's relationship with AI, examining the thin line between progress and peril. How do characters navigate this brave new world, and what sacrifices are they willing to make to reclaim control?`
}
