import { Component, OnInit } from '@angular/core';
import { HTTPLessonService } from '../user-screen/httplesson.service';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-feedback',
  templateUrl: './feedback.component.html',
  styleUrls: ['./feedback.component.scss']
})
export class FeedbackComponent implements OnInit {
  username = "";
  subject = [];
  feedback = [];

  constructor(private less : HTTPLessonService, private actRoute: ActivatedRoute) { }

  ngOnInit() {
    this.username = this.actRoute.snapshot.paramMap.get("username");
    var _this = this;
    this.less.infoUser(this.username).subscribe(data => {
      _this.subject = data.subject;
      _this.feedback = data.feedback;
    });
  }

}
