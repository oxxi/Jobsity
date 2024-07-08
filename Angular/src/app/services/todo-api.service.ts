import { HttpClient, HttpHeaders } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { catchError, Observable, throwError } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class TodoApiService {
  private baseUrl: string = environment.apiUrl;

  private http = inject(HttpClient);
  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
    }),
  };

  constructor() {}

  getAll(): Observable<any[]> {
    const url = `${this.baseUrl}tasks/`;
    return this.http.get<any[]>(url).pipe(
      catchError(() => {
        console.error('catch error in service');
        return throwError(() => {
          return new Error('Error retrieving data');
        });
      })
    );
  }

  getById(): Observable<any> {
    return this.http.get<any>('');
  }

  save(data: any): Observable<any> {
    const url = `${this.baseUrl}task`;
    return this.http
      .post<any>(url, JSON.stringify(data), this.httpOptions)
      .pipe(
        catchError(() => {
          console.error('catch error in service');
          return throwError(() => {
            return new Error('Erro while create new Task');
          });
        })
      );
  }

  update(id: number, data: any): Observable<any> {
    const url = `${this.baseUrl}task/${id}`;
    return this.http.put(url, JSON.stringify(data), this.httpOptions);
  }

  delete(id: number): Observable<any> {
    const url = `${this.baseUrl}task/${id}`;
    return this.http.delete(url).pipe(
      catchError(() => {
        console.error('catch error in service');
        return throwError(() => {
          return new Error('error while delete Task');
        });
      })
    );
  }
}
