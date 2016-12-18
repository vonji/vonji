import User from "../models/user.model";
import express from "express";
import {simpleRouting} from "./utils";

const router = express.Router();

simpleRouting(router, User);

export default router;
