-- drop table if exists items;

create table if not exists items
(
    id          bigserial primary key,
    --
    order_id    bigint not null references orders (id) on delete cascade,
    seller_id   bigint not null,
    --
    price       int    not null,
    check (price > 0),
    discount    int,
    check (discount >= 0),
    --
    name        text   not null,
    --
    created_at  timestamp with time zone default now() not null,
    updated_at  timestamp with time zone default now() not null
);

create index on items (seller_id);
