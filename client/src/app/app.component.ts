import { Component } from '@angular/core';

@Component({
  selector: 'CoatnCode-root',
  template: `
  <div class="container text-center">

    <div class="col">
    <div style="text-align:center" class="content">
   <coat-FamTable></coat-FamTable>

    </div>
    </div>







</div>
   

    <router-outlet></router-outlet>
  `,
  styles: []
})
export class AppComponent {
  title = 'code and code';
}


// <h1>{{title}}</h1>
// <div class="col">
// <div style="text-align:center" class="content">
// <h1>{{title}}</h1>
// <app-events></app-events>
// </div>
// </div>
