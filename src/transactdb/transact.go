package transactdb

import "transactdb/godbm"
import "encoding/binary"
import "sync"
import "bytes"


type record struct{
	Key string
	Ver uint64
	Update bool
	Value []byte
}

type Database struct{
	DB *godbm.HashDB
	transaction sync.RWMutex
}

// Opens an existing database
func OpenDB(path string) (*Database,error){
	db,err := godbm.Open(path)
	return &Database{DB:db},err
}
// Creates a new database
func CreateDB(path string, nbuckets uint32) (*Database,error){
	db,err := godbm.Create(path,nbuckets)
	return &Database{DB:db},err
}

func (db *Database) readRecord(k string) (*record,bool){
	db.transaction.RLock()
	defer db.transaction.RUnlock()
	return db.getRecord(k)
}

func (db *Database) getRecord(k string) (*record,bool){
	var ver uint64
	rec := new(record)
	rec.Key = k
	rec.Update = false
	v,err := db.DB.Get([]byte(k))
	if err!=nil { return nil,false }
	buf := bytes.NewBuffer(v)
	binary.Read(buf,binary.BigEndian,&ver)
	rec.Ver = ver
	rec.Value = buf.Bytes()
	return rec,true
}

func (db *Database) setRecord(rec *record) error{
	buf := new(bytes.Buffer)
	binary.Write(buf,binary.BigEndian,rec.Ver)
	buf.Write(rec.Value)
	return db.DB.Set([]byte(rec.Key),buf.Bytes())
}

func (db *Database) commit(cache map[string]*record) bool{
	db.transaction.Lock()
	defer db.transaction.Unlock()
	for _,r := range cache {
		or,ok := db.getRecord(r.Key)
		if ok && (or.Ver!=r.Ver) { return false }
	}
	for _,r := range cache {
		if r.Update {
			// TODO: error handling
			db.setRecord(r)
		}
	}
	return true
}

// Creates a new Transaction Object
func (db *Database) NewTransaction() *Transaction{
	return &Transaction{db,make(map[string]*record)}
}

// A Transaction On a Database
type Transaction struct{
	*Database
	cache map[string]*record
}

func (t *Transaction) readRecord2(k string) (rec *record,ok bool){
	rec,ok = t.cache[k]
	if !ok {
		rec,ok = t.readRecord(k)
		if ok { t.cache[k]=rec }
	}
	return
}

// Reads a Record
func (t *Transaction) Read(k string) ([]byte,bool) {
	rec,ok := t.readRecord2(k)
	if !ok { return nil,false }
	return rec.Value,true
}

// Writes a Record
func (t *Transaction) Write(k string,v []byte) {
	rec,ok := t.readRecord2(k)
	if !ok {
		rec = new(record)
		rec.Key = k
		rec.Ver = 0
		t.cache[k]=rec
	}
	rec.Update = true
	rec.Value = v
}

// Commits a Transaction and returns true on success
func (t *Transaction) Commit() bool{
	return t.commit(t.cache)
}

// Clears the Transaction Buffer. This is like Retry
func (t *Transaction) Clear() {
	t.cache = make(map[string]*record)
}

