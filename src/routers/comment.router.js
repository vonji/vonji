import Comment from "../models/comment.model";
import express from "express";
import {simpleRouting} from "./utils";

const router = express.Router();

simpleRouting(router, Comment);

export default router;
