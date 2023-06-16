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
