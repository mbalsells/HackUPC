import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface ScheduleResponse {
  end_time: string[];
  start_time: string[];
  subject: string[];
}

export interface InfoUserResponse {
  username: string;
  name: string;
  email: string;
  subjects: string[];
  feedbacks: number[];
}

export interface FeedbackResponse {
  success: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class HTTPLessonService {
  constructor(private http: HttpClient) {}

  schedule(_username: string): Observable<ScheduleResponse> {
    return this.http.get<ScheduleResponse>('http://localhost:8080/schedule', {
      params: {
        username: _username,
      }
    });
  }

  infoUser(_username: string): Observable<InfoUserResponse> {
    return this.http.get<InfoUserResponse>('http://localhost:8080/infouser', {
      params: {
        username: _username,
      }
    });
  }

  sendFeedback(_username: string, _punt: number, _assig: string): Observable<FeedbackResponse> {
    var _points = _punt as any as string;
    return this.http.get<FeedbackResponse>('http://localhost:8080/sendfeedback', {
      params: {
        username: _username,
        points: _points,
        subjectName: _assig,
      }
    });
  }
}