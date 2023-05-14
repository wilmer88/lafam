import { HttpClient, HttpErrorResponse, } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, tap, throwError } from 'rxjs';
import { Ifammembers } from './Imembers';
@Injectable({
  providedIn: 'root',
})

export class UserService {
  // private userUrl = "https://mifamily-app.herokuapp.com";
  private userUrl = "http://localhost:5000";

  constructor(private http: HttpClient) { };

  getUsers(): Observable<Ifammembers[]> {
    return this.http.get<Ifammembers[]>(this.userUrl).pipe(
      tap(data => console.log('ALL:', JSON.stringify(data))),
      catchError(this.handleError)
    );
}

private handleError(err: HttpErrorResponse) {

    if (err.status === 0) {

    console.error('An error occurred:', err.error);
  } else {

    console.error(
      `Backend returned code ${err.status}, body was: `, err.error);
  }

  return throwError(() => new Error('Something bad happened; please try again later.'));

}
}

