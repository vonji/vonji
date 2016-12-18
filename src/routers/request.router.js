import Request from "../models/request.model";
import express from "express";
import _ from "lodash";
import {simpleRouting} from "./utils";

const router = express.Router();

simpleRouting(router, Request);

export default router;
