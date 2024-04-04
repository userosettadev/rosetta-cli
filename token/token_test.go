package token_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/userosettadev/rosetta-cli/api"
	"github.com/userosettadev/rosetta-cli/token"
)

func TestCountTokens(t *testing.T) {

	require.Equal(t, 112, token.Count(getTestPrompt()))
}

func TestCountMultipleFiles(t *testing.T) {

	require.Equal(t, 5, token.CountMultipleFiles([]*api.File{{
		Path:    "internal/app.go",
		Content: []byte("package internal"),
	}, {
		Path:    "util/run.go",
		Content: []byte("package util\n"),
	}}))
}

func getTestPrompt() string {

	return `In a futuristic society where artificial intelligence governs every aspect of daily life, a brilliant scientist discovers a hidden flaw in the system. As chaos ensues, explore the consequences of this revelation on individuals and society. Consider the ethical dilemmas, interpersonal relationships, and the struggle for freedom in a world where technology dictates fate. Write a compelling narrative that delves into the intricacies of humanity's relationship with AI, examining the thin line between progress and peril. How do characters navigate this brave new world, and what sacrifices are they willing to make to reclaim control?`
}
