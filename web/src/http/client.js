import axios from "axios";

const instance = axios.create({
  baseURL: '/',
  timeout: 1000 * 5,
});

function client() {
  return instance;
}

export default client;
