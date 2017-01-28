import Ad from "../models/ad.model";
import express from "express";
import {simpleRouting} from "../utils/simpleRouting"

const router = express.Router();

simpleRouting(router, Ad);

export default router;
