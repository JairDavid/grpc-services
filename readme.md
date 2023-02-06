## GRPC Services
This repository contains 4 grpc services for learning purposes to understand how microservices work and the alternative (GRPC) instead of using message brokers like RabbitMQ or Kafka for microservices communication.
* Student service (unary)
  * Get students
  * Create students
* Exam service (unary)
  * Get exams
  * Create exams
* Question service (bidirectional)
  * Get questions by exams
  * Create questions for exams (streaming)
* Enrollment service (bidirectional)
  * Get enrollments by exams
  * Enroll student to exam

## What the services look like:
![GRPC Services](https://user-images.githubusercontent.com/67834146/216883949-b9f1d975-196d-48c9-af7a-e55b0621f256.png)
