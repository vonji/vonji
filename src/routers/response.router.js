import Response from "../models/response.model";
import express from "express";
import {simpleRouting} from "./utils";

const router = express.Router();

simpleRouting(router, Response);

export default router;
