# cache
cache package contains in-memory cache implementation

## Installation

~~~~
go get github.com/viking311/cache 
~~~~

## Usage

To create new cache instance you need to call the method NewCache().
The cache instanse provide next methods:

- Set(key string, value interface{}, ttl time.Duration)
- Get(key string) interface{}
- Delete(key string)

## Example of usage
~~~~
cache := cache.NewCache()
cache.Set("key1", "value1")
fmt.Println(cache.Get("key1"))
cache.Delete("key1")
~~~~

