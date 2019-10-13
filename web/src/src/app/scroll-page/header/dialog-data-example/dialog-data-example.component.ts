import {Component} from '@angular/core';
import {MatDialog} from '@angular/material/dialog';

@Component({
 selector: 'app-dialog-data-example',
 templateUrl: './dialog-data-example.component.html',
 styleUrls: ['./dialog-data-example.component.scss']
})
export class DialogDataExampleComponent {
  constructor(public dialog: MatDialog) {}

  openDialog() {
    const dialogRef = this.dialog.open(DialogDataExampleComponentDialog);

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }
}

@Component({
  selector: 'dialog-content-example-dialog',
  templateUrl: 'dialog-data-example-component-dialog.html',
})
export class DialogDataExampleComponentDialog {}
