// Package events defines the standardized, multi-tenant NATS message payloads.
// This package establishes the explicit communication contract between microservices,
// ensuring a consistent and trackable data flow for the entire document processing
// pipeline.
// Adherence to these structures is mandatory for inter-service communication.
package events

import "time"

// EventHeader is the mandatory metadata structure embedded in all NATS events.
// It provides the necessary context for multi-tenancy, security, and distributed tracing.
type EventHeader struct {
	Timestamp  time.Time `json:"timestamp"`
	WorkflowID string    `json:"workflowId"`
	UserID     string    `json:"userId"`
	TenantID   string    `json:"tenantId"`
	EventID    string    `json:"eventId"`
}

// PDFCreatedEvent is published by the API Gateway to initiate a new workflow.
// It signals that a new PDF has been uploaded and is ready for processing.
type PDFCreatedEvent struct {
	Header EventHeader `json:"header"`
	// PDFKey is the unique identifier for the uploaded PDF in the object store.
	PDFKey string `json:"pdfKey"`
	// Augmentation describes the narrator augmentation preferences supplied by the user.
	Augmentation *AugmentationPreferences `json:"augmentation,omitempty"`
}

// PNGCreatedEvent is published by the pdf-to-png-service for each successfully
// rendered and non-blank page of a PDF.
type PNGCreatedEvent struct {
	Header EventHeader `json:"header"`
	// PNGKey is the unique identifier for the generated PNG in the object store.
	PNGKey string `json:"pngKey"`
	// PageNumber is the 1-based index of this page within the original document.
	PageNumber int `json:"pageNumber"`
	// TotalPages is the total number of pages in the original document.
	TotalPages int `json:"totalPages"`
	// Augmentation carries the per-workflow augmentation preferences from the upload event.
	Augmentation *AugmentationPreferences `json:"augmentation,omitempty"`
}

// SummaryPlacement identifies where page-level summaries should be inserted relative to OCR text.
type SummaryPlacement string

const (
	// SummaryPlacementTop requests that summaries precede the raw OCR text.
	SummaryPlacementTop SummaryPlacement = "top"
	// SummaryPlacementBottom requests that summaries appear after the OCR text.
	SummaryPlacementBottom SummaryPlacement = "bottom"
)

// AugmentationCommentarySettings controls commentary augmentation.
type AugmentationCommentarySettings struct {
	Enabled            bool   `json:"enabled"`
	CustomInstructions string `json:"customInstructions,omitempty"`
}

// AugmentationSummarySettings controls summary augmentation.
type AugmentationSummarySettings struct {
	Enabled            bool             `json:"enabled"`
	Placement          SummaryPlacement `json:"placement"`
	CustomInstructions string           `json:"customInstructions,omitempty"`
}

// AugmentationPreferences aggregates narration augmentation options for a workflow.
type AugmentationPreferences struct {
	Commentary AugmentationCommentarySettings `json:"commentary"`
	Summary    AugmentationSummarySettings    `json:"summary"`
}

// TextProcessedEvent is published by the png-to-text-service after successfully
// performing OCR and optional augmentation on a single page.
type TextProcessedEvent struct {
	Header EventHeader `json:"header"`
	// PNGKey is the identifier of the source PNG image that was processed.
	PNGKey string `json:"pngKey"`
	// TextKey is the unique identifier for the processed text in the object store.
	TextKey string `json:"textKey"`
	// PageNumber is the 1-based index of this page within the original document.
	PageNumber int `json:"pageNumber"`
	// TotalPages is the total number of pages in the original document.
	TotalPages int `json:"totalPages"`
	// Voice is the voice to be used for the text-to-speech conversion.
	Voice string `json:"voice,omitempty"`
	// Seed is the seed for the random number generator.
	Seed int `json:"seed,omitempty"`
	// NGL is the number of layers to offload to the GPU.
	NGL int `json:"ngl,omitempty"`
	// TopP is the top-p sampling value.
	TopP float64 `json:"topP,omitempty"`
	// RepetitionPenalty is the repetition penalty.
	RepetitionPenalty float64 `json:"repetitionPenalty,omitempty"`
	// Temperature is the temperature for the sampling.
	Temperature float64 `json:"temperature,omitempty"`
}

// AudioChunkCreatedEvent is published by the text-to-speech service for each
// successfully generated audio segment.
type AudioChunkCreatedEvent struct {
	Header EventHeader `json:"header"`
	// AudioKey is the unique identifier for the generated audio chunk in the object
	// store.
	AudioKey string `json:"audioKey"`
	// PageNumber is the 1-based index of this page within the original document.
	PageNumber int `json:"pageNumber"`
	// TotalPages is the total number of pages in the original document.
	TotalPages int `json:"totalPages"`
}

// WavFileCreatedEvent is published by the pcm-to-wav-service for each
// successfully generated wav file.
type WavFileCreatedEvent struct {
	Header EventHeader `json:"header"`
	// WavKey is the unique identifier for the generated wav file in the object
	// store.
	WavKey string `json:"wavKey"`
	// PageNumber is the 1-based index of this page within the original document.
	PageNumber int `json:"pageNumber"`
	// TotalPages is the total number of pages in the original document.
	TotalPages int `json:"totalPages"`
}

// FinalAudioCreatedEvent is published by the wav-aggregator-service after
// successfully combining all wav files for a workflow.
type FinalAudioCreatedEvent struct {
	Header EventHeader `json:"header"`
	// FinalAudioKey is the unique identifier for the combined audio file in the
	// object store.
	FinalAudioKey string `json:"finalAudioKey"`
}
