package usecase

import (
	"testing"
	"tobialbertino/portfolio-be/internal/to_do/models/domain"
	"tobialbertino/portfolio-be/internal/to_do/models/entity"
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
		// 	Created_at: 0, // problem, maybe
		// 	Updated_at: 0,
		// }
		toDoRepo.Mock.On("Create", mock.Anything).Return(tt.want, nil)

		t.Run(tt.name, func(t *testing.T) {
			got, err := toDoUseCase.Create(tt.args.req)
			assert.Equal(t, tt.want, got, err)
		})
	}
}

func TestToDoUseCaseImpl_Delete(t *testing.T) {
	type args struct {
		id *int64
	}
	var Id *int64 = new(int64)
	*Id = 13
	tests := []struct {
		name    string
		args    args
		want    *domain.RowsAffected
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test-delete",
			args: args{
				id: Id,
			},
			want: &domain.RowsAffected{
				RowsAffected: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		// Program Mock
		toDoRepo.Mock.On("Delete", tt.args.id).Return(tt.want, nil)

		t.Run(tt.name, func(t *testing.T) {
			got, err := toDoUseCase.Delete(tt.args.id)
			assert.Equal(t, tt.want, got, err)
		})
	}
}

func TestToDoUseCaseImpl_DeleteAll(t *testing.T) {
	tests := []struct {
		name    string
		want    *domain.RowsAffected
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_deleteAll",
			want: &domain.RowsAffected{
				RowsAffected: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		// Program Mock
		toDoRepo.Mock.On("DeleteAll").Return(tt.want, nil)

		t.Run(tt.name, func(t *testing.T) {
			got, err := toDoUseCase.DeleteAll()
			assert.Equal(t, tt.want, got, err)
		})
	}
}

func TestToDoUseCaseImpl_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		want    *[]domain.ResponseToDo
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_GetAll",
			want: &[]domain.ResponseToDo{{
				Id:         0,
				Title:      "",
				Status:     false,
				Created_at: 0,
				Updated_at: 0,
			},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		// Program Mock
		var response = entity.ListToDo{
			{
				Id:         0,
				Title:      "",
				Status:     false,
				Created_at: 0,
				Updated_at: 0,
			},
		}

		toDoRepo.Mock.On("GetAll").Return(&response, nil)

		t.Run(tt.name, func(t *testing.T) {
			got, err := toDoUseCase.GetAll()
			assert.Equal(t, tt.want, got, err)
		})
	}
}

func TestToDoUseCaseImpl_Update(t *testing.T) {
	type args struct {
		req *domain.RequestUpdateToDo
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.RowsAffected
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_update",
			args: args{
				req: &domain.RequestUpdateToDo{
					Id:     12,
					Title:  "test",
					Status: false,
				},
			},
			want: &domain.RowsAffected{
				RowsAffected: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		// Program Mock
		// request := &entity.ToDo{
		// 	Id:         0,
		// 	Title:      "",
		// 	Status:     false,
		// 	Created_at: 0,
		// 	Updated_at: 0,
		// }
		toDoRepo.Mock.On("Update", mock.Anything).Return(1, nil)

		t.Run(tt.name, func(t *testing.T) {
			got, err := toDoUseCase.Update(tt.args.req)
			assert.Equal(t, tt.want, got, err)
		})
	}
}
