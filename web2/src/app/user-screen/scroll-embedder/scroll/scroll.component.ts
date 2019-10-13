import { Component, OnInit, Input } from '@angular/core';
import { HTTPLessonService, FeedbackResponse, InfoUserResponse, ScheduleResponse } from '../../httplesson.service'
import {Router} from '@angular/router';

@Component({
  selector: 'app-scroll',
  templateUrl: './scroll.component.html',
  styleUrls: ['./scroll.component.scss']
})
export class ScrollComponent implements OnInit {
  @Input() username: string;

  prod = [];
  ITER = [];
  cards = [];
  votes = [];
  it = 0;
  n = 0;

  constructor(
    private less: HTTPLessonService,
    private router: Router
    ) {}

  ngOnInit() {
    var _this = this;
    this.less.infoUser(this.username).subscribe(data => {
      _this.n = data.subject.length;
      _this.cards = data.subject;
      _this.votes = data.feedback;
      _this.prod = Array(_this.n).fill(0);
      _this.ITER = Array.from(Array(_this.n).keys());
    });
    console.log(this.ITER);
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
  community() {
    this.router.navigateByUrl("community/"+this.username);
  }
}
