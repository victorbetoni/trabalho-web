import { APIRequest, GET, POST, type Handler } from "./api";

export function login(username: string, password: string, handler: Handler<string>) {
  POST(new APIRequest("login", null, ({
    username: username,
    password: password,
  }), null), handler)
}

export function checkSession(handler: Handler<any>) {
  GET(new APIRequest("checkSession", null, null, null), handler)
}

export function logout(handler: Handler<any>) {
  GET(new APIRequest("logout", null, null ,null), handler)
} 