const moment = require('moment');
const Chance = require('chance');

exports.seed = (knex) => {
	const sknex = (resource) => ({
		del: function (options = {}) {
			return knex(resource).del()
				.then(() => {
					if (options.noSeq) {
						return;
					} else {
						return knex.raw(`ALTER SEQUENCE ${resource}_id_seq RESTART WITH 1`);
					}
				})
				.then(() => {
					console.log(`All ${resource} has been removed.`);
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
			sknex('comments').del(),
			sknex('responses').del(),
			sknex('requests_tags').del({ noSeq: true }),
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
					email: 'loup.peluso@vonji.fr',
					password: 'loup.peluso',
					displayed_name: 'Loup',
					real_name: 'Loup Peluso',
					description: 'Coucou les p\'tits lapins.',
					motto: 'J\'aime les pommes',
					birthday: moment('1985-04-03'),
					created_at: moment(),
					updated_at: moment(),
				},
				{
					email: 'david.lancar@vonji.fr',
					password: 'david.lancar',
					displayed_name: 'David',
					real_name: 'David Lancar',
					description: 'Vfjdksl fjdkqsm fjdq fieozm jfejf',
					motto: 'EIrezk fdjksf eziomqrem dskfd kl',
					birthday: moment('1985-04-03'),
					created_at: moment(),
					updated_at: moment(),
				},
			]),
			sknex('achievements').insert([
				{
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
			sknex('ads').insert([
				{
					created_at: moment(),
					updated_at: moment(),
					latitude: 51.12,
					longitude: 54.42,
					region: 'Paris',
					url: 'https://test.com',
					image_url: 'http://image.com',
					alt_text: 'Super text',
				},
				{
					created_at: moment(),
					updated_at: moment(),
					latitude: 47,
					longitude: 7.27,
					region: 'Rome',
					url: 'https://fdsfdsqfdsq.com',
					image_url: 'http://reuziorez.com',
					alt_text: 'JFKDLM jfdkls mfjdks fjieomFJKD',
				}
			]),
			sknex('tags').insert([
				{
					created_at: moment(),
					updated_at: moment(),
					name: 'Tag',
					description: 'Description of tag',
				},
				{
					created_at: moment(),
					updated_at: moment(),
					name: 'Tag2',
					description: 'Description of tag2',
				},
				{
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
					user_id: 1,
					created_at: moment(),
					updated_at: moment(),
					title: "FJK dsjf dksjmfd",
					message: "JKlfdskfd sjfkdls",
					read: true,
				},
				{
					user_id: 1,
					created_at: moment(),
					updated_at: moment(),
					title: "FJK dsjf dksjmfd",
					message: "JKlfdskfd sjfkdls",
					read: false,
				},
				{
					user_id: 1,
					created_at: moment(),
					updated_at: moment(),
					title: "Ikfkmfd",
					message: "Jskfd sjfkdls",
					read: false,
				},
			]),
		]);
	}).then(() => {
		return Promise.all([
			sknex('requests_tags').insert([
				{ tag_id: 1, request_id: 1, },
				{ tag_id: 2, request_id: 1, },
			]),
			sknex('responses').insert([
				{
					created_at: moment(),
					updated_at: moment(),
					user_id: 1,
					request_id: 1,
					content: "JFkdlsf djskqf ejkfm jdksqfj dksqlf jdksq fjdksmlqfd",
					value: 75,
					accepted: true,
					rating: 2,
				},
				{
					created_at: moment(),
					updated_at: moment(),
					user_id: 2,
					request_id: 1,
					content: "ureizo ifd jskqfmj dqimefzj fdm sjkfdsqfd",
					value: 5,
					accepted: false,
					rating: 2,
				},
			]),
			sknex('comments').insert([
				{
					created_at: moment(),
					updated_at: moment(),
					content: 'fjdkslqf djskqf djks',
					commentable_id: 1,
					commentable_type: 'response',
					user_id: 1,
					type: 'fdsjqkfd s',
				}
			]),
		]);
	});
};
