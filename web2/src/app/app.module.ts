import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { InitScreenComponent } from './init-screen/init-screen.component';
import { UserScreenComponent } from './user-screen/user-screen.component';
import { CentralContainerComponent } from './init-screen/central-container/central-container.component';
import { HeaderComponent } from './user-screen/header/header.component';
import { ScrollComponent } from './user-screen/scroll-embedder/scroll/scroll.component';
import { LoginBoxComponent } from './init-screen/central-container/login-box/login-box.component';
import { RegisterBoxComponent } from './init-screen/central-container/register-box/register-box.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MatButtonModule, MatCardModule, MatDialogModule, MatInputModule, MatTableModule, MatStepperModule, MatSliderModule,
  MatRadioModule, MatToolbarModule, MatMenuModule, MatIconModule, MatProgressSpinnerModule, MatTabsModule, MatListModule
} from '@angular/material';

import { HttpClientModule } from '@angular/common/http';
import { DialogDataExampleComponent, DialogDataExampleComponentDialog } from './user-screen/header/dialog-data-example/dialog-data-example.component';
import { DialogOutComponent, DialogOutComponentDialog } from './user-screen/header/dialog-out/dialog-out.component';
import { FeedbackComponent } from './feedback/feedback.component';
import { CommunityComponent } from './community/community.component';

@NgModule({
  declarations: [
    AppComponent,
    InitScreenComponent,
    UserScreenComponent,
    CentralContainerComponent,
    HeaderComponent,
    ScrollComponent,
    LoginBoxComponent,
    RegisterBoxComponent,
    DialogDataExampleComponent, DialogDataExampleComponentDialog,
    DialogOutComponent, DialogOutComponentDialog, FeedbackComponent, CommunityComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MatButtonModule, MatCardModule, MatDialogModule, MatInputModule, MatTableModule, MatStepperModule, MatSliderModule,
    MatRadioModule, MatToolbarModule, MatMenuModule, MatIconModule, MatProgressSpinnerModule, MatTabsModule, BrowserAnimationsModule,
    FormsModule, ReactiveFormsModule, HttpClientModule, MatListModule
  ],
  providers: [],
  bootstrap: [AppComponent],
  entryComponents: [
    DialogDataExampleComponent,
    DialogDataExampleComponentDialog,
    DialogOutComponent,
    DialogOutComponentDialog
  ],
})
export class AppModule { }
