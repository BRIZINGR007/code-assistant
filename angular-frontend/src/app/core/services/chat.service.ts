import { Injectable } from '@angular/core';
import { HttpClientService } from './http-client.service';
import { environment } from '../../../environments/environment';
import { catchError, firstValueFrom, Observable } from 'rxjs';
import { HttpErrorResponse } from '@angular/common/http';
import { IChat } from '../interfaces/chat.service.interface';

@Injectable({
  providedIn: 'root'
})
export class ChatService {
  private readonly CodeAssistantServiceEndPoint = environment.codeassistant_service_endpoint;


  constructor(private httpClientService: HttpClientService) { }

  UnSetCodeBaseChatSignal(): void {


  }
  RetriveSessionChatHistory(sessionId: string): Promise<IChat[]> {
    const observable = this.httpClientService.get<IChat[]>(
      `${this.CodeAssistantServiceEndPoint}/chat/get-session-chats`,
      { sessionId: sessionId },
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);
  }

  SessionChat(sessionId: string, userQuery: string, codeBaseIds: string[]): Promise<IChat> {
    const observable = this.httpClientService.post<IChat>(
      `${this.CodeAssistantServiceEndPoint}/code-assist/session-chat`,
      { session_id: sessionId, query: userQuery, codebase_ids: codeBaseIds },
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);

  }
  RetrieveSessionCodeBaseIds(sessionId: string): Observable<string[]> {
    return this.httpClientService.get<string[]>(
      `${this.CodeAssistantServiceEndPoint}/users/get-session-codebase-ids`,
      { sessionId: sessionId }
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
  }
  RetriveGeneralSessionChatHistory(): Promise<IChat[]> {
    const observable = this.httpClientService.get<IChat[]>(
      `${this.CodeAssistantServiceEndPoint}/chat/get-general-session-chats`,
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);
  }
  GeneralSessionChat(userQuery: string): Promise<IChat> {
    const observable = this.httpClientService.post<IChat>(
      `${this.CodeAssistantServiceEndPoint}/code-assist/general-session-chat`,
      { query: userQuery },
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);
  }
  DeleteGeneralChatSession(): Observable<any> {
    return this.httpClientService.delete(
      `${this.CodeAssistantServiceEndPoint}/users/delete-general-chat-session`,
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
  }
}
