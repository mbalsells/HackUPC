import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ScrollEmbedderComponent } from './scroll-embedder.component';

describe('ScrollEmbedderComponent', () => {
  let component: ScrollEmbedderComponent;
  let fixture: ComponentFixture<ScrollEmbedderComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ScrollEmbedderComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ScrollEmbedderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
