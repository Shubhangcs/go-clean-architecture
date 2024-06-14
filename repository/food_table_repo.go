package repository

import "database/sql"

type FoodTableRepo struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *FoodTableRepo{
	return &FoodTableRepo{
		db: db,
	}
}

func (f FoodTableRepo) CreateFoodTable() (string , error){
	_ , err := f.db.Exec("CREATE TABLE FOODSSS(ID VARCHAR(30))")
	if err != nil{
		return "" , err
	}
	return "Table Created Successfully" , err
}