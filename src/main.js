import express from "express";
import bodyParser from "body-parser";
import userRouter from "./routers/user.router";
import adRouter from "./routers/ad.router";
import achievementRouter from "./routers/achievement.router";
import requestRouter from "./routers/request.router";
import tagRouter from "./routers/tag.router";
import responseRouter from "./routers/response.router";
import notificationRouter from "./routers/notification.router";
import transactionRouter from "./routers/transaction.router";
import commentRouter from "./routers/comment.router";

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: true}));

app.use('/users', userRouter);
app.use('/achievements', achievementRouter);
app.use('/ads', adRouter);
app.use('/tags', tagRouter);
app.use('/requests', requestRouter);
app.use('/notifications', notificationRouter);
app.use('/responses', responseRouter);
app.use('/transactions', transactionRouter);
app.use('/comments', commentRouter);

app.listen(3000, () => {
	console.log('Application started');
});
