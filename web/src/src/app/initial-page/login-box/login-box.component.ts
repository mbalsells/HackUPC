import { Component, OnInit, Injectable } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login-box',
  templateUrl: './login-box.component.html',
  styleUrls: ['./login-box.component.scss']
})
export class LoginBoxComponent implements OnInit {
  loginInfo: FormGroup;

  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.loginInfo = new FormGroup({
      username: new FormControl('', [Validators.required]),
      password: new FormControl('', [Validators.required]),
    });
  }

  login() {
    const _username = this.loginInfo.get("username").value;
    const _password = this.loginInfo.get("password").value;

    this.http.get('http://localhost:8080/login', {
      params: {
        username: _username,
        password: _password,
      }
    }).subscribe();
  }
}
