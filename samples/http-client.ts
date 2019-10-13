import { Component, OnInit, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-sample-button',
  templateUrl: './sample-button.component.html',
  styleUrls: ['./sample-button.component.scss']
})
@Injectable()
export class SampleButtonComponent implements OnInit {
  message = null;

  constructor(private http: HttpClient) { }

  ngOnInit() {
  }

  buttonClicked() {
    this.message = this.http.get('http://localhost:8080/sampleEndpoint')
  }
}
