import { HttpClient, HttpErrorResponse, } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, tap, throwError } from 'rxjs';
import { Ifammembers } from './Imembers';
@Injectable({
  providedIn: 'root',
})

export class UserService {

  private userUrl = 'https://mifamily-app.herokuapp.com/lafamily';
 
  constructor(private http: HttpClient) { };

  getUsers(): Observable<Ifammembers[]> {
    return this.http.get<Ifammembers[]>(this.userUrl).pipe(
      tap(data => console.log('ALL:', JSON.stringify(data))),
      catchError(this.handleError)
    );
}

private handleError(err: HttpErrorResponse) {
//   let errorMassage = '';
//   if (err.error instanceof ErrorEvent) {
//     errorMassage = `An error occurred: ${err.error.message}`;
//   } else{
//     errorMassage = `Server returned code: ${err.status}, error message is: ${err.message}`
//   };
//   console.error(errorMassage);
//   return throwError(() => errorMassage)
// }


    if (err.status === 0) {
    // A client-side or network error occurred. Handle it accordingly.
    console.error('An error occurred:', err.error);
  } else {
    // The backend returned an unsuccessful response code.
    // The response body may contain clues as to what went wrong.
    console.error(
      `Backend returned code ${err.status}, body was: `, err.error);
  }
  // Return an observable with a user-facing error message.
  return throwError(() => new Error('Something bad happened; please try again later.'));

}
}



 // private userUrl = 'https://mifamily-app.herokuapp.com/lafamily';
  // constructor(private http: HttpClient) { };

  // getUsers(): Observable<Ifammembers[]> {
  //   console.log("getting fam called");
  //   return this.http.get<Ifammembers[]>(this.userUrl).pipe(
  //     tap(data => console.log('ALL:', JSON.stringify(data))),
  //     catchError(this.handleError)
  //   )
  
//   getUsers(): Observable<Ifammembers[]> {
//     console.log("getting fam called");
//     return this.http.get<Ifammembers[]>('https://localhost:8080/api/lafamily') 
   
// };

// private handleError<T> (operation = 'operation', result?: T) {
//   return (error:any): Observable<T> => {
//     console.error(error);
//     return of(result as T);
//   };
// };



