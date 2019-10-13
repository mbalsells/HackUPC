import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { InitialAppComponent } from './initial-page/initial-app.component'
import { ScrollAppComponent } from './scroll-page/scroll-app.component'

const routes: Routes = [
  { path: '', component: InitialAppComponent },
  { path: 'user', component: ScrollAppComponent },
  { path: '**', redirectTo: '' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
