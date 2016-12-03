const moment = require('moment');

exports.seed = function (knex, Promise) {
    return Promise.all([
        knex('users').del()
            .then(function () {
                return Promise.all([
                    knex('users').insert({
                        id: 1,
                        email: 'loup.peluso@vonji.fr',
                        password: 'loup.peluso',
                        displayed_name: 'Loup',
                        real_name: 'Loup Peluso',
                        description: 'Coucou les p\'tits lapins.',
                        motto: "J'aime les pommes",
                        birthday: moment('1985-04-03'),
                        created_at: moment(),
                        updated_at: moment()
                    }),
                ]);
            }),
        knex('achievements').del()
            .then(function () {
                return Promise.all([
                    knex('achievements').insert({
                        id: 1,
                        created_at: moment(),
                        updated_at: moment(),
                        award: 51,
                        name: "Blublub",
                        description: "fdjkslfds",
                        category: "tagzok",
                        check_id: 5,
                        check_data: 8,
                    }),
                ]);
            }),
    ]);
};
