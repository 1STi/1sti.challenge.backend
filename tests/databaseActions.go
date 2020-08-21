package tests

import (
	"github.com/YtaloWill/1sti.challenge.backend/database"
)

var createTables = `
CREATE TABLE IF NOT EXISTS tbUsers 
(id serial primary key, 
name varchar(50) not null, 
email varchar(70) not null);

CREATE TABLE IF NOT EXISTS tbStatus
(id serial primary key,
description varchar(20) not null);

CREATE TABLE IF NOT EXISTS tbTasks 
(id serial primary key, 
title varchar(140) not null, 
description varchar(250), 
iduser int references tbUsers(id), 
idStatus int references tbStatus(id));
`

var defaultStatus = `
INSERT INTO tbstatus VALUES 
(DEFAULT, 'A fazer'), 
(DEFAULT, 'Fazendo'), 
(DEFAULT, 'Feito');
`

func dropTables(){
	database.Db.QueryRow(`DROP TABLE tbtasks, tbusers, tbstatus;`)
}

func ClearTables(){
	database.Db.QueryRow(`DELETE FROM tbtasks, tbusers, tbstatus;`)
}

func setDefaulTableValues(){
	database.Db.QueryRow(defaultStatus)
}

func ensureTableExists(){
	database.Db.QueryRow(createTables)
}

func BuildTestDb(){
	database.ConnectDb()
	dropTables()
	ensureTableExists()
	setDefaulTableValues()
}