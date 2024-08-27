import axios from "axios";

export const clientIncrement = axios.create({
  baseURL: "http://localhost:80/api/v1/s1/",
});

export const clientRead = axios.create({
  baseURL: "http://localhost:80/api/v1/s2/",
});