import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-scroll',
  templateUrl: './scroll.component.html',
  styleUrls: ['./scroll.component.scss']
})
export class ScrollComponent implements OnInit {
  ITER = [0, 1, 2, 3, 4];
  cards = ["Algorithms", "Complex analysis", "Linear Optimization", "Parametrized Complexity", "Theory"];
  votes = [0,             5,                    9,                    3,                        5];
  it = 0;
  n = 5;

  constructor() { }

  ngOnInit() {
  }

  expandLess() {
    this.it--;
  }
  expandMore() {
    this.it++;
  }

}
