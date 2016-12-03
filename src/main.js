import express from "express";
import {User} from "../models/users.model";

const app = express();

app.get('/users', (req, res) => {
    User.fetchAll().then(models => {
        res.send(models);
    });
});

app.listen(3000, () => {
    console.log('Application started');
});