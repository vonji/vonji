import piper from './piper';
import {
	fetchOne,
	fetchAll,
	save,
	destroy,
} from './dbHelpers';

export default (router, resource) => {
	router.get('/:id', piper((req, res) => {
		return req.reply.ok(fetchOne(resource, req.params.id));
	}));

	router.get('/', piper((req, res) => {
		return req.reply.ok(fetchAll(resource));
	}));

	router.put('/:id', piper((req, res) => {
		return req.reply.empty(save(resource, req.body, req.params.id));
	}));

	router.post('/', piper((req, res) => {
		return req.reply.created(save(resource, req.body));
	}));

	router.delete('/:id', piper((req, res) => {
		return req.reply.empty(destroy(resource, req.params.id));
	}));
};
