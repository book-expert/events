package events_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/book-expert/events"
)

// TestAugmentationPreferencesRoundTrip ensures preferences survive JSON marshal/unmarshal.
func TestAugmentationPreferencesRoundTrip(t *testing.T) {
	t.Parallel()

	original := events.AugmentationPreferences{
		Commentary: events.AugmentationCommentarySettings{
			Enabled:            true,
			CustomInstructions: "Describe every chart.",
		},
		Summary: events.AugmentationSummarySettings{
			Enabled:            true,
			Placement:          events.SummaryPlacementBottom,
			CustomInstructions: "Provide two sentence overview.",
		},
	}

	input := events.PNGCreatedEvent{
		Header: events.EventHeader{
			Timestamp:  time.Now().UTC(),
			WorkflowID: "workflow-123",
			UserID:     "user-456",
			TenantID:   "tenant-789",
			EventID:    "event-000",
		},
		PNGKey:       "tenant/workflow/page.png",
		PageNumber:   1,
		TotalPages:   10,
		Augmentation: &original,
	}

	encoded, marshalErr := json.Marshal(input)
	if marshalErr != nil {
		t.Fatalf("marshal failed: %v", marshalErr)
	}

	var decoded events.PNGCreatedEvent

	unmarshalErr := json.Unmarshal(encoded, &decoded)
	if unmarshalErr != nil {
		t.Fatalf("unmarshal failed: %v", unmarshalErr)
	}

	if decoded.Augmentation == nil {
		t.Fatalf("expected augmentation preferences to be present")
	}

	got := *decoded.Augmentation
	if got != original {
		t.Fatalf("unexpected round-trip value: got %+v, want %+v", got, original)
	}
}
