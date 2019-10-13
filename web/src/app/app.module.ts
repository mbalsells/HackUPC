import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { InitialAppComponent } from './initial-page/initial-app.component';
import { AppComponent } from './app.component';
import { LoginBoxComponent } from './initial-page/login-box/login-box.component';
import { RegisterBoxComponent } from './initial-page/register-box/register-box.component';
import { CentralContainerComponent } from './initial-page/central-container/central-container.component';
import { ScrollAppComponent } from './scroll-page/scroll-app.component';
import { ScrollComponent } from './scroll-page/scroll/scroll.component';
import { ScrollEmbedderComponent } from './scroll-page/scroll-embedder/scroll-embedder.component';
import { HeaderComponent } from './scroll-page/header/header.component';
import { DialogDataExampleComponent, DialogDataExampleComponentDialog } from './scroll-page/header/dialog-data-example/dialog-data-example.component';

import {
  MatButtonModule, MatCardModule, MatDialogModule, MatInputModule, MatTableModule, MatStepperModule, MatSliderModule,
  MatRadioModule, MatToolbarModule, MatMenuModule, MatIconModule, MatProgressSpinnerModule, MatTabsModule
} from '@angular/material';

@NgModule({
  declarations: [
    AppComponent,
    LoginBoxComponent,
    RegisterBoxComponent,
    CentralContainerComponent,
    HeaderComponent,
    ScrollComponent,
    ScrollEmbedderComponent,
    ScrollAppComponent,
    ScrollAppComponent,
    InitialAppComponent,
    DialogDataExampleComponent,
    DialogDataExampleComponentDialog
 ],
  imports: [
    BrowserModule,
    AppRoutingModule, HttpClientModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatIconModule,
    MatCardModule,
    MatRadioModule,
    MatStepperModule,
    MatTabsModule,
    CommonModule,
    MatToolbarModule,
    MatInputModule,
    MatDialogModule,
    MatSliderModule,
    MatTableModule,
    MatMenuModule,
    MatProgressSpinnerModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent],
  exports: [
    CommonModule,
    MatToolbarModule,
    MatButtonModule,
    MatCardModule,
    MatInputModule,
    MatDialogModule,
    MatTableModule,
    MatMenuModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatSliderModule
  ],
  entryComponents: [
    DialogDataExampleComponent,
    DialogDataExampleComponentDialog
  ],
})
export class AppModule { }
