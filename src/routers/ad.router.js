import Ad from "../models/ad.model";
import express from "express";
import {fetchOne, fetchAll, remove, save} from "./utils";

const router = express.Router();

router.get('/:id', (req, res) => {
    fetchOne(res, new Ad({id: req.params.id}));
});

router.get('/', (req, res) => {
    fetchAll(res, Ad);
});

router.post('/', (req, res) => {
    save(res, req.body, new Ad());
});

router.put('/:id', (req, res) => {
    fetchOne(res, new Ad({id: req.params.id}), {
        done(model) {
            save(res, req.body, model);
        },
    });
});

router.delete('/:id', (req, res) => {
    fetchOne(res, new Ad({id: req.params.id}), {
        done(model) {
            remove(res, model);
        },
    });
});

export default router;