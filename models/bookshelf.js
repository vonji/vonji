import knexFactory from "knex";
import bookshelfFactory from "bookshelf";

export const knex = knexFactory({
    client: 'pg',
    connection: {
        host: '127.0.0.1',
        user: 'vonji-api',
        password: 'vonji-api',
        database: 'vonji-api-dev',
        charset: 'utf8'
    }
});

export const bs = bookshelfFactory(knex);