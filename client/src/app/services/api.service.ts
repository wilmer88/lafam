import { Injectable } from "@angular/core";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private port?: string;

  setPort(port: string) {
    this.port = port;
  }

  getBaseUrl() {
    if (this.port) {
      return `https://your-app-name.herokuapp.com:${this.port}/api`;
    } else {
      return `https://your-app-name.herokuapp.com/api`;
    }
  }

  // ...
}