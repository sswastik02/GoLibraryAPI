package models



type Book struct{
	ID 			uint		`gorm:"primary key;autoincrement" json:"id"`
	Title 		*string		`json:"title"`
	Author 		*string		`json:"author"`
	Publisher 	*string		`json:"publisher"`

}

// reason for using *string is That field might be nil. (Otherwise the default value is assigned "" for string)

