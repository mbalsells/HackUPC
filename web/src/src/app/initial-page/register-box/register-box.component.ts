import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-register-box',
  templateUrl: './register-box.component.html',
  styleUrls: ['./register-box.component.scss']
})
export class RegisterBoxComponent implements OnInit {
  nameGroup: FormGroup;
  emailGroup: FormGroup;
  usernameGroup: FormGroup;
  passwordGroup: FormGroup;

  constructor(private _formBuilder: FormBuilder,
    private http: HttpClient) {}

  ngOnInit() {
    this.nameGroup = this._formBuilder.group({
      name: ['', Validators.required]
    });
    this.emailGroup = this._formBuilder.group({
      email: ['', Validators.required]
    });
    this.usernameGroup = this._formBuilder.group({
      username: ['', Validators.required]
    });
    this.passwordGroup = this._formBuilder.group({
      password: ['', Validators.required]
    });
  }

  register() {
    const _name = this.nameGroup.get("name").value;
    const _email = this.emailGroup.get("email").value;
    const _username = this.usernameGroup.get("username").value;
    const _password = this.passwordGroup.get("password").value;
    
    this.http.get('http://localhost:8080/register', {
      params: {
        name: _name,
        email: _email,
        username: _username,
        password: _password,
      }
    }).subscribe();
  }
}
