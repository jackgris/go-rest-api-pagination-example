# Run the back-end


### Build container

```
docker build -t go-todo-api .
```


### Run container


This should be after our database is running.

```
 docker run -p 3000:3000 go-todo-api
```

### Fake Data

You can test your endpoint, and fake data with apps like Postman, curl, etc. In this case I will use [Hey with dynamic request body](https://github.com/adhocore/hey/tree/master) I clone and compile that fork of the original [Hey](https://github.com/rakyll/hey)

And you can run this command, you will generate 200 records:
```
hey -m POST -d '{"title":"{s:5:10}","description":"{s:5:10}"}' "http://localhost:3000/v2/todos"
```
