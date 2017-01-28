import Achievement from "../models/achievement.model";
import express from "express";
import {simpleRouting} from "../utils/simpleRouting"

const router = express.Router();

simpleRouting(router, Achievement);

export default router;
