import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { HTTPLoginService, SuccessMessage } from "../httplogin.service";

@Component({
  selector: 'app-register-box',
  templateUrl: './register-box.component.html',
  styleUrls: ['./register-box.component.scss'],
})
export class RegisterBoxComponent implements OnInit {
  nameGroup: FormGroup;
  emailGroup: FormGroup;
  usernameGroup: FormGroup;
  passwordGroup: FormGroup;

  constructor(
    private _formBuilder: FormBuilder,
    private router: Router,
    private logger: HTTPLoginService) {}

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
    
    var _this = this;
    this.logger.register(_name, _email, _username, _password)
    .subscribe({
      next(data: SuccessMessage) {
        if (data.success) {
          _this.router.navigateByUrl("/user/" + _username);
        }
      }
    });
  }
}