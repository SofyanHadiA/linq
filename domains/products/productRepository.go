package products

import (
	"fmt"

	. "github.com/SofyanHadiA/linq/core/database"
	. "github.com/SofyanHadiA/linq/core/repository"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type productRepository struct {
	db IDB
}

func NewProductRepository(db IDB) IRepository {
	return productRepository{
		db: db,
	}
}

func (repo productRepository) CountAll() (int, error) {
	countQuery := "SELECT COUNT(*) FROM products WHERE deleted = 0"

	var result int
	row, err := repo.db.ResolveSingle(countQuery)
	row.Scan(&result)
	utils.HandleWarn(err)
	return result, err
}

func (repo productRepository) IsExist(id uuid.UUID) (bool, error) {
	isExistQuery := "SELECT EXISTS(SELECT * FROM products WHERE uid=? AND deleted = 0)"

	var result bool
	row, err := repo.db.ResolveSingle(isExistQuery, id)
	row.Scan(&result)
	return result, err
}

func (repo productRepository) GetAll(paging utils.Paging) (IModels, error) {
	query := "SELECT * FROM products WHERE deleted=0 "

	if paging.Keyword != "" {
		query += ` AND (title LIKE '%?%' OR code LIKE '%?%' OR buy_price LIKE '%?%' OR sell_price LIKE '%?%') `
	}

	if paging.Order > 0 {
		var columnMap string
		switch paging.Order {
		case 1:
			columnMap = "title"
		case 2:
			columnMap = "code"
		case 3:
			columnMap = "sell_price"
		default:
			columnMap = "stock"
		}

		query += fmt.Sprintf(" ORDER BY %s %s ", columnMap, paging.OrderDir)
	}

	if paging.Length > 0 {
		query += fmt.Sprintf(" LIMIT %d ", paging.Length)
	} else {
		query += " LIMIT 25 "
	}

	rows := &sqlx.Rows{}
	var err error

	if paging.Keyword != "" {
		rows, err = repo.db.Resolve(query, paging.Keyword)
	} else {
		rows, err = repo.db.Resolve(query)
	}
	utils.HandleWarn(err)

	result := Products{}

	for rows.Next() {
		var product = &Product{}
		err := rows.StructScan(&product)
		utils.HandleWarn(err)
		result = append(result, (*product))
	}

	return &result, err
}

func (repo productRepository) Get(id uuid.UUID) (IModel, error) {
	selectQuery := "SELECT * FROM products WHERE uid = ? AND deleted= 0 "

	product := &Product{}
	rows, err := repo.db.ResolveSingle(selectQuery, id)
	utils.HandleWarn(err)
	rows.StructScan(product)

	return product, err
}

func (repo productRepository) Insert(model IModel) error {

	insertQuery := `INSERT INTO products 
		(uid, title, buy_price, sell_price, stock, code, created ) 
		VALUES(:uid, :title, :buy_price, :sell_price, :stock, :code, now())`

	product := *model.(*Product)
	product.Uid = uuid.NewV4()
	_, err := repo.db.Execute(insertQuery, &product)

	return err
}

func (repo productRepository) Update(model IModel) error {
	updateQuery := `UPDATE products SET 
		title=:title, buy_price=:buy_price, sell_price=:sell_price, 
		stock=:stock, code=:code, updated=now() WHERE uid=:uid`

	product, _ := model.(*Product)
	_, err := repo.db.Execute(updateQuery, product)

	return err
}

func (repo productRepository) UpdateProductPhoto(model IModel) error {
	updateQuery := "UPDATE products SET image=:image WHERE uid=:uid"

	product, _ := model.(*Product)
	_, err := repo.db.Execute(updateQuery, product)

	return err
}

func (repo productRepository) Delete(model IModel) error {
	deleteQuery := "UPDATE products SET deleted=1 WHERE uid=:uid"

	product, _ := model.(*Product)
	_, err := repo.db.Execute(deleteQuery, product)

	return err
}

func (repo productRepository) DeleteBulk(products []uuid.UUID) error {
	deleteQuery := "UPDATE products SET deleted=1 WHERE uid IN(?)"
	_, err := repo.db.ExecuteBulk(deleteQuery, products)

	return err
}
