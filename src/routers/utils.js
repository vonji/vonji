export const errorHandler = res => err => {
	console.error(err);
	res.sendStatus(500);
};

const fetch = (res, promise, {done} = {}) => {
	return promise.then(data => {
		if (data) {
			if (done) {
				done(data);
			} else {
				send(res, data);
			}
		} else {
			res.sendStatus(404);
		}
	}).catch(errorHandler(res));
};

export const send = (res, model) => {
	res.send(model.toJSON());
};

export const remove = (res, model) => {
	model.destroy()
		.then(() => res.sendStatus(204))
		.catch(errorHandler(res))
};

export const save = (res, attributes, model) => {
	model.save(attributes)
		.then((newModel) => {
			res.send(newModel.toJSON());
		})
		.catch(errorHandler(res));
};

export const fetchOne = (res, model, options) => {
	return fetch(res, model.fetch(), options);
};

export const fetchAll = (res, model, options) => {
	return fetch(res, model.fetchAll(), options);
};
