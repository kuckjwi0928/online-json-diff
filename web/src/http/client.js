import axios from "axios";

const instance = axios.create({
  baseURL: '/',
  timeout: 1000 * 3,
});

function client() {
  return instance;
}

export default client;
