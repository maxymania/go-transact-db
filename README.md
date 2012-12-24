go-transact-db
==============

Transactional DBM for Go based on godbm.

# Example

```go
db,err := transactdb.CreateDB(`db-file`,10) // create db-file with 2**10 buckets
// or open an existing db-file using: db,err := transactdb.OpenDB(`db-file`)

// Create a Transaction
t := db.NewTransaction()

// Do reads and writes
data,ok := t2.Read("A")
t.Write("A",[]byte("B"))

// Commit the work
ok = t.Commit()

// if ok is false the transaction failed.

```
