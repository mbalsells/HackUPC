import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ScrollAppComponent } from './scroll-app.component';

describe('ScrollAppComponent', () => {
  let component: ScrollAppComponent;
  let fixture: ComponentFixture<ScrollAppComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ScrollAppComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ScrollAppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
