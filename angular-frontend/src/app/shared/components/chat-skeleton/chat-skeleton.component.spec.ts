import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChatSkeletonComponent } from './chat-skeleton.component';

describe('ChatSkeletonComponent', () => {
  let component: ChatSkeletonComponent;
  let fixture: ComponentFixture<ChatSkeletonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ChatSkeletonComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ChatSkeletonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
