import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router'

@Component({
  selector: 'app-user-screen',
  templateUrl: './user-screen.component.html',
  styleUrls: ['./user-screen.component.scss']
})
export class UserScreenComponent implements OnInit {
  username = "";

  constructor(private actRoute: ActivatedRoute) {}

  ngOnInit() {
    this.username = this.actRoute.snapshot.paramMap.get("username");
  }
}
