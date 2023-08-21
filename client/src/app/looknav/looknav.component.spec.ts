import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LooknavComponent } from './looknav.component';

describe('LooknavComponent', () => {
  let component: LooknavComponent;
  let fixture: ComponentFixture<LooknavComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LooknavComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LooknavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
