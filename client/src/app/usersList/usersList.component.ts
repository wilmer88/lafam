import { Component, isStandalone, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { Ifammembers } from './Imembers';
import { UserService } from './user.service';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';

 @Component({
  selector: 'coat-users',
  templateUrl: './usersList.component.html',
  styleUrls: ['./userList.component.css']
})

export class UsersTable implements OnInit {
  cardTitle: string = "Family Members";
  imageWidth: number = 80;
  imageMargin: number = 5;
  showImage = false;
  errorMessage: string = '';
  sub!: Subscription;

  private _listFilter: string = '';

  get listFilter(): string {
    return this._listFilter;
  };

  set listFilter(value: string) {
    this._listFilter = value;
    console.log("in setter: ", value);
    this.filterMembers = this.performFilter(value);
  };

  filterMembers: Ifammembers[] = [];
  jfammembers:Ifammembers[]= [];

  

  performFilter(filterBy: string): Ifammembers[] {
    filterBy = filterBy.toLowerCase();
    return this.jfammembers.filter((member: Ifammembers) =>
      member.Firstname.toLocaleLowerCase().includes(filterBy));
  }

  toggleImage(): void {
    this.showImage = !this.showImage;
  }
  constructor(private httpUserService: UserService) { }
  ngOnInit(): void {
    this.sub = this.httpUserService.getUsers().subscribe({
      next: (mem) => {
        this.jfammembers = mem;
        this.filterMembers= this.jfammembers;
      },
      error: err => this.errorMessage = err,
    }); 
  }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }

  onRatingClicked(message: string): void {
    this.cardTitle = "product list: " + message;
  }

}