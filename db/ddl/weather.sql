create table weather
(
    id               serial
        constraint table_name_pk
            primary key
                deferrable,
    unix_date        integer,
    temperature      double precision,
    temperature_feel double precision,
    unix_sunrise     integer,
    unix_sunset      integer,
    sky              varchar(64)
);

