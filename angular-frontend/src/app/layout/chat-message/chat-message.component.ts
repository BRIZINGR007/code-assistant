import { Component, EventEmitter, Input, Output } from '@angular/core';
import { IChat } from '../../core/interfaces/chat.service.interface';
import { MarkdownComponent } from 'ngx-markdown';
import { ChatSkeletonComponent } from '../../shared/components/chat-skeleton/chat-skeleton.component';

@Component({
  selector: 'app-chat-message',
  imports: [ChatSkeletonComponent, MarkdownComponent],
  templateUrl: './chat-message.component.html',
  styleUrl: './chat-message.component.scss'
})
export class ChatMessageComponent {
  @Input() chat!: IChat;
  @Output() ReferncesTraversalEmitter = new EventEmitter<IChat>();

  showRefernces() {
    this.ReferncesTraversalEmitter.emit(this.chat);
  }

}
