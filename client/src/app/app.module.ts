import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import {HttpClientModule} from "@angular/common/http"
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';


import { UsersTable } from './usersList/usersList.component';
import { StarComponent } from './shared/star.component';
// import { EventsCard } from './events/events.component';


@NgModule({
  declarations: [
    AppComponent,
    // EventsCard,
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
