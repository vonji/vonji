/**
 * Created by loup on 19/11/16.
 */

import express from 'express';
import knexFactory from 'knex';
import bookshelfFactory from 'bookshelf';

const app = express();

const knex = knexFactory({
    client: 'pg',
    connection: {
        host: '127.0.0.1',
        user: 'vonji-api',
        password: 'vonji-api',
        database: 'vonji-api',
        charset: 'utf8'
    }
});

const bookshelf = bookshelfFactory(knex);

const User = bookshelf.Model.extend({
    tableName: 'users',
});

app.get('/', (req, res) => {
    res.send('Hello World!')
});

app.listen(3000, () => {
    console.log('Example app listening on port 3000!')
});