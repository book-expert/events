# Events

## Project Summary

This package defines the standardized, multi-tenant NATS message payloads for the book-expert microservices ecosystem.

## Detailed Description

This package establishes the explicit communication contract between microservices, ensuring a consistent and trackable data flow for the entire document processing pipeline. Adherence to these structures is mandatory for inter-service communication.

The following events are defined:

-   `EventHeader`: The mandatory metadata structure embedded in all NATS events. It provides the necessary context for multi-tenancy, security, and distributed tracing.
-   `PDFCreatedEvent`: Published by the API Gateway to initiate a new workflow. It signals that a new PDF has been uploaded and is ready for processing.
-   `PNGCreatedEvent`: Published by the `pdf-to-png-service` for each successfully rendered and non-blank page of a PDF.
-   `TextProcessedEvent`: Published by the `png-to-text-service` after successfully performing OCR and optional augmentation on a single page.
-   `AudioChunkCreatedEvent`: Published by the `text-to-speech` service for each successfully generated audio segment.

## Technology Stack

-   **Programming Language:** Go 1.25

## Getting Started

### Prerequisites

-   Go 1.25 or later.

### Installation

To use this library in your project, you can use `go get`:

```bash
go get github.com/book-expert/events
```

## Usage

To use the event structures, import the package into your Go project:

```go
import "github.com/book-expert/events"

// Example of creating a PDFCreatedEvent
event := events.PDFCreatedEvent{
    Header: events.EventHeader{
        Timestamp:  time.Now(),
        WorkflowID: "workflow-123",
        UserID:     "user-456",
        TenantID:   "tenant-789",
        EventID:    "event-abc",
    },
    PDFKey: "path/to/your/pdf.pdf",
}
```

## Testing

This package only contains data structures and does not have any specific tests.

## License

Distributed under the MIT License. See the `LICENSE` file for more information.