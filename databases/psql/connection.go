package psql

import (
	pg "github.com/jackc/pgx"
	"graphql_example/tools/convertions"
	"graphql_example/tools/viper"
	"sync"
	"log"
)

const PgDbName DbName = "PgDb"

type PgDb struct {
	loaded bool
	pool   *pg.ConnPool
	mutex  *sync.Mutex
}

// PgDbConn handles and stores the database connection
type PgDbConn struct {
	*pg.Conn
	pool *pg.ConnPool
}

// Get returns the database connection stored
func (d *PgDbConn) Get() interface{} {
	return d.Conn
}

// Close closes the database connection stored
func (d *PgDbConn) Close() error {
	d.pool.Release(d.Conn)
	return nil
}

// GetName returns the database name
func (d *PgDb) GetName() DbName {
	return PgDbName
}

func Conn() *PgDb {

	var (
		user     = viper.GetEnv("POSTGRES_USER")
		password = viper.GetEnv("POSTGRES_PASSWORD")
		port     = convertions.StringToInt64(viper.GetEnv("POSTGRES_PORT"))
		name     = viper.GetEnv("POSTGRES_NAME")
		host     = viper.GetEnv("POSTGRES_HOST")
		mxconn   = convertions.StringToInt64(viper.GetEnv("POSTGRES_MXCONN"))
		dbconn   = &PgDb{}
	)

	config := pg.ConnPoolConfig{ConnConfig: pg.ConnConfig{
		Host:     host,
		Port:     uint16(port),
		User:     user,
		Password: password,
		Database: name,
	}, MaxConnections: int(mxconn)}

	if sqldb, err := pg.NewConnPool(config); err != nil {
		log.Println(err)
		return nil
	} else {
		dbconn.pool  = sqldb
		dbconn.mutex = new(sync.Mutex)
		dbconn.loaded = true 
	}
	return dbconn
}

// Loaded returns if the database was loaded
func (d *PgDb) Loaded() bool {
	return d.loaded
}

// IsMaster returns if the connection is master
func (d *PgDb) IsMaster() bool {
	var isMaster bool

	var result string
	if err := d.pool.QueryRow("SELECT pg_is_in_recovery()").Scan(&result); err == nil {
		isMaster = (result == "f" || result == "false")
	}

	return isMaster
}

// Status returns if the connection is available
func (d *PgDb) Status() bool {
	var ok bool

	var result string
	if err := d.pool.QueryRow("SELECT pg_is_in_recovery()").Scan(&result); err == nil {
		ok = true
	}

	return ok
}

// Get returns a new connection of the database
func (d *PgDb) Get() dbConn {

	if _db, err := d.pool.Acquire(); err == nil {
		return &PgDbConn{_db, d.pool}
	}
	return nil
}