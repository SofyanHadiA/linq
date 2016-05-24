package sales

import (
	"fmt"

	. "github.com/SofyanHadiA/linq/core/database"
	. "github.com/SofyanHadiA/linq/core/repository"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type saleRepository struct {
	db IDB
}

func NewSaleRepository(db IDB) IRepository {
	return saleRepository{
		db: db,
	}
}

func (repo saleRepository) CountAll() (int, error) {
	countQuery := "SELECT COUNT(*) FROM sales WHERE deleted = 0"

	var result int
	row, err := repo.db.ResolveSingle(countQuery)
	row.Scan(&result)
	if err != nil {
		return -1, err
	}
	return result, err
}

func (repo saleRepository) IsExist(id uuid.UUID) (bool, error) {
	isExistQuery := "SELECT EXISTS(SELECT * FROM sales WHERE uid=? AND deleted = 0)"

	var result bool
	row, err := repo.db.ResolveSingle(isExistQuery, id)
	row.Scan(&result)
	return result, err
}

func (repo saleRepository) GetAll(paging utils.Paging) (IModels, error) {
	query := "SELECT * FROM sales WHERE deleted=0 "

	// if paging.Keyword != "" {
	// 	query += ` AND (title LIKE '%?%' OR code LIKE '%?%' OR buy_price LIKE '%?%' OR sell_price LIKE '%?%') `
	// }

	// if paging.Order > 0 {
	// 	var columnMap string
	// 	switch paging.Order {
	// 	case 1:
	// 		columnMap = "title"
	// 	case 2:
	// 		columnMap = "code"
	// 	case 3:
	// 		columnMap = "sell_price"
	// 	default:
	// 		columnMap = "created"
	// 	}

	// 	query += fmt.Sprintf(" ORDER BY %s %s ", columnMap, paging.OrderDir)
	// }

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
	if err != nil {
		return nil, err
	}

	result := Sales{}

	for rows.Next() {
		var sale = &Sale{}
		err := rows.StructScan(&sale)
		if err != nil {
			return nil, err
		}

		sale.Detail, err = repo.getDetail(sale.Uid)
		if err != nil {
			return nil, err
		}

		result = append(result, (*sale))
	}

	return &result, err
}

func (repo saleRepository) Get(id uuid.UUID) (IModel, error) {
	selectQuery := "SELECT * FROM sales WHERE uid = ? AND deleted= 0 "

	sale := &Sale{}
	rows, err := repo.db.ResolveSingle(selectQuery, id)
	if err != nil {
		return nil, err
	}
	rows.StructScan(sale)

	sale.Detail, err = repo.getDetail(sale.Uid)
	if err != nil {
		return nil, err
	}

	return sale, err
}

func (repo saleRepository) getDetail(uid uuid.UUID) (SaleDetails, error) {
	selectQuery := "SELECT * FROM sale_details WHERE uid = ? AND deleted= 0 "

	rows, err := repo.db.Resolve(selectQuery)
	if err != nil {
		return nil, err
	}

	details := SaleDetails{}

	for rows.Next() {
		var detail = &SaleDetail{}
		err := rows.StructScan(&detail)
		if err != nil {
			return nil, err
		}

		details = append(details, *detail)
	}

	return details, err
}

func (repo saleRepository) Insert(model IModel) error {
	insertQuery := `INSERT INTO sales 
		(uid, customer, user, discount, discount_type, total, total_payment, payment_type, note, created ) 
		VALUES(:uid, :customer, :user, :discount, :discount_type, :total, :total_payment, :payment_type, :note, now())`

	sale := model.(*Sale)
	sale.Uid = uuid.NewV4()

	_, err := repo.db.Execute(insertQuery, sale)

	return err
}

func (repo saleRepository) Update(model IModel) error {
	updateQuery := `UPDATE sales SET 
		customer=:customer, user=:user, discount=:discount, discount_type=:discount_type, total=:total, 
		total_payment=:total_payment, payment_type=:payment_type, note=:note, updated=now() WHERE uid=:uid`

	_, err := repo.db.Execute(updateQuery, model)

	return err
}


func (repo saleRepository) Delete(model IModel) error {
	deleteQuery := "UPDATE sales SET deleted=1 WHERE uid=:uid"

	_, err := repo.db.Execute(deleteQuery, model)

	return err
}

func (repo saleRepository) DeleteBulk(sales []uuid.UUID) error {
	deleteQuery := "UPDATE sales SET deleted=1 WHERE uid IN(?)"

	_, err := repo.db.ExecuteBulk(deleteQuery, sales)

	return err
}
