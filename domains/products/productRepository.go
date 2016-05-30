package products

import (
	"fmt"

	. "github.com/SofyanHadiA/linq/core"
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
	if err != nil {
		return -1, err
	}
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
	query := `SELECT products.*, product_categories.title as cat_title FROM products JOIN product_categories ON products.category = product_categories.uid WHERE products.deleted=0 `

	if paging.Keyword != "" {
		query += " AND (products.title LIKE ? OR product_categories.title LIKE ? OR code LIKE ? OR buy_price LIKE ? OR sell_price LIKE ?) "
	}

	if paging.Order > 0 {
		var columnMap string
		switch paging.Order {
		case 1:
			columnMap = "`code`"
		case 2:
			columnMap = "`code`"
		case 3:
			columnMap = "title"
		case 4:
			columnMap = "cat_title"
		case 5:
			columnMap = "sell_price"
		case 6:
			columnMap = "stock"
		default:
			columnMap = "created"
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
		keyword := "%" + paging.Keyword + "%"
		rows, err = repo.db.Resolve(query, keyword, keyword, keyword, keyword, keyword)
	} else {
		rows, err = repo.db.Resolve(query)
	}

	if err != nil {
		return nil, err
	}

	result := Products{}

	for rows.Next() {
		var product = &Product{}
		err := rows.StructScan(&product)
		if err != nil {
			return nil, err
		}

		product.Category, err = repo.getCategory(product.CategoryId)
		if err != nil {
			return nil, err
		}

		result = append(result, (*product))
	}

	return &result, err
}

func (repo productRepository) Get(id uuid.UUID) (IModel, error) {
	selectQuery := "SELECT * FROM products WHERE uid = ? AND deleted = 0 "

	product := &Product{}
	rows, err := repo.db.ResolveSingle(selectQuery, id)
	if err != nil {
		return nil, err
	}
	rows.StructScan(product)

	product.Category, err = repo.getCategory(product.CategoryId)
	if err != nil {
		return nil, err
	}

	return product, err
}

func (repo productRepository) getCategory(categoryid uuid.UUID) (ProductCategory, error) {
	selectQuery := "SELECT * FROM product_categories WHERE uid = ? AND deleted= 0 "

	productCategory := ProductCategory{}
	rows, err := repo.db.ResolveSingle(selectQuery, categoryid)
	if err != nil {
		return productCategory, err
	}
	rows.StructScan(&productCategory)

	return productCategory, err
}

func (repo productRepository) Insert(model IModel) error {

	insertQuery := `INSERT INTO products 
		(uid, title, buy_price, sell_price, stock, code, category, created ) 
		VALUES(:uid, :title, :buy_price, :sell_price, :stock, :code, :category, now())`

	product := model.(*Product)
	product.Uid = uuid.NewV4()

	_, err := repo.db.Execute(insertQuery, product)

	return err
}

func (repo productRepository) Update(model IModel) error {
	updateQuery := `UPDATE products SET 
		title=:title, buy_price=:buy_price, sell_price=:sell_price, 
		stock=:stock, code=:code, category=:category, updated=now() WHERE uid=:uid`

	_, err := repo.db.Execute(updateQuery, model)

	return err
}

func (repo productRepository) UpdateProductPhoto(model IModel) error {
	updateQuery := "UPDATE products SET image=:image WHERE uid=:uid"

	_, err := repo.db.Execute(updateQuery, model)

	return err
}

func (repo productRepository) Delete(model IModel) error {
	deleteQuery := "UPDATE products SET deleted=1 WHERE uid=:uid"

	_, err := repo.db.Execute(deleteQuery, model)

	return err
}

func (repo productRepository) DeleteBulk(products []uuid.UUID) error {
	deleteQuery := "UPDATE products SET deleted=1 WHERE uid IN(?)"

	_, err := repo.db.ExecuteBulk(deleteQuery, products)

	return err
}
