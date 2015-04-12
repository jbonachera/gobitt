# gobitt
Gobitt is a bittorrent tracker written in Go. 
This is still a work in progress, and should not be used in a production environnement.
Moreover, I used this project to learn Go, so beware, here be dragons. 

Ideas and suggestions are, of course, welcome !

# License
`Gobitt` is available under the [Beerware](http://en.wikipedia.org/wiki/Beerware) license.

If we meet some day, and you think this stuff is worth it, you can buy me a beer in return.

#How to run

Gobitt stores the peers and the hash they publish/want in a Database. Currently, two databases are supported: memory and mongodb. 

## In-Memory database
Memory is the default database, and is used if configuration does not include a "databaseplugin" directive, or if databaseplugin is set to "memory".
```
  go run cmd/gobitt/main.go
```

## MongoDB database
Gobitt can use a MongoDB server too: You will need a mongodb server running, and this example uses Docker to quickly start one.

```
cp config.ini.example config.ini 
cp mongodb.ini.example mongodb.ini
echo "databaseplugin = mongodb" >> config.ini
docker run  -d -p 27017:27017  mongo
go run cmd/gobitt/main.go
```

You can then use your favorite bittorent client to happily share your files !
