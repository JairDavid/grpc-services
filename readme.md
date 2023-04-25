## GRPC Services
This repository contains 2 grpc services for learning purposes, it's just for understand how protocol buffers works
* Student service (unary)
  * Get students
  * Create students
* Exam service (unary)
  * Get exams
  * Create exams
    - Question section (bidirectional)
      * Get questions by exams
      * Create questions for exams (streaming)
    - Enrollment section (bidirectional)
      * Get enrollments by exams
      * Enroll student to exam

## What the services look like (hypothetical gateway):
![GRPC Services (1)](https://user-images.githubusercontent.com/67834146/234155210-d43e971e-dbea-4005-bd14-431be1f36642.png)

## Generate code from proto files:
```bash
make gen
```

## Running both services
* Student service (unary)
* Exam service (unary, bidirectional, streaming)
```bash
make up
```
