import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DialogOutComponent } from './dialog-out.component';

describe('DialogOutComponent', () => {
  let component: DialogOutComponent;
  let fixture: ComponentFixture<DialogOutComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DialogOutComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DialogOutComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
