import { Component, OnInit, Input } from '@angular/core';
import { HTTPLessonService, FeedbackResponse, InfoUserResponse, ScheduleResponse } from '../../httplesson.service'

@Component({
  selector: 'app-scroll',
  templateUrl: './scroll.component.html',
  styleUrls: ['./scroll.component.scss']
})
export class ScrollComponent implements OnInit {
  @Input() username: string;

  prod = [0, 0, 0, 0, 0];
  ITER = [0, 1, 2, 3, 4];
  cards = ["Algorithms", "Complex analysis", "Linear Optimization", "Parametrized Complexity", "Theory"];
  votes = [0,             5,                    9,                    3,                        5];
  it = 0;
  n = 5;

  constructor(private less: HTTPLessonService) {}

  ngOnInit() {
    var _this = this;
    this.less.infoUser(this.username).subscribe(data => {
      _this.n = data.subject.length;
      _this.cards = data.subject;
      _this.votes = data.feedback;
      _this.prod = Array(_this.n).fill(0);
      _this.ITER = Array.from(Array(_this.n).keys());
    });
  }

  expandLess() {
    this.it--;
  }
  expandMore() {
    this.it++;
  }
  vote(i: number) {
    this.less.sendFeedback(this.username, this.prod[i], this.cards[i]).subscribe();
    this.votes[i] = this.prod[i];
    this.prod[i] = 0;
  }
  unvote(i: number) {
    this.less.sendFeedback(this.username, 0, this.cards[i]).subscribe();
    this.votes[i] = 0;
  }
  prop(i: number, e: any) {
    this.prod[i] = e.value;
  }
}
