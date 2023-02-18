import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { Ifammembers } from './Imembers';
import { UserService } from './user.service';

@Component({
  selector: 'coat-users',
  templateUrl: './usersList.component.html',
  styleUrls: ['./userList.component.css']
})
export class UsersTable implements OnDestroy {
  jfammembers:any[]= [
    // {
    //    "ID":1,
    // "Firstname": "felix",
    // "Happiness": 5,
    // "Urlstr": "assets/images/wilmer.jpg"
    // },

    // {
    //     "ID":2,
    //  "Firstname": "doris",
    //  "Happiness": 4,
    //  "Urlstr": "assets/images/wilmer.jpg"
    //  },

    //  {
    //     "ID":3,
    //  "Firstname": "wilmer",
    //  "Happiness": 2,
    //  "Urlstr": "assets/images/wilmer.jpg"
    //  },
]


  cardTitle: string = "Family Members";
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

  filterMembers: Ifammembers[] = [];
  fammembers: Ifammembers[] = [];

  constructor(private userService: UserService) { }

  performFilter(filterBy: string): Ifammembers[] {
    filterBy = filterBy.toLowerCase();
    return this.fammembers.filter((member: Ifammembers) =>
      member.Firstname.toLocaleLowerCase().includes(filterBy)
    );
  }

  // fam: string = "members";

  toggleImage(): void {
    this.showImage = !this.showImage;
  }

  // ngOnInit(): void {
  //   this.sub = this.userService.getUsers().subscribe({
  //     next: (mem) => {
  //       this.fammembers = mem;
  //       this.filterMembers = this.fammembers;
  //     },
  //     error: (err) => (this.errorMessage = err),
  //   });
  // }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }

  onRatingClicked(message: string): void {
    this.cardTitle = "product list: " + message;
  }
}