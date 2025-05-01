import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class HttpClientService {
  constructor(private http: HttpClient) { }

  post<T>(
    url: string,
    body?: any,
    params?: any,
    withCredentials: boolean = true
  ): Observable<T> {
    const httpParams = new HttpParams({
      fromObject: params || {},
    });
    return this.http.post<T>(url, body, {
      withCredentials,
      params: httpParams,
    });
  }
  get<T>(
    url: string,
    params?: any,
    withCredentials: boolean = true
  ): Observable<T> {
    const httpParams = new HttpParams({
      fromObject: params || {},
    });
    return this.http.get<T>(url, {
      withCredentials,
      params: httpParams,
    });
  }

  delete<T>(
    url: string,
    params?: any,
    withCredentials: boolean = true
  ): Observable<T> {
    const httpParams = new HttpParams({
      fromObject: params || {},
    });
    return this.http.delete<T>(url, {
      withCredentials,
      params: httpParams,
    });
  }

}
