# Get MySQL database ready

## Download image

First, from the command line pull the image:

```
docker pull mysql:8.0.33
```

## Run our container database

Now we can create a container with name `mysql-nextjs` with our MySQL database, with this command you will have the root user with the super secure password: 1234.
```
docker run --name mysql-nextjs -e MYSQL_ROOT_PASSWORD=1234 -d mysql:8.0.33
```

## Access to our MySQL command line

We can run this command to access the bash terminal of the container:
```
docker exec -it mysql-nextjs bash
```

After that the command `mysql -p` and the password we use before, in this case `1234`

Or a better option for that, go strait to the MySQL client with this command:

```
docker exec -it mysql-nextjs mysql -p
```

## Create user in MySQL and get privileges

All this from the MySQL client.

-Create user:
```
CREATE USER 'pagination'@'localhost' IDENTIFIED BY '1234';
```

- In order to grant all privileges of the database for a newly created user, execute the following command:

```
GRANT ALL PRIVILEGES ON * . * TO 'pagination'@'localhost' WITH GRANT OPTION;
```

- Create the same user for allow access from outside:

```
CREATE USER 'pagination'@'%' IDENTIFIED BY '1234';
```

- In order to grant all privileges of the database for a newly created user, execute the following command:

```
GRANT ALL PRIVILEGES ON *.* TO 'pagination'@'%' WITH GRANT OPTION;
```

Of course if you want manage the privileges in a better way, you can follow some tutorial like this:
[How to Create MySQL User and Grant Privileges: A Beginner’s Guide](https://www.hostinger.com/tutorials/mysql/how-create-mysql-user-and-grant-permissions-command-line)

- For changes to take effect immediately flush these privileges by typing in the command:

```
FLUSH PRIVILEGES;
```

Once that is done, your new user account has the same access to the database as the root user.

## Create database

With this command will create the database:

```
CREATE DATABASE todo_pagination;
```

## Problems Solved

### failed to connect database

Maybe the IP is not right, because you are running the database in Docker and you are using the IP `127.0.0.1`, you need to figure out the IP. For that you can use this command:

```
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql-nextjs
```
That could bring you this IP `172.17.0.2` so, you need to use that IP address.

## Note:

Many of this, will be done automatically, but at beginning, for testing purpose its nice have everything running.
