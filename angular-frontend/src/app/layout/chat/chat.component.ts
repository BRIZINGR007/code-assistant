import { Component, effect, ElementRef, ViewChild } from '@angular/core';
import { HeaderComponent } from '../header/header.component';
import { ChatService } from '../../core/services/chat.service';
import { ChatSkeletonComponent } from '../../shared/components/chat-skeleton/chat-skeleton.component';
import { ChatMessageComponent } from '../chat-message/chat-message.component';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RefrenceTraversalComponent } from '../refrence-traversal/refrence-traversal.component';
import { IChat, ReferencesWithSimilarity } from '../../core/interfaces/chat.service.interface';
import { ToastrService } from 'ngx-toastr';
import { ChatEnums } from '../../core/enums/chat.enums';

@Component({
  selector: 'app-chat',
  imports: [HeaderComponent, ChatSkeletonComponent, ChatMessageComponent, FormsModule, CommonModule, RefrenceTraversalComponent],
  templateUrl: './chat.component.html',
  styleUrl: './chat.component.scss'
})
export class ChatComponent {
  protected readonly ChatEnums = ChatEnums;
  protected isLoading: boolean = true;
  protected sessionId!: string;
  protected sessionChats: IChat[] = [];
  private sessionCodebaseIds: string[] = [];
  protected userQuestion: string = '';
  protected referencesForTraversal: ReferencesWithSimilarity[] = [];
  protected showReferencesTraversal: boolean = false;
  private initialTextareaHeight!: number;
  private initialBodyHeight!: number;
  @ViewChild('chatBody') chatBody!: ElementRef;
  @ViewChild('textareaInput') textareaInput!: ElementRef;
  ngAfterViewInit() {
    this.initialTextareaHeight = this.textareaInput.nativeElement.offsetHeight;
    this.initialBodyHeight = this.chatBody.nativeElement.offsetHeight;
  }
  constructor(
    private chatService: ChatService,
    private route: ActivatedRoute,
    private router: Router,
    private toastr: ToastrService
  ) { }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.sessionId = params['sessionId'];
      this.ComponentDataSetup();
    });
  }


  private async ComponentDataSetup() {
    this.isLoading = true;
    if (this.sessionId !== ChatEnums.GENERAL_CHAT_SESSION) {
      try {
        this.sessionChats = await this.chatService.RetriveSessionChatHistory(this.sessionId);
        this.chatService.RetrieveSessionCodeBaseIds(this.sessionId).subscribe({
          next: (codebaseIds: string[]) => {
            this.sessionCodebaseIds = codebaseIds;
            console.log('Codebase IDs:', this.sessionCodebaseIds);
            this.isLoading = false;
            if (!this.sessionChats || this.sessionChats.length === 0) {
              this.populateDummyChats();
            }
          },
          error: (error) => {
            console.error('Error fetching codebase IDs:', error);
            this.isLoading = false;
            this.router.navigate(["/"]);
          }
        });
      } catch (error) {
        console.error('Error fetching session chats:', error);
        this.isLoading = false;
      }
    } else {
      this.sessionChats = await this.chatService.RetriveGeneralSessionChatHistory();
      this.isLoading = false;
      if (!this.sessionChats || this.sessionChats.length === 0) {
        this.populateDummyChats();
      }
    }
  }
  protected async HandleSend() {
    if (this.userQuestion.trim()) {
      try {
        this.AppendUserQuestionToChatHistory(this.userQuestion);
        const userQuery = this.userQuestion;
        this.userQuestion = "";
        this.ResetBodyHeight();
        let chat;
        if (this.sessionId !== ChatEnums.GENERAL_CHAT_SESSION) {
          chat = await this.chatService.SessionChat(this.sessionId, userQuery, this.sessionCodebaseIds);
        } else {
          chat = await this.chatService.GeneralSessionChat(userQuery);
        }

        this.UpdateChat(chat);

      } catch (error) {
        this.toastr.error("Please start a new session as the codebase has been deleted.");
        this.sessionChats.pop();
      }
    }
  }
  private ResetBodyHeight() {
    setTimeout(() => {
      this.chatBody.nativeElement.style.height = this.initialBodyHeight + 'px';
      this.textareaInput.nativeElement.style.height = this.initialTextareaHeight + 'px';
    }, 0);
  }


  protected SetReferncesForTraversal(chat: IChat): void {
    if (!this.showReferencesTraversal) {
      this.referencesForTraversal = chat.references;
      this.showReferencesTraversal = true;
    } else {
      this.referencesForTraversal = chat.references;
    }
  }
  protected CloseRefernceTraversal() {
    this.showReferencesTraversal = false;
    this.referencesForTraversal = [];
  }

  private AppendUserQuestionToChatHistory(question: string) {
    const newChat: IChat = {
      chat_id: "123",
      user_id: "123",
      session_id: this.sessionId,
      ai_answer: "",
      user_question: question,
      references: []
    };
    this.sessionChats.push(newChat);
    console.log(this.sessionChats);
  }
  protected UpdateChat(chat: IChat) {
    if (this.sessionChats.length > 0) {
      this.sessionChats.pop();
    }
    this.sessionChats.push(chat);
  }

  protected OnEnter(event: Event): void {
    const keyboardEvent = event as KeyboardEvent;
    if (!keyboardEvent.shiftKey) {
      event.preventDefault();
      this.HandleSend();
    }
  }

  protected AutoResize(event: any) {
    const textarea = event.target;
    const maxHeight = 350;

    textarea.style.height = 'auto';
    const newHeight = Math.min(textarea.scrollHeight, maxHeight);
    textarea.style.height = newHeight + 'px';
    textarea.style.overflowY = textarea.scrollHeight > maxHeight ? 'auto' : 'hidden';
    const heightDifference = newHeight - this.initialTextareaHeight;
    if (heightDifference > 0) {
      this.chatBody.nativeElement.style.height =
        (this.initialBodyHeight - heightDifference) + 'px';
    } else {
      this.chatBody.nativeElement.style.height = this.initialBodyHeight + 'px';
    }
  }


  protected populateDummyChats() {

    this.sessionChats = [{
      chat_id: "123",
      user_id: "123",
      session_id: this.sessionId,
      ai_answer: "Thou canst click Reference Traversal to better understand how thy query traversed through the codebase.",
      user_question: "Welcome to Briznigr's Microservice Code Assistant. Feel free to chat and ask any coding questions thou mayest have.",
      references: []
    }]
  }

  protected DeleteGeneralSessionChats() {
    this.chatService.DeleteGeneralChatSession().subscribe({
      next: async () => {
        this.toastr.success("Chat has been succesfully deleted .")
        this.sessionChats = await this.chatService.RetriveGeneralSessionChatHistory();
        this.populateDummyChats();
      },
      error: (err) => {
        this.toastr.error("Error  in deleting  chats .", err);
      }
    })
  }


}
