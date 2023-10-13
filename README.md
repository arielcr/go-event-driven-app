# Go Event-Driven Microservices 

This monorepo contains a complete event-driven microservices system built in Go, showcasing concurrency and the use of AWS services like DynamoDB, SNS, and SQS. The project focuses on asynchronous event handling, allowing developers to understand and experiment with the power of Go for event-driven applications.

## Key Features:

- CLI for sending events to the microservices system.
- Microservice to store events in DynamoDB and publish them to an SNS topic.
- Multiple consumer services for processing events asynchronously, demonstrating Go's concurrency capabilities.
- Monorepo structure for easy management and collaboration.

## Getting Started:

Clone the repository.
Follow the setup instructions in the project documentation to configure AWS credentials and dependencies.
Use the CLI to send events to the microservices and see the system in action.
