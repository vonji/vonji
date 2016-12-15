const moment = require('moment');
const Chance = require('chance');

exports.seed = (knex) => {
  return Promise.all([
    knex('requests').del().catch(console.error),
  ]).then(Promise.all([
    knex('users').del().catch(console.error),
    knex('achievements').del().catch(console.error),
    knex('ads').del().catch(console.error),
    knex('tags').del().catch(console.error),
  ])).then(Promise.all([
    knex('users').insert({
      id: 1,
      email: 'loup.peluso@vonji.fr',
      password: 'loup.peluso',
      displayed_name: 'Loup',
      real_name: 'Loup Peluso',
      description: 'Coucou les p\'tits lapins.',
      motto: 'J\'aime les pommes',
      birthday: moment('1985-04-03'),
      created_at: moment(),
      updated_at: moment()
    }).catch(console.error),
    knex('achievements').insert({
      id: 1,
      created_at: moment(),
      updated_at: moment(),
      award: 51,
      name: 'Blublub',
      description: 'fdjkslfds',
      category: 'tagzok',
      check_id: 5,
      check_data: 8,
    }).catch(console.error),
    knex('ads').insert({
      id: 1,
      created_at: moment(),
      updated_at: moment(),
      latitude: 51.12,
      longitude: 54.42,
      region: 'Paris',
      url: 'https://test.com',
      image_url: 'http://image.com',
      alt_text: 'Super text',
    }).catch(console.error),
    knex('tags').insert([
      {
        id: 1,
        created_at: moment(),
        updated_at: moment(),
        name: 'Tag',
        description: 'Description of tag',
      },
      {
        id: 2,
        created_at: moment(),
        updated_at: moment(),
        name: 'Tag2',
        description: 'Description of tag2',
      }
    ]).catch(console.error),
  ])).then(Promise.all([
    knex('requests').insert({
      id: 1,
      created_at: moment(),
      updated_at: moment(),
      user_id: 1,
      content: 'My super request',
      title: 'Super request',
      views: 451,
      status: 'Open',
      duration: 5,
      frequency: 2,
      frequency_unit: 'days',
      period_start: moment(),
      period_end: moment(),
      location: 'Paris',
    }).catch(console.error),
  ]));
};
