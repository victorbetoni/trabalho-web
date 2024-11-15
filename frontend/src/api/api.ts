
export interface APIResponse<T> {
  status: number;
  message: string | null;
  body: T | null;
}

export class APIRequest<T> {
  public route: string;
  public endpoint: string | null = null; 
  public body: T | null;
  public query: T | null;

  constructor(route: string, endpoint: string | null, body: T | null, query: T | null) {
    console.log(import.meta.env.VITE_API_HOST)
    this.route = route;
    this.endpoint = endpoint ?? import.meta.env.VITE_API_HOST
    console.log("SDASDASD " + endpoint)
    this.body = body;
    this.query = query;  
  }
}

const headers = {
  "Content-Type": "application/json"
}

export type Handler<T> = (resp: APIResponse<T>) => void;

export function POST<A,T>(req: APIRequest<A>, handler: Handler<T>) {
  fetch(`${req.endpoint}/${req.route}`,{
    body: req.body == null ? "{}" : JSON.stringify(req.body),
    method: "POST",
    headers: headers
  }).then(x => x.json()).then(j => handler(j as APIResponse<T>))
}

export function GET<A,T>(req: APIRequest<A>, handler: Handler<T>) {
  console.log(`${req.endpoint}/${req.route}${req.query != null ? "?" + new URLSearchParams(req.query).toString() : ""}`)
  fetch(`${req.endpoint}/${req.route}${req.query != null ? "?" + new URLSearchParams(req.query).toString() : ""}`, {
    method: "GET",
    headers: headers
  }).then(x => x.json()).then(j => handler(j as APIResponse<T>))
}