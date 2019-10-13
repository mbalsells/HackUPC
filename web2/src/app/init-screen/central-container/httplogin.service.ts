import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface SuccessMessage {
  error: string;
  success: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class HTTPLoginService {
  constructor(private http: HttpClient) {}

  login(_username: string, _password: string): Observable<SuccessMessage> {
    return this.http.get<SuccessMessage>('http://localhost:8080/login', {
      params: {
        username: _username,
        password: _password,
      }
    });
  }

  register(_name: string, _email: string, _username: string, _password: string): Observable<SuccessMessage> {
    return this.http.get<SuccessMessage>('http://localhost:8080/register', {
      params: {
        name: _name,
        email: _email,
        username: _username,
        password: _password,
      }
    });
  }
}
