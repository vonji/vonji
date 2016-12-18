import Notification from "../models/notification.model";
import express from "express";
import {simpleRouting} from "./utils";

const router = express.Router();

simpleRouting(router, Notification);

export default router;
