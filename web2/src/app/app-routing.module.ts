import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {InitScreenComponent} from "./init-screen/init-screen.component";
import {UserScreenComponent} from "./user-screen/user-screen.component";
import {FeedbackComponent} from "./feedback/feedback.component";
import {CommunityComponent} from "./community/community.component";

const routes: Routes = [
  {path: '', component: InitScreenComponent},
  {path: 'user/:username', component: UserScreenComponent},
  {path: 'feedback/:username', component: FeedbackComponent},
  {path: 'community/:username', component: CommunityComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
