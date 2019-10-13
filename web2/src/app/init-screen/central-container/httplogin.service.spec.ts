import { TestBed } from '@angular/core/testing';

import { HTTPLoginService } from './httplogin.service';

describe('HTTPLoginService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: HTTPLoginService = TestBed.get(HTTPLoginService);
    expect(service).toBeTruthy();
  });
});
