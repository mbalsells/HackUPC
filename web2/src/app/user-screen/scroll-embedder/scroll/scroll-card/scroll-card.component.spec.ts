import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ScrollCardComponent } from './scroll-card.component';

describe('ScrollCardComponent', () => {
  let component: ScrollCardComponent;
  let fixture: ComponentFixture<ScrollCardComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ScrollCardComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ScrollCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
