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

# License

Copyright (c) 2012 maxymania

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
