import Tag from "../models/tag.model";
import express from "express";
import {simpleRouting} from "./utils";

const router = express.Router();

simpleRouting(router, Tag);

export default router;
