import { TestBed } from '@angular/core/testing';

import { HTTPLessonService } from './httplesson.service';

describe('HTTPLessonService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: HTTPLessonService = TestBed.get(HTTPLessonService);
    expect(service).toBeTruthy();
  });
});
