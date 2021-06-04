create schema if not exists transaction

    create table if not exists transaction.account
    (
        id              uuid not null primary key,
        document_number text,
        created_at      timestamp,
        updated_at      timestamp
    )

    create table if not exists transaction.transaction
    (
        id             uuid not null primary key,
        account_id     uuid references account (id),
        operation_type int references operation_types(id),
        amount         int,
        event_date     timestamp
    )

    create table if not exists transaction.operation_types
    (
        id          int,
        description text
    );

insert into operation_types
values (1, 'COMPRA A VISTA'),
       (2, 'COMPRA PARCELADA'),
       (3, 'SAQUE'),
       (4, 'PAGAMENTO')