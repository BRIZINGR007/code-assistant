import { Injectable } from '@angular/core';
import { ILoginPayload } from '../interfaces/auth.service.interface';
import { catchError, firstValueFrom, map, Observable, of } from 'rxjs';
import { HttpClientService } from './http-client.service';
import { environment } from '../../../environments/environment';
import { HttpErrorResponse } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly CodeAssistantServiceEndPoint = environment.codeassistant_service_endpoint;


  constructor(private httpClientService: HttpClientService) { }

  isAuthenticated() {
    return this.httpClientService
      .get<any>(`${this.CodeAssistantServiceEndPoint}/users/validate-session`)
      .pipe(
        map(() => true),
        catchError(() => {
          return of(false);
        })
      );
  }
  SignUp(payload: ILoginPayload) {
    return this.httpClientService.post(
      `${this.CodeAssistantServiceEndPoint}/users/signup`,
      payload
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
  }

  login(payload: ILoginPayload): Observable<any> {
    return this.httpClientService.post(
      `${this.CodeAssistantServiceEndPoint}/users/login`,
      payload
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
  }
  LogOut(): Promise<any> {
    const observable = this.httpClientService.post(
      `${this.CodeAssistantServiceEndPoint}/users/logout`,
    ).pipe(
      catchError((error: HttpErrorResponse) => { throw error; })
    );
    return firstValueFrom(observable);
  }

}
