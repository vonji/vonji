'use strict';

var _express = require('express');

var _express2 = _interopRequireDefault(_express);

var _knex = require('knex');

var _knex2 = _interopRequireDefault(_knex);

var _bookshelf = require('bookshelf');

var _bookshelf2 = _interopRequireDefault(_bookshelf);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

var app = (0, _express2.default)(); /**
                                     * Created by loup on 19/11/16.
                                     */

var knex = (0, _knex2.default)({
        client: 'pg',
        connection: {
            host: '127.0.0.1',
            user: 'vonji-api',
            password: 'vonji-api',
            database: 'vonji-api',
            charset: 'utf8'
        }
    });

var bookshelf = (0, _bookshelf2.default)(knex);

var User = bookshelf.Model.extend({
    tableName: 'users'
});

User.where('id', 1).fetch({withRelated: ['posts.tags']}).then(function (user) {
    console.log(user.related('posts').toJSON());
}).catch(function (err) {
    console.error(err);
});

app.get('/', function (req, res) {
    res.send('Hello World!');
});

app.listen(3000, function () {
    console.log('Example app listening on port 3000!');
});