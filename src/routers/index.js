import express from 'express';

import userRouter from "./user.router";
import adRouter from "./ad.router";
import achievementRouter from "./achievement.router";
import requestRouter from "./request.router";
import tagRouter from "./tag.router";
import responseRouter from "./response.router";
import notificationRouter from "./notification.router";
import transactionRouter from "./transaction.router";
import commentRouter from "./comment.router";

const router = express.Router();

router.use('/users', userRouter);
router.use('/achievements', achievementRouter);
router.use('/ads', adRouter);
router.use('/tags', tagRouter);
router.use('/requests', requestRouter);
router.use('/notifications', notificationRouter);
router.use('/responses', responseRouter);
router.use('/transactions', transactionRouter);
router.use('/comments', commentRouter);

export default router;
