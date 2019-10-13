import { Component, Inject, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { HTTPLessonService, FeedbackResponse, InfoUserResponse } from '../../httplesson.service';

@Component({
 selector: 'app-dialog-data-example',
 templateUrl: './dialog-data-example.component.html',
 styleUrls: ['./dialog-data-example.component.scss']
})
export class DialogDataExampleComponent {
  @Input() username: string;

  constructor(public dialog: MatDialog) {}

  openDialog() {
    const dialogRef = this.dialog.open(DialogDataExampleComponentDialog, {
        width:'600px',
        data: { username: this.username }
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }
}

@Component({
  selector: 'dialog-content-example-dialog',
  templateUrl: 'dialog-data-example-component-dialog.html',
})
export class DialogDataExampleComponentDialog implements OnInit {
  name: string;
  email: string;
  subjects: string[];
  feedbacks: number[];

  constructor(
    public dialogRef: MatDialogRef<DialogDataExampleComponentDialog>,
    @Inject(MAT_DIALOG_DATA) public data:any,
    private less: HTTPLessonService) {}

    ngOnInit() {
      var _this = this;
      this.less.infoUser(this.data.username)
      .subscribe(dt => {
        _this.name = dt.name;
        _this.email = dt.email;
        _this.subjects = dt.subjects;
        _this.feedbacks = dt.feedbacks;
      });
    }
}
