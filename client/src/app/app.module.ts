import { APP_INITIALIZER, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import {HttpClientModule} from "@angular/common/http";
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
// import {ApiService} from "./services/api.service"
import { FamTable } from './usersList/FamTable.component';
import { StarComponent } from './shared/star.component';
// import { EventsComponent } from './events/Events.component';
import { RouterModule } from '@angular/router';
import { EventsComponent1 } from './events1/Events1.component';
// import { EventsCard } from './events/events.component';



@NgModule({
  declarations: [
    AppComponent,
    // EventsCard,
    StarComponent,
    FamTable,
    EventsComponent1,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    RouterModule.forRoot([
      {path:"famTable", component: FamTable},
      {path: "events", component: EventsComponent1}
    ])
  ],
  providers: [
    // ApiService

  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
