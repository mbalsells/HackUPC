import { Component } from '@angular/core';
import {Router} from '@angular/router';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-dialog-out',
  templateUrl: './dialog-out.component.html',
  styleUrls: ['./dialog-out.component.scss']
})
export class DialogOutComponent {

  constructor(public dialog: MatDialog) { }

  openDialog() {
    const dialogRef = this.dialog.open(DialogOutComponentDialog);
    dialogRef.afterClosed().subscribe();
  }
}

@Component({
  selector: 'dialog-out-component-dialog',
  templateUrl: 'dialog-out.component.dialog.html',
})
export class DialogOutComponentDialog {
  constructor(private router: Router) {};

  goToLogin() {
    this.router.navigateByUrl("");
  }
}
