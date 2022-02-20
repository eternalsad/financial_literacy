CREATE TABLE "transactions" (
	"id"	INTEGER,
	"transaction_name"	TEXT,
	"amount"	NUMERIC,
	"transaction_date"	DATETIME,
	"isIncome"	BOOLEAN,
	"comment"	TEXT,
	"category_id"	INTEGER,
	FOREIGN KEY("category_id") REFERENCES "categories",
	PRIMARY KEY("id" AUTOINCREMENT)
);