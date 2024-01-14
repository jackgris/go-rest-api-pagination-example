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

### Run with a volume

If you want, can add a volume to your container.

#### Create volume

First we need to create the volume that you want to use if you didn't before. Using this command:

```
docker volume create pagination-volume
```

And you can see information with this command:

```
docker volume inspect pagination-volume
```

#### Run it

Now you can use that volume, in this way:
```
docker run --name database-server-pagination -p 3306:3306 -v pagination-volume:/pagination-volume -e MYSQL_ROOT_PASSWORD=1234 -d jackgris/pagination-todos-mysql:0.0.1
```

Note: You can use another volume that you create before, for see all the volumes you can use this command:
```
docker volume ls
```


## Access to our MySQL command line

We can run this command to access the bash terminal of the container:
```
docker exec -it mysql-nextjs bash
```

After that the command `mysql -p` and the password we use before, in this case `1234`

Or a better option for that, go straight to the MySQL client with this command:

```
docker exec -it mysql-nextjs mysql -p
```

## Create a user in MySQL and get privileges

All from the MySQL client.

-Create user:
```
CREATE USER 'pagination'@'localhost' IDENTIFIED BY '1234';
```

- In order to grant all privileges of the database to a newly created user, execute the following command:

```
GRANT ALL PRIVILEGES ON * . * TO 'pagination'@'localhost' WITH GRANT OPTION;
```

- Create the same user to allow access from outside:

```
CREATE USER 'pagination'@'%' IDENTIFIED BY '1234';
```

- In order to grant all privileges of the database to a newly created user, execute the following command:

```
GRANT ALL PRIVILEGES ON *.* TO 'pagination'@'%' WITH GRANT OPTION;
```

Of course, if you want to manage the privileges in a better way, you can follow some tutorials like this:
[How to Create MySQL User and Grant Privileges: A Beginner’s Guide](https://www.hostinger.com/tutorials/mysql/how-create-mysql-user-and-grant-permissions-command-line)

- For changes to take effect immediately flush these privileges by typing in the command:

```
FLUSH PRIVILEGES;
```

Once that is done, your new user account has the same access to the database as the root user.

### For example, the user can have remote access from anywhere, but they cannot drop any tables or databases.

####  Connect to the MySQL container
```bash
docker exec -it mysql-container mysql -u root -p
```
####  Create a new user with read and write privileges

```sql
CREATE USER 'your_remote_user'@'%' IDENTIFIED BY 'your_remote_password';
```

#### Grant SELECT, INSERT, UPDATE, and DELETE privileges on a specific database

```sql
GRANT SELECT, INSERT, UPDATE, DELETE ON your_database.* TO 'your_remote_user'@'%';
```

In MySQL, the `%` symbol in the context of user host specifications represents a wildcard character for the hostname. When used in the context of user creation or privilege assignment, % allows the user to connect from any host.

#### Flush privileges to apply changes
```sql
FLUSH PRIVILEGES;
```

## Create database

With this command will create the database:

```
CREATE DATABASE todo_pagination;
```

## Solve problems

### failed to connect database

Maybe the IP is not right, because you are running the database in Docker and you are using the IP `127.0.0.1`, you need to figure out the IP. For that, you can use this command:

```
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql-nextjs
```
That could bring you this IP `172.17.0.2` so, you need to use that IP address.

## Note:

Many of these steps will be done automatically, but at first, for the purpose of testing, it's nice to have everything running at least manually.
