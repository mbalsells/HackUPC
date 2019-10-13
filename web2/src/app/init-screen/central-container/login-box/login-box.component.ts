import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { HTTPLoginService, SuccessMessage } from '../httplogin.service';


@Component({
  selector: 'app-login-box',
  templateUrl: './login-box.component.html',
  styleUrls: ['./login-box.component.scss']
})
export class LoginBoxComponent implements OnInit {
  loginInfo: FormGroup;
  errMsg = false;
  
  constructor(
    private logger: HTTPLoginService,
    private router: Router) { }

  ngOnInit() {
    this.loginInfo = new FormGroup({
      username: new FormControl('', [Validators.required]),
      password: new FormControl('', [Validators.required]),
    });
  }

  login() {
    const _username = this.loginInfo.get("username").value;
    const _password = this.loginInfo.get("password").value;
    
    const _this = this;
    this.logger.login(_username, _password)
    .subscribe(
      data => {
        if (data.success) {
          _this.router.navigateByUrl("/user/" + _username);
        }
        else {
          _this.errMsg = true;
        }
      });
  }
}