import Ad from "../models/ad.model";
import express from "express";
import {fetchOne} from "./utils";

const router = express.Router();

router.get('/:id', (req, res) => {
    fetchOne(res, new Ad({id: req.params.id}));
});

router.get('/', (req, res) => {
    Ad.fetchAll()
        .then(resources => res.send(resources.toJSON()))
        .catch(err => {
            console.error(err);
            res.sendStatus(500);
        });
});

router.post('/', (req, res) => {
    new Ad().save(req.body)
        .then(resource => res.send(resource.toJSON()))
        .catch(err => {
            console.error(err);
            res.sendStatus(500);
        });
});

router.put('/:id', (req, res) => {
    new Ad({id: req.params.id})
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
    new Ad({id: req.params.id})
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