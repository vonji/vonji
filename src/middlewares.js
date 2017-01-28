export const replyMiddleware = (req, res, next) => {
	req.reply = {
		async ok(promise) {
			const result = await promise;
			res.status(200).send(result);
		},
		async created(promise) {
			const result = await promise;
			res.status(201).send(result);
		},
		async empty(promise) {
			await promise;
			res.sendStatus(204);
		},
	};
	next();
};

export const errorsHandlerMiddleware = (err, req, res, next) => {
  switch (err.name) {
    case 'NotFoundError':
      res.sendStatus(404);
      break;
    default:
      res.status(500).send(err);
      break;
  }
};

export const logErrorsMiddleware = (err, req, res, next) => {
  console.error(err);
  next(err);
};
