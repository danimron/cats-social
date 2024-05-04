package service

import (
	// "cats_social/exception"
	"cats_social/helper"
	"cats_social/model/domain"
	"cats_social/model/web"
	"cats_social/repository"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CatServiceImpl struct {
	CatRepository repository.CatRepository
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewCatService(catRepository repository.CatRepository, DB *sql.DB, validate *validator.Validate) CatService {
	return &CatServiceImpl{
		CatRepository: catRepository,
		DB:            DB,
		Validate:      validate,
	}
}

func (service *CatServiceImpl) Create(ctx context.Context, request web.CatCreateRequest) web.CatCreateResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	cat := domain.Cat{
		UserId:      request.UserId,
		Name:        request.Name,
		Race:        request.Race,
		Sex:         request.Sex,
		AgeInMonth:  request.AgeInMonth,
		Description: request.Description,
	}
	cat = service.CatRepository.Save(ctx, tx, cat)
	return helper.ToCategoryResponseCat(cat)
}

func (service *CatServiceImpl) Delete(ctx context.Context, CatId int) {
	db := service.DB
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	_, err = service.CatRepository.FindById(ctx, db, CatId)
	// if err != nil {
	// 	panic(exception.NewNotFoundError(err.Error()))
	// }
	// if book.Available == 0 {
	// 	message := errors.New("book is booked by someone cannot delete book")
	// 	panic(exception.NewFoundError(message.Error()))
	// }
	service.CatRepository.Delete(ctx, tx, CatId)
}

func (service *CatServiceImpl) Update(ctx context.Context, request web.CatCreateRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	cat := domain.Cat{
		Id:          request.Id,
		UserId:      request.UserId,
		Name:        request.Name,
		Race:        request.Race,
		Sex:         request.Sex,
		AgeInMonth:  request.AgeInMonth,
		Description: request.Description,
	}
	service.CatRepository.Update(ctx, tx, cat)
}

func (service *CatServiceImpl) FindAll(ctx context.Context, cat *web.CatGetParam) []web.CatGetResponse {
	db := service.DB
	sql := ""

	if cat.Id != "" {
		sql = AddCondition(sql) + "c.id = " + cat.Id
	}
	if cat.Race != "" {
		sql = AddCondition(sql) + "c.race = '" + cat.Race + "'"
	}
	if cat.Sex != "" {
		sql = AddCondition(sql) + "c.sex = '" + cat.Sex + "'"
	}
	if cat.Owned != "" {
		if strings.ToLower(cat.HasMatched) == "true" {
			sql = AddCondition(sql) + "c.user_id IS NOT NULL"
		} else {
			sql = AddCondition(sql) + "c.user_id IS NULL"
		}
	}
	if cat.HasMatched != "" {
		sql = " LEFT JOIN matchs m on (m.issuer_cat_id = c.id or m.receiver_cat_id = c.id) " + sql
		if strings.ToLower(cat.HasMatched) == "true" {
			sql = AddCondition(sql) + " m.id IS NOT NULL"
		} else {
			sql = AddCondition(sql) + " m.id IS NULL"
		}
	}
	if cat.AgeInMonth != "" {
		if strings.Contains(cat.AgeInMonth, ">") {
			sql = AddCondition(sql) + "c.age_in_months " + cat.AgeInMonth
		} else if strings.Contains(cat.AgeInMonth, "<") {
			sql = AddCondition(sql) + "c.age_in_months " + cat.AgeInMonth
		} else {
			sql = AddCondition(sql) + "c.age_in_months = " + cat.AgeInMonth
		}
	}
	if cat.Search != "" {
		sql = AddCondition(sql) + "c.name LIKE '%" + cat.Search + "%'"
	}
	if cat.Limit != "" {
		sql = sql + " LIMIT " + cat.Limit
	}
	if cat.Offset != "" {
		sql = sql + " OFFSET " + cat.Offset
	}
	sql = "SELECT c.id, c.name, c.race, c.sex, c.age_in_months, c.description, c.created_at from cats c" + sql
	fmt.Println(sql)

	cats := service.CatRepository.FindAll(ctx, db, sql)
	return helper.ToCategoryResponseCats(cats)
}

func AddCondition(sql string) string {
	finalSql := ""
	if sql == "" {
		finalSql = sql + " WHERE "
	} else {
		finalSql = sql + " AND "
	}
	return finalSql
}
