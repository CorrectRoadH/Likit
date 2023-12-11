import exp from "constants";
import { Configuration, DefaultApi } from "./openapi"
import axios from "axios";

const config = new Configuration({});

const instance = axios.create({
    baseURL: "",
    timeout: 60000,
    headers: {
      "Content-Type": "application/json",
    },
    withCredentials: false,
  });

const admin_api = new DefaultApi(config, "/admin/v1", instance)

export default admin_api;