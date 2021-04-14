create sequence if not exists dealer_id_seq;

create or replace function public.next_id(out result bigint) as
$$
declare
    our_epoch  bigint := 1590969600000; -- 01/06/2020 00:00:00
    seq_id     bigint;
    now_millis bigint;
    shard_id   int    := 1; -- shard #1 будет схема public для упрощения
begin
    select nextval('public.dealer_id_seq') % 1024 into seq_id;
    select floor(extract(epoch from clock_timestamp()) * 1000) into now_millis;
    result := (now_millis - our_epoch) << (61 - 42);
    result := result | (shard_id << 9);
    result := result | (seq_id);
end;
$$ language plpgsql;

-- drop table if exists orders;

create table if not exists orders
(
    id              bigint primary key           default public.next_id(),
    --
    status              text                     not null,
    workflow            text                     not null,
    --
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null
);

create index on orders (status);
