import User from "../models/user.model";
import express from "express";

const router = express.Router();

router.get('/:id', (req, res) => {
    new User({id: req.params.id})
        .fetch()
        .then(resource => {
            if (resource) {
                res.send(resource.toJSON());
            } else {
                res.sendStatus(404);
            }
        })
        .catch(err => {
            console.error(err);
            res.sendStatus(500);
        });
});

router.get('/', (req, res) => {
    User.fetchAll()
        .then(resources => res.send(resources.toJSON()));
});

router.post('/', (req, res) => {
    new User().save(req.body)
        .then(resource => res.send(resource.toJSON()))
        .catch(err => {
            console.error(err);
            res.sendStatus(500);
        });
});

router.put('/:id', (req, res) => {
    new User({id: req.params.id})
        .fetch()
        .then(resource => {
            if (resource) {
                resource.save(req.body).then(() => res.send(resource.toJSON()));
            } else {
                res.sendStatus(404);
            }
        })
        .catch(err => {
            console.error(err);
            res.sendStatus(500);
        });
});

router.delete('/:id', (req, res) => {
    new User({id: req.params.id})
        .fetch()
        .then(resource => {
            if (resource) {
                resource.destroy().then(() => res.sendStatus(204));
            } else {
                res.sendStatus(404);
            }
        })
        .catch(err => {
            console.error(err);
            res.sendStatus(500);
        });
});

export default router;