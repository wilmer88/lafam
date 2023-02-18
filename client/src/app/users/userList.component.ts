import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { Imembers } from './Iuser';
import { UserService } from './user.service';

@Component({
  selector: 'app-users',
  templateUrl: "./users_list.component.html",
  styleUrls: ["./user_list.component.css"]
})
export class UsersTable implements OnInit, OnDestroy {
  pageTitle: string = "Family Members";
  imageWidth: number = 80;
  imageMargin: number = 5;
  showImage = false;
  errorMessage: string = '';
  sub!: Subscription;
  private _listFilter: string = '';

  get listFilter(): string {
    return this._listFilter;
  }

  set listFilter(value: string) {
    this._listFilter = value;
    console.log("in setter: ", value);
    this.filterMembers = this.performFilter(value);
  }

  filterMembers: Imembers[] = [];
  members: Imembers[] = [];

  constructor(private userService: UserService) { }

  performFilter(filterBy: string): Imembers[] {
    filterBy = filterBy.toLowerCase();
    return this.members.filter((member: Imembers) =>
      member.Firstname.toLocaleLowerCase().includes(filterBy)
    );
  }

  // fam: string = "members";

  toggleImage(): void {
    this.showImage = !this.showImage;
  }

  ngOnInit(): void {
    this.sub = this.userService.getUsers().subscribe({
      next: (members) => {
        this.members = members;
        this.filterMembers = this.members;
      },
      error: (err) => (this.errorMessage = err),
    });
  }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }

  onRatingClicked(message: string): void {
    this.pageTitle = "product list: " + message;
  }
}