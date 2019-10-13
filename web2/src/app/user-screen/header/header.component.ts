import { Component, OnInit, Input } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  username = "";

  constructor(private router: Router, private actRoute: ActivatedRoute) {}

  ngOnInit() {
    this.username = this.actRoute.snapshot.paramMap.get("username");
  }

  feedbackClicked() {
    this.router.navigateByUrl('/feedback/' + this.username);
  }

  communityClicked() {
    this.router.navigateByUrl('/community/' + this.username);
  }

}
