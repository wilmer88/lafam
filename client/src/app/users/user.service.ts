import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, tap, throwError } from 'rxjs';
import { Imembers } from './Iuser';
@Injectable({
  providedIn: 'root',
})
export class UserService {
  private userUrl = '/';
  constructor(private http: HttpClient) {}

  getUsers(): Observable<Imembers[]> {
    return this.http.get<Imembers[]>(this.userUrl).pipe(
      tap((data) => console.log('ALL', JSON.stringify(data))),
      catchError(this.handleError)
    );
    // getUsers(): Imembers[]{
    // return [
    //         {ID: 0, FirstName:"wilmer", UrlStr: "assets/images/wilmer.jpg", Happiness:3.3},
    //         {ID: 1,FirstName:"doris", UrlStr: "assets/images/wilmer.jpg", Happiness:4.7},
    //         {ID: 10, FirstName:"felix", UrlStr: "assets/images/wilmer.jpg", Happiness:4.5 },
    //     ]
  }
  private handleError(err: HttpErrorResponse) {
    let errorMassage = '';
    if (err.error instanceof ErrorEvent) {
      errorMassage = `an error occurred: ${err.error.message}`;
    }
    console.error(errorMassage);
    return throwError(() => errorMassage);
  }
}
