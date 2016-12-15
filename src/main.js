import express from "express";
import bodyParser from "body-parser";
import userRouter from "./routers/user.router";
import adRouter from "./routers/ad.router";
import achievementRouter from "./routers/achievement.router";
import requestRouter from "./routers/request.router";
import tagRouter from "./routers/tag.router";

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: true}));

app.use('/users', userRouter);
app.use('/achievements', achievementRouter);
app.use('/ads', adRouter);
app.use('/tags', tagRouter);
app.use('/requests', requestRouter);

app.listen(3000, () => {
    console.log('Application started');
});
