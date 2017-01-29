import express from "express";
import bodyParser from "body-parser";
import {
	replyMiddleware,
	logErrorsMiddleware,
	errorsHandlerMiddleware,
} from "./utils/middlewares";
import usersCtrl from './ctrl/users';

const app = express();

app.use(replyMiddleware);

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: true}));

app.use('/api/users', usersCtrl);

app.use(logErrorsMiddleware);
app.use(errorsHandlerMiddleware);

app.listen(3000, () => {
	console.log('Application started');
});
