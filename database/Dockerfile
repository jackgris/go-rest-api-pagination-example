FROM mysql:8.0.33

ENV MYSQL_DATABASE=todo_pagination
ENV MYSQL_USER=pagination
ENV MYSQL_PASSWORD=1234
ENV MYSQL_ROOT_PASSWORD=1234
ENV MYSQL_PORT=3306

ADD ./database/data.sql /docker-entrypoint-initdb.d/

EXPOSE 3306
