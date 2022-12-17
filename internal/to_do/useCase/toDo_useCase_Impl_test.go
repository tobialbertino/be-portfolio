package usecase

import (
	"testing"
	"tobialbertino/portfolio-be/internal/to_do/models/domain"
	"tobialbertino/portfolio-be/internal/to_do/repository"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var validate *validator.Validate = validator.New()
var toDoRepo = repository.ToDoRepositoryMock{
	Mock: mock.Mock{},
}
var toDoUseCase = NewToDoUseCase(&toDoRepo, nil, validate)

func TestToDoUseCaseImpl_Create(t *testing.T) {
	type args struct {
		req *domain.RequestToDo
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.RowsAffected
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test-Input-Data",
			args: args{
				req: &domain.RequestToDo{
					Title: "New Title",
				},
			},
			want: &domain.RowsAffected{
				RowsAffected: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		// Program Mock -> entity input must be same (sensitive data must same)
		// entityRepo := entity.ToDo{
		// 	Id:         0,
		// 	Title:      tt.args.req.Title,
		// 	Status:     false,
		// 	Created_at: 0,
		// 	Updated_at: 0,
		// }
		toDoRepo.Mock.On("Create", mock.Anything).Return(tt.want, nil)

		t.Run(tt.name, func(t *testing.T) {
			got, err := toDoUseCase.Create(tt.args.req)
			assert.Equal(t, tt.want, got, err)
		})
	}
}
