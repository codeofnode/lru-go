# lru-go
lru implementation in golang

# API docs
1. First init a cache with size
```
c, er := lru.New(<size>)
```
c, the cache instance
er, if cache creation had any error

2. Add to cache
```
wasThereAnyEviction = c.Add(<any_string_value>)
```
wasThereAnyEviction, (bool) that tells whether there was an eviction for adding cache

3. Query to cache
```
val, ok = c.Query(<any_string_value>)
```
val, the value against the query (defaults to empty string "", if in case query not found)
ok, if query has corresponding value or not

# dir structure
* main.go -> the lru cache
* main_test.go -> the test files

# dependencies
* go 1.14+

## dev dependencies
* make 4.1+
* inotifywait 3.14

# testing
All modules has their `*_test` files as UT.
run tests with

```
make test
```

# TODO
* complete testing with full coverage
* use mutex locks for concurrency protection
* use separate key value for storing into in cache
* use persistent storage
* more to discuss
