import express from "express";
import simpleRouting from "../utils/simpleRouting";

const router = express.Router();

simpleRouting(router, 'ads');

export default router;
