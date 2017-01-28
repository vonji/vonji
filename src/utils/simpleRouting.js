import {piper} from './piper';

const fetch = async (Model, id) => {
	try {
		return await new Model({ id }).fetch({ require: true });
	} catch(err) {
		throw { name: 'NotFoundError' };
	}
}

const destroy = (Model, id) => {
	const futureModel = fetch(Model, id);
	return futureModel.then(model => model.destroy());
};

const save = (Model, attributes, id) => {
	const futureModel = id ? fetch(Model, id) : Promise.resolve(new Model());
	return futureModel.then(model => model.save(attributes));
};

const fetchAll = Model => Model.fetchAll();

const fetchOne = (Model, id) => {
	return fetch(Model, id).then(model => model.toJSON());
};

export const simpleRouting = (router, Model) => {
	router.get('/:id', piper((req, res) => {
		return req.reply.ok(fetchOne(Model, req.params.id));
	}));

	router.get('/', piper((req, res) => {
		return req.reply.ok(fetchAll(Model));
	}));

	router.post('/', piper((req, res) => {
		return req.reply.created(save(Model, req.body));
	}));

	router.put('/:id', piper((req, res) => {
		return req.reply.empty(save(Model, req.body, req.params.id));
	}));

	router.delete('/:id', piper((req, res) => {
		return req.reply.empty(destroy(Model, req.params.id));
	}));
};
