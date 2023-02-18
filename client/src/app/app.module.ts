import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import {HttpClientModule} from "@angular/common/http"

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
// import { JobsComponent } from './events/events.component';
// import { TodoComponent } from './person/todo.component';
import { UsersTable } from './users/userList.component';
import { StarComponent } from './shared/star.component';


@NgModule({
  declarations: [
    AppComponent,
    // JobsComponent,
    // TodoComponent,
    StarComponent,
    UsersTable
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
