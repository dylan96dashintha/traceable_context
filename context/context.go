package context

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"go.opentelemetry.io/otel/trace"
)

type traceableContext struct {
	context.Context
	traceId trace.TraceID
}

type TraceableContext interface {
	context.Context
	TraceId() trace.TraceID
}

func WithTraceId() TraceableContext {

	randomString, _ := generateRandomString(32)
	traceId, _ := trace.TraceIDFromHex(randomString)
	// need to handle the error
	return &traceableContext{
		Context: context.Background(),
		traceId: traceId,
	}
}

func (t *traceableContext) TraceId() trace.TraceID {
	return t.traceId
}

func generateRandomString(length int) (string, error) {
	// Create a byte slice of the desired length (half the length for hex encoding)
	bytes := make([]byte, length/2)

	// Generate random bytes
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Encode the bytes to a hex string
	return hex.EncodeToString(bytes), nil
}
