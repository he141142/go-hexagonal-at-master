package form

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/todo"
	"hex-base/internal/core/domain"
	"sync"
)

func (store *formStorage) GetById(ctx context.Context, formId uint) (*domain.Form, error) {
	_form, err := store.client.Form.
		Query().
		Where(form.ID(int64(formId))).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return &domain.Form{
		Id:        uint(_form.ID),
		Category:  _form.Category,
		IsDeleted: _form.IsDeleted,
		Status:    _form.Status,
		TodoID:    uint(_form.TodoID),
		Title:     _form.Title,
		Todo:      nil,
	}, nil
}

func (store *formStorage) ListByTodoId(ctx context.Context, todoId uint) (*domain.FormList, error) {
	FormDomain := make([]*domain.Form, 0)
	forms, err := store.client.Form.Query().Where(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(todo.FieldID), todoId))
	}).All(ctx)

	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	doneChan := make(chan interface{})
	wg.Add(len(forms))

	getTodo := func(done <-chan interface{}, form *ent.Form, wg *sync.WaitGroup) (<-chan *ent.Todo, <-chan error) {
		todoChan := make(chan *ent.Todo)
		errChan := make(chan error)
		go func() {
			defer wg.Done()
			todo, err_ := form.QueryTodo().Only(ctx)
			for {
				select {
				case <-done:
					return
				case todoChan <- todo:
				case errChan <- err_:
					fmt.Println(err_.Error())
				}
			}
		}()
		return todoChan, errChan
	}
	for _, form := range forms {
		_form := &domain.Form{}
		_form.Id = uint(form.ID)
		_form.Status = form.Status
		_form.Title = form.Title
		_form.IsDeleted = form.IsDeleted
		_form.Category = form.Category
		todoChan, errChan := getTodo(doneChan, form, &wg)
		if _err, ok := <-errChan; ok && _err != nil {
			return nil, _err
		}
		if _todo, ok := <-todoChan; ok {
			todoid := uint(_todo.ID)
			_form.TodoID = todoid
			_form.Todo = domain.NewTodo(todoid, _todo.Name, _todo.Task)
			FormDomain = append(FormDomain, _form)
		}
	}

	close(doneChan)
	wg.Wait()

	return &domain.FormList{
		Data: FormDomain,
	}, nil
}
