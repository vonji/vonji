export const fetchOne = (res, model, cb, error) => {
    model
        .fetch()
        .then(resource => {
            if (resource) {
                if (cb) {
                    cb(resource);
                } else {
                    res.send(resource.toJSON());
                }
            } else {
                res.sendStatus(404);
            }
        })
        .catch(error || (err => {
                console.log(err);
                res.sendStatus(500);
            }));
};