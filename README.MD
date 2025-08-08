# Job Queue Microservice

This project implements a job queue system using a clean architecture approach. It is designed to handle asynchronous tasks, such as sending emails after user registration etc.

## Services Overview

- **Adapter Service**: Connects to third-party services for sending and receiving data.
- **Auth Service**: Manages user authentication, including account operations, password resets, and OTP verification.
- **User Service**: Stores and manages user data.

## Workflow

1. **User Registration**: When a user registers, their data is stored via the User Service.
2. **Job Queue**: The system adds a job to the queue to send a confirmation email asynchronously.
3. **Email Sending**: The Adapter Service processes the job and interacts with third-party email providers to send the email.

## Features

- Decoupled architecture for scalability and maintainability.
- Asynchronous job processing for improved user experience.
- Secure authentication and user management.

## License

MIT