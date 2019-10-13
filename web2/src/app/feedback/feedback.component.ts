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
  average_me = [];
  average_all = [];

  constructor(private less : HTTPLessonService, private actRoute: ActivatedRoute) { }

  ngOnInit() {
    this.username = this.actRoute.snapshot.paramMap.get("username");
    var _this = this;
    this.less.infoUser(this.username).subscribe(data => {
      _this.subject = data.subject;
      _this.feedback = data.feedback;
      for (var i=0; i<this.subject.length; i++) {
        this.less.average(this.username, this.subject[i]).subscribe(data => {
          _this.average_me.push(data.average_me);
          _this.average_all.push(data.average_all);
        });
      }
    });
  }
}
