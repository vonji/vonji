const belongsToUser = (table, fieldName) => {
    table.integer(fieldName || 'user_id').references('users.id');
};

const belongsToRequest = (table) => {
    table.integer('request_id').references('requests.id');
};

exports.up = (knex) => {
    return Promise.all([
        knex.schema.createTable('users', table => {
            // has many achievements
            // has many tags
            table.increments().primary();
            table.timestamps();
            table.string('email').unique();
            table.string('password');
            table.string('displayed_name');
            table.string('real_name');
            table.string('description');
            table.string('motto');
            table.string('facebook_link');
            table.string('twitter_link');
            table.string('linkedin_link');
            table.string('phone');
            table.date('birthday');
            table.string('location');
            table.integer('vcoins');
            table.integer('vactions');
            table.string('avatar');
            table.string('gender');
        }),
        knex.schema.createTable('achievements', table => {
            table.increments().primary();
            table.timestamps();
            table.integer('award');
            table.string('name');
            table.string('description');
            table.string('category');
            table.integer('check_id');
            table.integer('check_data');
        }),
        knex.schema.createTable('ads', table => {
            table.increments().primary();
            table.timestamps();
            table.float('latitude');
            table.float('longitude');
            table.string('region');
            table.string('url');
            table.string('image-url');
            table.string('alt-text');
        }),
        knex.schema.createTable('tags', table => {
            table.increments().primary();
            table.timestamps();
            table.string('name');
            table.string('description');
        }),
    ]).then(() => {
        return Promise.all([
            knex.schema.createTable('comments', table => {
                // has many comments
                table.increments().primary();
                table.timestamps();
                belongsToUser(table);
                table.string('content');
                table.string('type');
                table.integer('commentable_id').references([
                    'comments.id',
                    'responses.id',
                    'requests.id',
                ]);
                table.enum('commentable_type', [
                    'comment',
                    'response',
                    'request',
                ]);
            }),
            knex.schema.createTable('notifications', table => {
                table.increments().primary();
                table.timestamps();
                belongsToUser(table);
                table.string('title');
                table.string('message');
                table.boolean('read');
            }),
            knex.schema.createTable('requests', table => {
                // has many responses
                // has many tags
                table.increments().primary();
                table.timestamps();
                belongsToUser(table);
                table.string('content');
                table.string('title');
                table.string('views');
                table.string('status');
                table.integer('duration');
                table.integer('frequency');
                table.string('frequency_unit');
                table.string('period_start');
                table.string('period_end');
                table.string('location');
            }),
            knex.schema.createTable('transactions', table => {
                table.increments().primary();
                table.timestamps();
                belongsToUser(table, 'from_user_id');
                belongsToUser(table, 'to_user_id');
                table.string('reason');
                table.string('source');
                table.enum('type', ['VCOIN', 'VACTION']);
                table.integer('amount');
            }),
        ]);
    }).then(() => {
        return Promise.all([
            knex.schema.createTable('tags_requests', table => {
                table.integer('tag_id').references('tags.id');
                table.integer('request_id').references('requests.id');
            }),
            knex.schema.createTable('responses', table => {
                // belongs to a request
                table.increments().primary();
                table.timestamps();
                belongsToUser(table);
                belongsToRequest(table);
                table.string('content');
                table.integer('value');
                table.boolean('accepted');
                table.integer('rating');
            }),
        ]);
    });
};

exports.down = function(knex) {
    return Promise.all([
        knex.schema.dropTable('tags_requests'),
        knex.schema.dropTable('responses'),
    ]).then(() => {
        return Promise.all([
            knex.schema.dropTable('comments'),
            knex.schema.dropTable('requests'),
            knex.schema.dropTable('transactions'),
            knex.schema.dropTable('notifications'),
        ]);
    }).then(() => {
        return Promise.all([
            knex.schema.dropTable('users'),
            knex.schema.dropTable('achievements'),
            knex.schema.dropTable('ads'),
            knex.schema.dropTable('tags'),
        ]);
    });
};
