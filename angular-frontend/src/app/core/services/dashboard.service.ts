import { Injectable } from '@angular/core';
import { HttpClientService } from './http-client.service';
import { catchError, delay, firstValueFrom, Observable, of } from 'rxjs';
import { HttpErrorResponse } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { CodebasePayload, ExtractCodePayload, IUserChatSessionsData } from '../interfaces/dashboard.interfaces';

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  private readonly CodeAssistantServiceEndPoint = environment.codeassistant_service_endpoint;

  constructor(private httpClientService: HttpClientService) { }

  retrieveCodeBaseData(): Promise<CodebasePayload[]> {
    const observable = this.httpClientService.get<CodebasePayload[]>(
      `${this.CodeAssistantServiceEndPoint}/users/get-code-bases`
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);
  }


  extractCode(payload: ExtractCodePayload): Observable<any> {
    return this.httpClientService.post(
      `${this.CodeAssistantServiceEndPoint}/code-assist/extract-code`,
      payload
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
  }
  DeleteCodeBasdeContext(codeBaseId: string): Observable<any> {
    return this.httpClientService.delete(
      `${this.CodeAssistantServiceEndPoint}/users/delete-codebase`,
      { codeBaseId: codeBaseId },
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
  }
  createChatSession(codeBaseIds: string[]): Observable<{ sessionId: string }> {
    const observable = this.httpClientService.post<{ sessionId: string }>(
      `${this.CodeAssistantServiceEndPoint}/users/create-session`,
      { codebase_ids: codeBaseIds }
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return observable;
  }
  RetrieveUserChatSessions(): Promise<IUserChatSessionsData[]> {
    const observable = this.httpClientService.get<IUserChatSessionsData[]>(
      `${this.CodeAssistantServiceEndPoint}/users/get-chat-sessions`
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);
  }
  DeleteChatSession(sessionId: string): Observable<any> {
    return this.httpClientService.delete(
      `${this.CodeAssistantServiceEndPoint}/users/delete-chat-session`,
      { sessionId: sessionId },
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
  }
  SynccodeBases(): Promise<any> {
    const observable = this.httpClientService.post(
      `${this.CodeAssistantServiceEndPoint}/users/sync-codebases`,
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);
  }

}
