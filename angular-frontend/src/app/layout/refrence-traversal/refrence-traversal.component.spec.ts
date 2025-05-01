import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RefrenceTraversalComponent } from './refrence-traversal.component';

describe('RefrenceTraversalComponent', () => {
  let component: RefrenceTraversalComponent;
  let fixture: ComponentFixture<RefrenceTraversalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RefrenceTraversalComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RefrenceTraversalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
