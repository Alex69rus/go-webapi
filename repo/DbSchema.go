package repo

var schema = `
create golang.table books(
    id SERIAL NOT NULL,
    title varchar(50) NOT NULL,
    primary key (id)
);

create table golang.authors(
    id SERIAL not null,
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    primary key (id)
);`
