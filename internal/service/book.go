package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/wynyga/gotoko/domain"
	"github.com/wynyga/gotoko/dto"
)

type bookService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBook(bookRepoitory domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookService {
	return &bookService{
		bookRepository:      bookRepoitory,
		bookStockRepository: bookStockRepository,
	}
}

func (b bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	result, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var data []dto.BookData
	for _, v := range result {
		data = append(data, dto.BookData{
			Id:          v.Id,
			Isbn:        v.Isbn,
			Title:       v.Title,
			Description: v.Description,
		})
	}
	return data, nil
}

func (b bookService) Show(ctx context.Context, id string) (dto.BookShowData, error) {
	data, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return dto.BookShowData{}, err
	}
	if data.Id == "" {
		return dto.BookShowData{}, domain.BookNotFound
	}
	stocks, err := b.bookStockRepository.FindByBookId(ctx, data.Id)
	if err != nil {
		return dto.BookShowData{}, err
	}
	stocksData := make([]dto.BookStockData, 0)
	for _, v := range stocks {
		stocksData = append(stocksData, dto.BookStockData{
			Code:   v.Code,
			Status: v.Status,
		})
	}
	return dto.BookShowData{
		BookData: dto.BookData{
			Id:          data.Id,
			Isbn:        data.Isbn,
			Title:       data.Title,
			Description: data.Description,
		},
		Stocks: stocksData,
	}, nil
}

func (b bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	book := domain.Book{
		Id:          uuid.NewString(),
		Isbn:        req.Isbn,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}
	return b.bookRepository.Save(ctx, &book)
}

func (b bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := b.bookRepository.FindById(ctx, req.Id)
	if err != nil {
		return err
	}
	if persisted.Id == "" {
		return errors.New("data buku tidak ditemukan")
	}
	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}
	return b.bookRepository.Update(ctx, &persisted)
}

func (b bookService) Delete(ctx context.Context, id string) error {
	persisted, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if persisted.Id == "" {
		return errors.New("data buku tidak ditemukan")
	}
	err = b.bookRepository.Delete(ctx, persisted.Id)
	if err != nil {
		return err
	}
	return b.bookStockRepository.DeleteByBookId(ctx, persisted.Id)
}
