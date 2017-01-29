import express from "express";
import bodyParser from "body-parser";
import passport from 'passport';
import apiRoutes from './routers';
import {
	BasicStrategy,
} from 'passport-http';
import {
	replyMiddleware,
	logErrorsMiddleware,
	errorsHandlerMiddleware,
} from "./utils/middlewares.js";

const app = express();

passport.use(new BasicStrategy((username, password, done) => done(null, { username })));

app.use(replyMiddleware);

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: true}));

app.use('/api', passport.authenticate('basic', { session: false }), apiRoutes);

app.use(logErrorsMiddleware);
app.use(errorsHandlerMiddleware);

app.listen(3000, () => {
	console.log('Application started');
});
