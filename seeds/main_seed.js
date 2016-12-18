const moment = require('moment');
const Chance = require('chance');

exports.seed = (knex) => {
	const sknex = (resource) => ({
		del: function () {
			return knex(resource).del()
				.then(() => {
					console.log(`All ${resource} has been removed.`)
				})
				.catch(console.error);
		},
		insert: function (data) {
			return knex(resource).insert(data)
				.then(() => {
					console.log(`All ${resource} has been inserted.`)
				})
				.catch(console.error);
		},
	});

	return Promise.resolve().then(() => {
		return Promise.all([
			sknex('responses').del(),
			sknex('tags_requests').del(),
		]);
	}).then(() => {
		return Promise.all([
			sknex('requests').del(),
			sknex('transactions').del(),
			sknex('notifications').del(),
		]);
	}).then(() => {
		return Promise.all([
			sknex('users').del(),
			sknex('achievements').del(),
			sknex('ads').del(),
			sknex('tags').del(),
		]);
	}).then(() => {
		return Promise.all([
			sknex('users').insert([
				{
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
				},
				{
					id: 2,
					email: 'david.lancar@vonji.fr',
					password: 'david.lancar',
					displayed_name: 'David',
					real_name: 'David Lancar',
					description: 'Vfjdksl fjdkqsm fjdq fieozm jfejf',
					motto: 'EIrezk fdjksf eziomqrem dskfd kl',
					birthday: moment('1985-04-03'),
					created_at: moment(),
					updated_at: moment()
				},
			]),
			sknex('achievements').insert([
				{
					id: 1,
					created_at: moment(),
					updated_at: moment(),
					award: 51,
					name: 'Rfjdks fezifmdsjk dfd',
					description: 'EI fkdsmf iOM FJKD fjdks fjkdlMJ FKDMS JFIdmsj kfl',
					category: 'TJK fdjsklf',
					check_id: 5,
					check_data: 8,
				},
				{
					id: 2,
					created_at: moment(),
					updated_at: moment(),
					award: 51,
					name: 'KFJMdks fjdkl jkldjfkdsl ',
					description: 'jfdmquiemozr kdlf sqfkd sqjfkd',
					category: 'JKFdm suirez',
					check_id: 5,
					check_data: 8,
				},
			]),
			sknex('ads').insert({
				id: 1,
				created_at: moment(),
				updated_at: moment(),
				latitude: 51.12,
				longitude: 54.42,
				region: 'Paris',
				url: 'https://test.com',
				image_url: 'http://image.com',
				alt_text: 'Super text',
			}),
			sknex('tags').insert([
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
				},
				{
					id: 3,
					created_at: moment(),
					updated_at: moment(),
					name: 'Tag3',
					description: 'Description of tag3',
				},
			]),
		]);
	}).then(() => {
		return Promise.all([
			sknex('requests').insert([
				{
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
				},
				{
					id: 2,
					created_at: moment(),
					updated_at: moment(),
					user_id: 1,
					content: 'kdljsmfezmif jksf',
					title: 'jfkdlmjieq fdkslf djkslqfmd',
					views: 4,
					status: 'Open',
					duration: 50,
					frequency: 21,
					frequency_unit: 'years',
					period_start: moment(),
					period_end: moment(),
					location: 'JFkds fkj',
				},
			]),
			sknex('notifications').insert([
				{
					id: 1,
					user_id: 1,
					title: "FJK dsjf dksjmfd",
					message: "JKlfdskfd sjfkdls",
					read: true,
				},
				{
					id: 2,
					user_id: 1,
					title: "FJK dsjf dksjmfd",
					message: "JKlfdskfd sjfkdls",
					read: false,
				},
				{
					id: 3,
					user_id: 1,
					title: "Ikfkmfd",
					message: "Jskfd sjfkdls",
					read: false,
				},
			]),
		]);
	}).then(() => {
		return Promise.all([
			sknex('tags_requests').insert([
				{
					tag_id: 1,
					request_id: 1,
				},
				{
					tag_id: 2,
					request_id: 1,
				},
			]),
		]);
	});
};
