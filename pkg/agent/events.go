package agent

import (
	"fmt"
	"time"
)

// EventKind identifies a structured agent-loop event.
//
// Deprecated: use github.com/sipeed/picoclaw/pkg/events.Kind for new runtime
// event consumers. This legacy kind exists only during the runtime event
// migration window.
type EventKind uint8

const (
	// EventKindTurnStart is emitted when a turn begins processing.
	EventKindTurnStart EventKind = iota
	// EventKindTurnEnd is emitted when a turn finishes, successfully or with an error.
	EventKindTurnEnd
	// EventKindLLMRequest is emitted before a provider chat request is made.
	EventKindLLMRequest
	// EventKindLLMDelta is emitted when a streaming provider yields a partial delta.
	EventKindLLMDelta
	// EventKindLLMResponse is emitted after a provider chat response is received.
	EventKindLLMResponse
	// EventKindLLMRetry is emitted when an LLM request is retried.
	EventKindLLMRetry
	// EventKindContextCompress is emitted when session history is forcibly compressed.
	EventKindContextCompress
	// EventKindSessionSummarize is emitted when asynchronous summarization completes.
	EventKindSessionSummarize
	// EventKindToolExecStart is emitted immediately before a tool executes.
	EventKindToolExecStart
	// EventKindToolExecEnd is emitted immediately after a tool finishes executing.
	EventKindToolExecEnd
	// EventKindToolExecSkipped is emitted when a queued tool call is skipped.
	EventKindToolExecSkipped
	// EventKindSteeringInjected is emitted when queued steering is injected into context.
	EventKindSteeringInjected
	// EventKindFollowUpQueued is emitted when an async tool queues a follow-up system message.
	EventKindFollowUpQueued
	// EventKindInterruptReceived is emitted when a soft interrupt message is accepted.
	EventKindInterruptReceived
	// EventKindSubTurnSpawn is emitted when a sub-turn is spawned.
	EventKindSubTurnSpawn
	// EventKindSubTurnEnd is emitted when a sub-turn finishes.
	EventKindSubTurnEnd
	// EventKindSubTurnResultDelivered is emitted when a sub-turn result is delivered.
	EventKindSubTurnResultDelivered
	// EventKindSubTurnOrphan is emitted when a sub-turn result cannot be delivered.
	EventKindSubTurnOrphan
	// EventKindError is emitted when a turn encounters an execution error.
	EventKindError

	eventKindCount
)

var eventKindNames = [...]string{
	"turn_start",
	"turn_end",
	"llm_request",
	"llm_delta",
	"llm_response",
	"llm_retry",
	"context_compress",
	"session_summarize",
	"tool_exec_start",
	"tool_exec_end",
	"tool_exec_skipped",
	"steering_injected",
	"follow_up_queued",
	"interrupt_received",
	"subturn_spawn",
	"subturn_end",
	"subturn_result_delivered",
	"subturn_orphan",
	"error",
}

// String returns the stable string form of an EventKind.
func (k EventKind) String() string {
	if k >= eventKindCount {
		return fmt.Sprintf("event_kind(%d)", k)
	}
	return eventKindNames[k]
}

// Event is the structured envelope broadcast by the agent EventBus.
//
// Deprecated: use github.com/sipeed/picoclaw/pkg/events.Event for new
// observation code. Agent payload types remain supported.
type Event struct {
	Kind    EventKind
	Time    time.Time
	Meta    EventMeta
	Context *TurnContext
	Payload any
}

// EventMeta contains correlation fields shared by all agent-loop events.
type EventMeta struct {
	AgentID      string
	TurnID       string
	ParentTurnID string
	SessionKey   string
	Iteration    int
	TracePath    string
	Source       string
	turnContext  *TurnContext
}
