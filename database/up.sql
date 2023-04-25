DROP DATABASE IF EXISTS school;

CREATE TABLE student(
id varchar(32) PRIMARY KEY,
name varchar(50) NOT NULL,
age INTEGER NOT NULL
);

CREATE TABLE exam(
    id varchar(32) PRIMARY KEY,
    name varchar(50) NOT NULL
);

CREATE TABLE question(
    id varchar(21) PRIMARY KEY,
    answer varchar (60) NOT NULL,
    question varchar (60) NOT NULL,
    exam_id varchar(32) NOT NULL,
    FOREIGN KEY (exam_id) REFERENCES exam(id)
);

CREATE TABLE enrollment(
    exam_id varchar(32) NOT NULL,
    student_id varchar(32) NOT NULL,
    FOREIGN KEY (exam_id) REFERENCES exam(id),
    FOREIGN KEY (student_id) REFERENCES student(id)
)
