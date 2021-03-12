package storage

var insertPlaceQuery string = `INSERT INTO places(place)
                                VALUES ($1)`

var createTableQuery string = ` CREATE TABLE IF NOT EXISTS %s (
                                burger VARCHAR ( 50 ) UNIQUE NOT NULL,
                                price double precision NOT NULL,
                                supply VARCHAR (50) NOT NULL,
                                rating double precision NOT NULL,
                                offered_from VARCHAR (50) NOT NULL
                                )`

var insertBurgerQuery string = ` INSERT INTO %s(burger, price, supply, rating, offered_from)
                                VALUES ($1, $2, $3, $4, $5)`

var updateBurgerQuery string = `UPDATE %s 
                                SET price = %v,
                                supply = %d,
                                offered_from = '%s',
                                rating = %v
                                WHERE burger = '%s';`

var selectRandomPlace string = `SELECT *
                                FROM places
                                ORDER BY random()
                                LIMIT 1;`

var selectRandomBurger string = `SELECT burger, price, supply, offered_from, rating
                                FROM %s
                                ORDER BY random()
                                LIMIT 1;`

var selectPageLimit string = `SELECT burger, price, supply, offered_from, rating 
                                FROM %s
                                LIMIT %d
                                OFFSET %d;`
