import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DialogDataExampleComponent } from './dialog-data-example.component';

describe('DialogDataExampleComponent', () => {
  let component: DialogDataExampleComponent;
  let fixture: ComponentFixture<DialogDataExampleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DialogDataExampleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DialogDataExampleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
