import express from "express";
import simpleRouting from "../utils/simpleRouting";

const router = express.Router();

simpleRouting(router, 'tags');

export default router;
