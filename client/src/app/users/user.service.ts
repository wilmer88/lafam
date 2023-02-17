
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, tap, throwError } from 'rxjs';
import { Imembers } from './Iuser';
@Injectable({
  providedIn: 'root',
})

export class UserService {
  private userUrl = 'http://localhost:8080/lafamily';

  constructor(private http: HttpClient) { };

  getUsers(): Observable<Imembers[]> {
    return this.http.get<Imembers[]>(this.userUrl).pipe(
      tap(data => console.log('ALL:', JSON.stringify(data))),
      catchError(this.handleError)
    );

}

  
  private handleError(err: HttpErrorResponse) {
    let errorMassage = '';
    if (err.error instanceof ErrorEvent) {
      errorMassage = `An error occurred: ${err.error.message}`;
    } else{
      errorMassage = `Server returned code: ${err.status}, error message is: ${err.message}`
      
    }
    console.error(errorMassage);
    return throwError(() => errorMassage)
  }
}

