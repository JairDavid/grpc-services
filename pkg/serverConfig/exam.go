package serverConfig

import (
	"context"
	"github.com/JairDavid/go-grpc-intro/pkg/domain"
	"github.com/JairDavid/go-grpc-intro/pkg/domain/exampb"
	"github.com/JairDavid/go-grpc-intro/pkg/domain/repository"
	"github.com/JairDavid/go-grpc-intro/pkg/domain/studentpb"
	"io"
)

type ExamServer struct {
	repository repository.Repository
	exampb.UnimplementedExamServiceServer
}

func NewExamServer(repo repository.Repository) *ExamServer {
	return &ExamServer{repository: repo}
}

func (s *ExamServer) GetExam(ctx context.Context, getExamRequest *exampb.GetExamRequest) (*exampb.Exam, error) {
	exam, err := s.repository.GetExam(ctx, getExamRequest.GetId())
	if err != nil {
		return nil, err
	}

	return &exampb.Exam{
		Id:   getExamRequest.GetId(),
		Name: exam.Name,
	}, nil
}

func (s *ExamServer) SetExam(ctx context.Context, examRequest *exampb.Exam) (*exampb.SetExamResponse, error) {
	exam := &domain.Exam{
		Id:   examRequest.GetId(),
		Name: examRequest.GetName(),
	}
	err := s.repository.SetExam(ctx, exam)
	if err != nil {
		return nil, err
	}

	return &exampb.SetExamResponse{
		Id:   examRequest.GetId(),
		Name: examRequest.GetName(),
	}, nil
}

func (s *ExamServer) SetQuestions(stream exampb.ExamService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&exampb.SetQuestionResponse{Ok: true})
		}
		if err != nil {
			return err
		}

		question := &domain.Question{
			Id:       msg.GetId(),
			Answer:   msg.GetAnswer(),
			Question: msg.GetQuestion(),
			ExamId:   msg.GetExamId(),
		}
		err = s.repository.SetQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&exampb.SetQuestionResponse{Ok: false})
		}
	}

}

func (e *ExamServer) SetEnrollment(stream exampb.ExamService_SetEnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&exampb.SetQuestionResponse{Ok: true})
		}
		if err != nil {
			return err
		}

		enrollment := &domain.Enrollment{
			StudentId: msg.GetStudentId(),
			ExamId:    msg.GetExamId(),
		}
		err = e.repository.SetEnrollment(context.Background(), enrollment)
		if err != nil {
			return stream.SendAndClose(&exampb.SetQuestionResponse{Ok: false})
		}
	}
}

func (e *ExamServer) GetStudentPerExam(req *exampb.EnrollmentRequest, stream exampb.ExamService_GetStudentsPerExamServer) error {
	students, err := e.repository.GetStudentPerExam(context.Background(), req.GetExamId())
	if err != nil {
		return err
	}

	for _, student := range students {
		student := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}
		err := stream.Send(student)
		if err != nil {
			return err
		}
	}
	return nil
}
